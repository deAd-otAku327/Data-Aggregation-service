package controller

import (
	"data-aggregation-service/internal/transport/rest/controller/subscription"
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/internal/validation"
	"log/slog"
)

func NewSubsController(s domain.SubscriptionService, v *validation.Validation, l *slog.Logger) domain.SubscriptionController {
	return subscription.New(s, v, l)
}
