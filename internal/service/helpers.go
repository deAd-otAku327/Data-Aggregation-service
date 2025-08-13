package service

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/repository/postgres/pgerrors"
	"data-aggregation-service/pkg/apperrors"
	"errors"
)

func wrapError(err error) error {
	if errors.Is(err, pgerrors.ErrsExclusionViolation[pgconsts.ConstraintExclusionNoOverlappingSubs]) {
		return apperrors.New(ErrSubscriptionActivePeriodInvalid, err)
	} else if errors.Is(err, pgerrors.ErrsCheckViolation[pgconsts.ConstraintCheckEndDateAfterStartDate]) {
		return apperrors.New(ErrSubscriptionEndDateInvalid, err)
	} else if errors.Is(err, pgerrors.ErrNoSubscription) {
		return apperrors.New(ErrSubscriptionNotFound, err)
	}

	return err
}
