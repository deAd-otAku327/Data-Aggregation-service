package postgres

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/repository/postgres/pgerrors"
	"data-aggregation-service/internal/types/models"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

func catchPQErrors(err error) error {
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case pgconsts.ErrExclusionConstraintViolation:
			return fmt.Errorf("%w: %w", pgerrors.ErrsExclusionViolation[pqErr.Constraint], err)
		case pgconsts.ErrCheckConstraintViolation:
			return fmt.Errorf("%w: %w", pgerrors.ErrsCheckViolation[pqErr.Constraint], err)
		}

	}
	return err
}

func applySubscriptionUpdateValues(query sq.UpdateBuilder, patch *models.SubscriptionPatch) sq.UpdateBuilder {
	if patch.Price != nil {
		query = query.Set(pgconsts.SubscriptionsPrice, patch.Price)
	}

	if patch.EndDate != nil {
		query = query.Set(pgconsts.SubscriptionsEndDate, patch.EndDate)
	}

	return query
}
