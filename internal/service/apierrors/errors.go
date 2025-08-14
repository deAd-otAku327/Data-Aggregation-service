package apierrors

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/repository/postgres/pgerrors"
	"data-aggregation-service/pkg/apperrors"
	"errors"
)

var (
	ErrParsingRequest     = errors.New("invalid request")
	ErrSomethingWentWrong = errors.New("sorry, something went wrong")

	ErrSubscriptionActivePeriodInvalid = errors.New("there is the same subscription with active period overlapping")
	ErrSubscriptionEndDateInvalid      = errors.New("end date invalid, value must be after start date")
	ErrSubscriptionNotFound            = errors.New("subscription with provided id not found")
)

func WrapWithApiError(err error) error {
	if errors.Is(err, pgerrors.ErrsExclusionViolation[pgconsts.ConstraintExclusionNoOverlappingSubs]) {
		return apperrors.New(ErrSubscriptionActivePeriodInvalid, err)
	} else if errors.Is(err, pgerrors.ErrsCheckViolation[pgconsts.ConstraintCheckEndDateAfterStartDate]) {
		return apperrors.New(ErrSubscriptionEndDateInvalid, err)
	} else if errors.Is(err, pgerrors.ErrNoSubscription) {
		return apperrors.New(ErrSubscriptionNotFound, err)
	}

	return err
}
