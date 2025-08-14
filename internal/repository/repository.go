package repository

import (
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/repository/postgres/subscription"
	"data-aggregation-service/internal/types/domain"
)

func NewSubsRepository(cfg config.SubsRepo) domain.SubscriptionsRepository {
	return subscription.New(cfg)
}
