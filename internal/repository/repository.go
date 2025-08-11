package repository

import (
	"context"
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/repository/postgres"
	"data-aggregation-service/internal/types/models"
)

type Repository interface {
	CreateSubscription(ctx context.Context, sub *models.Subscription) (*models.SubscriptionID, error)
	GetSubscription(ctx context.Context, subID *models.SubscriptionID) (*models.Subscription, error)
	UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error
	DeleteSubsription(ctx context.Context, subID *models.SubscriptionID) error
	ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) ([]*models.Subscription, error)
	GetSubscriptionsTotalCost(ctx context.Context, filters *models.SubscriptionsTotalCostFilters) (*models.SubscriptionsTotalCost, error)
}

type Repositories struct {
	Postgres Repository
}

func New(cfg config.PostgresDB) *Repositories {
	return &Repositories{
		Postgres: postgres.New(cfg),
	}
}
