package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dvcrn/go-rekordbox/internal/supbox"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to determine your home directory to location options.json")
	}

	dataPath := filepath.Join(homeDir, "/Library/Pioneer/rekordbox")

	databaseFilePath := supbox.GetDatabaseFilePath(dataPath)
	fmt.Print(databaseFilePath)
}
