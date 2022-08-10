package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DjmdLabel represents a row from 'djmdLabel'.
type DjmdLabel struct {
	ID                sql.NullString `json:"id"`                   // ID
	Name              sql.NullString `json:"name"`                 // Name
	UUID              sql.NullString `json:"uuid"`                 // UUID
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

// Exists returns true when the DjmdLabel exists in the database.
func (dl *DjmdLabel) Exists() bool {
	return dl._exists
}

// Deleted returns true when the DjmdLabel has been marked for deletion from
// the database.
func (dl *DjmdLabel) Deleted() bool {
	return dl._deleted
}

// Insert inserts the DjmdLabel to the database.
func (c *Client) InsertDjmdLabel(ctx context.Context, dl *DjmdLabel) error {
	db := c.db

	switch {
	case dl._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dl._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdLabel (` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)`
	// run
	logf(sqlstr, dl.ID, dl.Name, dl.UUID, dl.RbDataStatus, dl.RbLocalDataStatus, dl.RbLocalDeleted, dl.RbLocalSynced, dl.Usn, dl.RbLocalUsn, dl.CreatedAt, dl.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dl.ID, dl.Name, dl.UUID, dl.RbDataStatus, dl.RbLocalDataStatus, dl.RbLocalDeleted, dl.RbLocalSynced, dl.Usn, dl.RbLocalUsn, dl.CreatedAt, dl.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dl._exists = true
	return nil
}

// Update updates a DjmdLabel in the database.
func (c *Client) UpdateDjmdLabel(ctx context.Context, dl *DjmdLabel) error {
	db := c.db

	switch {
	case !dl._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dl._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdLabel SET ` +
		`Name = $1, UUID = $2, rb_data_status = $3, rb_local_data_status = $4, rb_local_deleted = $5, rb_local_synced = $6, usn = $7, rb_local_usn = $8, created_at = $9, updated_at = $10 ` +
		`WHERE ID = $11`
	// run
	logf(sqlstr, dl.Name, dl.UUID, dl.RbDataStatus, dl.RbLocalDataStatus, dl.RbLocalDeleted, dl.RbLocalSynced, dl.Usn, dl.RbLocalUsn, dl.CreatedAt, dl.UpdatedAt, dl.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dl.Name, dl.UUID, dl.RbDataStatus, dl.RbLocalDataStatus, dl.RbLocalDeleted, dl.RbLocalSynced, dl.Usn, dl.RbLocalUsn, dl.CreatedAt, dl.UpdatedAt, dl.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdLabel to the database.
func (c *Client) SaveDjmdLabel(ctx context.Context, dl *DjmdLabel) error {
	if dl.Exists() {
		return c.UpdateDjmdLabel(ctx, dl)
	}
	return c.InsertDjmdLabel(ctx, dl)
}

// Upsert performs an upsert for DjmdLabel.
func (c *Client) UpsertDjmdLabel(ctx context.Context, dl *DjmdLabel) error {
	db := c.db

	switch {
	case dl._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdLabel (` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`Name = EXCLUDED.Name, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dl.ID, dl.Name, dl.UUID, dl.RbDataStatus, dl.RbLocalDataStatus, dl.RbLocalDeleted, dl.RbLocalSynced, dl.Usn, dl.RbLocalUsn, dl.CreatedAt, dl.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dl.ID, dl.Name, dl.UUID, dl.RbDataStatus, dl.RbLocalDataStatus, dl.RbLocalDeleted, dl.RbLocalSynced, dl.Usn, dl.RbLocalUsn, dl.CreatedAt, dl.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dl._exists = true
	return nil
}

// Delete deletes the DjmdLabel from the database.
func (c *Client) DeleteDjmdLabel(ctx context.Context, dl *DjmdLabel) error {
	db := c.db

	switch {
	case !dl._exists: // doesn't exist
		return nil
	case dl._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdLabel ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dl.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dl.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dl._deleted = true
	return nil
}

func scanDjmdLabelRows(rows *sql.Rows) ([]*DjmdLabel, error) {
	var res []*DjmdLabel
	for rows.Next() {
		dl := DjmdLabel{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dl)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdLabel(ctx context.Context) ([]*DjmdLabel, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdLabel`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdLabelRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdLabelByName retrieves a row from 'djmdLabel' as a DjmdLabel.
//
// Generated from index 'djmd_label__name'.
func (c *Client) DjmdLabelByName(ctx context.Context, name sql.NullString) ([]*DjmdLabel, error) {
	// func DjmdLabelByName(ctx context.Context, db DB, name sql.NullString) ([]*DjmdLabel, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdLabel ` +
		`WHERE Name = $1`
	// run
	logf(sqlstr, name)
	rows, err := db.QueryContext(ctx, sqlstr, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdLabel
	for rows.Next() {
		dl := DjmdLabel{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdLabelByUUID retrieves a row from 'djmdLabel' as a DjmdLabel.
//
// Generated from index 'djmd_label__u_u_i_d'.
func (c *Client) DjmdLabelByUUID(ctx context.Context, uuid sql.NullString) ([]*DjmdLabel, error) {
	// func DjmdLabelByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdLabel, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdLabel ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdLabel
	for rows.Next() {
		dl := DjmdLabel{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdLabelByRbDataStatus retrieves a row from 'djmdLabel' as a DjmdLabel.
//
// Generated from index 'djmd_label_rb_data_status'.
func (c *Client) DjmdLabelByRbDataStatus(ctx context.Context, rbDataStatus sql.NullInt64) ([]*DjmdLabel, error) {
	// func DjmdLabelByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdLabel, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdLabel ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdLabel
	for rows.Next() {
		dl := DjmdLabel{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdLabelByRbLocalDataStatus retrieves a row from 'djmdLabel' as a DjmdLabel.
//
// Generated from index 'djmd_label_rb_local_data_status'.
func (c *Client) DjmdLabelByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus sql.NullInt64) ([]*DjmdLabel, error) {
	// func DjmdLabelByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdLabel, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdLabel ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdLabel
	for rows.Next() {
		dl := DjmdLabel{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdLabelByRbLocalDeleted retrieves a row from 'djmdLabel' as a DjmdLabel.
//
// Generated from index 'djmd_label_rb_local_deleted'.
func (c *Client) DjmdLabelByRbLocalDeleted(ctx context.Context, rbLocalDeleted sql.NullInt64) ([]*DjmdLabel, error) {
	// func DjmdLabelByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdLabel, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdLabel ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdLabel
	for rows.Next() {
		dl := DjmdLabel{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdLabelByRbLocalUsnID retrieves a row from 'djmdLabel' as a DjmdLabel.
//
// Generated from index 'djmd_label_rb_local_usn__i_d'.
func (c *Client) DjmdLabelByRbLocalUsnID(ctx context.Context, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdLabel, error) {
	// func DjmdLabelByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdLabel, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdLabel ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdLabel
	for rows.Next() {
		dl := DjmdLabel{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dl)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdLabelByID retrieves a row from 'djmdLabel' as a DjmdLabel.
//
// Generated from index 'sqlite_autoindex_djmdLabel_1'.
func (c *Client) DjmdLabelByID(ctx context.Context, id sql.NullString) (*DjmdLabel, error) {
	// func DjmdLabelByID(ctx context.Context, db DB, id sql.NullString) (*DjmdLabel, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdLabel ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dl := DjmdLabel{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dl.ID, &dl.Name, &dl.UUID, &dl.RbDataStatus, &dl.RbLocalDataStatus, &dl.RbLocalDeleted, &dl.RbLocalSynced, &dl.Usn, &dl.RbLocalUsn, &dl.CreatedAt, &dl.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dl, nil
}
