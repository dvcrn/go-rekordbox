package main

import (
	"fmt"

	"github.com/dvcrn/go-rekordbox/internal/supbox"
)

func main() {
	dataPath := supbox.GetDataPath()
	databaseFilePath := supbox.GetDatabaseFilePath(dataPath)
	fmt.Print(databaseFilePath)
}
