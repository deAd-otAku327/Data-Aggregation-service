package subscription

import (
	"context"
	"data-aggregation-service/internal/mappers/dtomap"
	"data-aggregation-service/internal/service/apierrors"
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/internal/types/dto"

	"github.com/google/uuid"
)

type subsServiceImpl struct {
	repo domain.SubscriptionRepository
}

func New(r domain.SubscriptionRepository) domain.SubscriptionService {
	return &subsServiceImpl{
		repo: r,
	}
}

func (s *subsServiceImpl) CreateSubscription(ctx context.Context, sub *domain.Subscription) (*dto.SubscriptionIDResponse, error) {
	sub.ID = uuid.New()

	subscriptionID, err := s.repo.Insert(ctx, sub)
	if err != nil {
		return nil, apierrors.WrapWithApiError(err)
	}

	return dtomap.MapToSubscriptionIDResponse(subscriptionID), nil
}

func (s *subsServiceImpl) GetSubscription(ctx context.Context, subID *domain.SubscriptionID) (*dto.SubscriptionResponse, error) {
	subscription, err := s.repo.SelectByID(ctx, subID)
	if err != nil {
		return nil, apierrors.WrapWithApiError(err)
	}

	return dtomap.MapToSubscriptionResponse(subscription), nil

}

func (s *subsServiceImpl) UpdateSubscription(ctx context.Context, subscriptionID *domain.SubscriptionID, patch *domain.SubscriptionPatch) error {
	err := s.repo.Update(ctx, subscriptionID, patch)
	if err != nil {
		return apierrors.WrapWithApiError(err)
	}
	return nil
}

func (s *subsServiceImpl) DeleteSubsription(ctx context.Context, subscriptionID *domain.SubscriptionID) error {
	err := s.repo.Delete(ctx, subscriptionID)
	if err != nil {
		return apierrors.WrapWithApiError(err)
	}

	return nil
}

func (s *subsServiceImpl) ListSubscriptions(ctx context.Context, filters *domain.SubscriptionFilters) (*dto.SubscriptionListResponse, error) {
	subs, err := s.repo.SelectList(ctx, filters)
	if err != nil {
		return nil, apierrors.WrapWithApiError(err)
	}

	return dtomap.MapToSubscriptionListResponse(subs), nil
}

func (s *subsServiceImpl) GetSubscriptionsTotalCost(ctx context.Context, filters *domain.SubscriptionsTotalCostFilters) (*dto.TotalCostResponse, error) {
	totalCost, err := s.repo.SelectTotalCost(ctx, filters)
	if err != nil {
		return nil, apierrors.WrapWithApiError(err)
	}

	return dtomap.MapToTotalCostResponse(totalCost), nil
}
