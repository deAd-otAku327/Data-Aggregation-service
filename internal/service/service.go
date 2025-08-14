package service

import (
	"data-aggregation-service/internal/service/subscription"
	"data-aggregation-service/internal/types/domain"
)

func NewSubsService(r domain.SubscriptionRepository) domain.SubscriptionService {
	return subscription.New(r)
}
