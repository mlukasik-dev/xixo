package permissions

import "context"

// Repository permissions repository
type Repository interface {
	CheckPermission(ctx context.Context, roleID, method string) (bool, error)
}
