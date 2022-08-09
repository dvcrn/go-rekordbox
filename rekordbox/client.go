package rekordbox

import (
	"encoding/base64"

	"github.com/dvcrn/go-rekordbox/internal/supbox"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

func NewClient(optionsFilePath string, asarFilePath string) (*Client, error) {
	// Files and paths
	rekordboxConfig := supbox.ReadRekordboxConfig(optionsFilePath)

	// {"options":[["db-path","/Users/david/Library/Pioneer/rekordbox/master.db"], ...}
	databaseFilePath := rekordboxConfig.Options[0][1]

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

func (c *Client) GetDB() *sqlx.DB {
	return c.db
}
