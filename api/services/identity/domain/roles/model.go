package roles

import (
	"database/sql"
	"time"
)

// Role model
type Role struct {
	ID          string           `db:"role_id"`
	AdminOnly   bool             `db:"admin_only"`
	DisplayName string           `db:"display_name"`
	Description string           `db:"description"`
	Permissions []sql.NullString `db:"permissions"` // alias of ARRAY_AGG function
	CreatedAt   time.Time        `db:"created_at"`
	UpdatedAt   time.Time        `db:"updated_at"`
}

// Name returns role's resource name
func (r *Role) Name() string {
	return Name{RoleID: r.ID}.String()
}
