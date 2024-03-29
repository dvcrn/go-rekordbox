package rekordbox

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	nulltype "github.com/mattn/go-nulltype"
)

// AgentRegistry represents a row from 'agentRegistry'.
type AgentRegistry struct {
	RegistryID nulltype.NullString `json:"registry_id"` // registry_id
	ID1        nulltype.NullString `json:"id1"`         // id_1
	ID2        nulltype.NullString `json:"id2"`         // id_2
	Int1       nulltype.NullInt64  `json:"int1"`        // int_1
	Int2       nulltype.NullInt64  `json:"int2"`        // int_2
	Str1       nulltype.NullString `json:"str1"`        // str_1
	Str2       nulltype.NullString `json:"str2"`        // str_2
	Date1      *Time               `json:"date1"`       // date_1
	Date2      *Time               `json:"date2"`       // date_2
	Text1      nulltype.NullString `json:"text1"`       // text_1
	Text2      nulltype.NullString `json:"text2"`       // text_2
	CreatedAt  Time                `json:"created_at"`  // created_at
	UpdatedAt  Time                `json:"updated_at"`  // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AgentRegistry exists in the database.
func (ar *AgentRegistry) Exists() bool {
	return ar._exists
}

// Deleted returns true when the AgentRegistry has been marked for deletion from
// the database.
func (ar *AgentRegistry) Deleted() bool {
	return ar._deleted
}

// Insert inserts the AgentRegistry to the database.
func (c *Client) InsertAgentRegistry(ctx context.Context, ar *AgentRegistry) error {
	db := c.db

	switch {
	case ar._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ar._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO agentRegistry (` +
		`registry_id, id_1, id_2, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)`
	// run
	logf(sqlstr, ar.RegistryID, ar.ID1, ar.ID2, ar.Int1, ar.Int2, ar.Str1, ar.Str2, ar.Date1, ar.Date2, ar.Text1, ar.Text2, ar.CreatedAt, ar.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, ar.RegistryID, ar.ID1, ar.ID2, ar.Int1, ar.Int2, ar.Str1, ar.Str2, ar.Date1, ar.Date2, ar.Text1, ar.Text2, ar.CreatedAt, ar.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	ar._exists = true
	return nil
}

// Update updates a AgentRegistry in the database.
func (c *Client) UpdateAgentRegistry(ctx context.Context, ar *AgentRegistry) error {
	db := c.db

	switch {
	case !ar._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ar._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE agentRegistry SET ` +
		`id_1 = $1, id_2 = $2, int_1 = $3, int_2 = $4, str_1 = $5, str_2 = $6, date_1 = $7, date_2 = $8, text_1 = $9, text_2 = $10, created_at = $11, updated_at = $12 ` +
		`WHERE registry_id = $13`
	// run
	logf(sqlstr, ar.ID1, ar.ID2, ar.Int1, ar.Int2, ar.Str1, ar.Str2, ar.Date1, ar.Date2, ar.Text1, ar.Text2, ar.CreatedAt, ar.UpdatedAt, ar.RegistryID)
	if _, err := db.ExecContext(ctx, sqlstr, ar.ID1, ar.ID2, ar.Int1, ar.Int2, ar.Str1, ar.Str2, ar.Date1, ar.Date2, ar.Text1, ar.Text2, ar.CreatedAt, ar.UpdatedAt, ar.RegistryID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AgentRegistry to the database.
func (c *Client) SaveAgentRegistry(ctx context.Context, ar *AgentRegistry) error {
	if ar.Exists() {
		return c.UpdateAgentRegistry(ctx, ar)
	}
	return c.InsertAgentRegistry(ctx, ar)
}

// Upsert performs an upsert for AgentRegistry.
func (c *Client) UpsertAgentRegistry(ctx context.Context, ar *AgentRegistry) error {
	db := c.db

	switch {
	case ar._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO agentRegistry (` +
		`registry_id, id_1, id_2, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (registry_id) DO ` +
		`UPDATE SET ` +
		`id_1 = EXCLUDED.id_1, id_2 = EXCLUDED.id_2, int_1 = EXCLUDED.int_1, int_2 = EXCLUDED.int_2, str_1 = EXCLUDED.str_1, str_2 = EXCLUDED.str_2, date_1 = EXCLUDED.date_1, date_2 = EXCLUDED.date_2, text_1 = EXCLUDED.text_1, text_2 = EXCLUDED.text_2, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, ar.RegistryID, ar.ID1, ar.ID2, ar.Int1, ar.Int2, ar.Str1, ar.Str2, ar.Date1, ar.Date2, ar.Text1, ar.Text2, ar.CreatedAt, ar.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, ar.RegistryID, ar.ID1, ar.ID2, ar.Int1, ar.Int2, ar.Str1, ar.Str2, ar.Date1, ar.Date2, ar.Text1, ar.Text2, ar.CreatedAt, ar.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	ar._exists = true
	return nil
}

// Delete deletes the AgentRegistry from the database.
func (c *Client) DeleteAgentRegistry(ctx context.Context, ar *AgentRegistry) error {
	db := c.db

	switch {
	case !ar._exists: // doesn't exist
		return nil
	case ar._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM agentRegistry ` +
		`WHERE registry_id = $1`
	// run
	logf(sqlstr, ar.RegistryID)
	if _, err := db.ExecContext(ctx, sqlstr, ar.RegistryID); err != nil {
		return logerror(err)
	}
	// set deleted
	ar._deleted = true
	return nil
}

func scanAgentRegistryRows(rows *sql.Rows) ([]*AgentRegistry, error) {
	var res []*AgentRegistry
	for rows.Next() {
		ar := AgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ar.RegistryID, &ar.ID1, &ar.ID2, &ar.Int1, &ar.Int2, &ar.Str1, &ar.Str2, &ar.Date1, &ar.Date2, &ar.Text1, &ar.Text2, &ar.CreatedAt, &ar.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ar)
	}

	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

func (c *Client) AllAgentRegistry(ctx context.Context) ([]*AgentRegistry, error) {
	db := c.db

	const sqlstr = `SELECT * FROM AgentRegistry`
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}

	defer rows.Close()
	res, err := scanAgentRegistryRows(rows)
	if err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AgentRegistryByID1ID2 retrieves a row from 'agentRegistry' as a AgentRegistry.
//
// Generated from index 'agent_registry_id_1_id_2'.
func (c *Client) AgentRegistryByID1ID2(ctx context.Context, id1, id2 nulltype.NullString) ([]*AgentRegistry, error) {
	// func AgentRegistryByID1ID2(ctx context.Context, db DB, id1 nulltype.NullString, id2 nulltype.NullString) ([]*AgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`registry_id, id_1, id_2, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, created_at, updated_at ` +
		`FROM agentRegistry ` +
		`WHERE id_1 = $1 AND id_2 = $2`
	// run
	logf(sqlstr, id1, id2)
	rows, err := db.QueryContext(ctx, sqlstr, id1, id2)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AgentRegistry
	for rows.Next() {
		ar := AgentRegistry{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ar.RegistryID, &ar.ID1, &ar.ID2, &ar.Int1, &ar.Int2, &ar.Str1, &ar.Str2, &ar.Date1, &ar.Date2, &ar.Text1, &ar.Text2, &ar.CreatedAt, &ar.UpdatedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ar)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AgentRegistryByRegistryID retrieves a row from 'agentRegistry' as a AgentRegistry.
//
// Generated from index 'sqlite_autoindex_agentRegistry_1'.
func (c *Client) AgentRegistryByRegistryID(ctx context.Context, registryID nulltype.NullString) (*AgentRegistry, error) {
	// func AgentRegistryByRegistryID(ctx context.Context, db DB, registryID nulltype.NullString) (*AgentRegistry, error) {
	db := c.db

	// query
	const sqlstr = `SELECT ` +
		`registry_id, id_1, id_2, int_1, int_2, str_1, str_2, date_1, date_2, text_1, text_2, created_at, updated_at ` +
		`FROM agentRegistry ` +
		`WHERE registry_id = $1`
	// run
	logf(sqlstr, registryID)
	ar := AgentRegistry{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, registryID).Scan(&ar.RegistryID, &ar.ID1, &ar.ID2, &ar.Int1, &ar.Int2, &ar.Str1, &ar.Str2, &ar.Date1, &ar.Date2, &ar.Text1, &ar.Text2, &ar.CreatedAt, &ar.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &ar, nil
}
