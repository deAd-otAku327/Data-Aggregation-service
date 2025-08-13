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
	}

	return err
}
