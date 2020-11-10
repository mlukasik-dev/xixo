package users

import (
	"context"

	"go.xixo.com/api/pkg/cursor"

	"github.com/google/uuid"
)

// UpdateMask .
type UpdateMask struct {
	FirstName   bool
	LastName    bool
	PhoneNumber bool
	Email       bool
	Roles       bool
}

// Repository user's repository.
type Repository interface {
	FindUsers(ctx context.Context, accountID uuid.UUID, cursor *cursor.Cursor, limit int32) ([]User, error)
	FindUserByID(ctx context.Context, accountID, userID uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, accountID uuid.UUID, input *User) (*User, error)
	CreateAccountAdmin(ctx context.Context, accountID uuid.UUID, input *User) (*User, error)
	UpdateUser(ctx context.Context, accountID, userID uuid.UUID, mask *UpdateMask, input *User) (*User, error)
	DeleteUser(ctx context.Context, accountID, userID uuid.UUID) error
	CountUsers(ctx context.Context, accountID uuid.UUID) (int32, error)
}
