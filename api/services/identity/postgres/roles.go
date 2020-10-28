package postgres

import (
	"context"
	"database/sql"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain/roles"
)

// verify interface compliance
var _ roles.Repository = (*repo)(nil)

func (r *repo) FindRoles(cursor *cursor.Cursor, limit int32, filter *roles.Filter) (roles []*roles.Role, err error) {
	if cursor == nil {
		// it's first request
		roles, err = findRolesTx(r.db, limit, filter)
	} else {
		// it's subsequent request
		roles, err = findRolesTxCursor(r.db, cursor, limit, filter)
	}
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *repo) FindRoleByID(roleID string) (*roles.Role, error) {
	role, err := findRoleByIDTx(r.db, roleID)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *repo) CreateRole(input *roles.CreateRoleInput) (*roles.Role, error) {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	roleID, err := createRoleTx(tx, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if len(input.Permissions) > 0 {
		// Granting permissions
		for _, method := range input.Permissions {
			if err = grantPermissionTx(tx, roleID, method); err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}
	role, err := findRoleByIDTx(tx, roleID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return role, nil
}

func (r *repo) UpdateRole(roleID string, mask *roles.UpdateMask, input *roles.UpdateRoleInput) (*roles.Role, error) {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	err = updateRoleTx(tx, roleID, mask, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if mask.Permissions {
		permissions, err := findRolePermissionsTx(tx, roleID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		// Grant new permissions
		for _, method := range input.Permissions {
			if !str.SliceContains(permissions, method) {
				err = grantPermissionTx(tx, roleID, method)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			}
		}
		// Deny permissions
		for _, method := range permissions {
			if !str.SliceContains(input.Permissions, method) {
				err = denyRoleFromUserTx(tx, roleID, method)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			}
		}
	}
	// Select updated role
	role, err := findRoleByIDTx(tx, roleID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return role, nil
}

func (r *repo) DeleteRole(roleID string) error {
	err := deleteRoleTx(r.db, roleID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CountRoles() (int32, error) {
	const query = `
		SELECT COUNT(*) FROM roles
	`
	var count int32
	err := r.db.QueryRowContext(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
