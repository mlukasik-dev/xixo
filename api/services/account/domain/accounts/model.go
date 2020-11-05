package accounts

import (
	"time"

	"github.com/google/uuid"
)

// Account account business model
type Account struct {
	ID          uuid.UUID
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Name returns account's resource name
func (a *Account) Name() string {
	return Name{AccountID: a.ID}.String()
}
