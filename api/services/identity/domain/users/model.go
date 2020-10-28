package users

import (
	"database/sql"
	"time"
)

// User database model of the user
type User struct {
	ID          string         `db:"user_id"`
	AccountID   string         `db:"account_id"`
	FirstName   string         `db:"first_name"`
	LastName    string         `db:"last_name"`
	Email       string         `db:"email"`
	PhoneNumber sql.NullString `db:"phone_number"`
	RoleIDs     []string       `db:"roles"` // alias of ARRAY_AGG function
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}

// Name returns user's resource name
func (u *User) Name() string {
	return Name{AccountID: u.AccountID, UserID: u.ID}.String()
}
