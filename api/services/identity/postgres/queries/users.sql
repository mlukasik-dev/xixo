-- name: FindUsers :many
SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
  FROM users_with_roles
    WHERE account_id = $1 ORDER BY created_at DESC, user_id DESC LIMIT $2;


-- name: FindUsersCursor :many
SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
  FROM users_with_roles
    WHERE account_id = $1 AND (
      created_at <= $2 AND (
        created_at < $2 OR user_id < $3
      )
    )
    ORDER BY created_at DESC, user_id DESC LIMIT $4;


-- name: FindUserByID :one
SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
  FROM users_with_roles WHERE account_id = $1 AND user_id = $2;


-- name: CreateUser :one
INSERT INTO
  users_with_roles(account_id, first_name, last_name, email, phone_number, roles)
VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles;


-- name: UpdateUser :one
UPDATE users_with_roles SET
	first_name = CASE WHEN $1 THEN $2 ELSE first_name END,
	last_name = CASE WHEN $3 THEN $4 ELSE last_name END,
	email = CASE WHEN $5 THEN $6 ELSE email END,
	phone_number = CASE WHEN $7 THEN $8 ELSE phone_number END,
	roles = CASE WHEN $9 THEN $10 ELSE roles END
WHERE user_id = $11
  RETURNING user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles;


-- name: DeleteUser :exec
DELETE FROM users WHERE account_id = $1 AND user_id = $2;


-- name: GrantRoleToUser :exec
INSERT INTO users_roles(user_id, role_id) VALUES ($1, $2);


-- name: GrantAccountAdminRoleToUser :exec
INSERT INTO users_roles(user_id, role_id) VALUES ($1, (
  SELECT role_id FROM roles WHERE account_admin
)) ON CONFLICT DO NOTHING;


-- name: DenyRoleFromUser :exec
DELETE FROM users_roles WHERE user_id = $1 AND role_id = $2;