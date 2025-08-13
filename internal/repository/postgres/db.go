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
		return nil, fmt.Errorf("%w: %w", pgerrors.ErrQueryExec, err)
	}

	return &subscription, nil
}

func (r *postgresRepo) UpdateSubscription(ctx context.Context, patch *models.SubscriptionPatch) error {
	return nil
}

func (r *postgresRepo) DeleteSubsription(ctx context.Context, subID *models.SubscriptionID) error {
	return nil
}

func (r *postgresRepo) ListSubscriptions(ctx context.Context, filters *models.SubscriptionFilters) ([]*models.Subscription, error) {
	return nil, nil
}

func (r *postgresRepo) GetSubscriptionsTotalCost(ctx context.Context, filters *models.SubscriptionsTotalCostFilters) (*models.SubscriptionsTotalCost, error) {
	return nil, nil
}
