package dtomap

import (
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/internal/types/dto"
)

func MapToSubscriptionIDResponse(model *domain.SubscriptionID) *dto.SubscriptionIDResponse {
	return &dto.SubscriptionIDResponse{
		SubID: model.SubID.String(),
	}
}

func MapToSubscriptionResponse(model *domain.Subscription) *dto.SubscriptionResponse {
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

func MapToSubscriptionListResponse(models []*domain.Subscription) *dto.SubscriptionListResponse {
	response := dto.SubscriptionListResponse{}
	for _, sub := range models {
		response.Subs = append(response.Subs, MapToSubscriptionResponse(sub))
	}
	return &response
}

func MapToTotalCostResponse(model *domain.SubscriptionsTotalCost) *dto.TotalCostResponse {
	return &dto.TotalCostResponse{
		TotalCost: model.TotalCost,
	}
}

func MapToErrorResponse(msgs []string, statusCode int) *dto.ErrorResponse {
	return &dto.ErrorResponse{
		Code:     statusCode,
		Messages: msgs,
	}
}
