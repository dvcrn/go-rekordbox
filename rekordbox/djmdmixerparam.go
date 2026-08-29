package rekordbox

import (
	"context"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdMixerParamByContentID returns every mixer parameter row for a track.
func (c *Client) DjmdMixerParamByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdMixerParam, error) {
	const query = `SELECT * FROM djmdMixerParam WHERE ContentID = $1`
	logf(query, contentID)

	rows, err := c.db.QueryContext(ctx, query, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	parameters, err := scanDjmdMixerParamRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return parameters, nil
}
