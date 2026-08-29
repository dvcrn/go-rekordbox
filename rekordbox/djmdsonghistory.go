package rekordbox

import (
	"context"

	nulltype "github.com/mattn/go-nulltype"
)

func (c *Client) RecentDjmdSongHistory(ctx context.Context, limit int) ([]*DjmdSongHistory, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdSongHistory ORDER BY created_at DESC LIMIT $1`
	rows, err := db.QueryContext(ctx, sqlstr, limit)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()

	res, err := scanDjmdSongHistoryRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHistoryByContentID returns every history entry for a track.
func (c *Client) DjmdSongHistoryByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdSongHistory, error) {
	const query = `SELECT * FROM djmdSongHistory WHERE ContentID = $1`
	logf(query, contentID)

	rows, err := c.db.QueryContext(ctx, query, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	history, err := scanDjmdSongHistoryRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return history, nil
}
