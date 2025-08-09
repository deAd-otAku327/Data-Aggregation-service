package dtomap

import (
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"

	"github.com/google/uuid"
)

const outputTimeFormat = "01-2006"

func MapToSubscriptionIDResponse(subID uuid.UUID) *dto.SubscriptionIDResponse {
	return &dto.SubscriptionIDResponse{
		SubID: subID.String(),
	}
}

func MapToSubscriptionResponse(model *models.Subscription) *dto.SubscriptionResponse {
	return &dto.SubscriptionResponse{
		ID:          model.ID.String(),
		ServiceName: model.ServiceName,
		Price:       model.Price,
		StartDate:   model.StartDate.Format(outputTimeFormat),
		UserID:      model.UserID.String(),
		EndDate: func() *string {
			if model.EndDate != nil {
				strDate := model.EndDate.Format(outputTimeFormat)
				return &strDate
			}
			return nil
		}(),
	}
}

func MapToSubscriptionListResponse(models []*models.Subscription) *dto.SubscriptionListResponse {
	response := dto.SubscriptionListResponse{}
	for _, sub := range models {
		response.Subs = append(response.Subs, MapToSubscriptionResponse(sub))
	}
	return &response
}

func MapToTotalCostResponse(totalCost int) *dto.TotalCostResponse {
	return &dto.TotalCostResponse{
		TotalCost: totalCost,
	}
}

func MapToErrorResponse(err error, code int) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		Message: err.Error(),
		Code:    code,
	}
}
