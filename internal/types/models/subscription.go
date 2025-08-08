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

type SubscriptionPatch struct {
	SubID   uuid.UUID
	Price   *int
	EndDate *time.Time
}

type SubscriptionFilters struct {
	UserID  *uuid.UUID
	Service *string
}
