// Code generated by go generate; DO NOT EDIT.
package db

import (
	stdsql "database/sql"
	"encoding/json"

	"agola.io/agola/internal/errors"
	"agola.io/agola/internal/sql"
	"agola.io/agola/services/configstore/types"

	sq "github.com/Masterminds/squirrel"
)

func (d *DB) fetchRemoteSources(tx *sql.Tx, q sq.Sqlizer) ([]*types.RemoteSource, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanRemoteSources(rows)
}

func (d *DB) scanRemoteSource(rows *stdsql.Rows, additionalFields []interface{}) (*types.RemoteSource, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.RemoteSource{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal RemoteSource")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanRemoteSources(rows *stdsql.Rows) ([]*types.RemoteSource, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.RemoteSource{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanRemoteSource(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchUsers(tx *sql.Tx, q sq.Sqlizer) ([]*types.User, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanUsers(rows)
}

func (d *DB) scanUser(rows *stdsql.Rows, additionalFields []interface{}) (*types.User, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.User{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal User")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanUsers(rows *stdsql.Rows) ([]*types.User, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.User{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanUser(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchUserTokens(tx *sql.Tx, q sq.Sqlizer) ([]*types.UserToken, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanUserTokens(rows)
}

func (d *DB) scanUserToken(rows *stdsql.Rows, additionalFields []interface{}) (*types.UserToken, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.UserToken{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal UserToken")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanUserTokens(rows *stdsql.Rows) ([]*types.UserToken, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.UserToken{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanUserToken(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchLinkedAccounts(tx *sql.Tx, q sq.Sqlizer) ([]*types.LinkedAccount, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanLinkedAccounts(rows)
}

func (d *DB) scanLinkedAccount(rows *stdsql.Rows, additionalFields []interface{}) (*types.LinkedAccount, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.LinkedAccount{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal LinkedAccount")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanLinkedAccounts(rows *stdsql.Rows) ([]*types.LinkedAccount, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.LinkedAccount{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanLinkedAccount(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchOrganizations(tx *sql.Tx, q sq.Sqlizer) ([]*types.Organization, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanOrganizations(rows)
}

func (d *DB) scanOrganization(rows *stdsql.Rows, additionalFields []interface{}) (*types.Organization, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.Organization{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal Organization")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanOrganizations(rows *stdsql.Rows) ([]*types.Organization, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.Organization{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanOrganization(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchOrganizationMembers(tx *sql.Tx, q sq.Sqlizer) ([]*types.OrganizationMember, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanOrganizationMembers(rows)
}

func (d *DB) scanOrganizationMember(rows *stdsql.Rows, additionalFields []interface{}) (*types.OrganizationMember, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.OrganizationMember{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal OrganizationMember")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanOrganizationMembers(rows *stdsql.Rows) ([]*types.OrganizationMember, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.OrganizationMember{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanOrganizationMember(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchProjectGroups(tx *sql.Tx, q sq.Sqlizer) ([]*types.ProjectGroup, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanProjectGroups(rows)
}

func (d *DB) scanProjectGroup(rows *stdsql.Rows, additionalFields []interface{}) (*types.ProjectGroup, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.ProjectGroup{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal ProjectGroup")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanProjectGroups(rows *stdsql.Rows) ([]*types.ProjectGroup, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.ProjectGroup{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanProjectGroup(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchProjects(tx *sql.Tx, q sq.Sqlizer) ([]*types.Project, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanProjects(rows)
}

func (d *DB) scanProject(rows *stdsql.Rows, additionalFields []interface{}) (*types.Project, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.Project{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal Project")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanProjects(rows *stdsql.Rows) ([]*types.Project, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.Project{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanProject(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchSecrets(tx *sql.Tx, q sq.Sqlizer) ([]*types.Secret, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanSecrets(rows)
}

func (d *DB) scanSecret(rows *stdsql.Rows, additionalFields []interface{}) (*types.Secret, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.Secret{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal Secret")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanSecrets(rows *stdsql.Rows) ([]*types.Secret, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.Secret{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanSecret(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}

func (d *DB) fetchVariables(tx *sql.Tx, q sq.Sqlizer) ([]*types.Variable, []string, error) {
	rows, err := d.query(tx, q)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	defer rows.Close()

	return d.scanVariables(rows)
}

func (d *DB) scanVariable(rows *stdsql.Rows, additionalFields []interface{}) (*types.Variable, string, error) {
	var id string
	var revision uint64
	var data []byte
	fields := append([]interface{}{&id, &revision, &data}, additionalFields...)
	if err := rows.Scan(fields...); err != nil {
		return nil, "", errors.Wrap(err, "failed to scan rows")
	}
	v := types.Variable{}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, "", errors.Wrap(err, "failed to unmarshal Variable")
		}
	}

	v.Revision = revision

	return &v, id, nil
}

func (d *DB) scanVariables(rows *stdsql.Rows) ([]*types.Variable, []string, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	fieldsNumber := len(cols)
	if fieldsNumber < 3 {
		return nil, nil, errors.Errorf("not enough columns (%d < 3)", len(cols))
	}
	var additionalFieldsPtr []interface{}
	if fieldsNumber > 3 {
		additionalFieldsNumber := fieldsNumber - 3
		additionalFields := make([]interface{}, additionalFieldsNumber)
		additionalFieldsPtr = make([]interface{}, additionalFieldsNumber)
		for i := 0; i < additionalFieldsNumber; i++ {
			additionalFieldsPtr[i] = &additionalFields[i]
		}
	}

	vs := []*types.Variable{}
	ids := []string{}
	for rows.Next() {
		v, id, err := d.scanVariable(rows, additionalFieldsPtr)
		if err != nil {
			rows.Close()
			return nil, nil, errors.WithStack(err)
		}
		vs = append(vs, v)
		ids = append(ids, id)
	}
	if err := rows.Err(); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return vs, ids, nil
}