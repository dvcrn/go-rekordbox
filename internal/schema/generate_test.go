package schema

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveConfig(t *testing.T) {
	t.Parallel()

	repository := newRepository(t)
	database := filepath.Join(t.TempDir(), "master.db")
	if err := os.WriteFile(database, []byte("encrypted"), 0o600); err != nil {
		t.Fatal(err)
	}

	var stdout bytes.Buffer
	config, err := resolveConfig(Config{
		DatabasePath: database,
		Repository:   repository,
		Stdout:       &stdout,
	})
	if err != nil {
		t.Fatalf("resolveConfig() error = %v", err)
	}
	if config.DatabasePath != database {
		t.Errorf("DatabasePath = %q, want %q", config.DatabasePath, database)
	}
	if config.Repository != repository {
		t.Errorf("Repository = %q, want %q", config.Repository, repository)
	}
	if config.Stdout != &stdout {
		t.Error("Stdout does not match configured writer")
	}
}

func TestResolveConfigRejectsInvalidDatabase(t *testing.T) {
	t.Parallel()

	_, err := resolveConfig(Config{
		DatabasePath: t.TempDir(),
		Repository:   newRepository(t),
	})
	if err == nil {
		t.Fatal("resolveConfig() error = nil, want invalid database error")
	}
}

func TestResolveConfigUsesEnvironmentDatabase(t *testing.T) {
	repository := newRepository(t)
	database := filepath.Join(t.TempDir(), "master.db")
	if err := os.WriteFile(database, []byte("encrypted"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("REKORDBOX_DATABASE", database)

	config, err := resolveConfig(Config{Repository: repository})
	if err != nil {
		t.Fatalf("resolveConfig() error = %v", err)
	}
	if config.DatabasePath != database {
		t.Errorf("DatabasePath = %q, want %q", config.DatabasePath, database)
	}
}

func TestResolveConfigRejectsIncompleteRepository(t *testing.T) {
	t.Parallel()

	database := filepath.Join(t.TempDir(), "master.db")
	if err := os.WriteFile(database, []byte("encrypted"), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := resolveConfig(Config{
		DatabasePath: database,
		Repository:   t.TempDir(),
	})
	if err == nil {
		t.Fatal("resolveConfig() error = nil, want missing directory error")
	}
}

func newRepository(t *testing.T) string {
	t.Helper()

	repository := t.TempDir()
	for _, name := range []string{"db", "rekordbox", "tpl"} {
		if err := os.Mkdir(filepath.Join(repository, name), 0o700); err != nil {
			t.Fatal(err)
		}
	}
	return repository
}
