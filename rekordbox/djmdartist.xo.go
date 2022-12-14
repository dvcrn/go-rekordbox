package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdArtist represents a row from 'djmdArtist'.
type DjmdArtist struct {
	ID                nulltype.NullString `json:"id"`                   // ID
	Name              nulltype.NullString `json:"name"`                 // Name
	SearchStr         nulltype.NullString `json:"search_str"`           // SearchStr
	UUID              nulltype.NullString `json:"uuid"`                 // UUID
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

// Exists returns true when the DjmdArtist exists in the database.
func (da *DjmdArtist) Exists() bool {
	return da._exists
}

// Deleted returns true when the DjmdArtist has been marked for deletion from
// the database.
func (da *DjmdArtist) Deleted() bool {
	return da._deleted
}

// Insert inserts the DjmdArtist to the database.
func (c *Client) InsertDjmdArtist(ctx context.Context, da *DjmdArtist) error {
	db := c.db

	switch {
	case da._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case da._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdArtist (` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`)`
	// run
	logf(sqlstr, da.ID, da.Name, da.SearchStr, da.UUID, da.RbDataStatus, da.RbLocalDataStatus, da.RbLocalDeleted, da.RbLocalSynced, da.Usn, da.RbLocalUsn, da.CreatedAt, da.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, da.ID, da.Name, da.SearchStr, da.UUID, da.RbDataStatus, da.RbLocalDataStatus, da.RbLocalDeleted, da.RbLocalSynced, da.Usn, da.RbLocalUsn, da.CreatedAt, da.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	da._exists = true
	return nil
}

// Update updates a DjmdArtist in the database.
func (c *Client) UpdateDjmdArtist(ctx context.Context, da *DjmdArtist) error {
	db := c.db

	switch {
	case !da._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case da._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdArtist SET ` +
		`Name = $1, SearchStr = $2, UUID = $3, rb_data_status = $4, rb_local_data_status = $5, rb_local_deleted = $6, rb_local_synced = $7, usn = $8, rb_local_usn = $9, created_at = $10, updated_at = $11 ` +
		`WHERE ID = $12`
	// run
	logf(sqlstr, da.Name, da.SearchStr, da.UUID, da.RbDataStatus, da.RbLocalDataStatus, da.RbLocalDeleted, da.RbLocalSynced, da.Usn, da.RbLocalUsn, da.CreatedAt, da.UpdatedAt, da.ID)
	if _, err := db.ExecContext(ctx, sqlstr, da.Name, da.SearchStr, da.UUID, da.RbDataStatus, da.RbLocalDataStatus, da.RbLocalDeleted, da.RbLocalSynced, da.Usn, da.RbLocalUsn, da.CreatedAt, da.UpdatedAt, da.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdArtist to the database.
func (c *Client) SaveDjmdArtist(ctx context.Context, da *DjmdArtist) error {
	if da.Exists() {
		return c.UpdateDjmdArtist(ctx, da)
	}
	return c.InsertDjmdArtist(ctx, da)
}

// Upsert performs an upsert for DjmdArtist.
func (c *Client) UpsertDjmdArtist(ctx context.Context, da *DjmdArtist) error {
	db := c.db

	switch {
	case da._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdArtist (` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`Name = EXCLUDED.Name, SearchStr = EXCLUDED.SearchStr, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, da.ID, da.Name, da.SearchStr, da.UUID, da.RbDataStatus, da.RbLocalDataStatus, da.RbLocalDeleted, da.RbLocalSynced, da.Usn, da.RbLocalUsn, da.CreatedAt, da.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, da.ID, da.Name, da.SearchStr, da.UUID, da.RbDataStatus, da.RbLocalDataStatus, da.RbLocalDeleted, da.RbLocalSynced, da.Usn, da.RbLocalUsn, da.CreatedAt, da.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	da._exists = true
	return nil
}

// Delete deletes the DjmdArtist from the database.
func (c *Client) DeleteDjmdArtist(ctx context.Context, da *DjmdArtist) error {
	db := c.db

	switch {
	case !da._exists: // doesn't exist
		return nil
	case da._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdArtist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, da.ID)
	if _, err := db.ExecContext(ctx, sqlstr, da.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	da._deleted = true
	return nil
}

func scanDjmdArtistRows(rows *sql.Rows) ([]*DjmdArtist, error) {
	var res []*DjmdArtist
	for rows.Next() {
		da := DjmdArtist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &da)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdArtist(ctx context.Context) ([]*DjmdArtist, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdArtist`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdArtistRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdArtistByName retrieves a row from 'djmdArtist' as a DjmdArtist.
//
// Generated from index 'djmd_artist__name'.
func (c *Client) DjmdArtistByName(ctx context.Context, name nulltype.NullString) ([]*DjmdArtist, error) {
	// func DjmdArtistByName(ctx context.Context, db DB, name nulltype.NullString) ([]*DjmdArtist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdArtist ` +
		`WHERE Name = $1`
	// run
	logf(sqlstr, name)
	rows, err := db.QueryContext(ctx, sqlstr, name)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdArtist
	for rows.Next() {
		da := DjmdArtist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &da)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdArtistByUUID retrieves a row from 'djmdArtist' as a DjmdArtist.
//
// Generated from index 'djmd_artist__u_u_i_d'.
func (c *Client) DjmdArtistByUUID(ctx context.Context, uuid nulltype.NullString) ([]*DjmdArtist, error) {
	// func DjmdArtistByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*DjmdArtist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdArtist ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdArtist
	for rows.Next() {
		da := DjmdArtist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &da)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdArtistByRbDataStatus retrieves a row from 'djmdArtist' as a DjmdArtist.
//
// Generated from index 'djmd_artist_rb_data_status'.
func (c *Client) DjmdArtistByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*DjmdArtist, error) {
	// func DjmdArtistByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*DjmdArtist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdArtist ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdArtist
	for rows.Next() {
		da := DjmdArtist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &da)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdArtistByRbLocalDataStatus retrieves a row from 'djmdArtist' as a DjmdArtist.
//
// Generated from index 'djmd_artist_rb_local_data_status'.
func (c *Client) DjmdArtistByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdArtist, error) {
	// func DjmdArtistByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdArtist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdArtist ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdArtist
	for rows.Next() {
		da := DjmdArtist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &da)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdArtistByRbLocalDeleted retrieves a row from 'djmdArtist' as a DjmdArtist.
//
// Generated from index 'djmd_artist_rb_local_deleted'.
func (c *Client) DjmdArtistByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*DjmdArtist, error) {
	// func DjmdArtistByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*DjmdArtist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdArtist ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdArtist
	for rows.Next() {
		da := DjmdArtist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &da)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdArtistByRbLocalUsnID retrieves a row from 'djmdArtist' as a DjmdArtist.
//
// Generated from index 'djmd_artist_rb_local_usn__i_d'.
func (c *Client) DjmdArtistByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdArtist, error) {
	// func DjmdArtistByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdArtist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdArtist ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdArtist
	for rows.Next() {
		da := DjmdArtist{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &da)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdArtistByID retrieves a row from 'djmdArtist' as a DjmdArtist.
//
// Generated from index 'sqlite_autoindex_djmdArtist_1'.
func (c *Client) DjmdArtistByID(ctx context.Context, id nulltype.NullString) (*DjmdArtist, error) {
	// func DjmdArtistByID(ctx context.Context, db DB, id nulltype.NullString) (*DjmdArtist, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Name, SearchStr, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdArtist ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	da := DjmdArtist{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&da.ID, &da.Name, &da.SearchStr, &da.UUID, &da.RbDataStatus, &da.RbLocalDataStatus, &da.RbLocalDeleted, &da.RbLocalSynced, &da.Usn, &da.RbLocalUsn, &da.CreatedAt, &da.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &da, nil
}
