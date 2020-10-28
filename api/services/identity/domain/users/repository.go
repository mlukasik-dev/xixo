package users

import (
	"go.xixo.com/api/pkg/cursor"

	"github.com/go-playground/validator/v10"
)

// UpdateMask .
type UpdateMask struct {
	FirstName   bool
	LastName    bool
	PhoneNumber bool
	Email       bool
	RoleIDs     bool
}

// Repository user's repository
type Repository interface {
	FindUsers(accountID string, cursor *cursor.Cursor, limit int32) ([]*User, error)
	FindUserByID(accountID, userID string) (*User, error)
	CreateUser(accountID string, input *CreateUserInput) (*User, error)
	CreateAccountAdmin(accountID string, input *CreateUserInput) (*User, error)
	UpdateUser(accountID, userID string, mask *UpdateMask, input *UpdateUserInput) (*User, error)
	DeleteUser(accountID, userID string) error
	CountUsers(accountID string) (int32, error)
}

// CreateUserInput .
type CreateUserInput struct {
	FirstName   string   `validate:"omitempty"`
	LastName    string   `validate:"omitempty"`
	PhoneNumber string   `validate:"omitempty"`
	Email       string   `validate:"required,email"`
	RoleIDs     []string `validate:"dive,uuid4,omitempty"`
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
	RoleIDs     []string `validate:"dive,uuid4,omitempty"`
}

// Validate .
func (i *UpdateUserInput) Validate(v *validator.Validate) error {
	return v.Struct(i)
}
