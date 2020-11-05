package users

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"go.xixo.com/api/services/identity/domain/roles"
)

// User database model of the user
type User struct {
	ID          uuid.UUID
	AccountID   uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber sql.NullString
	Roles       []uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Name returns user's resource name
func (u *User) Name() string {
	return Name{AccountID: u.AccountID, UserID: u.ID}.String()
}

// RoleNames returns slice of user's role resource names
func (u *User) RoleNames() []string {
	var names []string
	for _, id := range u.Roles {
		names = append(names, roles.Name{RoleID: id}.String())
	}
	return names
}
