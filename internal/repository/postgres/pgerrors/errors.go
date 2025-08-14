package pgerrors

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"errors"
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
