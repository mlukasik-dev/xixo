// When sqlc supports custom models,
// this file will be generated

package gen

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"go.xixo.com/api/services/identity/domain/admins"
)

const createAdmin = `-- name: CreateAdmin :one
INSERT INTO admins_with_roles(first_name, last_name, email, roles)
  VALUES ($1, $2, $3, $4)	RETURNING admin_id
`

type CreateAdminParams struct {
	FirstName string
	LastName  string
	Email     string
	Roles     []string
}

func (q *Queries) CreateAdmin(ctx context.Context, arg CreateAdminParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createAdmin,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		pq.Array(arg.Roles),
	)
	var admin_id uuid.UUID
	err := row.Scan(&admin_id)
	return admin_id, err
}

const deleteAdmin = `-- name: DeleteAdmin :exec
DELETE FROM admins WHERE admin_id = $1
`

func (q *Queries) DeleteAdmin(ctx context.Context, adminID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAdmin, adminID)
	return err
}

const denyRoleFromAdmin = `-- name: DenyRoleFromAdmin :exec
DELETE FROM admins_roles WHERE admin_id = $1 AND role_id = $2
`

type DenyRoleFromAdminParams struct {
	AdminID uuid.UUID
	RoleID  uuid.UUID
}

func (q *Queries) DenyRoleFromAdmin(ctx context.Context, arg DenyRoleFromAdminParams) error {
	_, err := q.db.ExecContext(ctx, denyRoleFromAdmin, arg.AdminID, arg.RoleID)
	return err
}

const findAdminByID = `-- name: FindAdminByID :one
SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
  FROM admins_with_roles WHERE admin_id = $1
`

type FindAdminByIDRow struct {
	AdminID   uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Roles     []string
}

func (q *Queries) FindAdminByID(ctx context.Context, adminID uuid.UUID) (FindAdminByIDRow, error) {
	row := q.db.QueryRowContext(ctx, findAdminByID, adminID)
	var i FindAdminByIDRow
	err := row.Scan(
		&i.AdminID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Roles),
	)
	return i, err
}

const findAdminRoles = `-- name: FindAdminRoles :one
SELECT roles FROM admins_with_roles WHERE admin_id = $1
`

func (q *Queries) FindAdminRoles(ctx context.Context, adminID uuid.UUID) ([]string, error) {
	row := q.db.QueryRowContext(ctx, findAdminRoles, adminID)
	var roles []string
	err := row.Scan(pq.Array(&roles))
	return roles, err
}

const findAdmins = `-- name: FindAdmins :many
SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
  FROM admins_with_roles
    ORDER BY created_at DESC, admin_id DESC LIMIT $1
`

type FindAdminsRow struct {
	AdminID   uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Roles     []string
}

func (q *Queries) FindAdmins(ctx context.Context, limit int32) ([]FindAdminsRow, error) {
	rows, err := q.db.QueryContext(ctx, findAdmins, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAdminsRow
	for rows.Next() {
		var i FindAdminsRow
		if err := rows.Scan(
			&i.AdminID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.Roles),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAdminsCursor = `-- name: FindAdminsCursor :many
SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
  FROM admins_with_roles WHERE (
    created_at <= $1 AND (
      created_at < $1 OR admin_id < $2
    )
  ) ORDER BY created_at DESC, admin_id DESC LIMIT $3
`

type FindAdminsCursorParams struct {
	CreatedAt time.Time
	AdminID   uuid.UUID
	Limit     int32
}

type FindAdminsCursorRow struct {
	AdminID   uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Roles     []string
}

func (q *Queries) FindAdminsCursor(ctx context.Context, arg FindAdminsCursorParams) ([]FindAdminsCursorRow, error) {
	rows, err := q.db.QueryContext(ctx, findAdminsCursor, arg.CreatedAt, arg.AdminID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAdminsCursorRow
	for rows.Next() {
		var i FindAdminsCursorRow
		if err := rows.Scan(
			&i.AdminID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
			pq.Array(&i.Roles),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const grantRoleToAdmin = `-- name: GrantRoleToAdmin :exec
INSERT INTO admins_roles(admin_id, role_id) VALUES ($1, $2)
`

type GrantRoleToAdminParams struct {
	AdminID uuid.UUID
	RoleID  uuid.UUID
}

func (q *Queries) GrantRoleToAdmin(ctx context.Context, arg GrantRoleToAdminParams) error {
	_, err := q.db.ExecContext(ctx, grantRoleToAdmin, arg.AdminID, arg.RoleID)
	return err
}

const updateAdmin = `-- name: UpdateAdmin :one
UPDATE admins_with_roles SET
	first_name = CASE WHEN $1 THEN $2 ELSE first_name END,
	last_name = CASE WHEN $3 THEN $4 ELSE last_name END,
	email = CASE WHEN $5 THEN $6 ELSE email END,
	roles = CASE WHEN $7 THEN $8 ELSE roles END
WHERE admin_id = $9
  RETURNING admin_id, first_name, last_name, email, created_at, updated_at, roles
`

type UpdateAdminParams struct {
	FirstName             string
	ShouldUpdateFirstName bool
	LastName              string
	ShouldUpdateLastName  bool
	Email                 string
	ShouldUpdateEmail     bool
	Roles                 []string
	ShouldUpdateRoles     bool
	AdminID               uuid.UUID
}

func (q *Queries) UpdateAdmin(ctx context.Context, arg UpdateAdminParams) (admins.Admin, error) {
	row := q.db.QueryRowContext(ctx, updateAdmin,
		arg.FirstName,
		arg.ShouldUpdateFirstName,
		arg.LastName,
		arg.ShouldUpdateLastName,
		arg.Email,
		arg.ShouldUpdateEmail,
		pq.Array(arg.Roles),
		arg.ShouldUpdateRoles,
		arg.AdminID,
	)
	var i admins.Admin
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Roles),
	)
	return i, err
}
