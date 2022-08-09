package main

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/dvcrn/go-rekordbox/internal/supbox"
)

func main() {
	// Files and paths
	rekordboxConfig := supbox.GetRekordboxConfig()
	// extract .app path
	langPath := rekordboxConfig.Options[5][1]
	rgp, err := regexp.Compile(".*/rekordbox.app")
	if err != nil {
		panic(err)
	}

	rekordboxPath := rgp.FindString(langPath)

	if len(rekordboxPath) < 5 {
		panic(fmt.Errorf("could not find rekordbox path"))
	}

	asarFilePath := supbox.GetAsarFilePath(rekordboxPath)

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
