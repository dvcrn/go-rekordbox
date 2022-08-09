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
	// process
	var res []*DjmdSongHistory
	for rows.Next() {
		dsh := DjmdSongHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dsh.ID, &dsh.HistoryID, &dsh.ContentID, &dsh.TrackNo, &dsh.UUID, &dsh.RbDataStatus, &dsh.RbLocalDataStatus, &dsh.RbLocalDeleted, &dsh.RbLocalSynced, &dsh.Usn, &dsh.RbLocalUsn, &dsh.CreatedAt, &dsh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dsh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
