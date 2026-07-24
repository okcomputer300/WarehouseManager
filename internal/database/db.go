package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() error {
	_ = godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "db"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "admin"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "lager_assets"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error validating DSN: %w", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	for i := 1; i <= 5; i++ {
		err = DB.Ping()
		if err == nil {
			log.Println("Successfully connected to db")

			if err := RunMigrations(DB); err != nil {
				return fmt.Errorf("migration error: %w", err)
			}

			return nil
		}
		log.Printf("Try %d: DB not ready yet\n", i)
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("couldn't connect to DB after 5 tries: %w", err)
}

func RunMigrations(db *sql.DB) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("could not create mysql driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("migration init failed: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		var dirtyErr migrate.ErrDirty
		if errors.As(err, &dirtyErr) {
			log.Printf("Detected dirty migration version %d. Forcing version to retry...\n", dirtyErr.Version)

			if forceErr := m.Force(dirtyErr.Version - 1); forceErr != nil {
				return fmt.Errorf("failed to force migration version: %w", forceErr)
			}

			if retryErr := m.Up(); retryErr != nil && !errors.Is(retryErr, migrate.ErrNoChange) {
				return fmt.Errorf("failed to re-apply migrations after force: %w", retryErr)
			}
		} else {
			return fmt.Errorf("failed to apply migrations: %w", err)
		}
	}

	log.Println("Database migrations applied successfully!")
	return nil
}
