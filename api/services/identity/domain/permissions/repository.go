package permissions

import (
	"context"

	"github.com/google/uuid"
)

// Repository permissions repository
type Repository interface {
	CheckPermission(context.Context, uuid.UUID, string) (bool, error)
}
