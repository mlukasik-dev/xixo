package admins

import (
	"time"

	"github.com/google/uuid"
	"go.xixo.com/api/services/identity/domain/roles"
)

// Admin database model of the admin
type Admin struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Roles     []uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Name returns admin's resource name
func (a *Admin) Name() string {
	return Name{AdminID: a.ID}.String()
}

// RoleNames returns slice of admin's role resource names
func (a *Admin) RoleNames() []string {
	var names []string
	for _, id := range a.Roles {
		names = append(names, roles.Name{RoleID: id}.String())
	}
	return names
}
