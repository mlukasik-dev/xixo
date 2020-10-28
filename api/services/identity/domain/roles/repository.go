package roles

import (
	"database/sql"
	"fmt"
	"net/url"

	"go.xixo.com/api/pkg/cursor"

	"github.com/go-playground/validator/v10"
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
	FindRoles(cursor *cursor.Cursor, limit int32, filter *Filter) ([]*Role, error)
	FindRoleByID(id string) (*Role, error)
	CreateRole(input *CreateRoleInput) (*Role, error)
	UpdateRole(id string, mask *UpdateMask, input *UpdateRoleInput) (*Role, error)
	DeleteRole(id string) error
	CountRoles() (int32, error)
}

// CreateRoleInput .
type CreateRoleInput struct {
	AdminOnly   bool
	DisplayName string
	Description string
	Permissions []string
}

// Validate .
func (i *CreateRoleInput) Validate(v *validator.Validate) error {
	return v.Struct(i)
}

// UpdateRoleInput .
type UpdateRoleInput struct {
	AdminOnly   bool
	DisplayName string
	Description string
	Permissions []string
}

// Validate .
func (i *UpdateRoleInput) Validate(v *validator.Validate) error {
	return v.Struct(i)
}
