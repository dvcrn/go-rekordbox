package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// ContentActiveCensor represents a row from 'contentActiveCensor'.
type ContentActiveCensor struct {
	ID                  nulltype.NullString `json:"ID"`                    // ID
	ContentID           nulltype.NullString `json:"ContentID"`             // ContentID
	ActiveCensors       nulltype.NullString `json:"ActiveCensors"`         // ActiveCensors
	RbActivecensorCount nulltype.NullInt64  `json:"rb_activecensor_count"` // rb_activecensor_count
	UUID                nulltype.NullString `json:"UUID"`                  // UUID
	RbDataStatus        nulltype.NullInt64  `json:"rb_data_status"`        // rb_data_status
	RbLocalDataStatus   nulltype.NullInt64  `json:"rb_local_data_status"`  // rb_local_data_status
	RbLocalDeleted      nulltype.NullInt64  `json:"rb_local_deleted"`      // rb_local_deleted
	RbLocalSynced       nulltype.NullInt64  `json:"rb_local_synced"`       // rb_local_synced
	Usn                 nulltype.NullInt64  `json:"usn"`                   // usn
	RbLocalUsn          nulltype.NullInt64  `json:"rb_local_usn"`          // rb_local_usn
	CreatedAt           Time                `json:"created_at"`            // created_at
	UpdatedAt           Time                `json:"updated_at"`            // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the ContentActiveCensor exists in the database.
func (cac *ContentActiveCensor) Exists() bool {
	return cac._exists
}

// Deleted returns true when the ContentActiveCensor has been marked for deletion from
// the database.
func (cac *ContentActiveCensor) Deleted() bool {
	return cac._deleted
}

// Insert inserts the ContentActiveCensor to the database.
func (c *Client) InsertContentActiveCensor(ctx context.Context, cac *ContentActiveCensor) error {
	db := c.db

	switch {
	case cac._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case cac._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO contentActiveCensor (` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)`
	// run
	logf(sqlstr, cac.ID, cac.ContentID, cac.ActiveCensors, cac.RbActivecensorCount, cac.UUID, cac.RbDataStatus, cac.RbLocalDataStatus, cac.RbLocalDeleted, cac.RbLocalSynced, cac.Usn, cac.RbLocalUsn, cac.CreatedAt, cac.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, cac.ID, cac.ContentID, cac.ActiveCensors, cac.RbActivecensorCount, cac.UUID, cac.RbDataStatus, cac.RbLocalDataStatus, cac.RbLocalDeleted, cac.RbLocalSynced, cac.Usn, cac.RbLocalUsn, cac.CreatedAt, cac.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	cac._exists = true
	return nil
}

// Update updates a ContentActiveCensor in the database.
func (c *Client) UpdateContentActiveCensor(ctx context.Context, cac *ContentActiveCensor) error {
	db := c.db

	switch {
	case !cac._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case cac._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE contentActiveCensor SET ` +
		`ContentID = $1, ActiveCensors = $2, rb_activecensor_count = $3, UUID = $4, rb_data_status = $5, rb_local_data_status = $6, rb_local_deleted = $7, rb_local_synced = $8, usn = $9, rb_local_usn = $10, created_at = $11, updated_at = $12 ` +
		`WHERE ID = $13`
	// run
	logf(sqlstr, cac.ContentID, cac.ActiveCensors, cac.RbActivecensorCount, cac.UUID, cac.RbDataStatus, cac.RbLocalDataStatus, cac.RbLocalDeleted, cac.RbLocalSynced, cac.Usn, cac.RbLocalUsn, cac.CreatedAt, cac.UpdatedAt, cac.ID)
	if _, err := db.ExecContext(ctx, sqlstr, cac.ContentID, cac.ActiveCensors, cac.RbActivecensorCount, cac.UUID, cac.RbDataStatus, cac.RbLocalDataStatus, cac.RbLocalDeleted, cac.RbLocalSynced, cac.Usn, cac.RbLocalUsn, cac.CreatedAt, cac.UpdatedAt, cac.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the ContentActiveCensor to the database.
func (c *Client) SaveContentActiveCensor(ctx context.Context, cac *ContentActiveCensor) error {
	if cac.Exists() {
		return c.UpdateContentActiveCensor(ctx, cac)
	}
	return c.InsertContentActiveCensor(ctx, cac)
}

// Upsert performs an upsert for ContentActiveCensor.
func (c *Client) UpsertContentActiveCensor(ctx context.Context, cac *ContentActiveCensor) error {
	db := c.db

	switch {
	case cac._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO contentActiveCensor (` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`ContentID = EXCLUDED.ContentID, ActiveCensors = EXCLUDED.ActiveCensors, rb_activecensor_count = EXCLUDED.rb_activecensor_count, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, cac.ID, cac.ContentID, cac.ActiveCensors, cac.RbActivecensorCount, cac.UUID, cac.RbDataStatus, cac.RbLocalDataStatus, cac.RbLocalDeleted, cac.RbLocalSynced, cac.Usn, cac.RbLocalUsn, cac.CreatedAt, cac.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, cac.ID, cac.ContentID, cac.ActiveCensors, cac.RbActivecensorCount, cac.UUID, cac.RbDataStatus, cac.RbLocalDataStatus, cac.RbLocalDeleted, cac.RbLocalSynced, cac.Usn, cac.RbLocalUsn, cac.CreatedAt, cac.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	cac._exists = true
	return nil
}

// Delete deletes the ContentActiveCensor from the database.
func (c *Client) DeleteContentActiveCensor(ctx context.Context, cac *ContentActiveCensor) error {
	db := c.db

	switch {
	case !cac._exists: // doesn't exist
		return nil
	case cac._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM contentActiveCensor ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, cac.ID)
	if _, err := db.ExecContext(ctx, sqlstr, cac.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	cac._deleted = true
	return nil
}

func scanContentActiveCensorRows(rows *sql.Rows) ([]*ContentActiveCensor, error) {
	var res []*ContentActiveCensor
	for rows.Next() {
		cac := ContentActiveCensor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cac)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllContentActiveCensor(ctx context.Context) ([]*ContentActiveCensor, error) {
	db := c.db

	const sqlstr = `SELECT * FROM ContentActiveCensor`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanContentActiveCensorRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ContentActiveCensorByContentID retrieves a row from 'contentActiveCensor' as a ContentActiveCensor.
//
// Generated from index 'content_active_censor__content_i_d'.
func (c *Client) ContentActiveCensorByContentID(ctx context.Context, contentID nulltype.NullString) ([]*ContentActiveCensor, error) {
	// func ContentActiveCensorByContentID(ctx context.Context, db DB, contentID nulltype.NullString) ([]*ContentActiveCensor, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM contentActiveCensor ` +
		`WHERE ContentID = $1`
	// run
	logf(sqlstr, contentID)
	rows, err := db.QueryContext(ctx, sqlstr, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ContentActiveCensor
	for rows.Next() {
		cac := ContentActiveCensor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cac)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ContentActiveCensorByUUID retrieves a row from 'contentActiveCensor' as a ContentActiveCensor.
//
// Generated from index 'content_active_censor__u_u_i_d'.
func (c *Client) ContentActiveCensorByUUID(ctx context.Context, uuid nulltype.NullString) ([]*ContentActiveCensor, error) {
	// func ContentActiveCensorByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*ContentActiveCensor, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM contentActiveCensor ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ContentActiveCensor
	for rows.Next() {
		cac := ContentActiveCensor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cac)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ContentActiveCensorByRbDataStatus retrieves a row from 'contentActiveCensor' as a ContentActiveCensor.
//
// Generated from index 'content_active_censor_rb_data_status'.
func (c *Client) ContentActiveCensorByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*ContentActiveCensor, error) {
	// func ContentActiveCensorByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*ContentActiveCensor, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM contentActiveCensor ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ContentActiveCensor
	for rows.Next() {
		cac := ContentActiveCensor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cac)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ContentActiveCensorByRbLocalDataStatus retrieves a row from 'contentActiveCensor' as a ContentActiveCensor.
//
// Generated from index 'content_active_censor_rb_local_data_status'.
func (c *Client) ContentActiveCensorByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*ContentActiveCensor, error) {
	// func ContentActiveCensorByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*ContentActiveCensor, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM contentActiveCensor ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ContentActiveCensor
	for rows.Next() {
		cac := ContentActiveCensor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cac)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ContentActiveCensorByRbLocalDeleted retrieves a row from 'contentActiveCensor' as a ContentActiveCensor.
//
// Generated from index 'content_active_censor_rb_local_deleted'.
func (c *Client) ContentActiveCensorByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*ContentActiveCensor, error) {
	// func ContentActiveCensorByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*ContentActiveCensor, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM contentActiveCensor ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ContentActiveCensor
	for rows.Next() {
		cac := ContentActiveCensor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cac)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ContentActiveCensorByRbLocalUsnID retrieves a row from 'contentActiveCensor' as a ContentActiveCensor.
//
// Generated from index 'content_active_censor_rb_local_usn__i_d'.
func (c *Client) ContentActiveCensorByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*ContentActiveCensor, error) {
	// func ContentActiveCensorByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*ContentActiveCensor, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM contentActiveCensor ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*ContentActiveCensor
	for rows.Next() {
		cac := ContentActiveCensor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &cac)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// ContentActiveCensorByID retrieves a row from 'contentActiveCensor' as a ContentActiveCensor.
//
// Generated from index 'sqlite_autoindex_contentActiveCensor_1'.
func (c *Client) ContentActiveCensorByID(ctx context.Context, id nulltype.NullString) (*ContentActiveCensor, error) {
	// func ContentActiveCensorByID(ctx context.Context, db DB, id nulltype.NullString) (*ContentActiveCensor, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, ActiveCensors, rb_activecensor_count, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM contentActiveCensor ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	cac := ContentActiveCensor{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&cac.ID, &cac.ContentID, &cac.ActiveCensors, &cac.RbActivecensorCount, &cac.UUID, &cac.RbDataStatus, &cac.RbLocalDataStatus, &cac.RbLocalDeleted, &cac.RbLocalSynced, &cac.Usn, &cac.RbLocalUsn, &cac.CreatedAt, &cac.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &cac, nil
}
