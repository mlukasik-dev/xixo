package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/account/postgres/gen"

	"github.com/google/uuid"
)

var _ accounts.Repository = (*Repository)(nil)

// FindAccounts .
func (r *Repository) FindAccounts(ctx context.Context, cursor *cursor.Cursor, limit int32) ([]accounts.Account, error) {
	var items []accounts.Account
	var err error
	if cursor == nil {
		// first request
		items, err = r.q.FindAccounts(ctx, limit)
	} else {
		items, err = r.q.FindAccountsCursor(ctx, gen.FindAccountsCursorParams{
			Limit:     limit,
			AccountID: cursor.UUID,
			CreatedAt: cursor.Timestamp,
		})
	}
	return items, err
}

// FindAccountByID .
func (r *Repository) FindAccountByID(ctx context.Context, accountID uuid.UUID) (*accounts.Account, error) {
	account, err := r.q.FindAccountByID(ctx, accountID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("account %w", err)
	}
	if err != nil {
		return nil, err
	}
	return &account, err
}

// CreateAccount .
func (r *Repository) CreateAccount(ctx context.Context, input *accounts.Account) (*accounts.Account, error) {
	account, err := r.q.CreateAccount(ctx, input.DisplayName)
	return &account, err
}

// UpdateAccount .
func (r *Repository) UpdateAccount(ctx context.Context, accountID uuid.UUID, mask *accounts.UpdateMask, input *accounts.Account) (*accounts.Account, error) {
	account, err := r.q.UpdateAccount(ctx, gen.UpdateAccountParams{
		AccountID:               accountID,
		DisplayName:             input.DisplayName,
		ShouldUpdateDisplayName: mask.DisplayName,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("account %w", err)
	}
	if err != nil {
		return nil, err
	}
	return &account, err
}

// DeleteAccount .
func (r *Repository) DeleteAccount(ctx context.Context, accountID uuid.UUID) error {
	err := r.q.DeleteAccount(ctx, accountID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("account %w", err)
	}
	return err
}

// CountAccounts .
func (r *Repository) CountAccounts(ctx context.Context) (int32, error) {
	count, err := r.q.CountAccounts(ctx)
	return int32(count), err
}
