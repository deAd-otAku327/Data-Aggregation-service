package controller

import (
	"data-aggregation-service/internal/service/apierrors"
	"data-aggregation-service/pkg/apperrors"
	"log/slog"
	"net/http"
)

func resolveError(err error, logger *slog.Logger) (int, error) {
	if err != nil {
		if apperr, ok := err.(*apperrors.AppError); ok {
			apierr := apperr.GetAPIErr()

			switch apierr {
			case apierrors.ErrSubscriptionActivePeriodInvalid, apierrors.ErrSubscriptionEndDateInvalid:
				return http.StatusBadRequest, apierr
			case apierrors.ErrSubscriptionNotFound:
				return http.StatusNotFound, apierr
			}
		}

		logger.Error(err.Error())
		return http.StatusInternalServerError, apierrors.ErrSomethingWentWrong
	}

	logger.Error("unable to resolve nil error")
	return http.StatusInternalServerError, apierrors.ErrSomethingWentWrong
}
