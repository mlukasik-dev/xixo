package auth

import (
	"context"

	"github.com/google/uuid"
)

// Repository .
type Repository interface {
	CheckUsersPassword(ctx context.Context, accountID uuid.UUID, email string, plainPassword string) (bool, error)
	UpdateUsersPassword(ctx context.Context, accountID uuid.UUID, email string, plainPassword string) error
	CheckAdminsPassword(ctx context.Context, email string, plainPassword string) (bool, error)
	UpdateAdminsPassword(ctx context.Context, email string, plainPassword string) error
	FindAdminInfoByEmail(ctx context.Context, email string) (*AdminInfo, error)
	FindUserInfoByEmail(ctx context.Context, accountID uuid.UUID, email string) (*UserInfo, error)
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
