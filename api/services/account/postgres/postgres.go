package postgres

import (
	"database/sql"

	"go.xixo.com/api/services/account/postgres/gen"

	"go.uber.org/zap"
)

// Repository .
type Repository struct {
	db     *sql.DB
	q      *gen.Queries
	logger *zap.Logger
}

// NewRepository initializes and returns new aggregated repository
func NewRepository(db *sql.DB, l *zap.Logger) *Repository {
	q := gen.New(db)
	return &Repository{db, q, l}
}
