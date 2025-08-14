package repository

import (
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/repository/postgres/subscriptions"
	"data-aggregation-service/internal/types/domain"
)

func NewSubsRepository(cfg config.SubsRepo) domain.SubscriptionsRepository {
	return subscriptions.New(cfg)
}
