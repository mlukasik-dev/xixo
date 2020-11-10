package roles

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"go.xixo.com/api/pkg/cursor"

	"github.com/google/uuid"
)

// UpdateMask .
type UpdateMask struct {
	AdminOnly   bool
	DisplayName bool
	Description bool
	Permissions bool
}

// Filter .
type Filter struct {
	AdminOnly sql.NullBool
}

// NewFilter .
func NewFilter(query string) (*Filter, error) {
	vals, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	var adminOnly sql.NullBool
	switch vals.Get("admin_only") {
	case "true":
		adminOnly.Bool = true
		adminOnly.Valid = true
		break
	case "":
		if len(vals["admin_only"]) > 0 {
			adminOnly.Bool = true
			adminOnly.Valid = true
		}
		break
	case "false":
		adminOnly.Bool = false
		adminOnly.Valid = true
		break
	default:
		return nil, fmt.Errorf("error") // TODO: change error message
	}
	return &Filter{
		AdminOnly: adminOnly,
	}, nil
}

// Repository .
type Repository interface {
	FindRoles(ctx context.Context, cursor *cursor.Cursor, limit int32, filter *Filter) ([]Role, error)
	FindRoleByID(ctx context.Context, roleID uuid.UUID) (*Role, error)
	CreateRole(ctx context.Context, input *Role) (*Role, error)
	UpdateRole(ctx context.Context, roleID uuid.UUID, mask *UpdateMask, input *Role) (*Role, error)
	DeleteRole(ctx context.Context, roleID uuid.UUID) error
	CountRoles(ctx context.Context) (int32, error)
}
