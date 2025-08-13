package pgerrors

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"errors"
)

var (
	ErrQueryBuilding = errors.New("sql query building failed")
	ErrQueryExec     = errors.New("sql query execution failed")

	ErrsCheckViolation = map[string]error{
		pgconsts.ConstraintCheckValidPriceValue:       errors.New("valid price value check violated"),
		pgconsts.ConstraintCheckEndDateAfterStartDate: errors.New("end date after start date check violated"),
	}

	ErrsExclusionViolation = map[string]error{
		pgconsts.ConstraintExclusionNoOverlappingSubs: errors.New("no overlapping subscriptions exclusion violated"),
	}
)
