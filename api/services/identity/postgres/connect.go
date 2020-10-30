package postgres

import (
	"flag"
	"fmt"
	"log"
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
	var envFile string
	flag.StringVar(&envFile, "env-file", "", "env file")
	flag.Parse()

	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Failed to load environment variables from: %s\n", envFile)
		} else {
			log.Printf("Loaded environment variables from: %s\n", envFile)
		}
	}

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
		host, port, user, password, dbname)

	return sqlx.MustConnect("pgx", psqlInfo)
}
