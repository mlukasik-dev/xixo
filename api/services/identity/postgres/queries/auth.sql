-- name: FindUsersPassword :one
SELECT password FROM users WHERE account_id = $1 AND email = $2;


-- name: UpdateUserPassword :exec
UPDATE users SET password = $1 WHERE account_id = $2 AND email = $3;


-- name: FindAdminPassword :one
SELECT password FROM admins WHERE email = $1;


-- name: UpdateAdminPassword :exec
UPDATE admins SET password = $1 WHERE email = $2;


-- name: FindUserInfoByEmail :one
SELECT user_id, registered, roles
  FROM users_with_roles
    WHERE account_id = $1 AND email = $2;


-- name: FindAdminInfoByEmail :one
SELECT admin_id, registered, roles
  FROM admins_with_roles WHERE email = $1;