package postgres

import (
	"go.xixo.com/api/services/account/domain/accounts"
	"go.xixo.com/api/services/account/domain/permissions"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Repository aggregated repository of domain repositories
type Repository interface {
	accounts.Repository
	permissions.Repository
}

type repo struct {
	db     *sqlx.DB
	logger *zap.Logger
}

// NewRepository initializes and returns new aggregated repository
func NewRepository(db *sqlx.DB, l *zap.Logger) Repository {
	return &repo{db, l}
}
