package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdSongTagList represents a row from 'djmdSongTagList'.
type DjmdSongTagList struct {
	ID                nulltype.NullString `json:"ID"`                   // ID
	ContentID         nulltype.NullString `json:"ContentID"`            // ContentID
	TrackNo           nulltype.NullInt64  `json:"TrackNo"`              // TrackNo
	UUID              nulltype.NullString `json:"UUID"`                 // UUID
	RbDataStatus      nulltype.NullInt64  `json:"rb_data_status"`       // rb_data_status
	RbLocalDataStatus nulltype.NullInt64  `json:"rb_local_data_status"` // rb_local_data_status
	RbLocalDeleted    nulltype.NullInt64  `json:"rb_local_deleted"`     // rb_local_deleted
	RbLocalSynced     nulltype.NullInt64  `json:"rb_local_synced"`      // rb_local_synced
	Usn               nulltype.NullInt64  `json:"usn"`                  // usn
	RbLocalUsn        nulltype.NullInt64  `json:"rb_local_usn"`         // rb_local_usn
	CreatedAt         Time                `json:"created_at"`           // created_at
	UpdatedAt         Time                `json:"updated_at"`           // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjmdSongTagList exists in the database.
func (dstl *DjmdSongTagList) Exists() bool {
	return dstl._exists
}

// Deleted returns true when the DjmdSongTagList has been marked for deletion from
// the database.
func (dstl *DjmdSongTagList) Deleted() bool {
	return dstl._deleted
}

// Insert inserts the DjmdSongTagList to the database.
func (c *Client) InsertDjmdSongTagList(ctx context.Context, dstl *DjmdSongTagList) error {
	db := c.db

	switch {
	case dstl._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dstl._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdSongTagList (` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`)`
	// run
	logf(sqlstr, dstl.ID, dstl.ContentID, dstl.TrackNo, dstl.UUID, dstl.RbDataStatus, dstl.RbLocalDataStatus, dstl.RbLocalDeleted, dstl.RbLocalSynced, dstl.Usn, dstl.RbLocalUsn, dstl.CreatedAt, dstl.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dstl.ID, dstl.ContentID, dstl.TrackNo, dstl.UUID, dstl.RbDataStatus, dstl.RbLocalDataStatus, dstl.RbLocalDeleted, dstl.RbLocalSynced, dstl.Usn, dstl.RbLocalUsn, dstl.CreatedAt, dstl.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dstl._exists = true
	return nil
}

// Update updates a DjmdSongTagList in the database.
func (c *Client) UpdateDjmdSongTagList(ctx context.Context, dstl *DjmdSongTagList) error {
	db := c.db

	switch {
	case !dstl._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dstl._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdSongTagList SET ` +
		`ContentID = $1, TrackNo = $2, UUID = $3, rb_data_status = $4, rb_local_data_status = $5, rb_local_deleted = $6, rb_local_synced = $7, usn = $8, rb_local_usn = $9, created_at = $10, updated_at = $11 ` +
		`WHERE ID = $12`
	// run
	logf(sqlstr, dstl.ContentID, dstl.TrackNo, dstl.UUID, dstl.RbDataStatus, dstl.RbLocalDataStatus, dstl.RbLocalDeleted, dstl.RbLocalSynced, dstl.Usn, dstl.RbLocalUsn, dstl.CreatedAt, dstl.UpdatedAt, dstl.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dstl.ContentID, dstl.TrackNo, dstl.UUID, dstl.RbDataStatus, dstl.RbLocalDataStatus, dstl.RbLocalDeleted, dstl.RbLocalSynced, dstl.Usn, dstl.RbLocalUsn, dstl.CreatedAt, dstl.UpdatedAt, dstl.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdSongTagList to the database.
func (c *Client) SaveDjmdSongTagList(ctx context.Context, dstl *DjmdSongTagList) error {
	if dstl.Exists() {
		return c.UpdateDjmdSongTagList(ctx, dstl)
	}
	return c.InsertDjmdSongTagList(ctx, dstl)
}

// Upsert performs an upsert for DjmdSongTagList.
func (c *Client) UpsertDjmdSongTagList(ctx context.Context, dstl *DjmdSongTagList) error {
	db := c.db

	switch {
	case dstl._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdSongTagList (` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`ContentID = EXCLUDED.ContentID, TrackNo = EXCLUDED.TrackNo, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dstl.ID, dstl.ContentID, dstl.TrackNo, dstl.UUID, dstl.RbDataStatus, dstl.RbLocalDataStatus, dstl.RbLocalDeleted, dstl.RbLocalSynced, dstl.Usn, dstl.RbLocalUsn, dstl.CreatedAt, dstl.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dstl.ID, dstl.ContentID, dstl.TrackNo, dstl.UUID, dstl.RbDataStatus, dstl.RbLocalDataStatus, dstl.RbLocalDeleted, dstl.RbLocalSynced, dstl.Usn, dstl.RbLocalUsn, dstl.CreatedAt, dstl.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dstl._exists = true
	return nil
}

// Delete deletes the DjmdSongTagList from the database.
func (c *Client) DeleteDjmdSongTagList(ctx context.Context, dstl *DjmdSongTagList) error {
	db := c.db

	switch {
	case !dstl._exists: // doesn't exist
		return nil
	case dstl._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdSongTagList ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dstl.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dstl.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dstl._deleted = true
	return nil
}

func scanDjmdSongTagListRows(rows *sql.Rows) ([]*DjmdSongTagList, error) {
	var res []*DjmdSongTagList
	for rows.Next() {
		dstl := DjmdSongTagList{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dstl)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdSongTagList(ctx context.Context) ([]*DjmdSongTagList, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdSongTagList`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdSongTagListRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongTagListByContentID retrieves a row from 'djmdSongTagList' as a DjmdSongTagList.
//
// Generated from index 'djmd_song_tag_list__content_i_d'.
func (c *Client) DjmdSongTagListByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdSongTagList, error) {
	// func DjmdSongTagListByContentID(ctx context.Context, db DB, contentID nulltype.NullString) ([]*DjmdSongTagList, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongTagList ` +
		`WHERE ContentID = $1`
	// run
	logf(sqlstr, contentID)
	rows, err := db.QueryContext(ctx, sqlstr, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongTagList
	for rows.Next() {
		dstl := DjmdSongTagList{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dstl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongTagListByUUID retrieves a row from 'djmdSongTagList' as a DjmdSongTagList.
//
// Generated from index 'djmd_song_tag_list__u_u_i_d'.
func (c *Client) DjmdSongTagListByUUID(ctx context.Context, uuid nulltype.NullString) ([]*DjmdSongTagList, error) {
	// func DjmdSongTagListByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*DjmdSongTagList, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongTagList ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongTagList
	for rows.Next() {
		dstl := DjmdSongTagList{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dstl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongTagListByRbDataStatus retrieves a row from 'djmdSongTagList' as a DjmdSongTagList.
//
// Generated from index 'djmd_song_tag_list_rb_data_status'.
func (c *Client) DjmdSongTagListByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*DjmdSongTagList, error) {
	// func DjmdSongTagListByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*DjmdSongTagList, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongTagList ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongTagList
	for rows.Next() {
		dstl := DjmdSongTagList{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dstl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongTagListByRbLocalDataStatus retrieves a row from 'djmdSongTagList' as a DjmdSongTagList.
//
// Generated from index 'djmd_song_tag_list_rb_local_data_status'.
func (c *Client) DjmdSongTagListByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdSongTagList, error) {
	// func DjmdSongTagListByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdSongTagList, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongTagList ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongTagList
	for rows.Next() {
		dstl := DjmdSongTagList{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dstl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongTagListByRbLocalDeleted retrieves a row from 'djmdSongTagList' as a DjmdSongTagList.
//
// Generated from index 'djmd_song_tag_list_rb_local_deleted'.
func (c *Client) DjmdSongTagListByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*DjmdSongTagList, error) {
	// func DjmdSongTagListByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*DjmdSongTagList, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongTagList ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongTagList
	for rows.Next() {
		dstl := DjmdSongTagList{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dstl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongTagListByRbLocalUsnID retrieves a row from 'djmdSongTagList' as a DjmdSongTagList.
//
// Generated from index 'djmd_song_tag_list_rb_local_usn__i_d'.
func (c *Client) DjmdSongTagListByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdSongTagList, error) {
	// func DjmdSongTagListByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdSongTagList, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongTagList ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSongTagList
	for rows.Next() {
		dstl := DjmdSongTagList{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dstl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSongTagListByID retrieves a row from 'djmdSongTagList' as a DjmdSongTagList.
//
// Generated from index 'sqlite_autoindex_djmdSongTagList_1'.
func (c *Client) DjmdSongTagListByID(ctx context.Context, id nulltype.NullString) (*DjmdSongTagList, error) {
	// func DjmdSongTagListByID(ctx context.Context, db DB, id nulltype.NullString) (*DjmdSongTagList, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, TrackNo, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSongTagList ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dstl := DjmdSongTagList{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dstl.ID, &dstl.ContentID, &dstl.TrackNo, &dstl.UUID, &dstl.RbDataStatus, &dstl.RbLocalDataStatus, &dstl.RbLocalDeleted, &dstl.RbLocalSynced, &dstl.Usn, &dstl.RbLocalUsn, &dstl.CreatedAt, &dstl.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dstl, nil
}
