package schema

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const encryptionKey = "402fd482c38817c35ffa8ffb8c7d93143b749e7d315df7a81732a1ff43608497"

type Config struct {
	DatabasePath string
	Repository   string
	Stdout       io.Writer
}

func Generate(ctx context.Context, config Config) error {
	resolved, err := resolveConfig(config)
	if err != nil {
		return err
	}

	tools, err := findTools()
	if err != nil {
		return err
	}

	workDir, err := os.MkdirTemp("", "go-rekordbox-")
	if err != nil {
		return fmt.Errorf("create temporary directory: %w", err)
	}
	defer os.RemoveAll(workDir)

	plaintextDatabase := filepath.Join(workDir, "plaintext.db")
	generatedDir := filepath.Join(workDir, "generated")

	if err := decrypt(ctx, tools.sqlcipher, resolved.DatabasePath, plaintextDatabase); err != nil {
		return err
	}
	if err := checkIntegrity(ctx, tools.sqlite, plaintextDatabase); err != nil {
		return err
	}
	if err := os.Mkdir(generatedDir, 0o700); err != nil {
		return fmt.Errorf("create generated model directory: %w", err)
	}
	if err := generateModels(ctx, tools.xo, resolved.Repository, plaintextDatabase, generatedDir); err != nil {
		return err
	}

	schemaPath := filepath.Join(workDir, "schema.sql")
	if err := dumpSchema(ctx, tools.sqlite, plaintextDatabase, schemaPath); err != nil {
		return err
	}
	if err := installOutputs(resolved.Repository, generatedDir, schemaPath); err != nil {
		return err
	}

	fmt.Fprintf(resolved.Stdout, "generated db/schema.sql and rekordbox/*.xo.go from %s\n", resolved.DatabasePath)
	return nil
}

type resolvedConfig struct {
	DatabasePath string
	Repository   string
	Stdout       io.Writer
}

func resolveConfig(config Config) (resolvedConfig, error) {
	repository := config.Repository
	if repository == "" {
		var err error
		repository, err = os.Getwd()
		if err != nil {
			return resolvedConfig{}, fmt.Errorf("get working directory: %w", err)
		}
	}

	databasePath := config.DatabasePath
	if databasePath == "" {
		databasePath = os.Getenv("REKORDBOX_DATABASE")
	}
	if databasePath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return resolvedConfig{}, fmt.Errorf("get home directory: %w", err)
		}
		databasePath = filepath.Join(homeDir, "Library", "Pioneer", "rekordbox", "master.db")
	}

	info, err := os.Stat(databasePath)
	if err != nil {
		return resolvedConfig{}, fmt.Errorf("inspect Rekordbox database %q: %w", databasePath, err)
	}
	if !info.Mode().IsRegular() {
		return resolvedConfig{}, fmt.Errorf("Rekordbox database is not a regular file: %s", databasePath)
	}
	for _, path := range []string{
		filepath.Join(repository, "db"),
		filepath.Join(repository, "rekordbox"),
		filepath.Join(repository, "tpl"),
	} {
		if info, err := os.Stat(path); err != nil || !info.IsDir() {
			return resolvedConfig{}, fmt.Errorf("required repository directory is unavailable: %s", path)
		}
	}

	stdout := config.Stdout
	if stdout == nil {
		stdout = os.Stdout
	}

	return resolvedConfig{
		DatabasePath: databasePath,
		Repository:   repository,
		Stdout:       stdout,
	}, nil
}

type toolPaths struct {
	sqlcipher string
	sqlite    string
	xo        string
}
