package pgerrors

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"errors"
)

var (
	ErrQueryBuilding = errors.New("sql query building failed")
	ErrQueryExec     = errors.New("sql query execution failed")

	ErrsExclusionViolation = map[string]error{
		pgconsts.ConstraintExclusionNoOverlappingSubs: errors.New("no overlapping subscriptions exclusion violated"),
	}
)
