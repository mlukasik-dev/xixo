package permissions

import (
	"context"

	"github.com/google/uuid"
)

// Repository permissions repository
type Repository interface {
	CheckPermission(ctx context.Context, roleID uuid.UUID, method string) (bool, error)
}
