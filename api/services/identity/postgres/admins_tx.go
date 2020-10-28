package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/pkg/str"
	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/admins"

	"github.com/lib/pq"
)

func findAdminsTx(tx Transaction, limit int32) (adminsSlice []*admins.Admin, err error) {
	const query = `
		SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
			FROM admins_with_roles
				ORDER BY created_at DESC, admin_id DESC LIMIT $1
	`
	rows, err := tx.QueryContext(context.Background(), query, limit)
	if errors.Is(err, sql.ErrNoRows) {
		return []*admins.Admin{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var nullRoles []sql.NullString
		var admin admins.Admin
		err := rows.
			Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt, pq.Array(&nullRoles))
		if err != nil {
			return nil, err
		}
		admin.RoleIDs = str.NullStringsToStrings(nullRoles)
		adminsSlice = append(adminsSlice, &admin)
	}
	return adminsSlice, rows.Err()
}

func findAdminsTxCursor(tx Transaction, cursor *cursor.Cursor, limit int32) (adminsSlice []*admins.Admin, err error) {
	const query = `
		SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
			FROM admins_with_roles WHERE (
				created_at <= $1 AND (
					created_at < $1 OR admin_id < $2
				)
			) ORDER BY created_at DESC, admin_id DESC LIMIT $3
	`
	rows, err := tx.QueryContext(context.Background(), query, cursor.Timestamp, cursor.UUID, limit)
	if errors.Is(err, sql.ErrNoRows) {
		return []*admins.Admin{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var nullRoles []sql.NullString
		var admin admins.Admin
		err := rows.
			Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt, pq.Array(&nullRoles))
		if err != nil {
			return nil, err
		}
		admin.RoleIDs = str.NullStringsToStrings(nullRoles)
		adminsSlice = append(adminsSlice, &admin)
	}
	return adminsSlice, rows.Err()
}

func findAdminByIDTx(tx Transaction, adminID string) (*admins.Admin, error) {
	const query = `
		SELECT admin_id, first_name, last_name, email, created_at, updated_at, roles
			FROM admins_with_roles WHERE admin_id = $1
	`
	var nullRoles []sql.NullString
	var admin admins.Admin
	err := tx.QueryRowContext(context.Background(), query, adminID).
		Scan(&admin.ID, &admin.FirstName, &admin.LastName, &admin.Email, &admin.CreatedAt, &admin.UpdatedAt, pq.Array(&nullRoles))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	admin.RoleIDs = str.NullStringsToStrings(nullRoles)
	return &admin, nil
}

func findAdminRolesTx(tx Transaction, adminID string) (roles []string, err error) {
	const query = `
		SELECT roles FROM admins_with_roles WHERE admin_id = $1
	`
	var nullRoles []sql.NullString
	err = tx.QueryRowContext(context.Background(), query, adminID).Scan(pq.Array(&nullRoles))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	roles = str.NullStringsToStrings(nullRoles)
	return roles, nil
}

func createAdminTx(tx Transaction, input *admins.CreateAdminInput) (adminID string, err error) {
	const query = `
		INSERT INTO admins(first_name, last_name, email)
			VALUES ($1, $2, $3)	RETURNING admin_id
	`
	err = tx.QueryRowContext(
		context.Background(), query,
		input.FirstName, input.LastName, input.Email,
	).Scan(&adminID)
	if err != nil {
		return "", err
	}
	return adminID, nil
}

func updateAdminTx(tx Transaction, adminID string, mask *admins.UpdateMask, input *admins.UpdateAdminInput) error {
	const query = `
		UPDATE admins SET
			first_name = COALESCE($1, first_name),
			last_name = COALESCE($2, last_name),
			email = COALESCE($3, email)
		WHERE admin_id = $4
	`
	var firstName, lastName, email sql.NullString
	if mask.FirstName {
		firstName = sql.NullString{String: input.FirstName, Valid: true}
	}
	if mask.LastName {
		lastName = sql.NullString{String: input.LastName, Valid: true}
	}
	if mask.Email {
		email = sql.NullString{String: input.Email, Valid: true}
	}
	_, err := tx.ExecContext(context.Background(), query, firstName, lastName, email, adminID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func deleteAdminTx(tx Transaction, adminID string) error {
	const query = `
		DELETE FROM admins WHERE admin_id = $1
	`
	_, err := tx.ExecContext(context.Background(), query, adminID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("admin %w", domain.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func grantRoleToAdminTx(tx Transaction, adminID, roleID string) error {
	const query = `
		INSERT INTO admins_roles(admin_id, role_id) VALUES ($1, $2)
	`
	_, err := tx.ExecContext(context.Background(), query, adminID, roleID)
	return err
}

func denyRoleFromAdminTx(tx Transaction, adminID, roleID string) error {
	const query = `
		DELETE FROM admins_roles WHERE admin_id = $1 AND role_id = $2
	`
	_, err := tx.ExecContext(context.Background(), query, adminID, roleID)
	return err
}
