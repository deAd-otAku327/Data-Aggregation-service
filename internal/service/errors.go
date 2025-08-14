package service

import "errors"

var (
	ErrSubscriptionActivePeriodInvalid = errors.New("there is the same subscription with active period overlapping")
	ErrSubscriptionEndDateInvalid      = errors.New("end date invalid, value must be after start date")
	ErrSubscriptionNotFound            = errors.New("subscription with provided id not found")
)
