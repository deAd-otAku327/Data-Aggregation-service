package service

import (
	"context"
	"data-aggregation-service/internal/repository"
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"
	"fmt"
)

type Service interface {
	CreateSubscription(ctx context.Context, sub *models.Subscription) (*dto.SubscriptionIDResponse, error)
	GetSubscription(ctx context.Context, subID *models.SubscriptionID) (*dto.SubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error
	DeleteSubsription(ctx context.Context, subID *models.SubscriptionID) error
	ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) (*dto.SubscriptionListResponse, error)
	GetTotalCost(ctx context.Context, filters *models.TotalCostFilters) (*dto.TotalCostResponse, error)
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
	fmt.Println(sub)
	return nil, nil
}

func (s *service) GetSubscription(ctx context.Context, subID *models.SubscriptionID) (*dto.SubscriptionResponse, error) {
	return nil, nil
}

func (s *service) UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error {
	return nil
}

func (s *service) DeleteSubsription(ctx context.Context, subID *models.SubscriptionID) error {
	return nil
}

func (s *service) ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) (*dto.SubscriptionListResponse, error) {
	return nil, nil
}

func (s *service) GetTotalCost(ctx context.Context, filters *models.TotalCostFilters) (*dto.TotalCostResponse, error) {
	return nil, nil
}
