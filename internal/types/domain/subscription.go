package domain

import (
	"context"
	"data-aggregation-service/internal/types/dto"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type SubscriptionController interface {
	HandleCreateSubscription() http.HandlerFunc
	HandleGetSubscription() http.HandlerFunc
	HandleUpdateSubscription() http.HandlerFunc
	HandleDeleteSubscription() http.HandlerFunc
	HandleListSubscriptions() http.HandlerFunc
	HandleGetSubscriptionsTotalCost() http.HandlerFunc
}

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, sub *Subscription) (*dto.SubscriptionIDResponse, error)
	GetSubscription(ctx context.Context, subID *SubscriptionID) (*dto.SubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, subID *SubscriptionID, patch *SubscriptionPatch) error
	DeleteSubsription(ctx context.Context, subID *SubscriptionID) error
	ListSubscriptions(ctx context.Context, filters *SubscriptionFilters) (*dto.SubscriptionListResponse, error)
	GetSubscriptionsTotalCost(ctx context.Context, filters *SubscriptionsTotalCostFilters) (*dto.TotalCostResponse, error)
}

type SubscriptionRepository interface {
	Insert(ctx context.Context, sub *Subscription) (*SubscriptionID, error)
	SelectByID(ctx context.Context, subID *SubscriptionID) (*Subscription, error)
	Update(ctx context.Context, subID *SubscriptionID, patch *SubscriptionPatch) error
	Delete(ctx context.Context, subID *SubscriptionID) error
	SelectList(ctx context.Context, filters *SubscriptionFilters) ([]*Subscription, error)
	SelectTotalCost(ctx context.Context, filters *SubscriptionsTotalCostFilters) (*SubscriptionsTotalCost, error)
}

type Subscription struct {
	ID          uuid.UUID
	ServiceName string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}

type SubscriptionID struct {
	SubID uuid.UUID
}

type SubscriptionPatch struct {
	Price   *int
	EndDate *time.Time
}

type SubscriptionFilters struct {
	UserID  *uuid.UUID
	Service *string
}

type SubscriptionsTotalCostFilters struct {
	FromDate   time.Time
	ToDate     time.Time
	SubFilters SubscriptionFilters
}

type SubscriptionsTotalCost struct {
	TotalCost int
}
