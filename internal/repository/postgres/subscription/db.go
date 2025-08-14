package subscription

import (
	"context"
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/repository/postgres/pgerrors"
	"data-aggregation-service/internal/types/domain"
	"data-aggregation-service/pkg/db"
	"data-aggregation-service/pkg/migrator"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type subsRepositoryImpl struct {
	db *sql.DB
}

func New(cfg config.SubsRepo) *subsRepositoryImpl {
	database := db.MustConnectDB(cfg.DriverName, cfg.URI, cfg.MaxOpenConns)
	migrator.MustApplyMigrations(database, cfg.MigrationsDir)

	return &subsRepositoryImpl{
		db: database,
	}
}

func (r *subsRepositoryImpl) Insert(ctx context.Context, sub *domain.Subscription) (*domain.SubscriptionID, error) {
	query, args, err := sq.Insert(pgconsts.SubscriptionsTable).
		Columns(
			pgconsts.SubscriptionsPublicID, pgconsts.SubscriptionsServiceName, pgconsts.SubscriptionsPrice,
			pgconsts.SubscriptionsUserID, pgconsts.SubscriptionsStartDate, pgconsts.SubscriptionsEndDate,
		).
		Values(sub.ID, sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate).
		Suffix(fmt.Sprintf("RETURNING %s", pgconsts.SubscriptionsPublicID)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	subscriptionID := domain.SubscriptionID{}

	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&subscriptionID.SubID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, pgerrors.CatchPQErrors(err))
	}

	return &subscriptionID, nil
}

func (r *subsRepositoryImpl) SelectByID(ctx context.Context, subscriptionID *domain.SubscriptionID) (*domain.Subscription, error) {
	query, args, err := sq.Select(
		pgconsts.SubscriptionsPublicID, pgconsts.SubscriptionsServiceName, pgconsts.SubscriptionsPrice,
		pgconsts.SubscriptionsUserID, pgconsts.SubscriptionsStartDate, pgconsts.SubscriptionsEndDate,
	).
		From(pgconsts.SubscriptionsTable).
		Where(sq.Eq{pgconsts.SubscriptionsPublicID: subscriptionID.SubID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	var subscription domain.Subscription

	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&subscription.ID, &subscription.ServiceName, &subscription.Price,
		&subscription.UserID, &subscription.StartDate, &subscription.EndDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pgerrors.ErrNoSubscription
		}
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, pgerrors.CatchPQErrors(err))
	}

	return &subscription, nil
}

func (r *subsRepositoryImpl) Update(ctx context.Context, subscriptionID *domain.SubscriptionID, patch *domain.SubscriptionPatch) error {
	queryCore := sq.Update(pgconsts.SubscriptionsTable)

	query, args, err := applySubscriptionUpdateValues(queryCore, patch).
		Where(sq.Eq{pgconsts.SubscriptionsPublicID: subscriptionID.SubID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, pgerrors.CatchPQErrors(err))
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, err)
	}

	if affected == 0 {
		return pgerrors.ErrNoSubscription
	}

	return nil
}

func (r *subsRepositoryImpl) Delete(ctx context.Context, subscriptionID *domain.SubscriptionID) error {
	query, args, err := sq.Delete(pgconsts.SubscriptionsTable).
		Where(sq.Eq{pgconsts.SubscriptionsPublicID: subscriptionID.SubID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, pgerrors.CatchPQErrors(err))
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, err)
	}

	if affected == 0 {
		return pgerrors.ErrNoSubscription
	}

	return nil
}

func (r *subsRepositoryImpl) SelectList(ctx context.Context, filters *domain.SubscriptionFilters) ([]*domain.Subscription, error) {
	queryCore := sq.Select(
		pgconsts.SubscriptionsPublicID, pgconsts.SubscriptionsServiceName, pgconsts.SubscriptionsPrice,
		pgconsts.SubscriptionsUserID, pgconsts.SubscriptionsStartDate, pgconsts.SubscriptionsEndDate,
	).From(pgconsts.SubscriptionsTable)

	query, args, err := applySubscriptionsListFilters(queryCore, filters).
		OrderBy(pgconsts.SubscriptionsStartDate).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, err)
	}
	defer rows.Close()

	subs := make([]*domain.Subscription, 0)
	for rows.Next() {
		subscription := domain.Subscription{}
		err := rows.Scan(&subscription.ID, &subscription.ServiceName, &subscription.Price,
			&subscription.UserID, &subscription.StartDate, &subscription.EndDate)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, err)
		}
		subs = append(subs, &subscription)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, err)
	}

	return subs, nil
}

func (r *subsRepositoryImpl) SelectTotalCost(ctx context.Context, filters *domain.SubscriptionsTotalCostFilters) (*domain.SubscriptionsTotalCost, error) {
	// $1 - EndDate param, $2 - StartDate param.
	// Ordering from WHERE statements.
	ageFunc := fmt.Sprintf("AGE(LEAST(%s, $1), GREATEST(%s, $2))",
		pgconsts.SubscriptionsEndDate, pgconsts.SubscriptionsStartDate,
	)
	sumFunc := fmt.Sprintf("COALESCE(SUM(%s * (EXTRACT(YEAR FROM %s) * 12 + EXTRACT(MONTH FROM %s))), 0) as total_cost",
		pgconsts.SubscriptionsPrice, ageFunc, ageFunc,
	)

	queryCore := sq.Select(sumFunc).
		From(pgconsts.SubscriptionsTable).
		// Visual expression: (period_start < end_date <= period_end) || (period_start <= start_date < period_end).
		Where(sq.Or{
			sq.And{
				sq.LtOrEq{pgconsts.SubscriptionsEndDate: filters.ToDate},
				sq.Gt{pgconsts.SubscriptionsEndDate: filters.FromDate},
			},
			sq.And{
				sq.GtOrEq{pgconsts.SubscriptionsStartDate: filters.FromDate},
				sq.Lt{pgconsts.SubscriptionsStartDate: filters.ToDate},
			},
		})

	query, args, err := applySubscriptionsListFilters(queryCore, &filters.SubFilters).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	subTotalCost := domain.SubscriptionsTotalCost{}

	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&subTotalCost.TotalCost)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, pgerrors.CatchPQErrors(err))
	}

	return &subTotalCost, nil
}
