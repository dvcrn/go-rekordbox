package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DjmdCloudProperty represents a row from 'djmdCloudProperty'.
type DjmdCloudProperty struct {
	ID                sql.NullString `json:"id"`                   // ID
	Reserved1         sql.NullString `json:"reserved1"`            // Reserved1
	Reserved2         sql.NullString `json:"reserved2"`            // Reserved2
	Reserved3         sql.NullString `json:"reserved3"`            // Reserved3
	Reserved4         sql.NullString `json:"reserved4"`            // Reserved4
	Reserved5         sql.NullString `json:"reserved5"`            // Reserved5
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

// Exists returns true when the DjmdCloudProperty exists in the database.
func (dcp *DjmdCloudProperty) Exists() bool {
	return dcp._exists
}

// Deleted returns true when the DjmdCloudProperty has been marked for deletion from
// the database.
func (dcp *DjmdCloudProperty) Deleted() bool {
	return dcp._deleted
}

// Insert inserts the DjmdCloudProperty to the database.
func (c *Client) InsertDjmdCloudProperty(ctx context.Context, dcp *DjmdCloudProperty) error {
	db := c.db

	switch {
	case dcp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dcp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdCloudProperty (` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)`
	// run
	logf(sqlstr, dcp.ID, dcp.Reserved1, dcp.Reserved2, dcp.Reserved3, dcp.Reserved4, dcp.Reserved5, dcp.UUID, dcp.RbDataStatus, dcp.RbLocalDataStatus, dcp.RbLocalDeleted, dcp.RbLocalSynced, dcp.Usn, dcp.RbLocalUsn, dcp.CreatedAt, dcp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dcp.ID, dcp.Reserved1, dcp.Reserved2, dcp.Reserved3, dcp.Reserved4, dcp.Reserved5, dcp.UUID, dcp.RbDataStatus, dcp.RbLocalDataStatus, dcp.RbLocalDeleted, dcp.RbLocalSynced, dcp.Usn, dcp.RbLocalUsn, dcp.CreatedAt, dcp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dcp._exists = true
	return nil
}

// Update updates a DjmdCloudProperty in the database.
func (c *Client) UpdateDjmdCloudProperty(ctx context.Context, dcp *DjmdCloudProperty) error {
	db := c.db

	switch {
	case !dcp._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dcp._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdCloudProperty SET ` +
		`Reserved1 = $1, Reserved2 = $2, Reserved3 = $3, Reserved4 = $4, Reserved5 = $5, UUID = $6, rb_data_status = $7, rb_local_data_status = $8, rb_local_deleted = $9, rb_local_synced = $10, usn = $11, rb_local_usn = $12, created_at = $13, updated_at = $14 ` +
		`WHERE ID = $15`
	// run
	logf(sqlstr, dcp.Reserved1, dcp.Reserved2, dcp.Reserved3, dcp.Reserved4, dcp.Reserved5, dcp.UUID, dcp.RbDataStatus, dcp.RbLocalDataStatus, dcp.RbLocalDeleted, dcp.RbLocalSynced, dcp.Usn, dcp.RbLocalUsn, dcp.CreatedAt, dcp.UpdatedAt, dcp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dcp.Reserved1, dcp.Reserved2, dcp.Reserved3, dcp.Reserved4, dcp.Reserved5, dcp.UUID, dcp.RbDataStatus, dcp.RbLocalDataStatus, dcp.RbLocalDeleted, dcp.RbLocalSynced, dcp.Usn, dcp.RbLocalUsn, dcp.CreatedAt, dcp.UpdatedAt, dcp.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdCloudProperty to the database.
func (c *Client) SaveDjmdCloudProperty(ctx context.Context, dcp *DjmdCloudProperty) error {
	if dcp.Exists() {
		return c.UpdateDjmdCloudProperty(ctx, dcp)
	}
	return c.InsertDjmdCloudProperty(ctx, dcp)
}

// Upsert performs an upsert for DjmdCloudProperty.
func (c *Client) UpsertDjmdCloudProperty(ctx context.Context, dcp *DjmdCloudProperty) error {
	db := c.db

	switch {
	case dcp._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdCloudProperty (` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`Reserved1 = EXCLUDED.Reserved1, Reserved2 = EXCLUDED.Reserved2, Reserved3 = EXCLUDED.Reserved3, Reserved4 = EXCLUDED.Reserved4, Reserved5 = EXCLUDED.Reserved5, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dcp.ID, dcp.Reserved1, dcp.Reserved2, dcp.Reserved3, dcp.Reserved4, dcp.Reserved5, dcp.UUID, dcp.RbDataStatus, dcp.RbLocalDataStatus, dcp.RbLocalDeleted, dcp.RbLocalSynced, dcp.Usn, dcp.RbLocalUsn, dcp.CreatedAt, dcp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dcp.ID, dcp.Reserved1, dcp.Reserved2, dcp.Reserved3, dcp.Reserved4, dcp.Reserved5, dcp.UUID, dcp.RbDataStatus, dcp.RbLocalDataStatus, dcp.RbLocalDeleted, dcp.RbLocalSynced, dcp.Usn, dcp.RbLocalUsn, dcp.CreatedAt, dcp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dcp._exists = true
	return nil
}

// Delete deletes the DjmdCloudProperty from the database.
func (c *Client) DeleteDjmdCloudProperty(ctx context.Context, dcp *DjmdCloudProperty) error {
	db := c.db

	switch {
	case !dcp._exists: // doesn't exist
		return nil
	case dcp._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdCloudProperty ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dcp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dcp.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dcp._deleted = true
	return nil
}

func scanDjmdCloudPropertyRows(rows *sql.Rows) ([]*DjmdCloudProperty, error) {
	var res []*DjmdCloudProperty
	for rows.Next() {
		dcp := DjmdCloudProperty{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dcp.ID, &dcp.Reserved1, &dcp.Reserved2, &dcp.Reserved3, &dcp.Reserved4, &dcp.Reserved5, &dcp.UUID, &dcp.RbDataStatus, &dcp.RbLocalDataStatus, &dcp.RbLocalDeleted, &dcp.RbLocalSynced, &dcp.Usn, &dcp.RbLocalUsn, &dcp.CreatedAt, &dcp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dcp)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdCloudProperty(ctx context.Context) ([]*DjmdCloudProperty, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdCloudProperty`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdCloudPropertyRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdCloudPropertyByUUID retrieves a row from 'djmdCloudProperty' as a DjmdCloudProperty.
//
// Generated from index 'djmd_cloud_property__u_u_i_d'.
func (c *Client) DjmdCloudPropertyByUUID(ctx context.Context, uuid sql.NullString) ([]*DjmdCloudProperty, error) {
	// func DjmdCloudPropertyByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdCloudProperty, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdCloudProperty ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdCloudProperty
	for rows.Next() {
		dcp := DjmdCloudProperty{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dcp.ID, &dcp.Reserved1, &dcp.Reserved2, &dcp.Reserved3, &dcp.Reserved4, &dcp.Reserved5, &dcp.UUID, &dcp.RbDataStatus, &dcp.RbLocalDataStatus, &dcp.RbLocalDeleted, &dcp.RbLocalSynced, &dcp.Usn, &dcp.RbLocalUsn, &dcp.CreatedAt, &dcp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dcp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdCloudPropertyByRbDataStatus retrieves a row from 'djmdCloudProperty' as a DjmdCloudProperty.
//
// Generated from index 'djmd_cloud_property_rb_data_status'.
func (c *Client) DjmdCloudPropertyByRbDataStatus(ctx context.Context, rbDataStatus sql.NullInt64) ([]*DjmdCloudProperty, error) {
	// func DjmdCloudPropertyByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdCloudProperty, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdCloudProperty ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdCloudProperty
	for rows.Next() {
		dcp := DjmdCloudProperty{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dcp.ID, &dcp.Reserved1, &dcp.Reserved2, &dcp.Reserved3, &dcp.Reserved4, &dcp.Reserved5, &dcp.UUID, &dcp.RbDataStatus, &dcp.RbLocalDataStatus, &dcp.RbLocalDeleted, &dcp.RbLocalSynced, &dcp.Usn, &dcp.RbLocalUsn, &dcp.CreatedAt, &dcp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dcp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdCloudPropertyByRbLocalDataStatus retrieves a row from 'djmdCloudProperty' as a DjmdCloudProperty.
//
// Generated from index 'djmd_cloud_property_rb_local_data_status'.
func (c *Client) DjmdCloudPropertyByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus sql.NullInt64) ([]*DjmdCloudProperty, error) {
	// func DjmdCloudPropertyByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdCloudProperty, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdCloudProperty ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdCloudProperty
	for rows.Next() {
		dcp := DjmdCloudProperty{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dcp.ID, &dcp.Reserved1, &dcp.Reserved2, &dcp.Reserved3, &dcp.Reserved4, &dcp.Reserved5, &dcp.UUID, &dcp.RbDataStatus, &dcp.RbLocalDataStatus, &dcp.RbLocalDeleted, &dcp.RbLocalSynced, &dcp.Usn, &dcp.RbLocalUsn, &dcp.CreatedAt, &dcp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dcp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdCloudPropertyByRbLocalDeleted retrieves a row from 'djmdCloudProperty' as a DjmdCloudProperty.
//
// Generated from index 'djmd_cloud_property_rb_local_deleted'.
func (c *Client) DjmdCloudPropertyByRbLocalDeleted(ctx context.Context, rbLocalDeleted sql.NullInt64) ([]*DjmdCloudProperty, error) {
	// func DjmdCloudPropertyByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdCloudProperty, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdCloudProperty ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdCloudProperty
	for rows.Next() {
		dcp := DjmdCloudProperty{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dcp.ID, &dcp.Reserved1, &dcp.Reserved2, &dcp.Reserved3, &dcp.Reserved4, &dcp.Reserved5, &dcp.UUID, &dcp.RbDataStatus, &dcp.RbLocalDataStatus, &dcp.RbLocalDeleted, &dcp.RbLocalSynced, &dcp.Usn, &dcp.RbLocalUsn, &dcp.CreatedAt, &dcp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dcp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdCloudPropertyByRbLocalUsnID retrieves a row from 'djmdCloudProperty' as a DjmdCloudProperty.
//
// Generated from index 'djmd_cloud_property_rb_local_usn__i_d'.
func (c *Client) DjmdCloudPropertyByRbLocalUsnID(ctx context.Context, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdCloudProperty, error) {
	// func DjmdCloudPropertyByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdCloudProperty, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdCloudProperty ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdCloudProperty
	for rows.Next() {
		dcp := DjmdCloudProperty{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dcp.ID, &dcp.Reserved1, &dcp.Reserved2, &dcp.Reserved3, &dcp.Reserved4, &dcp.Reserved5, &dcp.UUID, &dcp.RbDataStatus, &dcp.RbLocalDataStatus, &dcp.RbLocalDeleted, &dcp.RbLocalSynced, &dcp.Usn, &dcp.RbLocalUsn, &dcp.CreatedAt, &dcp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dcp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdCloudPropertyByID retrieves a row from 'djmdCloudProperty' as a DjmdCloudProperty.
//
// Generated from index 'sqlite_autoindex_djmdCloudProperty_1'.
func (c *Client) DjmdCloudPropertyByID(ctx context.Context, id sql.NullString) (*DjmdCloudProperty, error) {
	// func DjmdCloudPropertyByID(ctx context.Context, db DB, id sql.NullString) (*DjmdCloudProperty, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Reserved1, Reserved2, Reserved3, Reserved4, Reserved5, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdCloudProperty ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dcp := DjmdCloudProperty{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dcp.ID, &dcp.Reserved1, &dcp.Reserved2, &dcp.Reserved3, &dcp.Reserved4, &dcp.Reserved5, &dcp.UUID, &dcp.RbDataStatus, &dcp.RbLocalDataStatus, &dcp.RbLocalDeleted, &dcp.RbLocalSynced, &dcp.Usn, &dcp.RbLocalUsn, &dcp.CreatedAt, &dcp.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dcp, nil
}
