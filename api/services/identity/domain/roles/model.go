package roles

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Role model
type Role struct {
	ID          uuid.UUID
	AdminOnly   bool
	DisplayName string
	Description string
	Permissions []sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Name returns role's resource name
func (r *Role) Name() string {
	return Name{RoleID: r.ID}.String()
}
