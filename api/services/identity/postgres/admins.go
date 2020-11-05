package postgres

import (
	"github.com/google/uuid"
	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain/admins"
)

// verify interface compliance
var _ admins.Repository = (*repo)(nil)

func (r *repo) FindAdmins(cursor *cursor.Cursor, limit int32) ([]*admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) FindAdminByID(adminID uuid.UUID) (*admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CreateAdmin(input *admins.Admin) (*admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) UpdateAdmin(adminID uuid.UUID, mask *admins.UpdateMask, input *admins.Admin) (*admins.Admin, error) {
	panic("not implemented") // TODO: Implement
}

func (r *repo) DeleteAdmin(adminID uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}

func (r *repo) CountAdmins() (int32, error) {
	panic("not implemented") // TODO: Implement
}
