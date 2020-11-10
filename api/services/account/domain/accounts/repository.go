package accounts

import (
	"context"

	"go.xixo.com/api/pkg/cursor"

	"github.com/google/uuid"
)

// UpdateMask .
type UpdateMask struct {
	DisplayName bool
}

// Repository account's repository.
type Repository interface {
	FindAccounts(ctx context.Context, cursor *cursor.Cursor, limit int32) ([]Account, error)
	FindAccountByID(ctx context.Context, accountID uuid.UUID) (*Account, error)
	CreateAccount(ctx context.Context, input *Account) (*Account, error)
	UpdateAccount(ctx context.Context, accountID uuid.UUID, mask *UpdateMask, input *Account) (*Account, error)
	DeleteAccount(ctx context.Context, accountID uuid.UUID) error
	CountAccounts(ctx context.Context) (int32, error)
}
