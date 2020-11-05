package postgres

import (
	"github.com/google/uuid"
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain/users"
)

// verify interface compliance
var _ users.Repository = (*repo)(nil)

func (r *repo) FindUsers(accountID uuid.UUID, cursor *cursor.Cursor, limit int32) ([]*users.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) FindUserByID(accountID uuid.UUID, userID uuid.UUID) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CreateUser(accountID uuid.UUID, input *users.CreateUserInput) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CreateAccountAdmin(accountID uuid.UUID, input *users.CreateUserInput) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) UpdateUser(accountID uuid.UUID, userID uuid.UUID, mask *users.UpdateMask, input *users.UpdateUserInput) (*users.User, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) DeleteUser(accountID uuid.UUID, userID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CountUsers(accountID uuid.UUID) (int32, error) {
	panic("not implemented") // TODO: Implement
}
