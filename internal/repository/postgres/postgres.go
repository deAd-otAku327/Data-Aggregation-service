package postgres

import (
	"data-aggregation-service/internal/config"
	"data-aggregation-service/pkg/db"
	"data-aggregation-service/pkg/migrator"
	"database/sql"
)

type postgresRepo struct {
	db *sql.DB
}

func New(cfg *config.PostgresDB) *postgresRepo {
	database := db.MustConnectDB(cfg.DriverName, cfg.URI, cfg.MaxOpenConns)
	migrator.MustApplyMigrations(database, cfg.MigrationsDir)

	return &postgresRepo{
		db: database,
	}
}
