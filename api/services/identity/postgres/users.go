package postgres

import (
	"context"
	"database/sql"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain/users"
)

// verify interface compliance
var _ users.Repository = (*repo)(nil)

func (r *repo) FindUsers(accountID string, cursor *cursor.Cursor, limit int32) (users []*users.User, err error) {
	// first request
	if cursor == nil {
		users, err = findUsersTx(r.db, accountID, limit)
	} else {
		users, err = findUsersTxCursor(r.db, accountID, cursor, limit)
	}
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repo) FindUserByID(accountID string, userID string) (*users.User, error) {
	user, err := findUserByIDTx(r.db, accountID, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) CreateUser(accountID string, input *users.CreateUserInput) (*users.User, error) {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	userID, err := createUserTx(tx, accountID, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, roleID := range input.RoleIDs {
		if err = grantRoleToUserTx(tx, userID, roleID); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	user, err := findUserByIDTx(tx, accountID, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) CreateAccountAdmin(accountID string, input *users.CreateUserInput) (*users.User, error) {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	userID, err := createUserTx(tx, accountID, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if len(input.RoleIDs) > 0 {
		for _, roleID := range input.RoleIDs {
			if err = grantRoleToUserTx(tx, userID, roleID); err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}
	if err = grantAccountAdminRoleToUserTx(tx, userID); err != nil {
		tx.Rollback()
		return nil, err
	}
	user, err := findUserByIDTx(tx, accountID, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) UpdateUser(accountID string, userID string, mask *users.UpdateMask, input *users.UpdateUserInput) (*users.User, error) {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	err = updateUserTx(tx, accountID, userID, mask, input)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if mask.RoleIDs {
		roles, err := findUserRolesTx(tx, accountID, userID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		// Grant new roles
		for _, roleID := range input.RoleIDs {
			if !str.SliceContains(roles, roleID) {
				err = grantRoleToUserTx(tx, userID, roleID)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			}
		}
		// Deny roles
		for _, roleID := range roles {
			if !str.SliceContains(input.RoleIDs, roleID) {
				err = denyRoleFromUserTx(tx, userID, roleID)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			}
		}
	}
	// Select updated user
	user, err := findUserByIDTx(tx, accountID, userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) DeleteUser(accountID string, userID string) error {
	err := deleteUserTx(r.db, accountID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) CountUsers(accountID string) (int32, error) {
	const query = `
		SELECT COUNT(*) FROM users
	`
	var count int32
	err := r.db.QueryRowContext(context.Background(), query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
