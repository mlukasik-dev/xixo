package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type port uint16

func (p port) String() string {
	return ":" + strconv.Itoa(int(p))
}

// Postgres .
type Postgres struct {
	Host     string
	Port     port
	User     string
	Password string
	Name     string
}

func (pg *Postgres) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.User, pg.Password, pg.Name)
}

// Auth .
type Auth struct {
	Secret        string
	TokenDuration time.Duration
}

// Config .
type Config struct {
	Port     port
	Postgres Postgres
	Auth     Auth
}

// Global .
var Global Config

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

	var err error
	Global, err = parse()
	if err != nil {
		log.Fatalf("Failed to parse environment variables: %v\n", err)
	}
}

func parse() (Config, error) {
	p, err := strconv.Atoi(os.Getenv("PORT"))
	dbp, err := strconv.Atoi(os.Getenv("DB_PORT"))
	mins, err := strconv.Atoi(os.Getenv("AUTH_TOKEN_DURATION_MINUTES"))
	return Config{
		Port: port(p),
		Postgres: Postgres{
			os.Getenv("DB_HOST"),
			port(dbp),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		},
		Auth: Auth{
			os.Getenv("AUTH_SECRET"),
			time.Duration(mins) * time.Minute,
		},
	}, err
}
