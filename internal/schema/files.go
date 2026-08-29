package schema

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func installOutputs(repository, generatedDir, schemaPath string) error {
	generatedFiles, err := filepath.Glob(filepath.Join(generatedDir, "*.go"))
	if err != nil {
		return fmt.Errorf("list generated models: %w", err)
	}
	if len(generatedFiles) == 0 {
		return errors.New("install generated models: xo produced no Go files")
	}
	sort.Strings(generatedFiles)

	modelDir := filepath.Join(repository, "rekordbox")
	installed := make(map[string]struct{}, len(generatedFiles))
	for _, source := range generatedFiles {
		name := filepath.Base(source)
		if !strings.HasSuffix(name, ".xo.go") {
			return fmt.Errorf("install generated models: unexpected filename %q", name)
		}
		installed[name] = struct{}{}
	}
	for _, source := range generatedFiles {
		name := filepath.Base(source)
		destination := filepath.Join(modelDir, name)
		if err := replaceFile(source, destination, 0o644); err != nil {
			return fmt.Errorf("install generated model %q: %w", name, err)
		}
	}

	existingFiles, err := filepath.Glob(filepath.Join(modelDir, "*.xo.go"))
	if err != nil {
		return fmt.Errorf("list existing models: %w", err)
	}
	for _, existing := range existingFiles {
		if _, ok := installed[filepath.Base(existing)]; ok {
			continue
		}
		if err := os.Remove(existing); err != nil {
			return fmt.Errorf("remove stale generated model %q: %w", filepath.Base(existing), err)
		}
	}

	if err := replaceFile(schemaPath, filepath.Join(repository, "db", "schema.sql"), 0o644); err != nil {
		return fmt.Errorf("install schema: %w", err)
	}
	return nil
}

func replaceFile(source, destination string, mode fs.FileMode) (returnErr error) {
	contents, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	temporary, err := os.CreateTemp(filepath.Dir(destination), ".generate-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	closed := false
	defer func() {
		if !closed {
			if err := temporary.Close(); returnErr == nil {
				returnErr = err
			}
		}
		if err := os.Remove(temporaryPath); err != nil && !errors.Is(err, os.ErrNotExist) && returnErr == nil {
			returnErr = err
		}
	}()

	if err := temporary.Chmod(mode); err != nil {
		return err
	}
	if _, err := temporary.Write(contents); err != nil {
		return err
	}
	if err := temporary.Sync(); err != nil {
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	closed = true
	return os.Rename(temporaryPath, destination)
}
