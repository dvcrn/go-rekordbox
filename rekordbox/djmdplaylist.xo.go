package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DjmdPlaylist represents a row from 'djmdPlaylist'.
type DjmdPlaylist struct {
	ID                sql.NullString `json:"ID"`                   // ID
	Seq               sql.NullInt64  `json:"Seq"`                  // Seq
	Name              sql.NullString `json:"Name"`                 // Name
	ImagePath         sql.NullString `json:"ImagePath"`            // ImagePath
	Attribute         sql.NullInt64  `json:"Attribute"`            // Attribute
	ParentID          sql.NullString `json:"ParentID"`             // ParentID
	SmartList         sql.NullString `json:"SmartList"`            // SmartList
	UUID              sql.NullString `json:"UUID"`                 // UUID
	RbDataStatus      sql.NullInt64  `json:"rb_data_status"`       // rb_data_status
	RbLocalDataStatus sql.NullInt64  `json:"rb_local_data_status"` // rb_local_data_status
	RbLocalDeleted    sql.NullInt64  `json:"rb_local_deleted"`     // rb_local_deleted
	RbLocalSynced     sql.NullInt64  `json:"rb_local_synced"`      // rb_local_synced
	Usn               sql.NullInt64  `json:"usn"`                  // usn
	RbLocalUsn        sql.NullInt64  `json:"rb_local_usn"`         // rb_local_usn
	CreatedAt         Time           `json:"created_at"`           // created_at
	UpdatedAt         Time           `json:"updated_at"`           // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjmdPlaylist exists in the database.
func (dp *DjmdPlaylist) Exists() bool {
	return dp._exists
}

// Deleted returns true when the DjmdPlaylist has been marked for deletion from
// the database.
func (dp *DjmdPlaylist) Deleted() bool {
	return dp._deleted
}

// Insert inserts the DjmdPlaylist to the database.
func (c *Client) InsertDjmdPlaylist(ctx context.Context, dp *DjmdPlaylist) error {
	db := c.db

	switch {
	case dp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdPlaylist (` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`)`
	// run
	logf(sqlstr, dp.ID, dp.Seq, dp.Name, dp.ImagePath, dp.Attribute, dp.ParentID, dp.SmartList, dp.UUID, dp.RbDataStatus, dp.RbLocalDataStatus, dp.RbLocalDeleted, dp.RbLocalSynced, dp.Usn, dp.RbLocalUsn, dp.CreatedAt, dp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dp.ID, dp.Seq, dp.Name, dp.ImagePath, dp.Attribute, dp.ParentID, dp.SmartList, dp.UUID, dp.RbDataStatus, dp.RbLocalDataStatus, dp.RbLocalDeleted, dp.RbLocalSynced, dp.Usn, dp.RbLocalUsn, dp.CreatedAt, dp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dp._exists = true
	return nil
}

// Update updates a DjmdPlaylist in the database.
func (c *Client) UpdateDjmdPlaylist(ctx context.Context, dp *DjmdPlaylist) error {
	db := c.db

	switch {
	case !dp._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dp._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdPlaylist SET ` +
		`Seq = $1, Name = $2, ImagePath = $3, Attribute = $4, ParentID = $5, SmartList = $6, UUID = $7, rb_data_status = $8, rb_local_data_status = $9, rb_local_deleted = $10, rb_local_synced = $11, usn = $12, rb_local_usn = $13, created_at = $14, updated_at = $15 ` +
		`WHERE ID = $16`
	// run
	logf(sqlstr, dp.Seq, dp.Name, dp.ImagePath, dp.Attribute, dp.ParentID, dp.SmartList, dp.UUID, dp.RbDataStatus, dp.RbLocalDataStatus, dp.RbLocalDeleted, dp.RbLocalSynced, dp.Usn, dp.RbLocalUsn, dp.CreatedAt, dp.UpdatedAt, dp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dp.Seq, dp.Name, dp.ImagePath, dp.Attribute, dp.ParentID, dp.SmartList, dp.UUID, dp.RbDataStatus, dp.RbLocalDataStatus, dp.RbLocalDeleted, dp.RbLocalSynced, dp.Usn, dp.RbLocalUsn, dp.CreatedAt, dp.UpdatedAt, dp.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdPlaylist to the database.
func (c *Client) SaveDjmdPlaylist(ctx context.Context, dp *DjmdPlaylist) error {
	if dp.Exists() {
		return c.UpdateDjmdPlaylist(ctx, dp)
	}
	return c.InsertDjmdPlaylist(ctx, dp)
}

// Upsert performs an upsert for DjmdPlaylist.
func (c *Client) UpsertDjmdPlaylist(ctx context.Context, dp *DjmdPlaylist) error {
	db := c.db

	switch {
	case dp._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdPlaylist (` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`Seq = EXCLUDED.Seq, Name = EXCLUDED.Name, ImagePath = EXCLUDED.ImagePath, Attribute = EXCLUDED.Attribute, ParentID = EXCLUDED.ParentID, SmartList = EXCLUDED.SmartList, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dp.ID, dp.Seq, dp.Name, dp.ImagePath, dp.Attribute, dp.ParentID, dp.SmartList, dp.UUID, dp.RbDataStatus, dp.RbLocalDataStatus, dp.RbLocalDeleted, dp.RbLocalSynced, dp.Usn, dp.RbLocalUsn, dp.CreatedAt, dp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dp.ID, dp.Seq, dp.Name, dp.ImagePath, dp.Attribute, dp.ParentID, dp.SmartList, dp.UUID, dp.RbDataStatus, dp.RbLocalDataStatus, dp.RbLocalDeleted, dp.RbLocalSynced, dp.Usn, dp.RbLocalUsn, dp.CreatedAt, dp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dp._exists = true
	return nil
}

// Delete deletes the DjmdPlaylist from the database.
func (c *Client) DeleteDjmdPlaylist(ctx context.Context, dp *DjmdPlaylist) error {
	db := c.db

	switch {
	case !dp._exists: // doesn't exist
		return nil
	case dp._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdPlaylist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dp.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dp._deleted = true
	return nil
}

func (c *Client) AllDjmdPlaylist(ctx context.Context) ([]*DjmdPlaylist, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdPlaylist`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByAttribute retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist__attribute'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByAttribute(ctx context.Context, db DB, attribute sql.NullInt64) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByAttribute {
// func DjmdPlaylistByAttribute(db DB, attribute sql.NullInt64) ([]*DjmdPlaylist, error) {
// true
// attribute
// DjmdPlaylist
// DjmdPlaylistByAttribute
// false
// false
func (c *Client) DjmdPlaylistByAttribute(ctx context.Context, attribute sql.NullInt64) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByAttribute(ctx context.Context, db DB, attribute sql.NullInt64) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE Attribute = $1`
	// run
	logf(sqlstr, attribute)
	rows, err := db.QueryContext(ctx, sqlstr, attribute)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByName retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist__name'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByName(ctx context.Context, db DB, name sql.NullString) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByName {
// func DjmdPlaylistByName(db DB, name sql.NullString) ([]*DjmdPlaylist, error) {
// true
// name
// DjmdPlaylist
// DjmdPlaylistByName
// false
// false
func (c *Client) DjmdPlaylistByName(ctx context.Context, name sql.NullString) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByName(ctx context.Context, db DB, name sql.NullString) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE Name = $1`
	// run
	logf(sqlstr, name)
	rows, err := db.QueryContext(ctx, sqlstr, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByParentID retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist__parent_i_d'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByParentID(ctx context.Context, db DB, parentID sql.NullString) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByParentID {
// func DjmdPlaylistByParentID(db DB, parentID sql.NullString) ([]*DjmdPlaylist, error) {
// true
// parentID
// DjmdPlaylist
// DjmdPlaylistByParentID
// false
// false
func (c *Client) DjmdPlaylistByParentID(ctx context.Context, parentID sql.NullString) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByParentID(ctx context.Context, db DB, parentID sql.NullString) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE ParentID = $1`
	// run
	logf(sqlstr, parentID)
	rows, err := db.QueryContext(ctx, sqlstr, parentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistBySeq retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist__seq'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistBySeq(ctx context.Context, db DB, seq sql.NullInt64) ([]*DjmdPlaylist, error) {
// DjmdPlaylistBySeq {
// func DjmdPlaylistBySeq(db DB, seq sql.NullInt64) ([]*DjmdPlaylist, error) {
// true
// seq
// DjmdPlaylist
// DjmdPlaylistBySeq
// false
// false
func (c *Client) DjmdPlaylistBySeq(ctx context.Context, seq sql.NullInt64) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistBySeq(ctx context.Context, db DB, seq sql.NullInt64) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE Seq = $1`
	// run
	logf(sqlstr, seq)
	rows, err := db.QueryContext(ctx, sqlstr, seq)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByUUID retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist__u_u_i_d'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByUUID {
// func DjmdPlaylistByUUID(db DB, uuid sql.NullString) ([]*DjmdPlaylist, error) {
// true
// uuid
// DjmdPlaylist
// DjmdPlaylistByUUID
// false
// false
func (c *Client) DjmdPlaylistByUUID(ctx context.Context, uuid sql.NullString) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByRbDataStatus retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist_rb_data_status'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByRbDataStatus {
// func DjmdPlaylistByRbDataStatus(db DB, rbDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
// true
// rbDataStatus
// DjmdPlaylist
// DjmdPlaylistByRbDataStatus
// false
// false
func (c *Client) DjmdPlaylistByRbDataStatus(ctx context.Context, rbDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByRbLocalDataStatus retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist_rb_local_data_status'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByRbLocalDataStatus {
// func DjmdPlaylistByRbLocalDataStatus(db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
// true
// rbLocalDataStatus
// DjmdPlaylist
// DjmdPlaylistByRbLocalDataStatus
// false
// false
func (c *Client) DjmdPlaylistByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByRbLocalDeleted retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist_rb_local_deleted'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByRbLocalDeleted {
// func DjmdPlaylistByRbLocalDeleted(db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdPlaylist, error) {
// true
// rbLocalDeleted
// DjmdPlaylist
// DjmdPlaylistByRbLocalDeleted
// false
// false
func (c *Client) DjmdPlaylistByRbLocalDeleted(ctx context.Context, rbLocalDeleted sql.NullInt64) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByRbLocalUsnID retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'djmd_playlist_rb_local_usn__i_d'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdPlaylist, error) {
// DjmdPlaylistByRbLocalUsnID {
// func DjmdPlaylistByRbLocalUsnID(db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdPlaylist, error) {
// true
// rbLocalUsn, id
// DjmdPlaylist
// DjmdPlaylistByRbLocalUsnID
// false
// false
func (c *Client) DjmdPlaylistByRbLocalUsnID(ctx context.Context, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdPlaylist, error) {
	// func DjmdPlaylistByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdPlaylist
	for rows.Next() {
		dp := DjmdPlaylist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdPlaylistByID retrieves a row from 'djmdPlaylist' as a DjmdPlaylist.
//
// Generated from index 'sqlite_autoindex_djmdPlaylist_1'.
// func (dp *DjmdPlaylist) djmdPlaylist(db DB) (error)
// func DjmdPlaylistByID(ctx context.Context, db DB, id sql.NullString) (*DjmdPlaylist, error) {
// DjmdPlaylistByID {
// func DjmdPlaylistByID(db DB, id sql.NullString) (*DjmdPlaylist, error) {
// true
// id
// DjmdPlaylist
// DjmdPlaylistByID
// true
// true
func (c *Client) DjmdPlaylistByID(ctx context.Context, id sql.NullString) (*DjmdPlaylist, error) {
	// func DjmdPlaylistByID(ctx context.Context, db DB, id sql.NullString) (*DjmdPlaylist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, ImagePath, Attribute, ParentID, SmartList, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdPlaylist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dp := DjmdPlaylist{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dp.ID, &dp.Seq, &dp.Name, &dp.ImagePath, &dp.Attribute, &dp.ParentID, &dp.SmartList, &dp.UUID, &dp.RbDataStatus, &dp.RbLocalDataStatus, &dp.RbLocalDeleted, &dp.RbLocalSynced, &dp.Usn, &dp.RbLocalUsn, &dp.CreatedAt, &dp.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dp, nil
}
