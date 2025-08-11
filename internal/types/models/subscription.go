package models

import (
	"time"

	"github.com/google/uuid"
)

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
	SubID   uuid.UUID
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
