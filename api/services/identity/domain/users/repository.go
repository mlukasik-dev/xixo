package users

import (
	"go.xixo.com/api/pkg/cursor"

	"github.com/go-playground/validator/v10"
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

// Repository user's repository
type Repository interface {
	FindUsers(accountID uuid.UUID, cursor *cursor.Cursor, limit int32) ([]*User, error)
	FindUserByID(accountID, userID uuid.UUID) (*User, error)
	CreateUser(accountID uuid.UUID, input *CreateUserInput) (*User, error)
	CreateAccountAdmin(accountID uuid.UUID, input *CreateUserInput) (*User, error)
	UpdateUser(accountID, userID uuid.UUID, mask *UpdateMask, input *UpdateUserInput) (*User, error)
	DeleteUser(accountID, userID uuid.UUID) error
	CountUsers(accountID uuid.UUID) (int32, error)
}

// CreateUserInput .
type CreateUserInput struct {
	FirstName   string   `validate:"omitempty"`
	LastName    string   `validate:"omitempty"`
	PhoneNumber string   `validate:"omitempty"`
	Email       string   `validate:"required,email"`
	Roles       []string `validate:"dive,uuid4,omitempty"`
}

// Validate .
func (i *CreateUserInput) Validate(v *validator.Validate) error {
	return v.Struct(i)
}

// UpdateUserInput .
type UpdateUserInput struct {
	FirstName   string   `validate:"omitempty"`
	LastName    string   `validate:"omitempty"`
	PhoneNumber string   `validate:"omitempty"`
	Email       string   `validate:"omitempty,email"`
	Roles       []string `validate:"dive,uuid4,omitempty"`
}

// Validate .
func (i *UpdateUserInput) Validate(v *validator.Validate) error {
	return v.Struct(i)
}
