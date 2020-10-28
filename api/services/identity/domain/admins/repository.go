package admins

import (
	"go.xixo.com/api/pkg/cursor"

	"github.com/go-playground/validator/v10"
)

// UpdateMask .
type UpdateMask struct {
	FirstName bool
	LastName  bool
	Email     bool
	RoleIDs   bool
}

// Repository Admin's repository
type Repository interface {
	FindAdmins(cursor *cursor.Cursor, limit int32) ([]*Admin, error)
	FindAdminByID(id string) (*Admin, error)
	CreateAdmin(input *CreateAdminInput) (*Admin, error)
	UpdateAdmin(id string, mask *UpdateMask, input *UpdateAdminInput) (*Admin, error)
	DeleteAdmin(id string) error
	CountAdmins() (int32, error)
}

// CreateAdminInput .
type CreateAdminInput struct {
	FirstName string
	LastName  string
	Email     string
	RoleIDs   []string `validate:"dive,uuid4,omitempty"`
}

// Validate .
func (i *CreateAdminInput) Validate(v *validator.Validate) error {
	return v.Struct(i)
}

// UpdateAdminInput .
type UpdateAdminInput struct {
	FirstName string
	LastName  string
	Email     string
	RoleIDs   []string `validate:"dive,uuid4,omitempty"`
}

// Validate .
func (i *UpdateAdminInput) Validate(v *validator.Validate) error {
	return v.Struct(i)
}
