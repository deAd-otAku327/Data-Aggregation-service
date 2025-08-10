package dtomap

import (
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"
)

func MapToSubscriptionIDResponse(model *models.SubscriptionID) *dto.SubscriptionIDResponse {
	return &dto.SubscriptionIDResponse{
		SubID: model.SubID.String(),
	}
}

func MapToSubscriptionResponse(model *models.Subscription) *dto.SubscriptionResponse {
	return &dto.SubscriptionResponse{
		ID:          model.ID.String(),
		ServiceName: model.ServiceName,
		Price:       model.Price,
		StartDate:   model.StartDate.Format(dto.DateTimeLayout),
		UserID:      model.UserID.String(),
		EndDate: func() *string {
			if model.EndDate != nil {
				strDate := model.EndDate.Format(dto.DateTimeLayout)
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

func MapToErrorResponse(msgs []string, statusCode int) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		Code:     statusCode,
		Messages: msgs,
	}
}
