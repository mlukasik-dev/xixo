package postgres

import (
	"database/sql"

	"go.xixo.com/api/services/identity/domain/admins"
	"go.xixo.com/api/services/identity/domain/auth"
	"go.xixo.com/api/services/identity/domain/permissions"
	"go.xixo.com/api/services/identity/domain/roles"
	"go.xixo.com/api/services/identity/domain/users"
	"go.xixo.com/api/services/identity/postgres/gen"

	"go.uber.org/zap"
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
	db     *sql.DB
	q      *gen.Queries
	logger *zap.Logger
}

// NewRepository initializes and returns new app repository
func NewRepository(db *sql.DB, l *zap.Logger) Repository {
	q := gen.New(db)
	return &repo{db, q, l}
}
