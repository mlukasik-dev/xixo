package postgres

import (
	"database/sql"

	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/account/domain/permissions"
	"go.xixo.com/api/services/identity/postgres/gen"

	"go.uber.org/zap"
)

// Repository aggregated repository of domain repositories
type Repository interface {
	accounts.Repository
	permissions.Repository
}

type repo struct {
	db     *sql.DB
	q      *gen.Queries
	logger *zap.Logger
}

// NewRepository initializes and returns new aggregated repository
func NewRepository(db *sql.DB, l *zap.Logger) Repository {
	q := gen.New(db)
	return &repo{db, q, l}
}
