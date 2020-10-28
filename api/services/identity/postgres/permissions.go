package postgres

import "context"

// Checks if this role has permissions to use this method
func (r *repo) CheckPermission(ctx context.Context, roleID, method string) (hasPermission bool, err error) {
	const query = `
		SELECT EXISTS (
			SELECT permission_id FROM permissions
				WHERE role_id = $1 AND method = $2
		)
	`
	err = r.db.QueryRowContext(ctx, query, roleID, method).Scan(&hasPermission)
	if err != nil {
		return false, err
	}
	return hasPermission, nil
}