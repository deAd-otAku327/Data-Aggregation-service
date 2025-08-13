package postgres

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/repository/postgres/pgerrors"

	"github.com/lib/pq"
)

func catchPQErrors(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case pgconsts.ErrExclusionConstraintViolation:
			return pgerrors.ErrsExclusionViolation[pqErr.Constraint]
		}

	}
	return err
}
