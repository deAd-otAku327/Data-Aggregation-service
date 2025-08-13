package modelmap

import (
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"
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

func MapToSubscription(request *dto.CreateSubscriptionRequest) *models.Subscription {
	return &models.Subscription{
		ServiceName: request.ServiceName,
		Price:       request.Price,
		UserID:      *mustParseUUID(&request.UserID),
		StartDate:   *mustParseTime(&request.StartDate),
		EndDate:     mustParseTime(request.EndDate),
	}
}

func MapGetSubscriptionToSubscriptionID(request *dto.GetSubscriptionRequest) *models.SubscriptionID {
	return &models.SubscriptionID{
		SubID: *mustParseUUID(&request.SubID),
	}
}

func MapDeleteSubscriptionToSubscriptionID(request *dto.DeleteSubscriptionRequest) *models.SubscriptionID {
	return &models.SubscriptionID{
		SubID: *mustParseUUID(&request.SubID),
	}
}

func MapUpdateSubscriptionToSubscriptionID(request *dto.UpdateSubscriptionRequest) *models.SubscriptionID {
	return &models.SubscriptionID{
		SubID: *mustParseUUID(&request.SubID),
	}
}

func MapToSubscriptionPatch(request *dto.UpdateSubscriptionRequest) *models.SubscriptionPatch {
	return &models.SubscriptionPatch{
		Price:   request.Price,
		EndDate: mustParseTime(request.EndDate),
	}
}

func MapToSubscriptionFilterParams(request *dto.ListSubscriptionsRequest) *models.SubscriptionFilters {
	return &models.SubscriptionFilters{
		UserID:  mustParseUUID(request.UserID),
		Service: request.ServiceName,
	}
}

func MapToSubscriptionsTotalCostFilters(request *dto.GetTotalCostRequest) *models.SubscriptionsTotalCostFilters {
	return &models.SubscriptionsTotalCostFilters{
		FromDate: *mustParseTime(&request.FromDate),
		ToDate:   *mustParseTime(&request.ToDate),
		SubFilters: models.SubscriptionFilters{
			UserID:  mustParseUUID(request.UserID),
			Service: request.ServiceName,
		},
	}
}
