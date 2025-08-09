package service

import (
	"context"
	"data-aggregation-service/internal/repository"
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateSubscription(ctx context.Context, sub *models.Subscription) (*dto.SubscriptionIDResponse, error)
	GetSubscription(ctx context.Context, subID uuid.UUID) (*dto.SubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error
	DeleteSubsription(ctx context.Context, subID uuid.UUID) error
	ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) (*dto.SubscriptionListResponse, error)
	GetTotalCost(ctx context.Context, periodStart, periodEnd time.Time, filters *models.SubscriptionFilters) (*dto.TotalCostResponse, error)
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
	return nil, nil
}

func (s *service) GetSubscription(ctx context.Context, subID uuid.UUID) (*dto.SubscriptionResponse, error) {
	return nil, nil
}

func (s *service) UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error {
	return nil
}

func (s *service) DeleteSubsription(ctx context.Context, subID uuid.UUID) error {
	return nil
}

func (s *service) ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) (*dto.SubscriptionListResponse, error) {
	return nil, nil
}

func (s *service) GetTotalCost(ctx context.Context, periodStart, periodEnd time.Time, filters *models.SubscriptionFilters) (*dto.TotalCostResponse, error) {
	return nil, nil
}
