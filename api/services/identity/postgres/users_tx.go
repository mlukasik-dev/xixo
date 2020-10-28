package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.xixo.com/api/services/identity/domain"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain/users"

	"github.com/lib/pq"
)

func findUsersTx(tx Transaction, accountID string, limit int32) (usersSlice []*users.User, err error) {
	const query = `
		SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
			FROM users_with_roles
				WHERE account_id = $1 ORDER BY created_at DESC, user_id DESC LIMIT $2
	`
	rows, err := tx.QueryContext(context.Background(), query, accountID, limit)
	if errors.Is(err, sql.ErrNoRows) {
		// check if account exists if not return an error
		return []*users.User{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nullRoles []sql.NullString
		var user users.User
		err := rows.
			Scan(
				&user.ID, &user.AccountID, &user.FirstName, &user.LastName, &user.Email,
				&user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt, pq.Array(&nullRoles),
			)
		if err != nil {
			return nil, err
		}
		user.RoleIDs = str.NullStringsToStrings(nullRoles)
		usersSlice = append(usersSlice, &user)
	}

	return usersSlice, rows.Err()
}

func findUsersTxCursor(tx Transaction, accountID string, cursor *cursor.Cursor, limit int32) (usersSlice []*users.User, err error) {
	const query = `
		SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
			FROM users_with_roles
				WHERE account_id = $1 AND (
					created_at <= $2 AND (
						created_at < $2 OR user_id < $3
					)
				)
				ORDER BY created_at DESC, user_id DESC LIMIT $4
	`
	rows, err := tx.QueryContext(context.Background(), query, accountID, cursor.Timestamp, cursor.UUID, limit)
	if errors.Is(err, sql.ErrNoRows) {
		// check if account exists if not return an error
		return []*users.User{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nullRoles []sql.NullString
		var user users.User
		err := rows.
			Scan(
				&user.ID, &user.AccountID, &user.FirstName, &user.LastName, &user.Email,
				&user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt, pq.Array(&nullRoles),
			)
		if err != nil {
			return nil, err
		}
		user.RoleIDs = str.NullStringsToStrings(nullRoles)
		usersSlice = append(usersSlice, &user)
	}
	return usersSlice, rows.Err()
}

func findUserByIDTx(tx Transaction, accountID string, userID string) (*users.User, error) {
	const query = `
		SELECT user_id, account_id, first_name, last_name, email, phone_number, created_at, updated_at, roles
			FROM users_with_roles WHERE account_id = $1 AND user_id = $2
	`
	var nullRoles []sql.NullString
	var user users.User
	err := tx.QueryRowContext(context.Background(), query, accountID, userID).Scan(
		&user.ID, &user.AccountID, &user.FirstName, &user.LastName, &user.Email,
		&user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt, pq.Array(&nullRoles),
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	user.RoleIDs = str.NullStringsToStrings(nullRoles)
	return &user, nil
}

func findUserRolesTx(tx Transaction, accountID string, userID string) (roles []string, err error) {
	const query = `
		SELECT roles FROM users_with_roles WHERE account_id = $1 AND user_id = $2
	`
	var nullRoles []sql.NullString
	err = tx.QueryRowContext(context.Background(), query, accountID, userID).Scan(pq.Array(&nullRoles))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	roles = str.NullStringsToStrings(nullRoles)
	return roles, nil
}

func createUserTx(tx Transaction, accountID string, input *users.CreateUserInput) (id string, err error) {
	const query = `
		INSERT INTO users(first_name, last_name, email, phone_number, account_id)
			VALUES ($1, $2, $3, $4, $5) RETURNING user_id
	`
	err = tx.QueryRowContext(
		context.Background(), query,
		input.FirstName, input.LastName, input.Email, input.PhoneNumber, accountID,
	).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func updateUserTx(tx Transaction, accountID string, userID string, mask *users.UpdateMask, input *users.UpdateUserInput) error {
	const query = `
		UPDATE users SET
			first_name = COALESCE($1, first_name),
			last_name = COALESCE($2, last_name),
			email = COALESCE($3, email),
			phone_number = COALESCE($4, phone_number)
		WHERE account_id = $5 AND user_id = $6
	`
	var firstName, lastName, email, phoneNumber sql.NullString
	if mask.FirstName {
		firstName = sql.NullString{String: input.FirstName, Valid: true}
	}
	if mask.LastName {
		lastName = sql.NullString{String: input.LastName, Valid: true}
	}
	if mask.Email {
		email = sql.NullString{String: input.Email, Valid: true}
	}
	if mask.PhoneNumber {
		phoneNumber = sql.NullString{String: input.PhoneNumber, Valid: true}
	}
	_, err := tx.ExecContext(context.Background(), query,
		firstName, lastName, email, phoneNumber, accountID, userID,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user %w", domain.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func deleteUserTx(tx Transaction, accountID, userID string) error {
	const query = `
		DELETE FROM users WHERE account_id = $1 AND user_id = $2
	`
	_, err := tx.ExecContext(context.Background(), query, accountID, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user %w", domain.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func grantRoleToUserTx(tx Transaction, userID, roleID string) error {
	const query = `
		INSERT INTO users_roles(user_id, role_id) VALUES ($1, $2)
	`
	_, err := tx.ExecContext(context.Background(), query, userID, roleID)
	return err
}

func grantAccountAdminRoleToUserTx(tx Transaction, userID string) error {
	const query = `
		INSERT INTO users_roles(user_id, role_id) VALUES ($1, (
			SELECT role_id FROM roles WHERE account_admin
		)) ON CONFLICT DO NOTHING
	`
	_, err := tx.ExecContext(context.Background(), query, userID)
	return err
}

func denyRoleFromUserTx(tx Transaction, userID, roleID string) error {
	const query = `
		DELETE FROM users_roles WHERE user_id = $1 AND role_id = $2
	`
	_, err := tx.ExecContext(context.Background(), query, userID, roleID)
	return err
}
