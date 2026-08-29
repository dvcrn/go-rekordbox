package rekordbox

import (
	"context"

	nulltype "github.com/mattn/go-nulltype"
)

// ImageFileByTableNameTargetUUID returns every image for a database record.
func (c *Client) ImageFileByTableNameTargetUUID(ctx context.Context, tableName, targetUUID nulltype.NullString) ([]*ImageFile, error) {
	const query = `SELECT * FROM imageFile WHERE TableName = $1 AND TargetUUID = $2`
	logf(query, tableName, targetUUID)

	rows, err := c.db.QueryContext(ctx, query, tableName, targetUUID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	images, err := scanImageFileRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return images, nil
}
