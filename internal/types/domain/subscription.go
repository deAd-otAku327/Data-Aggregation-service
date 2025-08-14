package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type SubscriptionsRepository interface {
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
