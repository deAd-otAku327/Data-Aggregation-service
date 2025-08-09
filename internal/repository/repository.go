package repository

import (
	"context"
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/repository/postgres"
	"data-aggregation-service/internal/types/models"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	CreateSubscription(ctx context.Context, sub *models.Subscription) (*uuid.UUID, error)
	GetSubscription(ctx context.Context, subID uuid.UUID) (*models.Subscription, error)
	UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error
	DeleteSubsription(ctx context.Context, subID uuid.UUID) error
	ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) ([]*models.Subscription, error)
	GetTotalCost(ctx context.Context, periodStart, periodEnd time.Time, filters *models.SubscriptionFilters) (*int, error)
}

type Repositories struct {
	Postgres Repository
}

func New(cfg config.PostgresDB) *Repositories {
	return &Repositories{
		Postgres: postgres.New(cfg),
	}
}
