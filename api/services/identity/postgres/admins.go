package postgres

import (
	"context"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain/admins"

	"github.com/google/uuid"
)

// verify interface compliance
var _ admins.Repository = (*Repository)(nil)

// FindAdmins .
func (r *Repository) FindAdmins(ctx context.Context, cursor *cursor.Cursor, limit int32) ([]admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

// FindAdminByID .
func (r *Repository) FindAdminByID(ctx context.Context, adminID uuid.UUID) (*admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

// CreateAdmin .
func (r *Repository) CreateAdmin(ctx context.Context, input *admins.Admin) (*admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

// UpdateAdmin .
func (r *Repository) UpdateAdmin(ctx context.Context, adminID uuid.UUID, mask *admins.UpdateMask, input *admins.Admin) (*admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

// DeleteAdmin .
func (r *Repository) DeleteAdmin(ctx context.Context, adminID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

// CountAdmins .
func (r *Repository) CountAdmins(ctx context.Context) (int32, error) {
	panic("not implemented") // TODO: Implement
}
