package service

import (
	"context"
	"data-aggregation-service/internal/mappers/dtomap"
	"data-aggregation-service/internal/repository"
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"

	"github.com/google/uuid"
)

type Service interface {
	CreateSubscription(ctx context.Context, sub *models.Subscription) (*dto.SubscriptionIDResponse, error)
	GetSubscription(ctx context.Context, subID *models.SubscriptionID) (*dto.SubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error
	DeleteSubsription(ctx context.Context, subID *models.SubscriptionID) error
	ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) (*dto.SubscriptionListResponse, error)
	GetSubscriptionsTotalCost(ctx context.Context, filters *models.SubscriptionsTotalCostFilters) (*dto.TotalCostResponse, error)
}

type service struct {
	repo repository.Repository
}

func New(r repository.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateSubscription(ctx context.Context, sub *models.Subscription) (*dto.SubscriptionIDResponse, error) {
	sub.ID = uuid.New()

	subscriptionID, err := s.repo.CreateSubscription(ctx, sub)
	if err != nil {
		return nil, wrapError(err)
	}

	return dtomap.MapToSubscriptionIDResponse(subscriptionID), nil
}

func (s *service) GetSubscription(ctx context.Context, subID *models.SubscriptionID) (*dto.SubscriptionResponse, error) {
	subscription, err := s.repo.GetSubscription(ctx, subID)
	if err != nil {
		return nil, wrapError(err)
	}

	if subscription != nil {
		return dtomap.MapToSubscriptionResponse(subscription), nil
	}

	return nil, nil
}

func (s *service) UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error {
	return nil
}

func (s *service) DeleteSubsription(ctx context.Context, subID *models.SubscriptionID) error {
	err := s.repo.DeleteSubsription(ctx, subID)
	if err != nil {
		return wrapError(err)
	}
	return nil
}

func (s *service) ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) (*dto.SubscriptionListResponse, error) {
	return nil, nil
}

func (s *service) GetSubscriptionsTotalCost(ctx context.Context, filters *models.SubscriptionsTotalCostFilters) (*dto.TotalCostResponse, error) {
	return nil, nil
}
