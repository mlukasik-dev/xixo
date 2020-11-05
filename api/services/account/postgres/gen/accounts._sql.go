// When sqlc supports custom models,
// this file will be generated

package gen

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.xixo.com/api/services/account/domain/accounts"
)

const count = `-- name: Count :one
SELECT COUNT(*) FROM accounts
`

func (q *Queries) Count(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, count)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts(display_name)
  VALUES ($1) RETURNING account_id, display_name, created_at, updated_at
`

func (q *Queries) CreateAccount(ctx context.Context, displayName string) (accounts.Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, displayName)
	var i accounts.Account
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE account_id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, accountID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, accountID)
	return err
}

const findAccountByID = `-- name: FindAccountByID :one
SELECT account_id, display_name, created_at, updated_at
  FROM accounts WHERE account_id = $1
`

func (q *Queries) FindAccountByID(ctx context.Context, accountID uuid.UUID) (accounts.Account, error) {
	row := q.db.QueryRowContext(ctx, findAccountByID, accountID)
	var i accounts.Account
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findAccounts = `-- name: FindAccounts :many
SELECT account_id, display_name, created_at, updated_at
	FROM accounts ORDER BY created_at DESC, account_id DESC
		LIMIT $1
`

func (q *Queries) FindAccounts(ctx context.Context, limit int32) ([]accounts.Account, error) {
	rows, err := q.db.QueryContext(ctx, findAccounts, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []accounts.Account
	for rows.Next() {
		var i accounts.Account
		if err := rows.Scan(
			&i.ID,
			&i.DisplayName,
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

const findAccountsCursor = `-- name: FindAccountsCursor :many
SELECT account_id, display_name, created_at, updated_at FROM accounts
  WHERE created_at <= $1
    AND ( created_at < $1 OR account_id < $2 )
  ORDER BY created_at DESC,
            account_id DESC
    LIMIT $3
`

type FindAccountsCursorParams struct {
	CreatedAt time.Time
	AccountID uuid.UUID
	Limit     int32
}

func (q *Queries) FindAccountsCursor(ctx context.Context, arg FindAccountsCursorParams) ([]accounts.Account, error) {
	rows, err := q.db.QueryContext(ctx, findAccountsCursor, arg.CreatedAt, arg.AccountID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []accounts.Account
	for rows.Next() {
		var i accounts.Account
		if err := rows.Scan(
			&i.ID,
			&i.DisplayName,
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

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts SET
	display_name = CASE WHEN $1 THEN $2 ELSE display_name END
WHERE account_id = $3
  RETURNING account_id, display_name, created_at, updated_at
`

type UpdateAccountParams struct {
	DisplayName             string
	ShouldUpdateDisplayName string
	AccountID               uuid.UUID
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (accounts.Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.DisplayName, arg.ShouldUpdateDisplayName, arg.AccountID)
	var i accounts.Account
	err := row.Scan(
		&i.ID,
		&i.DisplayName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
