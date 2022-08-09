package rekordbox

import (
	"encoding/base64"
	"fmt"
	"regexp"

	"github.com/dvcrn/go-rekordbox/internal/supbox"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func NewClient() (*Client, error) {
	// Files and paths
	rekordboxConfig := supbox.GetRekordboxConfig()

	// extract .app path
	langPath := rekordboxConfig.Options[5][1]
	rgp, err := regexp.Compile(".*/rekordbox.app")
	if err != nil {
		return nil, err
	}

	rekordboxPath := rgp.FindString(langPath)

	if len(rekordboxPath) < 5 {
		return nil, fmt.Errorf("could not find rekordbox path")
	}

	asarFilePath := supbox.GetAsarFilePath(rekordboxPath)
	dataPath := supbox.GetDataPath()
	databaseFilePath := supbox.GetDatabaseFilePath(dataPath)

	// Database decryption
	encodedPasswordData := supbox.GetEncryptedPasswordDataFromConfig(rekordboxConfig)
	decodedPasswordData, err := base64.StdEncoding.DecodeString(encodedPasswordData)
	if err != nil {
		return nil, err
	}

	passwordString := supbox.GetEncryptedPassword(asarFilePath)
	password := []byte(passwordString)
	decryptedBytes := supbox.Decrypt(decodedPasswordData, password)
	encryptionKey := string(decryptedBytes)

	// Open the Database
	dsn := supbox.GetDatabaseDSN(databaseFilePath, encryptionKey)
	db, err := sqlx.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	return &Client{db: db}, nil
}

func (c *Client) Close() {
	c.db.Close()
}
