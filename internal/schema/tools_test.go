package schema

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestDecryptIncludesWAL(t *testing.T) {
	sqlcipher, err := exec.LookPath("sqlcipher")
	if err != nil {
		t.Skip("sqlcipher is not installed")
	}
	sqlite, err := exec.LookPath("sqlite3")
	if err != nil {
		t.Skip("sqlite3 is not installed")
	}

	database := filepath.Join(t.TempDir(), "master.db")
	command := exec.Command(sqlcipher, "-batch", database)
	stdin, err := command.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout, err := command.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	var stderr bytes.Buffer
	command.Stderr = &stderr
	if err := command.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		_, _ = fmt.Fprintln(stdin, ".quit")
		_ = stdin.Close()
		_ = command.Wait()
	}()

	_, err = fmt.Fprintf(stdin, `PRAGMA key = '%s';
PRAGMA journal_mode = WAL;
PRAGMA wal_autocheckpoint = 0;
CREATE TABLE wal_only (value TEXT NOT NULL);
INSERT INTO wal_only VALUES ('included');
SELECT 'ready';
`, encryptionKey)
	if err != nil {
		t.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	ready := false
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "ready" {
			ready = true
			break
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	if !ready {
		t.Fatalf("create WAL fixture: %s", stderr.String())
	}
	if info, err := os.Stat(database + "-wal"); err != nil || info.Size() == 0 {
		t.Fatalf("WAL fixture is unavailable: %v", err)
	}

	plaintext := filepath.Join(t.TempDir(), "plaintext.db")
	if err := decrypt(context.Background(), sqlcipher, database, plaintext); err != nil {
		t.Fatal(err)
	}
	output, err := exec.Command(sqlite, plaintext, "SELECT value FROM wal_only;").CombinedOutput()
	if err != nil {
		t.Fatalf("query exported database: %v: %s", err, output)
	}
	if got, want := strings.TrimSpace(string(output)), "included"; got != want {
		t.Fatalf("exported value = %q, want %q", got, want)
	}
}

func TestCheckIntegrityRejectsEmptyDatabase(t *testing.T) {
	t.Parallel()

	database := filepath.Join(t.TempDir(), "empty.db")
	if err := os.WriteFile(database, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if err := checkIntegrity(context.Background(), "sqlite3", database); err == nil {
		t.Fatal("checkIntegrity() error = nil, want empty export error")
	}
}

func TestSQLiteQuote(t *testing.T) {
	t.Parallel()

	if got, want := sqliteQuote("it's/here"), "it''s/here"; got != want {
		t.Errorf("sqliteQuote() = %q, want %q", got, want)
	}
}
