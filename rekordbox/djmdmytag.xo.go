package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdMyTag represents a row from 'djmdMyTag'.
type DjmdMyTag struct {
	ID                nulltype.NullString `json:"id"`                   // ID
	Seq               nulltype.NullInt64  `json:"seq"`                  // Seq
	Name              nulltype.NullString `json:"name"`                 // Name
	Attribute         nulltype.NullInt64  `json:"attribute"`            // Attribute
	ParentID          nulltype.NullString `json:"parent_id"`            // ParentID
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

// Exists returns true when the DjmdMyTag exists in the database.
func (dmt *DjmdMyTag) Exists() bool {
	return dmt._exists
}

// Deleted returns true when the DjmdMyTag has been marked for deletion from
// the database.
func (dmt *DjmdMyTag) Deleted() bool {
	return dmt._deleted
}

// Insert inserts the DjmdMyTag to the database.
func (c *Client) InsertDjmdMyTag(ctx context.Context, dmt *DjmdMyTag) error {
	db := c.db

	switch {
	case dmt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dmt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdMyTag (` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`)`
	// run
	logf(sqlstr, dmt.ID, dmt.Seq, dmt.Name, dmt.Attribute, dmt.ParentID, dmt.UUID, dmt.RbDataStatus, dmt.RbLocalDataStatus, dmt.RbLocalDeleted, dmt.RbLocalSynced, dmt.Usn, dmt.RbLocalUsn, dmt.CreatedAt, dmt.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dmt.ID, dmt.Seq, dmt.Name, dmt.Attribute, dmt.ParentID, dmt.UUID, dmt.RbDataStatus, dmt.RbLocalDataStatus, dmt.RbLocalDeleted, dmt.RbLocalSynced, dmt.Usn, dmt.RbLocalUsn, dmt.CreatedAt, dmt.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dmt._exists = true
	return nil
}

// Update updates a DjmdMyTag in the database.
func (c *Client) UpdateDjmdMyTag(ctx context.Context, dmt *DjmdMyTag) error {
	db := c.db

	switch {
	case !dmt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dmt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdMyTag SET ` +
		`Seq = $1, Name = $2, Attribute = $3, ParentID = $4, UUID = $5, rb_data_status = $6, rb_local_data_status = $7, rb_local_deleted = $8, rb_local_synced = $9, usn = $10, rb_local_usn = $11, created_at = $12, updated_at = $13 ` +
		`WHERE ID = $14`
	// run
	logf(sqlstr, dmt.Seq, dmt.Name, dmt.Attribute, dmt.ParentID, dmt.UUID, dmt.RbDataStatus, dmt.RbLocalDataStatus, dmt.RbLocalDeleted, dmt.RbLocalSynced, dmt.Usn, dmt.RbLocalUsn, dmt.CreatedAt, dmt.UpdatedAt, dmt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dmt.Seq, dmt.Name, dmt.Attribute, dmt.ParentID, dmt.UUID, dmt.RbDataStatus, dmt.RbLocalDataStatus, dmt.RbLocalDeleted, dmt.RbLocalSynced, dmt.Usn, dmt.RbLocalUsn, dmt.CreatedAt, dmt.UpdatedAt, dmt.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdMyTag to the database.
func (c *Client) SaveDjmdMyTag(ctx context.Context, dmt *DjmdMyTag) error {
	if dmt.Exists() {
		return c.UpdateDjmdMyTag(ctx, dmt)
	}
	return c.InsertDjmdMyTag(ctx, dmt)
}

// Upsert performs an upsert for DjmdMyTag.
func (c *Client) UpsertDjmdMyTag(ctx context.Context, dmt *DjmdMyTag) error {
	db := c.db

	switch {
	case dmt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdMyTag (` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`Seq = EXCLUDED.Seq, Name = EXCLUDED.Name, Attribute = EXCLUDED.Attribute, ParentID = EXCLUDED.ParentID, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dmt.ID, dmt.Seq, dmt.Name, dmt.Attribute, dmt.ParentID, dmt.UUID, dmt.RbDataStatus, dmt.RbLocalDataStatus, dmt.RbLocalDeleted, dmt.RbLocalSynced, dmt.Usn, dmt.RbLocalUsn, dmt.CreatedAt, dmt.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dmt.ID, dmt.Seq, dmt.Name, dmt.Attribute, dmt.ParentID, dmt.UUID, dmt.RbDataStatus, dmt.RbLocalDataStatus, dmt.RbLocalDeleted, dmt.RbLocalSynced, dmt.Usn, dmt.RbLocalUsn, dmt.CreatedAt, dmt.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dmt._exists = true
	return nil
}

// Delete deletes the DjmdMyTag from the database.
func (c *Client) DeleteDjmdMyTag(ctx context.Context, dmt *DjmdMyTag) error {
	db := c.db

	switch {
	case !dmt._exists: // doesn't exist
		return nil
	case dmt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdMyTag ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dmt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dmt.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dmt._deleted = true
	return nil
}

func scanDjmdMyTagRows(rows *sql.Rows) ([]*DjmdMyTag, error) {
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdMyTag(ctx context.Context) ([]*DjmdMyTag, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdMyTag`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdMyTagRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagByParentID retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'djmd_my_tag__parent_i_d'.
func (c *Client) DjmdMyTagByParentID(ctx context.Context, parentID nulltype.NullString) ([]*DjmdMyTag, error) {
	// func DjmdMyTagByParentID(ctx context.Context, db DB, parentID nulltype.NullString) ([]*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE ParentID = $1`
	// run
	logf(sqlstr, parentID)
	rows, err := db.QueryContext(ctx, sqlstr, parentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagBySeq retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'djmd_my_tag__seq'.
func (c *Client) DjmdMyTagBySeq(ctx context.Context, seq nulltype.NullInt64) ([]*DjmdMyTag, error) {
	// func DjmdMyTagBySeq(ctx context.Context, db DB, seq nulltype.NullInt64) ([]*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE Seq = $1`
	// run
	logf(sqlstr, seq)
	rows, err := db.QueryContext(ctx, sqlstr, seq)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagByUUID retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'djmd_my_tag__u_u_i_d'.
func (c *Client) DjmdMyTagByUUID(ctx context.Context, uuid nulltype.NullString) ([]*DjmdMyTag, error) {
	// func DjmdMyTagByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagByRbDataStatus retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'djmd_my_tag_rb_data_status'.
func (c *Client) DjmdMyTagByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*DjmdMyTag, error) {
	// func DjmdMyTagByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagByRbLocalDataStatus retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'djmd_my_tag_rb_local_data_status'.
func (c *Client) DjmdMyTagByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdMyTag, error) {
	// func DjmdMyTagByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagByRbLocalDeleted retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'djmd_my_tag_rb_local_deleted'.
func (c *Client) DjmdMyTagByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*DjmdMyTag, error) {
	// func DjmdMyTagByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagByRbLocalUsnID retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'djmd_my_tag_rb_local_usn__i_d'.
func (c *Client) DjmdMyTagByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdMyTag, error) {
	// func DjmdMyTagByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMyTag
	for rows.Next() {
		dmt := DjmdMyTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMyTagByID retrieves a row from 'djmdMyTag' as a DjmdMyTag.
//
// Generated from index 'sqlite_autoindex_djmdMyTag_1'.
func (c *Client) DjmdMyTagByID(ctx context.Context, id nulltype.NullString) (*DjmdMyTag, error) {
	// func DjmdMyTagByID(ctx context.Context, db DB, id nulltype.NullString) (*DjmdMyTag, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, Seq, Name, Attribute, ParentID, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMyTag ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dmt := DjmdMyTag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dmt.ID, &dmt.Seq, &dmt.Name, &dmt.Attribute, &dmt.ParentID, &dmt.UUID, &dmt.RbDataStatus, &dmt.RbLocalDataStatus, &dmt.RbLocalDeleted, &dmt.RbLocalSynced, &dmt.Usn, &dmt.RbLocalUsn, &dmt.CreatedAt, &dmt.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dmt, nil
}
