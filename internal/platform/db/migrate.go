package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migrateSqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(sqlDB *sql.DB, migrationsDir string) error {
	driver, err := migrateSqlite3.WithInstance(sqlDB, &migrateSqlite3.Config{
		// MigrationsTable: "schema_migrations", // optionnel
	})
	if err != nil {
		return fmt.Errorf("migrate: create sqlite driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDir,
		"sqlite3",
		driver,
	)
	if err != nil {
		return fmt.Errorf("migrate: new instance: %w", err)
	}

	// Applique toutes les migrations "up" non appliquées
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate: up: %w", err)
	}

	return nil
}
