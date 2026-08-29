package rekordbox

import (
	"context"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdCueByContentID returns every cue for a track.
func (c *Client) DjmdCueByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdCue, error) {
	const query = `SELECT * FROM djmdCue WHERE ContentID = $1`
	logf(query, contentID)

	rows, err := c.db.QueryContext(ctx, query, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	cues, err := scanDjmdCueRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return cues, nil
}
