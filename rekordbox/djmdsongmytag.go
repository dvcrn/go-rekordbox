package rekordbox

import (
	"context"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdSongMyTagByContentID returns every tag assignment for a track.
func (c *Client) DjmdSongMyTagByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdSongMyTag, error) {
	const query = `SELECT * FROM djmdSongMyTag WHERE ContentID = $1`
	return c.djmdSongMyTags(ctx, query, contentID)
}

// DjmdSongMyTagByMyTagID returns every track assignment for a tag.
func (c *Client) DjmdSongMyTagByMyTagID(ctx context.Context, myTagID nulltype.NullString) ([]*DjmdSongMyTag, error) {
	const query = `SELECT * FROM djmdSongMyTag WHERE MyTagID = $1`
	return c.djmdSongMyTags(ctx, query, myTagID)
}

func (c *Client) djmdSongMyTags(ctx context.Context, query string, value nulltype.NullString) ([]*DjmdSongMyTag, error) {
	logf(query, value)

	rows, err := c.db.QueryContext(ctx, query, value)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	tags, err := scanDjmdSongMyTagRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return tags, nil
}
