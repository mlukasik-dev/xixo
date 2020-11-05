// When sqlc supports custom models,
// this file will be generated

package gen

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"go.xixo.com/api/services/identity/domain/roles"
)

const createRole = `-- name: CreateRole :one
INSERT INTO
  roles_with_permissions(admin_only, display_name, description, permissions)
VALUES ($1, $2, $3, $4)
  RETURNING role_id, admin_only, display_name, description, permissions, created_at, updated_at
`

type CreateRoleParams struct {
	AdminOnly   bool
	DisplayName string
	Description string
	Permissions []string
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (roles.Role, error) {
	row := q.db.QueryRowContext(ctx, createRole,
		arg.AdminOnly,
		arg.DisplayName,
		arg.Description,
		pq.Array(arg.Permissions),
	)
	var i roles.Role
	err := row.Scan(
		&i.ID,
		&i.AdminOnly,
		&i.DisplayName,
		&i.Description,
		pq.Array(&i.Permissions),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles WHERE role_id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, roleID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteRole, roleID)
	return err
}

const denyPermission = `-- name: DenyPermission :exec
DELETE FROM permissions WHERE role_id = $1 AND method = $2
`

type DenyPermissionParams struct {
	RoleID uuid.UUID
	Method string
}

func (q *Queries) DenyPermission(ctx context.Context, arg DenyPermissionParams) error {
	_, err := q.db.ExecContext(ctx, denyPermission, arg.RoleID, arg.Method)
	return err
}

const findRoleByID = `-- name: FindRoleByID :one
SELECT
  role_id, admin_only, display_name, description, permissions, created_at, updated_at
FROM roles_with_permissions WHERE role_id = $1
`

func (q *Queries) FindRoleByID(ctx context.Context, roleID uuid.UUID) (roles.Role, error) {
	row := q.db.QueryRowContext(ctx, findRoleByID, roleID)
	var i roles.Role
	err := row.Scan(
		&i.ID,
		&i.AdminOnly,
		&i.DisplayName,
		&i.Description,
		pq.Array(&i.Permissions),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findRolePermissions = `-- name: FindRolePermissions :one
SELECT
  permissions
FROM roles_with_permissions WHERE role_id = $1
`

func (q *Queries) FindRolePermissions(ctx context.Context, roleID uuid.UUID) ([]string, error) {
	row := q.db.QueryRowContext(ctx, findRolePermissions, roleID)
	var permissions []string
	err := row.Scan(pq.Array(&permissions))
	return permissions, err
}

const findRoles = `-- name: FindRoles :many
SELECT
  role_id, admin_only, display_name, description, permissions, created_at, updated_at
    FROM roles_with_permissions WHERE
      admin_only = COALESCE($1, admin_only)
    ORDER BY created_at DESC, role_id DESC LIMIT $2
`

type FindRolesParams struct {
	AdminOnly bool
	Limit     int32
}

func (q *Queries) FindRoles(ctx context.Context, arg FindRolesParams) ([]roles.Role, error) {
	rows, err := q.db.QueryContext(ctx, findRoles, arg.AdminOnly, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []roles.Role
	for rows.Next() {
		var i roles.Role
		if err := rows.Scan(
			&i.ID,
			&i.AdminOnly,
			&i.DisplayName,
			&i.Description,
			pq.Array(&i.Permissions),
			&i.CreatedAt,
			&i.UpdatedAt,
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

const findRolesCursor = `-- name: FindRolesCursor :many
SELECT
  role_id, admin_only, display_name, description, permissions, created_at, updated_at
    FROM roles_with_permissions WHERE
      created_at <= $1 AND (
        created_at < $1 OR role_id < $2
      ) AND
      admin_only = COALESCE($3, admin_only)
    ORDER BY created_at DESC, role_id DESC LIMIT $4
`

type FindRolesCursorParams struct {
	CreatedAt time.Time
	RoleID    uuid.UUID
	AdminOnly bool
	Limit     int32
}

func (q *Queries) FindRolesCursor(ctx context.Context, arg FindRolesCursorParams) ([]roles.Role, error) {
	rows, err := q.db.QueryContext(ctx, findRolesCursor,
		arg.CreatedAt,
		arg.RoleID,
		arg.AdminOnly,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []roles.Role
	for rows.Next() {
		var i roles.Role
		if err := rows.Scan(
			&i.ID,
			&i.AdminOnly,
			&i.DisplayName,
			&i.Description,
			pq.Array(&i.Permissions),
			&i.CreatedAt,
			&i.UpdatedAt,
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

const grantPermission = `-- name: GrantPermission :exec
INSERT INTO permissions(role_id, method)
  VALUES ($1, $2) ON CONFLICT DO NOTHING
`

type GrantPermissionParams struct {
	RoleID uuid.UUID
	Method string
}

func (q *Queries) GrantPermission(ctx context.Context, arg GrantPermissionParams) error {
	_, err := q.db.ExecContext(ctx, grantPermission, arg.RoleID, arg.Method)
	return err
}

const updateRole = `-- name: UpdateRole :one
UPDATE roles_with_permissions SET
	admin_only = CASE WHEN $1 THEN $2 ELSE admin_only END,
	display_name = CASE WHEN $3 THEN $4 ELSE display_name END,
	description = CASE WHEN $5 THEN $6 ELSE description END,
	permissions = CASE WHEN $7 THEN $8 ELSE permissions END
WHERE role_id = $9
  RETURNING role_id, admin_only, display_name, description, permissions, created_at, updated_at
`

type UpdateRoleParams struct {
	AdminOnly               bool
	ShouldUpdateAdminOnly   bool
	DisplayName             string
	ShouldUpdateDisplayName bool
	Description             string
	ShouldUpdateDescription bool
	Permissions             []string
	ShouldUpdatePermissions bool
	RoleID                  uuid.UUID
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) (roles.Role, error) {
	row := q.db.QueryRowContext(ctx, updateRole,
		arg.AdminOnly,
		arg.ShouldUpdateAdminOnly,
		arg.DisplayName,
		arg.ShouldUpdateDisplayName,
		arg.Description,
		arg.ShouldUpdateDescription,
		pq.Array(arg.Permissions),
		arg.ShouldUpdatePermissions,
		arg.RoleID,
	)
	var i roles.Role
	err := row.Scan(
		&i.ID,
		&i.AdminOnly,
		&i.DisplayName,
		&i.Description,
		pq.Array(&i.Permissions),
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
