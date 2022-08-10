package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// DjmdSort represents a row from 'djmdSort'.
type DjmdSort struct {
	ID                sql.NullString `json:"id"`                   // ID
	MenuItemID        sql.NullString `json:"menu_item_id"`         // MenuItemID
	Seq               sql.NullInt64  `json:"seq"`                  // Seq
	Disable           sql.NullInt64  `json:"disable"`              // Disable
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

// Exists returns true when the DjmdSort exists in the database.
func (ds *DjmdSort) Exists() bool {
	return ds._exists
}

// Deleted returns true when the DjmdSort has been marked for deletion from
// the database.
func (ds *DjmdSort) Deleted() bool {
	return ds._deleted
}

// Insert inserts the DjmdSort to the database.
func (c *Client) InsertDjmdSort(ctx context.Context, ds *DjmdSort) error {
	db := c.db

	switch {
	case ds._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ds._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdSort (` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)`
	// run
	logf(sqlstr, ds.ID, ds.MenuItemID, ds.Seq, ds.Disable, ds.UUID, ds.RbDataStatus, ds.RbLocalDataStatus, ds.RbLocalDeleted, ds.RbLocalSynced, ds.Usn, ds.RbLocalUsn, ds.CreatedAt, ds.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, ds.ID, ds.MenuItemID, ds.Seq, ds.Disable, ds.UUID, ds.RbDataStatus, ds.RbLocalDataStatus, ds.RbLocalDeleted, ds.RbLocalSynced, ds.Usn, ds.RbLocalUsn, ds.CreatedAt, ds.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	ds._exists = true
	return nil
}

// Update updates a DjmdSort in the database.
func (c *Client) UpdateDjmdSort(ctx context.Context, ds *DjmdSort) error {
	db := c.db

	switch {
	case !ds._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ds._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdSort SET ` +
		`MenuItemID = $1, Seq = $2, Disable = $3, UUID = $4, rb_data_status = $5, rb_local_data_status = $6, rb_local_deleted = $7, rb_local_synced = $8, usn = $9, rb_local_usn = $10, created_at = $11, updated_at = $12 ` +
		`WHERE ID = $13`
	// run
	logf(sqlstr, ds.MenuItemID, ds.Seq, ds.Disable, ds.UUID, ds.RbDataStatus, ds.RbLocalDataStatus, ds.RbLocalDeleted, ds.RbLocalSynced, ds.Usn, ds.RbLocalUsn, ds.CreatedAt, ds.UpdatedAt, ds.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ds.MenuItemID, ds.Seq, ds.Disable, ds.UUID, ds.RbDataStatus, ds.RbLocalDataStatus, ds.RbLocalDeleted, ds.RbLocalSynced, ds.Usn, ds.RbLocalUsn, ds.CreatedAt, ds.UpdatedAt, ds.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdSort to the database.
func (c *Client) SaveDjmdSort(ctx context.Context, ds *DjmdSort) error {
	if ds.Exists() {
		return c.UpdateDjmdSort(ctx, ds)
	}
	return c.InsertDjmdSort(ctx, ds)
}

// Upsert performs an upsert for DjmdSort.
func (c *Client) UpsertDjmdSort(ctx context.Context, ds *DjmdSort) error {
	db := c.db

	switch {
	case ds._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdSort (` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`MenuItemID = EXCLUDED.MenuItemID, Seq = EXCLUDED.Seq, Disable = EXCLUDED.Disable, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, ds.ID, ds.MenuItemID, ds.Seq, ds.Disable, ds.UUID, ds.RbDataStatus, ds.RbLocalDataStatus, ds.RbLocalDeleted, ds.RbLocalSynced, ds.Usn, ds.RbLocalUsn, ds.CreatedAt, ds.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, ds.ID, ds.MenuItemID, ds.Seq, ds.Disable, ds.UUID, ds.RbDataStatus, ds.RbLocalDataStatus, ds.RbLocalDeleted, ds.RbLocalSynced, ds.Usn, ds.RbLocalUsn, ds.CreatedAt, ds.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	ds._exists = true
	return nil
}

// Delete deletes the DjmdSort from the database.
func (c *Client) DeleteDjmdSort(ctx context.Context, ds *DjmdSort) error {
	db := c.db

	switch {
	case !ds._exists: // doesn't exist
		return nil
	case ds._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdSort ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, ds.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ds.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	ds._deleted = true
	return nil
}

func scanDjmdSortRows(rows *sql.Rows) ([]*DjmdSort, error) {
	var res []*DjmdSort
	for rows.Next() {
		ds := DjmdSort{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ds.ID, &ds.MenuItemID, &ds.Seq, &ds.Disable, &ds.UUID, &ds.RbDataStatus, &ds.RbLocalDataStatus, &ds.RbLocalDeleted, &ds.RbLocalSynced, &ds.Usn, &ds.RbLocalUsn, &ds.CreatedAt, &ds.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ds)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdSort(ctx context.Context) ([]*DjmdSort, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdSort`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdSortRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSortByUUID retrieves a row from 'djmdSort' as a DjmdSort.
//
// Generated from index 'djmd_sort__u_u_i_d'.
func (c *Client) DjmdSortByUUID(ctx context.Context, uuid sql.NullString) ([]*DjmdSort, error) {
	// func DjmdSortByUUID(ctx context.Context, db DB, uuid sql.NullString) ([]*DjmdSort, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSort ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSort
	for rows.Next() {
		ds := DjmdSort{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ds.ID, &ds.MenuItemID, &ds.Seq, &ds.Disable, &ds.UUID, &ds.RbDataStatus, &ds.RbLocalDataStatus, &ds.RbLocalDeleted, &ds.RbLocalSynced, &ds.Usn, &ds.RbLocalUsn, &ds.CreatedAt, &ds.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ds)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSortByRbDataStatus retrieves a row from 'djmdSort' as a DjmdSort.
//
// Generated from index 'djmd_sort_rb_data_status'.
func (c *Client) DjmdSortByRbDataStatus(ctx context.Context, rbDataStatus sql.NullInt64) ([]*DjmdSort, error) {
	// func DjmdSortByRbDataStatus(ctx context.Context, db DB, rbDataStatus sql.NullInt64) ([]*DjmdSort, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSort ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSort
	for rows.Next() {
		ds := DjmdSort{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ds.ID, &ds.MenuItemID, &ds.Seq, &ds.Disable, &ds.UUID, &ds.RbDataStatus, &ds.RbLocalDataStatus, &ds.RbLocalDeleted, &ds.RbLocalSynced, &ds.Usn, &ds.RbLocalUsn, &ds.CreatedAt, &ds.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ds)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSortByRbLocalDataStatus retrieves a row from 'djmdSort' as a DjmdSort.
//
// Generated from index 'djmd_sort_rb_local_data_status'.
func (c *Client) DjmdSortByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus sql.NullInt64) ([]*DjmdSort, error) {
	// func DjmdSortByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus sql.NullInt64) ([]*DjmdSort, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSort ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSort
	for rows.Next() {
		ds := DjmdSort{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ds.ID, &ds.MenuItemID, &ds.Seq, &ds.Disable, &ds.UUID, &ds.RbDataStatus, &ds.RbLocalDataStatus, &ds.RbLocalDeleted, &ds.RbLocalSynced, &ds.Usn, &ds.RbLocalUsn, &ds.CreatedAt, &ds.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ds)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSortByRbLocalDeleted retrieves a row from 'djmdSort' as a DjmdSort.
//
// Generated from index 'djmd_sort_rb_local_deleted'.
func (c *Client) DjmdSortByRbLocalDeleted(ctx context.Context, rbLocalDeleted sql.NullInt64) ([]*DjmdSort, error) {
	// func DjmdSortByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted sql.NullInt64) ([]*DjmdSort, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSort ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSort
	for rows.Next() {
		ds := DjmdSort{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ds.ID, &ds.MenuItemID, &ds.Seq, &ds.Disable, &ds.UUID, &ds.RbDataStatus, &ds.RbLocalDataStatus, &ds.RbLocalDeleted, &ds.RbLocalSynced, &ds.Usn, &ds.RbLocalUsn, &ds.CreatedAt, &ds.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ds)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSortByRbLocalUsnID retrieves a row from 'djmdSort' as a DjmdSort.
//
// Generated from index 'djmd_sort_rb_local_usn__i_d'.
func (c *Client) DjmdSortByRbLocalUsnID(ctx context.Context, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdSort, error) {
	// func DjmdSortByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn sql.NullInt64, id sql.NullString) ([]*DjmdSort, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSort ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdSort
	for rows.Next() {
		ds := DjmdSort{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ds.ID, &ds.MenuItemID, &ds.Seq, &ds.Disable, &ds.UUID, &ds.RbDataStatus, &ds.RbLocalDataStatus, &ds.RbLocalDeleted, &ds.RbLocalSynced, &ds.Usn, &ds.RbLocalUsn, &ds.CreatedAt, &ds.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ds)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdSortByID retrieves a row from 'djmdSort' as a DjmdSort.
//
// Generated from index 'sqlite_autoindex_djmdSort_1'.
func (c *Client) DjmdSortByID(ctx context.Context, id sql.NullString) (*DjmdSort, error) {
	// func DjmdSortByID(ctx context.Context, db DB, id sql.NullString) (*DjmdSort, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, MenuItemID, Seq, Disable, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdSort ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	ds := DjmdSort{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&ds.ID, &ds.MenuItemID, &ds.Seq, &ds.Disable, &ds.UUID, &ds.RbDataStatus, &ds.RbLocalDataStatus, &ds.RbLocalDeleted, &ds.RbLocalSynced, &ds.Usn, &ds.RbLocalUsn, &ds.CreatedAt, &ds.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &ds, nil
}
