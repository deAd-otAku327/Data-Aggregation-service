package modelmap

import (
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"
	"time"

	"github.com/google/uuid"
)

// DTO validation exepts parse errors.

func MapToSubscription(request *dto.CreateSubscriptionRequest) *models.Subscription {
	userID, _ := uuid.Parse(request.UserID)
	stDate, _ := time.Parse("01-2006", request.StartDate)
	var endDate *time.Time
	if request.EndDate != nil {
		parsed, _ := time.Parse("01-2006", *request.EndDate)
		endDate = &parsed
	}
	return &models.Subscription{
		ServiceName: request.ServiceName,
		Price:       request.Price,
		UserID:      userID,
		StartDate:   stDate,
		EndDate:     endDate,
	}
}

func MapToSubscriptionPatch(request *dto.UpdateSubscriptionRequest) *models.SubscriptionPatch {
	subID, _ := uuid.Parse(request.SubID)
	var endDate *time.Time
	if request.EndDate != nil {
		parsed, _ := time.Parse("01-2006", *request.EndDate)
		endDate = &parsed
	}
	return &models.SubscriptionPatch{
		SubID:   subID,
		Price:   request.Price,
		EndDate: endDate,
	}
}

func MapToSubscriptionFilterParams(request *dto.ListSubscriptionsRequest) *models.SubscriptionFilterParams {
	var userID *uuid.UUID
	if request.UserID != nil {
		parsed, _ := uuid.Parse(*request.UserID)
		userID = &parsed
	}
	return &models.SubscriptionFilterParams{
		UserID:  userID,
		Service: request.ServiceName,
	}
}
