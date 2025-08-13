package postgres

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/repository/postgres/pgerrors"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

func toNullableTime(t *time.Time) *sql.NullTime {
	nullable := sql.NullTime{}
	if t != nil {
		nullable.Time = *t
		nullable.Valid = true
	}

	return &nullable
}

func catchPQErrors(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case pgconsts.ErrCheckConstraintViolation:
			return pgerrors.ErrsCheckViolation[pqErr.Constraint]

		case pgconsts.ErrExclusionConstraintViolation:
			return pgerrors.ErrsExclusionViolation[pqErr.Constraint]
		}

	}
	return err
}
