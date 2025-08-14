package service

import (
	"context"
	"data-aggregation-service/internal/mappers/dtomap"
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/internal/types/dto"

	"github.com/google/uuid"
)

type Service interface {
	CreateSubscription(ctx context.Context, sub *domain.Subscription) (*dto.SubscriptionIDResponse, error)
	GetSubscription(ctx context.Context, subID *domain.SubscriptionID) (*dto.SubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, subID *domain.SubscriptionID, patch *domain.SubscriptionPatch) error
	DeleteSubsription(ctx context.Context, subID *domain.SubscriptionID) error
	ListSubscriptions(ctx context.Context, filters *domain.SubscriptionFilters) (*dto.SubscriptionListResponse, error)
	GetSubscriptionsTotalCost(ctx context.Context, filters *domain.SubscriptionsTotalCostFilters) (*dto.TotalCostResponse, error)
}

type service struct {
	repo domain.SubscriptionsRepository
}

func New(r domain.SubscriptionsRepository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateSubscription(ctx context.Context, sub *domain.Subscription) (*dto.SubscriptionIDResponse, error) {
	sub.ID = uuid.New()

	subscriptionID, err := s.repo.Insert(ctx, sub)
	if err != nil {
		return nil, wrapError(err)
	}

	return dtomap.MapToSubscriptionIDResponse(subscriptionID), nil
}

func (s *service) GetSubscription(ctx context.Context, subID *domain.SubscriptionID) (*dto.SubscriptionResponse, error) {
	subscription, err := s.repo.SelectByID(ctx, subID)
	if err != nil {
		return nil, wrapError(err)
	}

	return dtomap.MapToSubscriptionResponse(subscription), nil

}

func (s *service) UpdateSubscription(ctx context.Context, subscriptionID *domain.SubscriptionID, patch *domain.SubscriptionPatch) error {
	err := s.repo.Update(ctx, subscriptionID, patch)
	if err != nil {
		return wrapError(err)
	}
	return nil
}

func (s *service) DeleteSubsription(ctx context.Context, subscriptionID *domain.SubscriptionID) error {
	err := s.repo.Delete(ctx, subscriptionID)
	if err != nil {
		return wrapError(err)
	}

	return nil
}

func (s *service) ListSubscriptions(ctx context.Context, filters *domain.SubscriptionFilters) (*dto.SubscriptionListResponse, error) {
	subs, err := s.repo.SelectList(ctx, filters)
	if err != nil {
		return nil, wrapError(err)
	}

	return dtomap.MapToSubscriptionListResponse(subs), nil
}

func (s *service) GetSubscriptionsTotalCost(ctx context.Context, filters *domain.SubscriptionsTotalCostFilters) (*dto.TotalCostResponse, error) {
	totalCost, err := s.repo.SelectTotalCost(ctx, filters)
	if err != nil {
		return nil, wrapError(err)
	}

	return dtomap.MapToTotalCostResponse(totalCost), nil
}
