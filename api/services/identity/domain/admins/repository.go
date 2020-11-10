package admins

import (
	"context"

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

// Repository Admin's repository.
type Repository interface {
	FindAdmins(ctx context.Context, cursor *cursor.Cursor, limit int32) ([]Admin, error)
	FindAdminByID(ctx context.Context, adminID uuid.UUID) (*Admin, error)
	CreateAdmin(ctx context.Context, input *Admin) (*Admin, error)
	UpdateAdmin(ctx context.Context, adminID uuid.UUID, mask *UpdateMask, input *Admin) (*Admin, error)
	DeleteAdmin(ctx context.Context, adminID uuid.UUID) error
	CountAdmins(ctx context.Context) (int32, error)
}
