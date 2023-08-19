package rekordbox

import (
	"context"

	nulltype "github.com/mattn/go-nulltype"
)

func (c *Client) DjmdRecommendLikeByAnyContentId(ctx context.Context, contentId nulltype.NullString) ([]*DjmdRecommendLike, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID1, ContentID2, LikeRate, DataCreatedH, DataCreatedL, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdRecommendLike ` +
		`WHERE ContentID1 = $1 OR ContentID2 = $1`
	// run
	logf(sqlstr, contentId)
	rows, err := db.QueryContext(ctx, sqlstr, contentId)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdRecommendLike
	for rows.Next() {
		drl := DjmdRecommendLike{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&drl.ID, &drl.ContentId1, &drl.ContentId2, &drl.LikeRate, &drl.DataCreatedH, &drl.DataCreatedL, &drl.UUID, &drl.RbDataStatus, &drl.RbLocalDataStatus, &drl.RbLocalDeleted, &drl.RbLocalSynced, &drl.Usn, &drl.RbLocalUsn, &drl.CreatedAt, &drl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &drl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
