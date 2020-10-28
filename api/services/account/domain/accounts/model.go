package accounts

import "time"

// Account account business model
type Account struct {
	ID          string    `db:"account_id"`
	DisplayName string    `db:"display_name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Name returns account's resource name
func (a *Account) Name() string {
	return Name{AccountID: a.ID}.String()
}
