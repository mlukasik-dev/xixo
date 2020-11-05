package auth

import "github.com/google/uuid"

// Repository .
type Repository interface {
	CheckUsersPassword(accountID uuid.UUID, email string, plainPassword string) (bool, error)
	UpdateUsersPassword(accountID uuid.UUID, email string, plainPassword string) error
	CheckAdminsPassword(email string, plainPassword string) (bool, error)
	UpdateAdminsPassword(email string, plainPassword string) error
	FindAdminInfoByEmail(email string) (*AdminInfo, error)
	FindUserInfoByEmail(accountID uuid.UUID, email string) (*UserInfo, error)
}

// AdminInfo .
type AdminInfo struct {
	ID         uuid.UUID
	Roles      []uuid.UUID
	Registered bool
}

// UserInfo .
type UserInfo struct {
	ID         uuid.UUID
	Roles      []uuid.UUID
	Registered bool
}
