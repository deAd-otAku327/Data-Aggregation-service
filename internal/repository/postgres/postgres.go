package postgres

import (
	"context"
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/types/models"
	"data-aggregation-service/pkg/db"
	"data-aggregation-service/pkg/migrator"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type postgresRepo struct {
	db *sql.DB
}

func New(cfg config.PostgresDB) *postgresRepo {
	database := db.MustConnectDB(cfg.DriverName, cfg.URI, cfg.MaxOpenConns)
	migrator.MustApplyMigrations(database, cfg.MigrationsDir)

	return &postgresRepo{
		db: database,
	}
}

func (r *postgresRepo) CreateSubscription(ctx context.Context, sub *models.Subscription) (*uuid.UUID, error) {
	return nil, nil
}

func (r *postgresRepo) GetSubscription(ctx context.Context, subID uuid.UUID) (*models.Subscription, error) {
	return nil, nil
}

func (r *postgresRepo) UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error {
	return nil
}

func (r *postgresRepo) DeleteSubsription(ctx context.Context, subID uuid.UUID) error {
	return nil
}

func (r *postgresRepo) ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) ([]*models.Subscription, error) {
	return nil, nil
}

func (r *postgresRepo) GetTotalCost(ctx context.Context, periodStart, periodEnd time.Time, filters *models.SubscriptionFilters) (*int, error) {
	return nil, nil
}
