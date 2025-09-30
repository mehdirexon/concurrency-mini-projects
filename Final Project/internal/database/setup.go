package database

import (
	"database/sql"
	config "final-project/internal/config"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var appConfig *config.AppConfig

func Register(app *config.AppConfig) {
	appConfig = app
}

func Init() *sql.DB {
	conn := connectDB()
	if conn != nil {
		// panic
	}

	return conn
}

func connectDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")
	if dsn == "" {
		appConfig.ErrorLogger.Fatalln("🚫 🚫 DSN environment variable not set 🚫 🚫 ")
	}

	for {
		connection, err := Open(dsn)
		if err != nil {
			appConfig.ErrorLogger.Printf("🚫 Couldn't connect to database")

		} else {
			appConfig.InfoLogger.Printf("✅ Connected to database")
			return connection
		}

		if counts > 10 {
			return nil
		}

		appConfig.InfoLogger.Printf("🔸 Backing off for 1 second")
		<-time.After(1 * time.Second)
		counts++
		continue
	}
}

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
