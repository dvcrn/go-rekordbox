package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdSongHotCueBanklist represents a row from 'djmdSongHotCueBanklist'.
type DjmdSongHotCueBanklist struct {
	ID                 nulltype.NullString `json:"id"`                    // ID
	HotCueBanklistID   nulltype.NullString `json:"hot_cue_banklist_id"`   // HotCueBanklistID
	ContentID          nulltype.NullString `json:"content_id"`            // ContentID
	TrackNo            nulltype.NullInt64  `json:"track_no"`              // TrackNo
	CueID              nulltype.NullString `json:"cue_id"`                // CueID
	InMsec             nulltype.NullInt64  `json:"in_msec"`               // InMsec
	InFrame            nulltype.NullInt64  `json:"in_frame"`              // InFrame
	InMpegFrame        nulltype.NullInt64  `json:"in_mpeg_frame"`         // InMpegFrame
	InMpegAbs          nulltype.NullInt64  `json:"in_mpeg_abs"`           // InMpegAbs
	OutMsec            nulltype.NullInt64  `json:"out_msec"`              // OutMsec
	OutFrame           nulltype.NullInt64  `json:"out_frame"`             // OutFrame
	OutMpegFrame       nulltype.NullInt64  `json:"out_mpeg_frame"`        // OutMpegFrame
	OutMpegAbs         nulltype.NullInt64  `json:"out_mpeg_abs"`          // OutMpegAbs
	Color              nulltype.NullInt64  `json:"color"`                 // Color
	ColorTableIndex    nulltype.NullInt64  `json:"color_table_index"`     // ColorTableIndex
	ActiveLoop         nulltype.NullInt64  `json:"active_loop"`           // ActiveLoop
	Comment            nulltype.NullString `json:"comment"`               // Comment
	BeatLoopSize       nulltype.NullInt64  `json:"beat_loop_size"`        // BeatLoopSize
	CueMicrosec        nulltype.NullInt64  `json:"cue_microsec"`          // CueMicrosec
	InPointSeekInfo    nulltype.NullString `json:"in_point_seek_info"`    // InPointSeekInfo
	OutPointSeekInfo   nulltype.NullString `json:"out_point_seek_info"`   // OutPointSeekInfo
	HotCueBanklistUUID nulltype.NullString `json:"hot_cue_banklist_uuid"` // HotCueBanklistUUID
	UUID               nulltype.NullString `json:"uuid"`                  // UUID
	RbDataStatus       nulltype.NullInt64  `json:"rb_data_status"`        // rb_data_status
	RbLocalDataStatus  nulltype.NullInt64  `json:"rb_local_data_status"`  // rb_local_data_status
	RbLocalDeleted     nulltype.NullInt64  `json:"rb_local_deleted"`      // rb_local_deleted
	RbLocalSynced      nulltype.NullInt64  `json:"rb_local_synced"`       // rb_local_synced
	Usn                nulltype.NullInt64  `json:"usn"`                   // usn
	RbLocalUsn         nulltype.NullInt64  `json:"rb_local_usn"`          // rb_local_usn
	CreatedAt          Time                `json:"created_at"`            // created_at
	UpdatedAt          Time                `json:"updated_at"`            // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjmdSongHotCueBanklist exists in the database.
func (dshcb *DjmdSongHotCueBanklist) Exists() bool {
	return dshcb._exists
}

// Deleted returns true when the DjmdSongHotCueBanklist has been marked for deletion from
// the database.
func (dshcb *DjmdSongHotCueBanklist) Deleted() bool {
	return dshcb._deleted
}

// Insert inserts the DjmdSongHotCueBanklist to the database.
func (c *Client) InsertDjmdSongHotCueBanklist(ctx context.Context, dshcb *DjmdSongHotCueBanklist) error {
	db := c.db

	switch {
	case dshcb._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dshcb._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdSongHotCueBanklist (` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31` +
		`)`
	// run
	logf(sqlstr, dshcb.ID, dshcb.HotCueBanklistID, dshcb.ContentID, dshcb.TrackNo, dshcb.CueID, dshcb.InMsec, dshcb.InFrame, dshcb.InMpegFrame, dshcb.InMpegAbs, dshcb.OutMsec, dshcb.OutFrame, dshcb.OutMpegFrame, dshcb.OutMpegAbs, dshcb.Color, dshcb.ColorTableIndex, dshcb.ActiveLoop, dshcb.Comment, dshcb.BeatLoopSize, dshcb.CueMicrosec, dshcb.InPointSeekInfo, dshcb.OutPointSeekInfo, dshcb.HotCueBanklistUUID, dshcb.UUID, dshcb.RbDataStatus, dshcb.RbLocalDataStatus, dshcb.RbLocalDeleted, dshcb.RbLocalSynced, dshcb.Usn, dshcb.RbLocalUsn, dshcb.CreatedAt, dshcb.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dshcb.ID, dshcb.HotCueBanklistID, dshcb.ContentID, dshcb.TrackNo, dshcb.CueID, dshcb.InMsec, dshcb.InFrame, dshcb.InMpegFrame, dshcb.InMpegAbs, dshcb.OutMsec, dshcb.OutFrame, dshcb.OutMpegFrame, dshcb.OutMpegAbs, dshcb.Color, dshcb.ColorTableIndex, dshcb.ActiveLoop, dshcb.Comment, dshcb.BeatLoopSize, dshcb.CueMicrosec, dshcb.InPointSeekInfo, dshcb.OutPointSeekInfo, dshcb.HotCueBanklistUUID, dshcb.UUID, dshcb.RbDataStatus, dshcb.RbLocalDataStatus, dshcb.RbLocalDeleted, dshcb.RbLocalSynced, dshcb.Usn, dshcb.RbLocalUsn, dshcb.CreatedAt, dshcb.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dshcb._exists = true
	return nil
}

// Update updates a DjmdSongHotCueBanklist in the database.
func (c *Client) UpdateDjmdSongHotCueBanklist(ctx context.Context, dshcb *DjmdSongHotCueBanklist) error {
	db := c.db

	switch {
	case !dshcb._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dshcb._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdSongHotCueBanklist SET ` +
		`HotCueBanklistID = $1, ContentID = $2, TrackNo = $3, CueID = $4, InMsec = $5, InFrame = $6, InMpegFrame = $7, InMpegAbs = $8, OutMsec = $9, OutFrame = $10, OutMpegFrame = $11, OutMpegAbs = $12, Color = $13, ColorTableIndex = $14, ActiveLoop = $15, Comment = $16, BeatLoopSize = $17, CueMicrosec = $18, InPointSeekInfo = $19, OutPointSeekInfo = $20, HotCueBanklistUUID = $21, UUID = $22, rb_data_status = $23, rb_local_data_status = $24, rb_local_deleted = $25, rb_local_synced = $26, usn = $27, rb_local_usn = $28, created_at = $29, updated_at = $30 ` +
		`WHERE ID = $31`
	// run
	logf(sqlstr, dshcb.HotCueBanklistID, dshcb.ContentID, dshcb.TrackNo, dshcb.CueID, dshcb.InMsec, dshcb.InFrame, dshcb.InMpegFrame, dshcb.InMpegAbs, dshcb.OutMsec, dshcb.OutFrame, dshcb.OutMpegFrame, dshcb.OutMpegAbs, dshcb.Color, dshcb.ColorTableIndex, dshcb.ActiveLoop, dshcb.Comment, dshcb.BeatLoopSize, dshcb.CueMicrosec, dshcb.InPointSeekInfo, dshcb.OutPointSeekInfo, dshcb.HotCueBanklistUUID, dshcb.UUID, dshcb.RbDataStatus, dshcb.RbLocalDataStatus, dshcb.RbLocalDeleted, dshcb.RbLocalSynced, dshcb.Usn, dshcb.RbLocalUsn, dshcb.CreatedAt, dshcb.UpdatedAt, dshcb.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dshcb.HotCueBanklistID, dshcb.ContentID, dshcb.TrackNo, dshcb.CueID, dshcb.InMsec, dshcb.InFrame, dshcb.InMpegFrame, dshcb.InMpegAbs, dshcb.OutMsec, dshcb.OutFrame, dshcb.OutMpegFrame, dshcb.OutMpegAbs, dshcb.Color, dshcb.ColorTableIndex, dshcb.ActiveLoop, dshcb.Comment, dshcb.BeatLoopSize, dshcb.CueMicrosec, dshcb.InPointSeekInfo, dshcb.OutPointSeekInfo, dshcb.HotCueBanklistUUID, dshcb.UUID, dshcb.RbDataStatus, dshcb.RbLocalDataStatus, dshcb.RbLocalDeleted, dshcb.RbLocalSynced, dshcb.Usn, dshcb.RbLocalUsn, dshcb.CreatedAt, dshcb.UpdatedAt, dshcb.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdSongHotCueBanklist to the database.
func (c *Client) SaveDjmdSongHotCueBanklist(ctx context.Context, dshcb *DjmdSongHotCueBanklist) error {
	if dshcb.Exists() {
		return c.UpdateDjmdSongHotCueBanklist(ctx, dshcb)
	}
	return c.InsertDjmdSongHotCueBanklist(ctx, dshcb)
}

// Upsert performs an upsert for DjmdSongHotCueBanklist.
func (c *Client) UpsertDjmdSongHotCueBanklist(ctx context.Context, dshcb *DjmdSongHotCueBanklist) error {
	db := c.db

	switch {
	case dshcb._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdSongHotCueBanklist (` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`HotCueBanklistID = EXCLUDED.HotCueBanklistID, ContentID = EXCLUDED.ContentID, TrackNo = EXCLUDED.TrackNo, CueID = EXCLUDED.CueID, InMsec = EXCLUDED.InMsec, InFrame = EXCLUDED.InFrame, InMpegFrame = EXCLUDED.InMpegFrame, InMpegAbs = EXCLUDED.InMpegAbs, OutMsec = EXCLUDED.OutMsec, OutFrame = EXCLUDED.OutFrame, OutMpegFrame = EXCLUDED.OutMpegFrame, OutMpegAbs = EXCLUDED.OutMpegAbs, Color = EXCLUDED.Color, ColorTableIndex = EXCLUDED.ColorTableIndex, ActiveLoop = EXCLUDED.ActiveLoop, Comment = EXCLUDED.Comment, BeatLoopSize = EXCLUDED.BeatLoopSize, CueMicrosec = EXCLUDED.CueMicrosec, InPointSeekInfo = EXCLUDED.InPointSeekInfo, OutPointSeekInfo = EXCLUDED.OutPointSeekInfo, HotCueBanklistUUID = EXCLUDED.HotCueBanklistUUID, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dshcb.ID, dshcb.HotCueBanklistID, dshcb.ContentID, dshcb.TrackNo, dshcb.CueID, dshcb.InMsec, dshcb.InFrame, dshcb.InMpegFrame, dshcb.InMpegAbs, dshcb.OutMsec, dshcb.OutFrame, dshcb.OutMpegFrame, dshcb.OutMpegAbs, dshcb.Color, dshcb.ColorTableIndex, dshcb.ActiveLoop, dshcb.Comment, dshcb.BeatLoopSize, dshcb.CueMicrosec, dshcb.InPointSeekInfo, dshcb.OutPointSeekInfo, dshcb.HotCueBanklistUUID, dshcb.UUID, dshcb.RbDataStatus, dshcb.RbLocalDataStatus, dshcb.RbLocalDeleted, dshcb.RbLocalSynced, dshcb.Usn, dshcb.RbLocalUsn, dshcb.CreatedAt, dshcb.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dshcb.ID, dshcb.HotCueBanklistID, dshcb.ContentID, dshcb.TrackNo, dshcb.CueID, dshcb.InMsec, dshcb.InFrame, dshcb.InMpegFrame, dshcb.InMpegAbs, dshcb.OutMsec, dshcb.OutFrame, dshcb.OutMpegFrame, dshcb.OutMpegAbs, dshcb.Color, dshcb.ColorTableIndex, dshcb.ActiveLoop, dshcb.Comment, dshcb.BeatLoopSize, dshcb.CueMicrosec, dshcb.InPointSeekInfo, dshcb.OutPointSeekInfo, dshcb.HotCueBanklistUUID, dshcb.UUID, dshcb.RbDataStatus, dshcb.RbLocalDataStatus, dshcb.RbLocalDeleted, dshcb.RbLocalSynced, dshcb.Usn, dshcb.RbLocalUsn, dshcb.CreatedAt, dshcb.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dshcb._exists = true
	return nil
}

// Delete deletes the DjmdSongHotCueBanklist from the database.
func (c *Client) DeleteDjmdSongHotCueBanklist(ctx context.Context, dshcb *DjmdSongHotCueBanklist) error {
	db := c.db

	switch {
	case !dshcb._exists: // doesn't exist
		return nil
	case dshcb._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdSongHotCueBanklist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dshcb.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dshcb.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dshcb._deleted = true
	return nil
}

func scanDjmdSongHotCueBanklistRows(rows *sql.Rows) ([]*DjmdSongHotCueBanklist, error) {
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdSongHotCueBanklist(ctx context.Context) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdSongHotCueBanklist`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdSongHotCueBanklistRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByContentID retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist__content_i_d'.
func (c *Client) DjmdSongHotCueBanklistByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByContentID(ctx context.Context, db DB, contentID nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE ContentID = $1`
	// run
	logf(sqlstr, contentID)
	rows, err := db.QueryContext(ctx, sqlstr, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByHotCueBanklistID retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist__hot_cue_banklist_i_d'.
func (c *Client) DjmdSongHotCueBanklistByHotCueBanklistID(ctx context.Context, hotCueBanklistID nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByHotCueBanklistID(ctx context.Context, db DB, hotCueBanklistID nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE HotCueBanklistID = $1`
	// run
	logf(sqlstr, hotCueBanklistID)
	rows, err := db.QueryContext(ctx, sqlstr, hotCueBanklistID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByHotCueBanklistUUID retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist__hot_cue_banklist_u_u_i_d'.
func (c *Client) DjmdSongHotCueBanklistByHotCueBanklistUUID(ctx context.Context, hotCueBanklistUUID nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByHotCueBanklistUUID(ctx context.Context, db DB, hotCueBanklistUUID nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE HotCueBanklistUUID = $1`
	// run
	logf(sqlstr, hotCueBanklistUUID)
	rows, err := db.QueryContext(ctx, sqlstr, hotCueBanklistUUID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByUUID retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist__u_u_i_d'.
func (c *Client) DjmdSongHotCueBanklistByUUID(ctx context.Context, uuid nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByRbDataStatus retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist_rb_data_status'.
func (c *Client) DjmdSongHotCueBanklistByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByRbLocalDataStatus retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist_rb_local_data_status'.
func (c *Client) DjmdSongHotCueBanklistByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByRbLocalDeleted retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist_rb_local_deleted'.
func (c *Client) DjmdSongHotCueBanklistByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByRbLocalUsnID retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'djmd_song_hot_cue_banklist_rb_local_usn__i_d'.
func (c *Client) DjmdSongHotCueBanklistByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongHotCueBanklist
	for rows.Next() {
		dshcb := DjmdSongHotCueBanklist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dshcb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongHotCueBanklistByID retrieves a row from 'djmdSongHotCueBanklist' as a DjmdSongHotCueBanklist.
//
// Generated from index 'sqlite_autoindex_djmdSongHotCueBanklist_1'.
func (c *Client) DjmdSongHotCueBanklistByID(ctx context.Context, id nulltype.NullString) (*DjmdSongHotCueBanklist, error) {
	// func DjmdSongHotCueBanklistByID(ctx context.Context, db DB, id nulltype.NullString) (*DjmdSongHotCueBanklist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, HotCueBanklistID, ContentID, TrackNo, CueID, InMsec, InFrame, InMpegFrame, InMpegAbs, OutMsec, OutFrame, OutMpegFrame, OutMpegAbs, Color, ColorTableIndex, ActiveLoop, Comment, BeatLoopSize, CueMicrosec, InPointSeekInfo, OutPointSeekInfo, HotCueBanklistUUID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongHotCueBanklist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dshcb := DjmdSongHotCueBanklist{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dshcb.ID, &dshcb.HotCueBanklistID, &dshcb.ContentID, &dshcb.TrackNo, &dshcb.CueID, &dshcb.InMsec, &dshcb.InFrame, &dshcb.InMpegFrame, &dshcb.InMpegAbs, &dshcb.OutMsec, &dshcb.OutFrame, &dshcb.OutMpegFrame, &dshcb.OutMpegAbs, &dshcb.Color, &dshcb.ColorTableIndex, &dshcb.ActiveLoop, &dshcb.Comment, &dshcb.BeatLoopSize, &dshcb.CueMicrosec, &dshcb.InPointSeekInfo, &dshcb.OutPointSeekInfo, &dshcb.HotCueBanklistUUID, &dshcb.UUID, &dshcb.RbDataStatus, &dshcb.RbLocalDataStatus, &dshcb.RbLocalDeleted, &dshcb.RbLocalSynced, &dshcb.Usn, &dshcb.RbLocalUsn, &dshcb.CreatedAt, &dshcb.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dshcb, nil
}
