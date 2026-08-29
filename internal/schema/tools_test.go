package schema

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

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
