package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/dvcrn/go-rekordbox/internal/supbox"
)

func main() {
	// Files and paths
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Unable to determine your home directory to location options.json")
	}

	optionsFilePath := filepath.Join(homeDir, "/Library/Application Support/Pioneer/rekordboxAgent/storage/", "options.json")

	rekordboxConfig := supbox.ReadRekordboxConfig(optionsFilePath)

	// extract .app path from options.json, we need this to construct
	// the path to the .asar file
	langPath := rekordboxConfig.Options[5][1]
	rgp, err := regexp.Compile(".*/rekordbox.app")
	if err != nil {
		panic(err)
	}

	rekordboxPath := rgp.FindString(langPath)

	if len(rekordboxPath) < 5 {
		panic(fmt.Errorf("could not find rekordbox path"))
	}

	asarFilePath := filepath.Join(rekordboxPath, "/Contents/MacOS/rekordboxAgent.app/Contents/Resources/app.asar")

	// Database decryption
	encodedPasswordData := supbox.GetEncryptedPasswordDataFromConfig(rekordboxConfig)
	decodedPasswordData, err := base64.StdEncoding.DecodeString(encodedPasswordData)
	if err != nil {
		panic(err)
	}

	passwordString := supbox.GetEncryptedPassword(asarFilePath)
	password := []byte(passwordString)
	decryptedBytes := supbox.Decrypt(decodedPasswordData, password)
	encryptionKey := string(decryptedBytes)

	fmt.Print(encryptionKey)
}
