package rekordbox

import "context"

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
