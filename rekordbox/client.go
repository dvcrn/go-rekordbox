package rekordbox

import (
	"github.com/dvcrn/go-rekordbox/internal/supbox"
	"github.com/jmoiron/sqlx"
	"github.com/nurcahyaari/sqlabst"
)

type Client struct {
	db *sqlabst.SqlAbst
}

func NewClient(optionsFilePath string) (*Client, error) {
	// Files and paths
	rekordboxConfig := supbox.ReadRekordboxConfig(optionsFilePath)

	// {"options":[["db-path","/Users/david/Library/Pioneer/rekordbox/master.db"], ...}
	databaseFilePath := rekordboxConfig.Options[0][1]

	// Using a static password like done on https://github.com/mganss/CueGen
	encryptionKey := string("402fd482c38817c35ffa8ffb8c7d93143b749e7d315df7a81732a1ff43608497")

	// Open the Database
	dsn := supbox.GetDatabaseDSN(databaseFilePath, encryptionKey)
	db, err := sqlx.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	return &Client{db: sqlabst.NewSqlAbst(db)}, nil
}

func (c *Client) Close() {
	c.db.Close()
}

func (c *Client) GetSqlAbst() *sqlabst.SqlAbst {
	return c.db
}

func (c *Client) GetDB() *sqlx.DB {
	return c.db.GetDB()
}
