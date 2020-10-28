package postgres

import (
	"context"
	"database/sql"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain/admins"
)

// verify interface compliance
var _ admins.Repository = (*repo)(nil)

func (r *repo) FindAdmins(cursor *cursor.Cursor, limit int32) (admins []*admins.Admin, err error) {
	// first request
	if cursor == nil {
		admins, err = findAdminsTx(r.db, limit)
	} else {
		admins, err = findAdminsTxCursor(r.db, cursor, limit)
	}
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *repo) FindAdminByID(adminID string) (*admins.Admin, error) {
	admin, err := findAdminByIDTx(r.db, adminID)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *repo) CreateAdmin(input *admins.CreateAdminInput) (*admins.Admin, error) {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	adminID, err := createAdminTx(tx, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, roleID := range input.RoleIDs {
		if err = grantRoleToAdminTx(tx, adminID, roleID); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	admin, err := findAdminByIDTx(tx, adminID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *repo) UpdateAdmin(adminID string, mask *admins.UpdateMask, input *admins.UpdateAdminInput) (*admins.Admin, error) {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	err = updateAdminTx(tx, adminID, mask, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if mask.RoleIDs {
		roles, err := findAdminRolesTx(tx, adminID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		// Grant new roles
		for _, roleID := range input.RoleIDs {
			if !str.SliceContains(roles, roleID) {
				err = grantRoleToAdminTx(tx, adminID, roleID)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			}
		}
		// Deny roles
		for _, roleID := range roles {
			if !str.SliceContains(input.RoleIDs, roleID) {
				err = denyRoleFromAdminTx(tx, adminID, roleID)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			}
		}
	}
	// Select updated admin
	admin, err := findAdminByIDTx(tx, adminID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *repo) DeleteAdmin(adminID string) error {
	err := deleteAdminTx(r.db, adminID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CountAdmins() (int32, error) {
	const query = `
		SELECT COUNT(*) FROM admins
	`
	var count int32
	err := r.db.QueryRowContext(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
