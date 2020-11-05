package postgres

import (
	"github.com/google/uuid"
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/account/domain/accounts"
)

var _ accounts.Repository = (*repo)(nil)

func (r *repo) FindAccounts(cursor *cursor.Cursor, limit int32) ([]*accounts.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) FindAccountByID(accoundID uuid.UUID) (*accounts.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CreateAccount(input *accounts.CreateAccountInput) (*accounts.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) UpdateAccount(accoundID uuid.UUID, mask *accounts.UpdateMask, input *accounts.UpdateAccountInput) (*accounts.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) DeleteAccount(accoundID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (r *repo) Count() (int32, error) {
	panic("not implemented") // TODO: Implement
}
