package postgres

import (
	"context"
	"data-aggregation-service/internal/config"
	"data-aggregation-service/internal/repository/postgres/pgconsts"
	"data-aggregation-service/internal/repository/postgres/pgerrors"
	"data-aggregation-service/internal/types/models"
	"data-aggregation-service/pkg/db"
	"data-aggregation-service/pkg/migrator"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type postgresRepo struct {
	db *sql.DB
}

func New(cfg config.PostgresDB) *postgresRepo {
	database := db.MustConnectDB(cfg.DriverName, cfg.URI, cfg.MaxOpenConns)
	migrator.MustApplyMigrations(database, cfg.MigrationsDir)

	return &postgresRepo{
		db: database,
	}
}

func (r *postgresRepo) CreateSubscription(ctx context.Context, sub *models.Subscription) (*models.SubscriptionID, error) {
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

	subscriptionID := models.SubscriptionID{}

	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&subscriptionID.SubID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, catchPQErrors(err))
	}

	return &subscriptionID, nil
}

func (r *postgresRepo) GetSubscription(ctx context.Context, subscriptionID *models.SubscriptionID) (*models.Subscription, error) {
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

	var subscription models.Subscription

	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&subscription.ID, &subscription.ServiceName, &subscription.Price,
		&subscription.UserID, &subscription.StartDate, &subscription.EndDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pgerrors.ErrNoSubscription
		}
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, catchPQErrors(err))
	}

	return &subscription, nil
}

func (r *postgresRepo) UpdateSubscription(ctx context.Context, subscriptionID *models.SubscriptionID, patch *models.SubscriptionPatch) error {
	queryCore := sq.Update(pgconsts.SubscriptionsTable)

	query, args, err := applySubscriptionUpdateValues(queryCore, patch).
		Where(sq.Eq{pgconsts.SubscriptionsPublicID: subscriptionID.SubID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, catchPQErrors(err))
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

func (r *postgresRepo) DeleteSubsription(ctx context.Context, subscriptionID *models.SubscriptionID) error {
	query, args, err := sq.Delete(pgconsts.SubscriptionsTable).
		Where(sq.Eq{pgconsts.SubscriptionsPublicID: subscriptionID.SubID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryBuilding, err)
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, catchPQErrors(err))
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

func (r *postgresRepo) ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) ([]*models.Subscription, error) {
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

	subs := make([]*models.Subscription, 0)
	for rows.Next() {
		subscription := models.Subscription{}
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

func (r *postgresRepo) GetSubscriptionsTotalCost(ctx context.Context, filters *models.SubscriptionsTotalCostFilters) (*models.SubscriptionsTotalCost, error) {
	// $1 - EndDate param, $2 - StartDate param.
	// Ordering from WHERE statements.
	ageFunc := fmt.Sprintf("AGE(LEAST(%s, $1), GREATEST(%s, $2))",
		pgconsts.SubscriptionsEndDate, pgconsts.SubscriptionsStartDate,
	)
	sumFunc := fmt.Sprintf("SUM(%s * (EXTRACT(YEAR FROM %s) * 12 + EXTRACT(MONTH FROM %s))) as total_cost",
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

	subTotalCost := models.SubscriptionsTotalCost{}

	row := r.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&subTotalCost.TotalCost)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, catchPQErrors(err))
	}

	return &subTotalCost, nil
}
