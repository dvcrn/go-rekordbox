package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// CloudAgentRegistry represents a row from 'cloudAgentRegistry'.
type CloudAgentRegistry struct {
	ID                nulltype.NullString `json:"id"`                   // ID
	Int1              nulltype.NullInt64  `json:"int1"`                 // int_1
	Int2              nulltype.NullInt64  `json:"int2"`                 // int_2
	Str1              nulltype.NullString `json:"str1"`                 // str_1
	Str2              nulltype.NullString `json:"str2"`                 // str_2
	Date1             *Time               `json:"date1"`                // date_1
	Date2             *Time               `json:"date2"`                // date_2
	Text1             nulltype.NullString `json:"text1"`                // text_1
	Text2             nulltype.NullString `json:"text2"`                // text_2
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

// Exists returns true when the CloudAgentRegistry exists in the database.
func (car *CloudAgentRegistry) Exists() bool {
	return car._exists
}

// Deleted returns true when the CloudAgentRegistry has been marked for deletion from
// the database.
func (car *CloudAgentRegistry) Deleted() bool {
	return car._deleted
}

// Insert inserts the CloudAgentRegistry to the database.
func (c *Client) InsertCloudAgentRegistry(ctx context.Context, car *CloudAgentRegistry) error {
	db := c.db

	switch {
	case car._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case car._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO cloudAgentRegistry (` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`)`
	// run
	logf(sqlstr, car.ID, car.Int1, car.Int2, car.Str1, car.Str2, car.Date1, car.Date2, car.Text1, car.Text2, car.UUID, car.RbDataStatus, car.RbLocalDataStatus, car.RbLocalDeleted, car.RbLocalSynced, car.Usn, car.RbLocalUsn, car.CreatedAt, car.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, car.ID, car.Int1, car.Int2, car.Str1, car.Str2, car.Date1, car.Date2, car.Text1, car.Text2, car.UUID, car.RbDataStatus, car.RbLocalDataStatus, car.RbLocalDeleted, car.RbLocalSynced, car.Usn, car.RbLocalUsn, car.CreatedAt, car.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	car._exists = true
	return nil
}

// Update updates a CloudAgentRegistry in the database.
func (c *Client) UpdateCloudAgentRegistry(ctx context.Context, car *CloudAgentRegistry) error {
	db := c.db

	switch {
	case !car._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case car._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE cloudAgentRegistry SET ` +
		`int_1 = $1, int_2 = $2, str_1 = $3, str_2 = $4, date_1 = $5, date_2 = $6, text_1 = $7, text_2 = $8, UUID = $9, rb_data_status = $10, rb_local_data_status = $11, rb_local_deleted = $12, rb_local_synced = $13, usn = $14, rb_local_usn = $15, created_at = $16, updated_at = $17 ` +
		`WHERE ID = $18`
	// run
	logf(sqlstr, car.Int1, car.Int2, car.Str1, car.Str2, car.Date1, car.Date2, car.Text1, car.Text2, car.UUID, car.RbDataStatus, car.RbLocalDataStatus, car.RbLocalDeleted, car.RbLocalSynced, car.Usn, car.RbLocalUsn, car.CreatedAt, car.UpdatedAt, car.ID)
	if _, err := db.ExecContext(ctx, sqlstr, car.Int1, car.Int2, car.Str1, car.Str2, car.Date1, car.Date2, car.Text1, car.Text2, car.UUID, car.RbDataStatus, car.RbLocalDataStatus, car.RbLocalDeleted, car.RbLocalSynced, car.Usn, car.RbLocalUsn, car.CreatedAt, car.UpdatedAt, car.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the CloudAgentRegistry to the database.
func (c *Client) SaveCloudAgentRegistry(ctx context.Context, car *CloudAgentRegistry) error {
	if car.Exists() {
		return c.UpdateCloudAgentRegistry(ctx, car)
	}
	return c.InsertCloudAgentRegistry(ctx, car)
}

// Upsert performs an upsert for CloudAgentRegistry.
func (c *Client) UpsertCloudAgentRegistry(ctx context.Context, car *CloudAgentRegistry) error {
	db := c.db

	switch {
	case car._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO cloudAgentRegistry (` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`int_1 = EXCLUDED.int_1, int_2 = EXCLUDED.int_2, str_1 = EXCLUDED.str_1, str_2 = EXCLUDED.str_2, date_1 = EXCLUDED.date_1, date_2 = EXCLUDED.date_2, text_1 = EXCLUDED.text_1, text_2 = EXCLUDED.text_2, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, car.ID, car.Int1, car.Int2, car.Str1, car.Str2, car.Date1, car.Date2, car.Text1, car.Text2, car.UUID, car.RbDataStatus, car.RbLocalDataStatus, car.RbLocalDeleted, car.RbLocalSynced, car.Usn, car.RbLocalUsn, car.CreatedAt, car.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, car.ID, car.Int1, car.Int2, car.Str1, car.Str2, car.Date1, car.Date2, car.Text1, car.Text2, car.UUID, car.RbDataStatus, car.RbLocalDataStatus, car.RbLocalDeleted, car.RbLocalSynced, car.Usn, car.RbLocalUsn, car.CreatedAt, car.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	car._exists = true
	return nil
}

// Delete deletes the CloudAgentRegistry from the database.
func (c *Client) DeleteCloudAgentRegistry(ctx context.Context, car *CloudAgentRegistry) error {
	db := c.db

	switch {
	case !car._exists: // doesn't exist
		return nil
	case car._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM cloudAgentRegistry ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, car.ID)
	if _, err := db.ExecContext(ctx, sqlstr, car.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	car._deleted = true
	return nil
}

func scanCloudAgentRegistryRows(rows *sql.Rows) ([]*CloudAgentRegistry, error) {
	var res []*CloudAgentRegistry
	for rows.Next() {
		car := CloudAgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&car.ID, &car.Int1, &car.Int2, &car.Str1, &car.Str2, &car.Date1, &car.Date2, &car.Text1, &car.Text2, &car.UUID, &car.RbDataStatus, &car.RbLocalDataStatus, &car.RbLocalDeleted, &car.RbLocalSynced, &car.Usn, &car.RbLocalUsn, &car.CreatedAt, &car.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &car)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllCloudAgentRegistry(ctx context.Context) ([]*CloudAgentRegistry, error) {
	db := c.db

	const sqlstr = `SELECT * FROM CloudAgentRegistry`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanCloudAgentRegistryRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CloudAgentRegistryByUUID retrieves a row from 'cloudAgentRegistry' as a CloudAgentRegistry.
//
// Generated from index 'cloud_agent_registry__u_u_i_d'.
func (c *Client) CloudAgentRegistryByUUID(ctx context.Context, uuid nulltype.NullString) ([]*CloudAgentRegistry, error) {
	// func CloudAgentRegistryByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*CloudAgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM cloudAgentRegistry ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*CloudAgentRegistry
	for rows.Next() {
		car := CloudAgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&car.ID, &car.Int1, &car.Int2, &car.Str1, &car.Str2, &car.Date1, &car.Date2, &car.Text1, &car.Text2, &car.UUID, &car.RbDataStatus, &car.RbLocalDataStatus, &car.RbLocalDeleted, &car.RbLocalSynced, &car.Usn, &car.RbLocalUsn, &car.CreatedAt, &car.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &car)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CloudAgentRegistryByRbDataStatus retrieves a row from 'cloudAgentRegistry' as a CloudAgentRegistry.
//
// Generated from index 'cloud_agent_registry_rb_data_status'.
func (c *Client) CloudAgentRegistryByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*CloudAgentRegistry, error) {
	// func CloudAgentRegistryByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*CloudAgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM cloudAgentRegistry ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*CloudAgentRegistry
	for rows.Next() {
		car := CloudAgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&car.ID, &car.Int1, &car.Int2, &car.Str1, &car.Str2, &car.Date1, &car.Date2, &car.Text1, &car.Text2, &car.UUID, &car.RbDataStatus, &car.RbLocalDataStatus, &car.RbLocalDeleted, &car.RbLocalSynced, &car.Usn, &car.RbLocalUsn, &car.CreatedAt, &car.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &car)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CloudAgentRegistryByRbLocalDataStatus retrieves a row from 'cloudAgentRegistry' as a CloudAgentRegistry.
//
// Generated from index 'cloud_agent_registry_rb_local_data_status'.
func (c *Client) CloudAgentRegistryByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*CloudAgentRegistry, error) {
	// func CloudAgentRegistryByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*CloudAgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM cloudAgentRegistry ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*CloudAgentRegistry
	for rows.Next() {
		car := CloudAgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&car.ID, &car.Int1, &car.Int2, &car.Str1, &car.Str2, &car.Date1, &car.Date2, &car.Text1, &car.Text2, &car.UUID, &car.RbDataStatus, &car.RbLocalDataStatus, &car.RbLocalDeleted, &car.RbLocalSynced, &car.Usn, &car.RbLocalUsn, &car.CreatedAt, &car.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &car)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CloudAgentRegistryByRbLocalDeleted retrieves a row from 'cloudAgentRegistry' as a CloudAgentRegistry.
//
// Generated from index 'cloud_agent_registry_rb_local_deleted'.
func (c *Client) CloudAgentRegistryByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*CloudAgentRegistry, error) {
	// func CloudAgentRegistryByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*CloudAgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM cloudAgentRegistry ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*CloudAgentRegistry
	for rows.Next() {
		car := CloudAgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&car.ID, &car.Int1, &car.Int2, &car.Str1, &car.Str2, &car.Date1, &car.Date2, &car.Text1, &car.Text2, &car.UUID, &car.RbDataStatus, &car.RbLocalDataStatus, &car.RbLocalDeleted, &car.RbLocalSynced, &car.Usn, &car.RbLocalUsn, &car.CreatedAt, &car.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &car)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CloudAgentRegistryByRbLocalUsnID retrieves a row from 'cloudAgentRegistry' as a CloudAgentRegistry.
//
// Generated from index 'cloud_agent_registry_rb_local_usn__i_d'.
func (c *Client) CloudAgentRegistryByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*CloudAgentRegistry, error) {
	// func CloudAgentRegistryByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*CloudAgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM cloudAgentRegistry ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*CloudAgentRegistry
	for rows.Next() {
		car := CloudAgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&car.ID, &car.Int1, &car.Int2, &car.Str1, &car.Str2, &car.Date1, &car.Date2, &car.Text1, &car.Text2, &car.UUID, &car.RbDataStatus, &car.RbLocalDataStatus, &car.RbLocalDeleted, &car.RbLocalSynced, &car.Usn, &car.RbLocalUsn, &car.CreatedAt, &car.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &car)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// CloudAgentRegistryByID retrieves a row from 'cloudAgentRegistry' as a CloudAgentRegistry.
//
// Generated from index 'sqlite_autoindex_cloudAgentRegistry_1'.
func (c *Client) CloudAgentRegistryByID(ctx context.Context, id nulltype.NullString) (*CloudAgentRegistry, error) {
	// func CloudAgentRegistryByID(ctx context.Context, db DB, id nulltype.NullString) (*CloudAgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM cloudAgentRegistry ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	car := CloudAgentRegistry{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&car.ID, &car.Int1, &car.Int2, &car.Str1, &car.Str2, &car.Date1, &car.Date2, &car.Text1, &car.Text2, &car.UUID, &car.RbDataStatus, &car.RbLocalDataStatus, &car.RbLocalDeleted, &car.RbLocalSynced, &car.Usn, &car.RbLocalUsn, &car.CreatedAt, &car.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &car, nil
}
