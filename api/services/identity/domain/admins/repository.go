package admins

import (
	"go.xixo.com/api/pkg/cursor"

	"github.com/google/uuid"
)

// UpdateMask .
type UpdateMask struct {
	FirstName bool
	LastName  bool
	Email     bool
	Roles     bool
}

// Repository Admin's repository
type Repository interface {
	FindAdmins(cursor *cursor.Cursor, limit int32) ([]*Admin, error)
	FindAdminByID(adminID uuid.UUID) (*Admin, error)
	CreateAdmin(input *Admin) (*Admin, error)
	UpdateAdmin(adminID uuid.UUID, mask *UpdateMask, input *Admin) (*Admin, error)
	DeleteAdmin(adminID uuid.UUID) error
	CountAdmins() (int32, error)
}
