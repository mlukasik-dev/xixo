-- name: FindAccounts :many
SELECT account_id, display_name, created_at, updated_at
	FROM accounts ORDER BY created_at DESC, account_id DESC
		LIMIT $1;

-- name: FindAccountsCursor :many
SELECT account_id, display_name, created_at, updated_at FROM accounts
  WHERE created_at <= $1
    AND ( created_at < $1 OR account_id < $2 )
  ORDER BY created_at DESC,
            account_id DESC
    LIMIT $3;

-- name: FindAccountByID :one
SELECT account_id, display_name, created_at, updated_at
  FROM accounts WHERE account_id = $1;

-- name: CreateAccount :one
INSERT INTO accounts(display_name)
  VALUES ($1) RETURNING account_id, display_name, created_at, updated_at;

-- name: UpdateAccount :one
UPDATE accounts SET
	display_name = CASE WHEN $1 THEN $2 ELSE display_name END
WHERE account_id = $3
  RETURNING account_id, display_name, created_at, updated_at;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE account_id = $1;

-- name: Count :one
SELECT COUNT(*) FROM accounts;