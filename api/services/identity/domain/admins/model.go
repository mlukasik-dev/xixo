package admins

import (
	"time"
)

// Admin database model of the admin
type Admin struct {
	ID        string    `db:"admin_id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	RoleIDs   []string  `db:"roles"` // alias of ARRAY_AGG function
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Name returns admin's resource name
func (a *Admin) Name() string {
	return Name{AdminID: a.ID}.String()
}
