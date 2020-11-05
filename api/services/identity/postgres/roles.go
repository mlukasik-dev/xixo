package postgres

import (
	"github.com/google/uuid"
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain/roles"
)

// verify interface compliance
var _ roles.Repository = (*repo)(nil)

func (r *repo) FindRoles(cursor *cursor.Cursor, limit int32, filter *roles.Filter) ([]*roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) FindRoleByID(roleID uuid.UUID) (*roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CreateRole(input *roles.CreateRoleInput) (*roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) UpdateRole(roleID uuid.UUID, mask *roles.UpdateMask, input *roles.UpdateRoleInput) (*roles.Role, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) DeleteRole(roleID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CountRoles() (int32, error) {
	panic("not implemented") // TODO: Implement
}
