package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// UUIDIDMap represents a row from 'uuidIDMap'.
type UUIDIDMap struct {
	ID                nulltype.NullString `json:"ID"`                   // ID
	TableName         nulltype.NullString `json:"TableName"`            // TableName
	TargetUUID        nulltype.NullString `json:"TargetUUID"`           // TargetUUID
	CurrentID         nulltype.NullString `json:"CurrentID"`            // CurrentID
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

// Exists returns true when the UUIDIDMap exists in the database.
func (um *UUIDIDMap) Exists() bool {
	return um._exists
}

// Deleted returns true when the UUIDIDMap has been marked for deletion from
// the database.
func (um *UUIDIDMap) Deleted() bool {
	return um._deleted
}

// Insert inserts the UUIDIDMap to the database.
func (c *Client) InsertUUIDIDMap(ctx context.Context, um *UUIDIDMap) error {
	db := c.db

	switch {
	case um._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case um._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO uuidIDMap (` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)`
	// run
	logf(sqlstr, um.ID, um.TableName, um.TargetUUID, um.CurrentID, um.UUID, um.RbDataStatus, um.RbLocalDataStatus, um.RbLocalDeleted, um.RbLocalSynced, um.Usn, um.RbLocalUsn, um.CreatedAt, um.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, um.ID, um.TableName, um.TargetUUID, um.CurrentID, um.UUID, um.RbDataStatus, um.RbLocalDataStatus, um.RbLocalDeleted, um.RbLocalSynced, um.Usn, um.RbLocalUsn, um.CreatedAt, um.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	um._exists = true
	return nil
}

// Update updates a UUIDIDMap in the database.
func (c *Client) UpdateUUIDIDMap(ctx context.Context, um *UUIDIDMap) error {
	db := c.db

	switch {
	case !um._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case um._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE uuidIDMap SET ` +
		`TableName = $1, TargetUUID = $2, CurrentID = $3, UUID = $4, rb_data_status = $5, rb_local_data_status = $6, rb_local_deleted = $7, rb_local_synced = $8, usn = $9, rb_local_usn = $10, created_at = $11, updated_at = $12 ` +
		`WHERE ID = $13`
	// run
	logf(sqlstr, um.TableName, um.TargetUUID, um.CurrentID, um.UUID, um.RbDataStatus, um.RbLocalDataStatus, um.RbLocalDeleted, um.RbLocalSynced, um.Usn, um.RbLocalUsn, um.CreatedAt, um.UpdatedAt, um.ID)
	if _, err := db.ExecContext(ctx, sqlstr, um.TableName, um.TargetUUID, um.CurrentID, um.UUID, um.RbDataStatus, um.RbLocalDataStatus, um.RbLocalDeleted, um.RbLocalSynced, um.Usn, um.RbLocalUsn, um.CreatedAt, um.UpdatedAt, um.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the UUIDIDMap to the database.
func (c *Client) SaveUUIDIDMap(ctx context.Context, um *UUIDIDMap) error {
	if um.Exists() {
		return c.UpdateUUIDIDMap(ctx, um)
	}
	return c.InsertUUIDIDMap(ctx, um)
}

// Upsert performs an upsert for UUIDIDMap.
func (c *Client) UpsertUUIDIDMap(ctx context.Context, um *UUIDIDMap) error {
	db := c.db

	switch {
	case um._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO uuidIDMap (` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`TableName = EXCLUDED.TableName, TargetUUID = EXCLUDED.TargetUUID, CurrentID = EXCLUDED.CurrentID, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, um.ID, um.TableName, um.TargetUUID, um.CurrentID, um.UUID, um.RbDataStatus, um.RbLocalDataStatus, um.RbLocalDeleted, um.RbLocalSynced, um.Usn, um.RbLocalUsn, um.CreatedAt, um.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, um.ID, um.TableName, um.TargetUUID, um.CurrentID, um.UUID, um.RbDataStatus, um.RbLocalDataStatus, um.RbLocalDeleted, um.RbLocalSynced, um.Usn, um.RbLocalUsn, um.CreatedAt, um.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	um._exists = true
	return nil
}

// Delete deletes the UUIDIDMap from the database.
func (c *Client) DeleteUUIDIDMap(ctx context.Context, um *UUIDIDMap) error {
	db := c.db

	switch {
	case !um._exists: // doesn't exist
		return nil
	case um._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM uuidIDMap ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, um.ID)
	if _, err := db.ExecContext(ctx, sqlstr, um.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	um._deleted = true
	return nil
}

func scanUUIDIDMapRows(rows *sql.Rows) ([]*UUIDIDMap, error) {
	var res []*UUIDIDMap
	for rows.Next() {
		um := UUIDIDMap{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&um.ID, &um.TableName, &um.TargetUUID, &um.CurrentID, &um.UUID, &um.RbDataStatus, &um.RbLocalDataStatus, &um.RbLocalDeleted, &um.RbLocalSynced, &um.Usn, &um.RbLocalUsn, &um.CreatedAt, &um.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &um)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllUUIDIDMap(ctx context.Context) ([]*UUIDIDMap, error) {
	db := c.db

	const sqlstr = `SELECT * FROM UUIDIDMap`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanUUIDIDMapRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UUIDIDMapByID retrieves a row from 'uuidIDMap' as a UUIDIDMap.
//
// Generated from index 'sqlite_autoindex_uuidIDMap_1'.
func (c *Client) UUIDIDMapByID(ctx context.Context, id nulltype.NullString) (*UUIDIDMap, error) {
	// func UUIDIDMapByID(ctx context.Context, db DB, id nulltype.NullString) (*UUIDIDMap, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM uuidIDMap ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	um := UUIDIDMap{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&um.ID, &um.TableName, &um.TargetUUID, &um.CurrentID, &um.UUID, &um.RbDataStatus, &um.RbLocalDataStatus, &um.RbLocalDeleted, &um.RbLocalSynced, &um.Usn, &um.RbLocalUsn, &um.CreatedAt, &um.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &um, nil
}

// UUIDIDMapByUUID retrieves a row from 'uuidIDMap' as a UUIDIDMap.
//
// Generated from index 'uuid_i_d_map__u_u_i_d'.
func (c *Client) UUIDIDMapByUUID(ctx context.Context, uuid nulltype.NullString) ([]*UUIDIDMap, error) {
	// func UUIDIDMapByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*UUIDIDMap, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM uuidIDMap ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*UUIDIDMap
	for rows.Next() {
		um := UUIDIDMap{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&um.ID, &um.TableName, &um.TargetUUID, &um.CurrentID, &um.UUID, &um.RbDataStatus, &um.RbLocalDataStatus, &um.RbLocalDeleted, &um.RbLocalSynced, &um.Usn, &um.RbLocalUsn, &um.CreatedAt, &um.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &um)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UUIDIDMapByRbDataStatus retrieves a row from 'uuidIDMap' as a UUIDIDMap.
//
// Generated from index 'uuid_i_d_map_rb_data_status'.
func (c *Client) UUIDIDMapByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*UUIDIDMap, error) {
	// func UUIDIDMapByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*UUIDIDMap, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM uuidIDMap ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*UUIDIDMap
	for rows.Next() {
		um := UUIDIDMap{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&um.ID, &um.TableName, &um.TargetUUID, &um.CurrentID, &um.UUID, &um.RbDataStatus, &um.RbLocalDataStatus, &um.RbLocalDeleted, &um.RbLocalSynced, &um.Usn, &um.RbLocalUsn, &um.CreatedAt, &um.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &um)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UUIDIDMapByRbLocalDataStatus retrieves a row from 'uuidIDMap' as a UUIDIDMap.
//
// Generated from index 'uuid_i_d_map_rb_local_data_status'.
func (c *Client) UUIDIDMapByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*UUIDIDMap, error) {
	// func UUIDIDMapByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*UUIDIDMap, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM uuidIDMap ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*UUIDIDMap
	for rows.Next() {
		um := UUIDIDMap{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&um.ID, &um.TableName, &um.TargetUUID, &um.CurrentID, &um.UUID, &um.RbDataStatus, &um.RbLocalDataStatus, &um.RbLocalDeleted, &um.RbLocalSynced, &um.Usn, &um.RbLocalUsn, &um.CreatedAt, &um.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &um)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UUIDIDMapByRbLocalDeleted retrieves a row from 'uuidIDMap' as a UUIDIDMap.
//
// Generated from index 'uuid_i_d_map_rb_local_deleted'.
func (c *Client) UUIDIDMapByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*UUIDIDMap, error) {
	// func UUIDIDMapByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*UUIDIDMap, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM uuidIDMap ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*UUIDIDMap
	for rows.Next() {
		um := UUIDIDMap{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&um.ID, &um.TableName, &um.TargetUUID, &um.CurrentID, &um.UUID, &um.RbDataStatus, &um.RbLocalDataStatus, &um.RbLocalDeleted, &um.RbLocalSynced, &um.Usn, &um.RbLocalUsn, &um.CreatedAt, &um.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &um)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UUIDIDMapByRbLocalUsnID retrieves a row from 'uuidIDMap' as a UUIDIDMap.
//
// Generated from index 'uuid_i_d_map_rb_local_usn__i_d'.
func (c *Client) UUIDIDMapByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*UUIDIDMap, error) {
	// func UUIDIDMapByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*UUIDIDMap, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, TableName, TargetUUID, CurrentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM uuidIDMap ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*UUIDIDMap
	for rows.Next() {
		um := UUIDIDMap{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&um.ID, &um.TableName, &um.TargetUUID, &um.CurrentID, &um.UUID, &um.RbDataStatus, &um.RbLocalDataStatus, &um.RbLocalDeleted, &um.RbLocalSynced, &um.Usn, &um.RbLocalUsn, &um.CreatedAt, &um.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &um)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
