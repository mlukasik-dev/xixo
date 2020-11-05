package accounts

import (
	"go.xixo.com/api/pkg/cursor"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// UpdateMask .
type UpdateMask struct {
	DisplayName bool
}

// Repository account's repository
type Repository interface {
	FindAccounts(cursor *cursor.Cursor, limit int32) ([]*Account, error)
	FindAccountByID(accoundID uuid.UUID) (*Account, error)
	CreateAccount(input *CreateAccountInput) (*Account, error)
	UpdateAccount(accoundID uuid.UUID, mask *UpdateMask, input *UpdateAccountInput) (*Account, error)
	DeleteAccount(accoundID uuid.UUID) error
	Count() (int32, error)
}

// CreateAccountInput .
type CreateAccountInput struct {
	DisplayName string `validate:"required"`
}

// Validate .
func (i *CreateAccountInput) Validate(validate *validator.Validate) error {
	errs := validate.Struct(i)
	return errs
}

// UpdateAccountInput .
type UpdateAccountInput struct {
	DisplayName string `validate:"omitempty"`
}

// Validate .
func (i *UpdateAccountInput) Validate(validate *validator.Validate) error {
	return nil
}
