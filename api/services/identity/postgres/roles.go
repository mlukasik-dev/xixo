package postgres

import (
	"context"

	"github.com/google/uuid"
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain/roles"
)

// verify interface compliance
var _ roles.Repository = (*Repository)(nil)

// FindRoles .
func (r *Repository) FindRoles(ctx context.Context, cursor *cursor.Cursor, limit int32, filter *roles.Filter) ([]roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

// FindRoleByID .
func (r *Repository) FindRoleByID(ctx context.Context, roleID uuid.UUID) (*roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

// CreateRole ..
func (r *Repository) CreateRole(ctx context.Context, input *roles.Role) (*roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

// UpdateRole .
func (r *Repository) UpdateRole(ctx context.Context, roleID uuid.UUID, mask *roles.UpdateMask, input *roles.Role) (*roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

// DeleteRole .
func (r *Repository) DeleteRole(ctx context.Context, roleID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

// CountRoles .
func (r *Repository) CountRoles(ctx context.Context) (int32, error) {
	panic("not implemented") // TODO: Implement
}
