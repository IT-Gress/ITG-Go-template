package database

import (
	"embed"
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

var db *sqlx.DB

// Init initializes the database connection using sqlx and returns the database instance
// after running a migration.
func Init(postgresDSN string) (*sqlx.DB, error) {
	// Check if the database connection is already established
	if db != nil {
		slog.Debug("Database connection already established")
		return db, nil
	}

	// Connect to the database
	db, err := sqlx.Connect("postgres", postgresDSN)
	if err != nil {
		return nil, err
	}

	if err := migrateDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	// Create iofs (in-memory) driver for the embedded migrations
	sourceDriver, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance(
		"iofs", sourceDriver,
		"postgres", driver)
	if err != nil {
		return err
	}

	// Run migrations up
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
