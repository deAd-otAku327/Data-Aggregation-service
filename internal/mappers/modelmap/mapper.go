package modelmap

import (
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/internal/types/dto"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

// DTO validation must exepts parse errors.

func mustParseTime(val *string) *time.Time {
	if val == nil {
		return nil
	}
	date, err := time.Parse(dto.DateTimeLayout, *val)
	if err != nil {
		slog.Error("validation and mapping contract violated")
		panic(err)
	}
	return &date
}

func mustParseUUID(val *string) *uuid.UUID {
	if val == nil {
		return nil
	}
	uuid, err := uuid.Parse(*val)
	if err != nil {
		slog.Error("validation and mapping contract violated")
		panic(err)
	}
	return &uuid
}

func MapToSubscription(request *dto.CreateSubscriptionRequest) *domain.Subscription {
	return &domain.Subscription{
		ServiceName: request.ServiceName,
		Price:       request.Price,
		UserID:      *mustParseUUID(&request.UserID),
		StartDate:   *mustParseTime(&request.StartDate),
		EndDate:     mustParseTime(request.EndDate),
	}
}

func MapGetSubscriptionToSubscriptionID(request *dto.GetSubscriptionRequest) *domain.SubscriptionID {
	return &domain.SubscriptionID{
		SubID: *mustParseUUID(&request.SubID),
	}
}

func MapDeleteSubscriptionToSubscriptionID(request *dto.DeleteSubscriptionRequest) *domain.SubscriptionID {
	return &domain.SubscriptionID{
		SubID: *mustParseUUID(&request.SubID),
	}
}

func MapUpdateSubscriptionToSubscriptionID(request *dto.UpdateSubscriptionRequest) *domain.SubscriptionID {
	return &domain.SubscriptionID{
		SubID: *mustParseUUID(&request.SubID),
	}
}

func MapToSubscriptionPatch(request *dto.UpdateSubscriptionRequest) *domain.SubscriptionPatch {
	return &domain.SubscriptionPatch{
		Price:   request.Price,
		EndDate: mustParseTime(request.EndDate),
	}
}

func MapToSubscriptionFilterParams(request *dto.ListSubscriptionsRequest) *domain.SubscriptionFilters {
	return &domain.SubscriptionFilters{
		UserID:  mustParseUUID(request.UserID),
		Service: request.ServiceName,
	}
}

func MapToSubscriptionsTotalCostFilters(request *dto.GetTotalCostRequest) *domain.SubscriptionsTotalCostFilters {
	return &domain.SubscriptionsTotalCostFilters{
		FromDate: *mustParseTime(&request.FromDate),
		ToDate:   *mustParseTime(&request.ToDate),
		SubFilters: domain.SubscriptionFilters{
			UserID:  mustParseUUID(request.UserID),
			Service: request.ServiceName,
		},
	}
}
