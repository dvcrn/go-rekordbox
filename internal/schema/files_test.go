package schema

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInstallOutputs(t *testing.T) {
	t.Parallel()

	repository := newRepository(t)
	modelDir := filepath.Join(repository, "rekordbox")
	generatedDir := t.TempDir()
	schemaPath := filepath.Join(t.TempDir(), "schema.sql")

	writeTestFile(t, filepath.Join(modelDir, "keep.go"), "package rekordbox\n")
	writeTestFile(t, filepath.Join(modelDir, "stale.xo.go"), "stale\n")
	writeTestFile(t, filepath.Join(generatedDir, "fresh.xo.go"), "fresh\n")
	writeTestFile(t, schemaPath, "CREATE TABLE fresh;\n")

	if err := installOutputs(repository, generatedDir, schemaPath); err != nil {
		t.Fatalf("installOutputs() error = %v", err)
	}
	assertTestFile(t, filepath.Join(modelDir, "keep.go"), "package rekordbox\n")
	assertTestFile(t, filepath.Join(modelDir, "fresh.xo.go"), "fresh\n")
	assertTestFile(t, filepath.Join(repository, "db", "schema.sql"), "CREATE TABLE fresh;\n")
	if _, err := os.Stat(filepath.Join(modelDir, "stale.xo.go")); !os.IsNotExist(err) {
		t.Errorf("stale generated model still exists: %v", err)
	}
}

func TestInstallOutputsRejectsUnexpectedFile(t *testing.T) {
	t.Parallel()

	repository := newRepository(t)
	generatedDir := t.TempDir()
	schemaPath := filepath.Join(t.TempDir(), "schema.sql")
	writeTestFile(t, filepath.Join(generatedDir, "unexpected.go"), "package rekordbox\n")
	writeTestFile(t, schemaPath, "CREATE TABLE fresh;\n")

	if err := installOutputs(repository, generatedDir, schemaPath); err == nil {
		t.Fatal("installOutputs() error = nil, want unexpected filename error")
	}
}

func TestInstallOutputsRejectsEmptyDirectory(t *testing.T) {
	t.Parallel()

	schemaPath := filepath.Join(t.TempDir(), "schema.sql")
	writeTestFile(t, schemaPath, "CREATE TABLE fresh;\n")
	if err := installOutputs(newRepository(t), t.TempDir(), schemaPath); err == nil {
		t.Fatal("installOutputs() error = nil, want empty model directory error")
	}
}

func writeTestFile(t *testing.T, path, contents string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
}

func assertTestFile(t *testing.T, path, want string) {
	t.Helper()
	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(contents); got != want {
		t.Errorf("contents of %s = %q, want %q", path, got, want)
	}
}
