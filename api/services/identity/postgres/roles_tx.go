package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.xixo.com/api/pkg/cursor"
	"go.xixo.com/api/services/identity/domain"
	"go.xixo.com/api/services/identity/domain/roles"

	"github.com/lib/pq"
)

func findRolesTx(tx Transaction, limit int32, filter *roles.Filter) ([]*roles.Role, error) {
	const query = `
		SELECT
			role_id, admin_only, display_name, description, permissions, created_at, updated_at
				FROM roles_with_permissions WHERE
					admin_only = COALESCE($1, admin_only)
				ORDER BY created_at DESC, role_id DESC LIMIT $2
	`
	rows, err := tx.QueryContext(context.Background(), query, filter.AdminOnly, limit)
	if errors.Is(err, sql.ErrNoRows) {
		return []*roles.Role{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolesSlice []*roles.Role
	for rows.Next() {
		var role roles.Role
		err := rows.Scan(&role.ID, &role.AdminOnly, &role.DisplayName, &role.Description, pq.Array(&role.Permissions), &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rolesSlice = append(rolesSlice, &role)
	}
	return rolesSlice, rows.Err()
}

func findRolesTxCursor(tx Transaction, cursor *cursor.Cursor, limit int32, filter *roles.Filter) ([]*roles.Role, error) {
	const query = `
		SELECT
			role_id, admin_only, display_name, description, permissions, created_at, updated_at
				FROM roles_with_permissions WHERE
					created_at <= $1 AND (
						created_at < $1 OR role_id < $2
					) AND
					admin_only = COALESCE($3, admin_only)
				ORDER BY created_at DESC, role_id DESC LIMIT $4
	`
	rows, err := tx.QueryContext(context.Background(), query, cursor.Timestamp, cursor.UUID, filter.AdminOnly, limit)
	if errors.Is(err, sql.ErrNoRows) {
		return []*roles.Role{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rolesSlice []*roles.Role
	for rows.Next() {
		var role roles.Role
		err := rows.Scan(&role.ID, &role.AdminOnly, &role.DisplayName, &role.Description, pq.Array(&role.Permissions), &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rolesSlice = append(rolesSlice, &role)
	}

	return rolesSlice, rows.Err()
}

func findRoleByIDTx(tx Transaction, roleID string) (*roles.Role, error) {
	const query = `
		SELECT
			role_id, admin_only, display_name, description, permissions, created_at, updated_at
				FROM roles_with_permissions WHERE role_id = $1
	`
	var role roles.Role
	err := tx.QueryRowContext(context.Background(), query, roleID).
		Scan(&role.ID, &role.AdminOnly, &role.DisplayName, &role.Description, pq.Array(&role.Permissions), &role.CreatedAt, &role.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("role %w", domain.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func findRolePermissionsTx(tx Transaction, roleID string) (permissions []string, err error) {
	const query = `
		SELECT
			permissions
				FROM roles_with_permissions WHERE role_id = $1
	`
	err = tx.QueryRowContext(context.Background(), query, roleID).
		Scan(pq.Array(&permissions))
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func createRoleTx(tx Transaction, input *roles.CreateRoleInput) (roleID string, err error) {
	const query = `
		INSERT INTO roles(admin_only, display_name, description) VALUES ($1, $2, $3)
			RETURNING role_id
	`
	err = tx.QueryRowContext(context.Background(), query, input.AdminOnly, input.DisplayName, input.Description).Scan(&roleID)
	if err != nil {
		return "", err
	}
	return
}

func updateRoleTx(tx Transaction, roleID string, mask *roles.UpdateMask, input *roles.UpdateRoleInput) error {
	const query = `
		UPDATE roles SET
			admin_only = COALESCE($1, admin_only),
			display_name = COALESCE($2, display_name),
			description = COALESCE($3, description)
		WHERE role_id = $4
	`
	var adminOnly sql.NullBool
	var displayName, description sql.NullString
	if mask.AdminOnly {
		adminOnly = sql.NullBool{Bool: input.AdminOnly, Valid: true}
	}
	if mask.DisplayName {
		displayName = sql.NullString{String: input.DisplayName, Valid: true}
	}
	if mask.Description {
		description = sql.NullString{String: input.Description, Valid: true}
	}
	_, err := tx.ExecContext(context.Background(), query, adminOnly, displayName, description, roleID)
	return err
}

func deleteRoleTx(tx Transaction, roleID string) error {
	const query = `
		DELETE FROM roles WHERE role_id = $1
	`
	_, err := tx.ExecContext(context.Background(), query, roleID)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("role %w", domain.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func grantPermissionTx(tx Transaction, roleID, method string) error {
	const query = `
		INSERT INTO permissions(role_id, method)
		VALUES ($1, $2)
			ON CONFLICT ON CONSTRAINT permissions_unique_role_id_and_method
				DO UPDATE SET updated_at = NOW()
	`
	_, err := tx.ExecContext(context.Background(), query, roleID, method)
	return err
}

func denyPermissionTx(tx Transaction, roleID, method string) error {
	const query = `
		DELETE FROM permissions WHERE role_id = $1 AND method = $2
	`
	_, err := tx.ExecContext(context.Background(), query, roleID, method)
	return err
}
