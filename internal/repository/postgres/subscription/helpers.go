package subscription

import (
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/types/domain"

	sq "github.com/Masterminds/squirrel"
)

func applySubscriptionUpdateValues(query sq.UpdateBuilder, patch *domain.SubscriptionPatch) sq.UpdateBuilder {
	if patch.Price != nil {
		query = query.Set(pgconsts.SubscriptionsPrice, patch.Price)
	}

	if patch.EndDate != nil {
		query = query.Set(pgconsts.SubscriptionsEndDate, patch.EndDate)
	}

	return query
}

func applySubscriptionsListFilters(query sq.SelectBuilder, filters *domain.SubscriptionFilters) sq.SelectBuilder {
	if filters.Service != nil {
		query = query.Where(sq.Eq{pgconsts.SubscriptionsServiceName: filters.Service})
	}

	if filters.UserID != nil {
		query = query.Where(sq.Eq{pgconsts.SubscriptionsUserID: filters.UserID})
	}

	return query
}
