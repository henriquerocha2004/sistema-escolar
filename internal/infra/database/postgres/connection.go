package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type PgConfig struct {
	User               string `json:"user"`
	Password           string `json:"password"`
	Port               string `json:"port"`
	Database           string `json:"database"`
	Host               string `json:"host"`
	AttemptsConnection int    `json:"attempts_connection"`
}

var connectionAttempts int

func Connect() *sql.DB {
	config := loadConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
	)

	for {
		connection, err := connect(dsn)
		if err != nil {
			log.Println("postgres is not ready")
			connectionAttempts++
		} else {
			log.Println("Connected to postgres")
			return connection
		}

		if connectionAttempts > config.AttemptsConnection {
			log.Println(err)
			panic(err)
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(50)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadConfig() *PgConfig {
	attempts, _ := strconv.Atoi(os.Getenv("ATTEMPTS_CONNECTION"))

	pgConfig := PgConfig{
		User:               os.Getenv("DB_USER"),
		Password:           os.Getenv("DB_PASSWORD"),
		Port:               os.Getenv("DB_PORT"),
		Host:               os.Getenv("DB_HOST"),
		Database:           os.Getenv("DB_NAME"),
		AttemptsConnection: attempts,
	}

	return &pgConfig
}
