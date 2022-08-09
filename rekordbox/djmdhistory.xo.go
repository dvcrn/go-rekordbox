package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DjmdHistory represents a row from 'djmdHistory'.
type DjmdHistory struct {
	ID                sql.NullString `json:"ID"`                   // ID
	Seq               sql.NullInt64  `json:"Seq"`                  // Seq
	Name              sql.NullString `json:"Name"`                 // Name
	Attribute         sql.NullInt64  `json:"Attribute"`            // Attribute
	ParentID          sql.NullString `json:"ParentID"`             // ParentID
	DateCreated       sql.NullString `json:"DateCreated"`          // DateCreated
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

// Exists returns true when the DjmdHistory exists in the database.
func (dh *DjmdHistory) Exists() bool {
	return dh._exists
}

// Deleted returns true when the DjmdHistory has been marked for deletion from
// the database.
func (dh *DjmdHistory) Deleted() bool {
	return dh._deleted
}

// Insert inserts the DjmdHistory to the database.
func (c *Client) InsertDjmdHistory(ctx context.Context, dh *DjmdHistory) error {
	db := c.db

	switch {
	case dh._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dh._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdHistory (` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)`
	// run
	logf(sqlstr, dh.ID, dh.Seq, dh.Name, dh.Attribute, dh.ParentID, dh.DateCreated, dh.UUID, dh.RbDataStatus, dh.RbLocalDataStatus, dh.RbLocalDeleted, dh.RbLocalSynced, dh.Usn, dh.RbLocalUsn, dh.CreatedAt, dh.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dh.ID, dh.Seq, dh.Name, dh.Attribute, dh.ParentID, dh.DateCreated, dh.UUID, dh.RbDataStatus, dh.RbLocalDataStatus, dh.RbLocalDeleted, dh.RbLocalSynced, dh.Usn, dh.RbLocalUsn, dh.CreatedAt, dh.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dh._exists = true
	return nil
}

// Update updates a DjmdHistory in the database.
func (c *Client) UpdateDjmdHistory(ctx context.Context, dh *DjmdHistory) error {
	db := c.db

	switch {
	case !dh._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dh._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdHistory SET ` +
		`Seq = $1, Name = $2, Attribute = $3, ParentID = $4, DateCreated = $5, UUID = $6, rb_data_status = $7, rb_local_data_status = $8, rb_local_deleted = $9, rb_local_synced = $10, usn = $11, rb_local_usn = $12, created_at = $13, updated_at = $14 ` +
		`WHERE ID = $15`
	// run
	logf(sqlstr, dh.Seq, dh.Name, dh.Attribute, dh.ParentID, dh.DateCreated, dh.UUID, dh.RbDataStatus, dh.RbLocalDataStatus, dh.RbLocalDeleted, dh.RbLocalSynced, dh.Usn, dh.RbLocalUsn, dh.CreatedAt, dh.UpdatedAt, dh.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dh.Seq, dh.Name, dh.Attribute, dh.ParentID, dh.DateCreated, dh.UUID, dh.RbDataStatus, dh.RbLocalDataStatus, dh.RbLocalDeleted, dh.RbLocalSynced, dh.Usn, dh.RbLocalUsn, dh.CreatedAt, dh.UpdatedAt, dh.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdHistory to the database.
func (c *Client) SaveDjmdHistory(ctx context.Context, dh *DjmdHistory) error {
	if dh.Exists() {
		return c.UpdateDjmdHistory(ctx, dh)
	}
	return c.InsertDjmdHistory(ctx, dh)
}

// Upsert performs an upsert for DjmdHistory.
func (c *Client) UpsertDjmdHistory(ctx context.Context, dh *DjmdHistory) error {
	db := c.db

	switch {
	case dh._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdHistory (` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`Seq = EXCLUDED.Seq, Name = EXCLUDED.Name, Attribute = EXCLUDED.Attribute, ParentID = EXCLUDED.ParentID, DateCreated = EXCLUDED.DateCreated, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dh.ID, dh.Seq, dh.Name, dh.Attribute, dh.ParentID, dh.DateCreated, dh.UUID, dh.RbDataStatus, dh.RbLocalDataStatus, dh.RbLocalDeleted, dh.RbLocalSynced, dh.Usn, dh.RbLocalUsn, dh.CreatedAt, dh.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dh.ID, dh.Seq, dh.Name, dh.Attribute, dh.ParentID, dh.DateCreated, dh.UUID, dh.RbDataStatus, dh.RbLocalDataStatus, dh.RbLocalDeleted, dh.RbLocalSynced, dh.Usn, dh.RbLocalUsn, dh.CreatedAt, dh.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dh._exists = true
	return nil
}

// Delete deletes the DjmdHistory from the database.
func (c *Client) DeleteDjmdHistory(ctx context.Context, dh *DjmdHistory) error {
	db := c.db

	switch {
	case !dh._exists: // doesn't exist
		return nil
	case dh._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdHistory ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dh.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dh.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dh._deleted = true
	return nil
}

func (c *Client) AllDjmdHistory(ctx context.Context) ([]*DjmdHistory, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdHistory`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByName retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'djmd_history__name'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByName(ctx context.Context, db DB, name sql.NullString) ([]*DjmdHistory, error) {
// DjmdHistoryByName {
// func DjmdHistoryByName(db DB, name sql.NullString) ([]*DjmdHistory, error) {
// true
// name
// DjmdHistory
// DjmdHistoryByName
// false
// false
func (c *Client) DjmdHistoryByName(ctx context.Context, name sql.NullString) ([]*DjmdHistory, error) {
	// func DjmdHistoryByName(ctx context.Context, db DB, name sql.NullString) ([]*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE Name = $1`
	// run
	logf(sqlstr, name)
	rows, err := db.QueryContext(ctx, sqlstr, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByParentID retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'djmd_history__parent_i_d'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByParentID(ctx context.Context, db DB, parentID sql.NullString) ([]*DjmdHistory, error) {
// DjmdHistoryByParentID {
// func DjmdHistoryByParentID(db DB, parentID sql.NullString) ([]*DjmdHistory, error) {
// true
// parentID
// DjmdHistory
// DjmdHistoryByParentID
// false
// false
func (c *Client) DjmdHistoryByParentID(ctx context.Context, parentID sql.NullString) ([]*DjmdHistory, error) {
	// func DjmdHistoryByParentID(ctx context.Context, db DB, parentID sql.NullString) ([]*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE ParentID = $1`
	// run
	logf(sqlstr, parentID)
	rows, err := db.QueryContext(ctx, sqlstr, parentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByUUID retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'djmd_history__u_u_i_d'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdHistory, error) {
// DjmdHistoryByUUID {
// func DjmdHistoryByUUID(db DB, uuid sql.NullString) ([]*DjmdHistory, error) {
// true
// uuid
// DjmdHistory
// DjmdHistoryByUUID
// false
// false
func (c *Client) DjmdHistoryByUUID(ctx context.Context, uuid sql.NullString) ([]*DjmdHistory, error) {
	// func DjmdHistoryByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByRbDataStatus retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'djmd_history_rb_data_status'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
// DjmdHistoryByRbDataStatus {
// func DjmdHistoryByRbDataStatus(db DB, rbDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
// true
// rbDataStatus
// DjmdHistory
// DjmdHistoryByRbDataStatus
// false
// false
func (c *Client) DjmdHistoryByRbDataStatus(ctx context.Context, rbDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
	// func DjmdHistoryByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByRbLocalDataStatus retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'djmd_history_rb_local_data_status'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
// DjmdHistoryByRbLocalDataStatus {
// func DjmdHistoryByRbLocalDataStatus(db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
// true
// rbLocalDataStatus
// DjmdHistory
// DjmdHistoryByRbLocalDataStatus
// false
// false
func (c *Client) DjmdHistoryByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
	// func DjmdHistoryByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByRbLocalDeleted retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'djmd_history_rb_local_deleted'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdHistory, error) {
// DjmdHistoryByRbLocalDeleted {
// func DjmdHistoryByRbLocalDeleted(db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdHistory, error) {
// true
// rbLocalDeleted
// DjmdHistory
// DjmdHistoryByRbLocalDeleted
// false
// false
func (c *Client) DjmdHistoryByRbLocalDeleted(ctx context.Context, rbLocalDeleted sql.NullInt64) ([]*DjmdHistory, error) {
	// func DjmdHistoryByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByRbLocalUsnID retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'djmd_history_rb_local_usn__i_d'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdHistory, error) {
// DjmdHistoryByRbLocalUsnID {
// func DjmdHistoryByRbLocalUsnID(db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdHistory, error) {
// true
// rbLocalUsn, id
// DjmdHistory
// DjmdHistoryByRbLocalUsnID
// false
// false
func (c *Client) DjmdHistoryByRbLocalUsnID(ctx context.Context, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdHistory, error) {
	// func DjmdHistoryByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdHistory
	for rows.Next() {
		dh := DjmdHistory{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dh)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdHistoryByID retrieves a row from 'djmdHistory' as a DjmdHistory.
//
// Generated from index 'sqlite_autoindex_djmdHistory_1'.
// func (dh *DjmdHistory) djmdHistory(db DB) (error)
// func DjmdHistoryByID(ctx context.Context, db DB, id sql.NullString) (*DjmdHistory, error) {
// DjmdHistoryByID {
// func DjmdHistoryByID(db DB, id sql.NullString) (*DjmdHistory, error) {
// true
// id
// DjmdHistory
// DjmdHistoryByID
// true
// true
func (c *Client) DjmdHistoryByID(ctx context.Context, id sql.NullString) (*DjmdHistory, error) {
	// func DjmdHistoryByID(ctx context.Context, db DB, id sql.NullString) (*DjmdHistory, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, DateCreated, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdHistory ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dh := DjmdHistory{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dh.ID, &dh.Seq, &dh.Name, &dh.Attribute, &dh.ParentID, &dh.DateCreated, &dh.UUID, &dh.RbDataStatus, &dh.RbLocalDataStatus, &dh.RbLocalDeleted, &dh.RbLocalSynced, &dh.Usn, &dh.RbLocalUsn, &dh.CreatedAt, &dh.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dh, nil
}