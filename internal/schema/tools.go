package schema

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func findTools() (toolPaths, error) {
	paths := toolPaths{}
	for name, destination := range map[string]*string{
		"sqlcipher": &paths.sqlcipher,
		"sqlite3":   &paths.sqlite,
		"xo":        &paths.xo,
	} {
		path, err := exec.LookPath(name)
		if err != nil {
			return toolPaths{}, fmt.Errorf("find required command %q: %w", name, err)
		}
		*destination = path
	}
	return paths, nil
}

func decrypt(ctx context.Context, sqlcipher, encryptedDatabase, plaintextDatabase string) error {
	script := strings.Join([]string{
		"PRAGMA key = '" + encryptionKey + "';",
		"ATTACH DATABASE '" + sqliteQuote(plaintextDatabase) + "' AS plaintext KEY '';",
		"SELECT sqlcipher_export('plaintext');",
		"DETACH DATABASE plaintext;",
	}, "\n")

	command := exec.CommandContext(ctx, sqlcipher, "-bail", encryptedDatabase)
	command.Stdin = strings.NewReader(script)
	if output, err := command.CombinedOutput(); err != nil {
		return commandError("decrypt database", err, output)
	}
	return nil
}

func checkIntegrity(ctx context.Context, sqlite, database string) error {
	info, err := os.Stat(database)
	if err != nil {
		return fmt.Errorf("inspect decrypted database: %w", err)
	}
	if info.Size() == 0 {
		return errors.New("inspect decrypted database: export is empty")
	}

	output, err := exec.CommandContext(ctx, sqlite, database, "PRAGMA integrity_check;").CombinedOutput()
	if err != nil {
		return commandError("check decrypted database integrity", err, output)
	}
	if result := strings.TrimSpace(string(output)); result != "ok" {
		return fmt.Errorf("decrypted database failed integrity check: %s", result)
	}
	return nil
}

func generateModels(ctx context.Context, xo, repository, database, destination string) error {
	databaseURL := (&url.URL{Scheme: "file", Path: database}).String()
	command := exec.CommandContext(
		ctx,
		xo,
		"schema",
		databaseURL,
		"-o", destination,
		"--src", filepath.Join(repository, "tpl"),
		"--tpl-pkg", "rekordbox",
	)
	if output, err := command.CombinedOutput(); err != nil {
		return commandError("generate Go models", err, output)
	}
	return nil
}

func dumpSchema(ctx context.Context, sqlite, database, destination string) error {
	output, err := exec.CommandContext(ctx, sqlite, database, ".schema --nosys").CombinedOutput()
	if err != nil {
		return commandError("dump database schema", err, output)
	}
	if len(bytes.TrimSpace(output)) == 0 {
		return errors.New("dump database schema: sqlite3 returned an empty schema")
	}
	return os.WriteFile(destination, output, 0o600)
}

func sqliteQuote(value string) string {
	return strings.ReplaceAll(value, "'", "''")
}

func commandError(action string, err error, output []byte) error {
	detail := strings.TrimSpace(string(output))
	if detail == "" {
		return fmt.Errorf("%s: %w", action, err)
	}
	return fmt.Errorf("%s: %w: %s", action, err, detail)
}
