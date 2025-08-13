package migrator

import (
	"database/sql"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MustApplyMigrations(db *sql.DB, migrationsDir string) {
	postgres, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	migrator, err := migrate.NewWithDatabaseInstance(migrationsDir, "db", postgres)
	if err != nil {
		panic(err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	slog.Info("database migrated and ready")
}
