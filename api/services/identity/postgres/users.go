package postgres

import (
	"context"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain/users"

	"github.com/google/uuid"
)

// verify interface compliance
var _ users.Repository = (*Repository)(nil)

// FindUsers .
func (r *Repository) FindUsers(ctx context.Context, accountID uuid.UUID, cursor *cursor.Cursor, limit int32) ([]users.User, error) {
	panic("not implemented") // TODO: Implement
}

// FindUserByID .
func (r *Repository) FindUserByID(ctx context.Context, accountID uuid.UUID, userID uuid.UUID) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

// CreateUser .
func (r *Repository) CreateUser(ctx context.Context, accountID uuid.UUID, input *users.User) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

// CreateAccountAdmin .
func (r *Repository) CreateAccountAdmin(ctx context.Context, accountID uuid.UUID, input *users.User) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

// UpdateUser .
func (r *Repository) UpdateUser(ctx context.Context, accountID uuid.UUID, userID uuid.UUID, mask *users.UpdateMask, input *users.User) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

// DeleteUser .
func (r *Repository) DeleteUser(ctx context.Context, accountID uuid.UUID, userID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

// CountUsers .
func (r *Repository) CountUsers(ctx context.Context, accountID uuid.UUID) (int32, error) {
	panic("not implemented") // TODO: Implement
}
