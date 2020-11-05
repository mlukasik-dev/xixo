-- name: FindAdmins :many
SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
  FROM admins_with_roles
    ORDER BY created_at DESC, admin_id DESC LIMIT $1;


-- name: FindAdminsCursor :many
SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
  FROM admins_with_roles WHERE (
    created_at <= $1 AND (
      created_at < $1 OR admin_id < $2
    )
  ) ORDER BY created_at DESC, admin_id DESC LIMIT $3;


-- name: FindAdminByID :one
SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
  FROM admins_with_roles WHERE admin_id = $1;


-- name: FindAdminRoles :one
SELECT roles FROM admins_with_roles WHERE admin_id = $1;


-- name: CreateAdmin :one
INSERT INTO
  admins_with_roles(first_name, last_name, email, roles)
VALUES ($1, $2, $3, $4)
  RETURNING admin_id, first_name, last_name, email, created_at, updated_at, roles;


-- name: UpdateAdmin :one
UPDATE admins_with_roles SET
	first_name = CASE WHEN $1 THEN $2 ELSE first_name END,
	last_name = CASE WHEN $3 THEN $4 ELSE last_name END,
	email = CASE WHEN $5 THEN $6 ELSE email END,
	roles = CASE WHEN $7 THEN $8 ELSE roles END
WHERE admin_id = $9
  RETURNING admin_id, first_name, last_name, email, created_at, updated_at, roles;

-- name: DeleteAdmin :exec
DELETE FROM admins WHERE admin_id = $1;


-- name: GrantRoleToAdmin :exec
INSERT INTO admins_roles(admin_id, role_id) VALUES ($1, $2);


-- name: DenyRoleFromAdmin :exec
DELETE FROM admins_roles WHERE admin_id = $1 AND role_id = $2;