package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// DjmdMixerParam represents a row from 'djmdMixerParam'.
type DjmdMixerParam struct {
	ID                nulltype.NullString `json:"id"`                   // ID
	ContentID         nulltype.NullString `json:"content_id"`           // ContentID
	GainHigh          nulltype.NullInt64  `json:"gain_high"`            // GainHigh
	GainLow           nulltype.NullInt64  `json:"gain_low"`             // GainLow
	PeakHigh          nulltype.NullInt64  `json:"peak_high"`            // PeakHigh
	PeakLow           nulltype.NullInt64  `json:"peak_low"`             // PeakLow
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

// Exists returns true when the DjmdMixerParam exists in the database.
func (dmp *DjmdMixerParam) Exists() bool {
	return dmp._exists
}

// Deleted returns true when the DjmdMixerParam has been marked for deletion from
// the database.
func (dmp *DjmdMixerParam) Deleted() bool {
	return dmp._deleted
}

// Insert inserts the DjmdMixerParam to the database.
func (c *Client) InsertDjmdMixerParam(ctx context.Context, dmp *DjmdMixerParam) error {
	db := c.db

	switch {
	case dmp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dmp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO djmdMixerParam (` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)`
	// run
	logf(sqlstr, dmp.ID, dmp.ContentID, dmp.GainHigh, dmp.GainLow, dmp.PeakHigh, dmp.PeakLow, dmp.UUID, dmp.RbDataStatus, dmp.RbLocalDataStatus, dmp.RbLocalDeleted, dmp.RbLocalSynced, dmp.Usn, dmp.RbLocalUsn, dmp.CreatedAt, dmp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dmp.ID, dmp.ContentID, dmp.GainHigh, dmp.GainLow, dmp.PeakHigh, dmp.PeakLow, dmp.UUID, dmp.RbDataStatus, dmp.RbLocalDataStatus, dmp.RbLocalDeleted, dmp.RbLocalSynced, dmp.Usn, dmp.RbLocalUsn, dmp.CreatedAt, dmp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dmp._exists = true
	return nil
}

// Update updates a DjmdMixerParam in the database.
func (c *Client) UpdateDjmdMixerParam(ctx context.Context, dmp *DjmdMixerParam) error {
	db := c.db

	switch {
	case !dmp._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dmp._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE djmdMixerParam SET ` +
		`ContentID = $1, GainHigh = $2, GainLow = $3, PeakHigh = $4, PeakLow = $5, UUID = $6, rb_data_status = $7, rb_local_data_status = $8, rb_local_deleted = $9, rb_local_synced = $10, usn = $11, rb_local_usn = $12, created_at = $13, updated_at = $14 ` +
		`WHERE ID = $15`
	// run
	logf(sqlstr, dmp.ContentID, dmp.GainHigh, dmp.GainLow, dmp.PeakHigh, dmp.PeakLow, dmp.UUID, dmp.RbDataStatus, dmp.RbLocalDataStatus, dmp.RbLocalDeleted, dmp.RbLocalSynced, dmp.Usn, dmp.RbLocalUsn, dmp.CreatedAt, dmp.UpdatedAt, dmp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dmp.ContentID, dmp.GainHigh, dmp.GainLow, dmp.PeakHigh, dmp.PeakLow, dmp.UUID, dmp.RbDataStatus, dmp.RbLocalDataStatus, dmp.RbLocalDeleted, dmp.RbLocalSynced, dmp.Usn, dmp.RbLocalUsn, dmp.CreatedAt, dmp.UpdatedAt, dmp.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjmdMixerParam to the database.
func (c *Client) SaveDjmdMixerParam(ctx context.Context, dmp *DjmdMixerParam) error {
	if dmp.Exists() {
		return c.UpdateDjmdMixerParam(ctx, dmp)
	}
	return c.InsertDjmdMixerParam(ctx, dmp)
}

// Upsert performs an upsert for DjmdMixerParam.
func (c *Client) UpsertDjmdMixerParam(ctx context.Context, dmp *DjmdMixerParam) error {
	db := c.db

	switch {
	case dmp._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO djmdMixerParam (` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15` +
		`)` +
		` ON CONFLICT (ID) DO ` +
		`UPDATE SET ` +
		`ContentID = EXCLUDED.ContentID, GainHigh = EXCLUDED.GainHigh, GainLow = EXCLUDED.GainLow, PeakHigh = EXCLUDED.PeakHigh, PeakLow = EXCLUDED.PeakLow, UUID = EXCLUDED.UUID, rb_data_status = EXCLUDED.rb_data_status, rb_local_data_status = EXCLUDED.rb_local_data_status, rb_local_deleted = EXCLUDED.rb_local_deleted, rb_local_synced = EXCLUDED.rb_local_synced, usn = EXCLUDED.usn, rb_local_usn = EXCLUDED.rb_local_usn, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, dmp.ID, dmp.ContentID, dmp.GainHigh, dmp.GainLow, dmp.PeakHigh, dmp.PeakLow, dmp.UUID, dmp.RbDataStatus, dmp.RbLocalDataStatus, dmp.RbLocalDeleted, dmp.RbLocalSynced, dmp.Usn, dmp.RbLocalUsn, dmp.CreatedAt, dmp.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, dmp.ID, dmp.ContentID, dmp.GainHigh, dmp.GainLow, dmp.PeakHigh, dmp.PeakLow, dmp.UUID, dmp.RbDataStatus, dmp.RbLocalDataStatus, dmp.RbLocalDeleted, dmp.RbLocalSynced, dmp.Usn, dmp.RbLocalUsn, dmp.CreatedAt, dmp.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	dmp._exists = true
	return nil
}

// Delete deletes the DjmdMixerParam from the database.
func (c *Client) DeleteDjmdMixerParam(ctx context.Context, dmp *DjmdMixerParam) error {
	db := c.db

	switch {
	case !dmp._exists: // doesn't exist
		return nil
	case dmp._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM djmdMixerParam ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, dmp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dmp.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dmp._deleted = true
	return nil
}

func scanDjmdMixerParamRows(rows *sql.Rows) ([]*DjmdMixerParam, error) {
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllDjmdMixerParam(ctx context.Context) ([]*DjmdMixerParam, error) {
	db := c.db

	const sqlstr = `SELECT * FROM DjmdMixerParam`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanDjmdMixerParamRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByContentID retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'djmd_mixer_param__content_i_d'.
func (c *Client) DjmdMixerParamByContentID(ctx context.Context, contentID nulltype.NullString) ([]*DjmdMixerParam, error) {
	// func DjmdMixerParamByContentID(ctx context.Context, db DB, contentID nulltype.NullString) ([]*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE ContentID = $1`
	// run
	logf(sqlstr, contentID)
	rows, err := db.QueryContext(ctx, sqlstr, contentID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByContentIDRbLocalDeleted retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'djmd_mixer_param__content_i_d_rb_local_deleted'.
func (c *Client) DjmdMixerParamByContentIDRbLocalDeleted(ctx context.Context, contentID nulltype.NullString, rbLocalDeleted nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	// func DjmdMixerParamByContentIDRbLocalDeleted(ctx context.Context, db DB, contentID nulltype.NullString, rbLocalDeleted nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE ContentID = $1 AND rb_local_deleted = $2`
	// run
	logf(sqlstr, contentID, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, contentID, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByUUID retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'djmd_mixer_param__u_u_i_d'.
func (c *Client) DjmdMixerParamByUUID(ctx context.Context, uuid nulltype.NullString) ([]*DjmdMixerParam, error) {
	// func DjmdMixerParamByUUID(ctx context.Context, db DB, uuid nulltype.NullString) ([]*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE UUID = $1`
	// run
	logf(sqlstr, uuid)
	rows, err := db.QueryContext(ctx, sqlstr, uuid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByRbDataStatus retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'djmd_mixer_param_rb_data_status'.
func (c *Client) DjmdMixerParamByRbDataStatus(ctx context.Context, rbDataStatus nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	// func DjmdMixerParamByRbDataStatus(ctx context.Context, db DB, rbDataStatus nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE rb_data_status = $1`
	// run
	logf(sqlstr, rbDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByRbLocalDataStatus retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'djmd_mixer_param_rb_local_data_status'.
func (c *Client) DjmdMixerParamByRbLocalDataStatus(ctx context.Context, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	// func DjmdMixerParamByRbLocalDataStatus(ctx context.Context, db DB, rbLocalDataStatus nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE rb_local_data_status = $1`
	// run
	logf(sqlstr, rbLocalDataStatus)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDataStatus)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByRbLocalDeleted retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'djmd_mixer_param_rb_local_deleted'.
func (c *Client) DjmdMixerParamByRbLocalDeleted(ctx context.Context, rbLocalDeleted nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	// func DjmdMixerParamByRbLocalDeleted(ctx context.Context, db DB, rbLocalDeleted nulltype.NullInt64) ([]*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE rb_local_deleted = $1`
	// run
	logf(sqlstr, rbLocalDeleted)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalDeleted)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByRbLocalUsnID retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'djmd_mixer_param_rb_local_usn__i_d'.
func (c *Client) DjmdMixerParamByRbLocalUsnID(ctx context.Context, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdMixerParam, error) {
	// func DjmdMixerParamByRbLocalUsnID(ctx context.Context, db DB, rbLocalUsn nulltype.NullInt64, id nulltype.NullString) ([]*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE rb_local_usn = $1 AND ID = $2`
	// run
	logf(sqlstr, rbLocalUsn, id)
	rows, err := db.QueryContext(ctx, sqlstr, rbLocalUsn, id)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjmdMixerParam
	for rows.Next() {
		dmp := DjmdMixerParam{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &dmp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// DjmdMixerParamByID retrieves a row from 'djmdMixerParam' as a DjmdMixerParam.
//
// Generated from index 'sqlite_autoindex_djmdMixerParam_1'.
func (c *Client) DjmdMixerParamByID(ctx context.Context, id nulltype.NullString) (*DjmdMixerParam, error) {
	// func DjmdMixerParamByID(ctx context.Context, db DB, id nulltype.NullString) (*DjmdMixerParam, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`ID, ContentID, GainHigh, GainLow, PeakHigh, PeakLow, UUID, rb_data_status, rb_local_data_status, rb_local_deleted, rb_local_synced, usn, rb_local_usn, created_at, updated_at ` +
		`FROM djmdMixerParam ` +
		`WHERE ID = $1`
	// run
	logf(sqlstr, id)
	dmp := DjmdMixerParam{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dmp.ID, &dmp.ContentID, &dmp.GainHigh, &dmp.GainLow, &dmp.PeakHigh, &dmp.PeakLow, &dmp.UUID, &dmp.RbDataStatus, &dmp.RbLocalDataStatus, &dmp.RbLocalDeleted, &dmp.RbLocalSynced, &dmp.Usn, &dmp.RbLocalUsn, &dmp.CreatedAt, &dmp.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &dmp, nil
}
