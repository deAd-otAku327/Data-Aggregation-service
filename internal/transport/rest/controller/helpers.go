package controller

import (
	"data-aggregation-service/internal/service"
	"data-aggregation-service/pkg/apperrors"
	"log/slog"
	"net/http"
)

func resolveError(err error, logger *slog.Logger) (int, error) {
	if err != nil {
		if apperr, ok := err.(*apperrors.AppError); ok {
			apierr := apperr.GetAPIErr()

			switch apierr {
			case service.ErrInvalidPrice, service.ErrInvalidEndDate:
				logger.Warn("invalid value detected on repository layer: " + apierr.Error())
				return http.StatusBadRequest, apierr
			case service.ErrSubscriptionActivePeriodInvalid:
				return http.StatusBadRequest, apierr
			}
		}

		logger.Error(err.Error())
		return http.StatusInternalServerError, ErrSomethingWentWrong
	}

	logger.Error("unable to resolve nil error")
	return http.StatusInternalServerError, ErrSomethingWentWrong
}
