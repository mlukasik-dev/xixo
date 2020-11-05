package postgres

import (
	"database/sql"

	"go.xixo.com/api/services/account/config"

	// importing postgres driver
	_ "github.com/jackc/pgx/v4/stdlib"
)

// MustConnect tries to connect to postgreSQL DB
// if success returns db connection overwise panics
func MustConnect() *sql.DB {
	db, err := sql.Open("pgx", config.Global.Postgres.String())
	if err != nil {
		panic(err)
	}
	return db
}
