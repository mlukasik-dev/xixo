package postgres

import (
	"fmt"
	"os"

	// importing postgres driver
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	host,
	port,
	user,
	password,
	dbname string
)

func init() {
	godotenv.Load("identity-service/.env")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
}

// MustConnect tries to connect to postgreSQL DB
// if success returns db connection overwise panics
func MustConnect() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "root", "identity")

	return sqlx.MustConnect("pgx", psqlInfo)
}
