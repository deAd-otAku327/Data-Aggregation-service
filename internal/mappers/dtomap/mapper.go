package dtomap

import (
	"data-aggregation-service/internal/types/dto"
	"data-aggregation-service/internal/types/models"
)

const outputTimeFormat = "01-2006"

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
