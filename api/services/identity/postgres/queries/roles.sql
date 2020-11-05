-- name: FindRoles :many
SELECT
  role_id, admin_only, display_name, description, permissions, created_at, updated_at
    FROM roles_with_permissions WHERE
      admin_only = COALESCE($1, admin_only)
    ORDER BY created_at DESC, role_id DESC LIMIT $2;


-- name: FindRolesCursor :many
SELECT
  role_id, admin_only, display_name, description, permissions, created_at, updated_at
    FROM roles_with_permissions WHERE
      created_at <= $1 AND (
        created_at < $1 OR role_id < $2
      ) AND
      admin_only = COALESCE($3, admin_only)
    ORDER BY created_at DESC, role_id DESC LIMIT $4;


-- name: FindRoleByID :one
SELECT
  role_id, admin_only, display_name, description, permissions, created_at, updated_at
FROM roles_with_permissions WHERE role_id = $1;


-- name: FindRolePermissions :one
SELECT
  permissions
FROM roles_with_permissions WHERE role_id = $1;


-- name: CreateRole :one
INSERT INTO
  roles_with_permissions(admin_only, display_name, description, permissions)
VALUES ($1, $2, $3, $4)
  RETURNING role_id, admin_only, display_name, description, permissions, created_at, updated_at;


-- name: UpdateRole :one
UPDATE roles_with_permissions SET
	admin_only = CASE WHEN $1 THEN $2 ELSE admin_only END,
	display_name = CASE WHEN $3 THEN $4 ELSE display_name END,
	description = CASE WHEN $5 THEN $6 ELSE description END,
	permissions = CASE WHEN $7 THEN $8 ELSE permissions END
WHERE role_id = $9
  RETURNING role_id, admin_only, display_name, description, permissions, created_at, updated_at;


-- name: DeleteRole :exec
DELETE FROM roles WHERE role_id = $1;


-- name: GrantPermission :exec
INSERT INTO permissions(role_id, method)
  VALUES ($1, $2) ON CONFLICT DO NOTHING;


-- name: DenyPermission :exec
DELETE FROM permissions WHERE role_id = $1 AND method = $2;
