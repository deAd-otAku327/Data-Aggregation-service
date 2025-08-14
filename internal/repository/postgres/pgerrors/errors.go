package pgerrors

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

var (
	ErrQueryBuilding = errors.New("sql query building failed")
	ErrQueryExec     = errors.New("sql query execution failed")

	ErrsExclusionViolation = map[string]error{
		pgconsts.ConstraintExclusionNoOverlappingSubs: errors.New("no_overlapping_subscriptions exclusion violated"),
	}

	ErrsCheckViolation = map[string]error{
		pgconsts.ConstraintCheckEndDateAfterStartDate: errors.New("end_date_after_start_date check violated"),
	}

	ErrNoSubscription = errors.New("no subscription with provided id")
)

func CatchPQErrors(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case pgconsts.ErrExclusionConstraintViolation:
			return fmt.Errorf("%w: %w", ErrsExclusionViolation[pqErr.Constraint], err)
		case pgconsts.ErrCheckConstraintViolation:
			return fmt.Errorf("%w: %w", ErrsCheckViolation[pqErr.Constraint], err)
		}

	}
	return err
}
