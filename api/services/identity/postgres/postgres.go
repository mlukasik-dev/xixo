package postgres

import (
	"go.xixo.com/api/services/identity/domain/admins"
	"go.xixo.com/api/services/identity/domain/auth"
	"go.xixo.com/api/services/identity/domain/permissions"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/api/services/identity/domain/users"

	"github.com/jmoiron/sqlx"
)

// Repository aggregate repository of domain repositories
type Repository interface {
	auth.Repository
	roles.Repository
	admins.Repository
	users.Repository
	permissions.Repository
}

type repo struct {
	db *sqlx.DB
}

// NewRepository initializes and returns new app repository
func NewRepository(db *sqlx.DB) Repository {
	return &repo{db}
}
