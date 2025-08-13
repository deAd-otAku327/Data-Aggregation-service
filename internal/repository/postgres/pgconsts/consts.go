package pgconsts

const (
	// Violation code translations.
	// From http://www.postgresql.org/docs/9.3/static/errcodes-appendix.html
	ErrCheckConstraintViolation     = "check_violation"
	ErrExclusionConstraintViolation = "exclusion_violation"

	// "Subscriptions" table constraints.
	ConstraintCheckValidPriceValue       = "valid_price_value"
	ConstraintCheckEndDateAfterStartDate = "end_date_after_start_date"
	ConstraintExclusionNoOverlappingSubs = "no_overlapping_subscriptions"

	// "Subscriptions" table namings.
	SubscriptionsTable = "subscriptions"

	SubscriptionsPublicID    = "public_id"
	SubscriptionsServiceName = "service_name"
	SubscriptionsPrice       = "price"
	SubscriptionsUserID      = "user_id"
	SubscriptionsStartDate   = "start_date"
	SubscriptionsEndDate     = "end_date"
)
