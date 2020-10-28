package postgres

import (
	"context"
	"database/sql"
	"errors"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/account/domain/accounts"
)

// verify interface compliance
var _ accounts.Repository = (*repo)(nil)

func (r *repo) FindAccounts(cursor *cursor.Cursor, limit int32) ([]*accounts.Account, error) {
	var accountsSlice []*accounts.Account
	var err error
	// first request
	if cursor == nil {
		const q = `
			SELECT account_id, display_name, created_at, updated_at
				FROM accounts ORDER BY created_at DESC, account_id DESC
					LIMIT $1;
		`
		err = r.db.SelectContext(context.Background(), &accountsSlice, q, limit)
	} else {
		const q = `
			SELECT account_id, display_name, created_at, updated_at FROM accounts
				WHERE created_at <= $1
					AND ( created_at < $1 OR account_id < $2 )
				ORDER BY created_at DESC,
								 account_id DESC
					LIMIT $3;
		`
		err = r.db.SelectContext(context.Background(), &accountsSlice, q, cursor.Timestamp, cursor.UUID, limit)
	}
	if errors.Is(err, sql.ErrNoRows) {
		return []*accounts.Account{}, nil
	}
	if err != nil {
		// return unknown error to the caller
		return nil, err
	}

	return accountsSlice, nil
}

func (r *repo) FindAccountByID(id string) (*accounts.Account, error) {
	const query = `
		SELECT account_id, display_name, created_at, updated_at
			FROM accounts WHERE account_id = $1
	`

	var account accounts.Account
	err := r.db.GetContext(context.Background(), &account, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, accounts.ErrNotFound
	}
	if err != nil {
		// return unknown error to the caller
		return nil, err
	}

	return &account, nil
}

func (r *repo) CreateAccount(input *accounts.CreateAccountInput) (*accounts.Account, error) {
	const query = `
		INSERT INTO accounts(display_name)
			VALUES ($1) RETURNING account_id, display_name, created_at, updated_at
	`

	var account accounts.Account
	err := r.db.GetContext(context.Background(), &account, query, input.DisplayName)
	if err != nil {
		// return unknown error to the caller
		return nil, err
	}

	return &account, nil
}

func (r *repo) UpdateAccount(id string, mask *accounts.UpdateMask, input *accounts.UpdateAccountInput) (*accounts.Account, error) {
	const query = `
		UPDATE accounts SET
			display_name = COALESCE($1, display_name)
		WHERE account_id = $2
			RETURNING account_id, display_name, created_at, updated_at
	`

	var displayName sql.NullString
	if mask.DisplayName {
		displayName = sql.NullString{String: input.DisplayName, Valid: true}
	}
	var account accounts.Account
	err := r.db.GetContext(context.Background(), &account, query,
		displayName, id,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, accounts.ErrNotFound
	}
	if err != nil {
		// return unknown error to the caller
		return nil, err
	}

	return &account, nil
}

func (r *repo) DeleteAccount(id string) error {
	const query = `
		DELETE FROM accounts WHERE account_id = $1
	`
	_, err := r.db.ExecContext(context.Background(), query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return accounts.ErrNotFound
	}
	if err != nil {
		// return unknown error to the caller
		return err
	}
	return nil
}

func (r *repo) Count() (int32, error) {
	const query = `
		SELECT COUNT(*) FROM accounts
	`
	var count int32
	err := r.db.QueryRowContext(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
