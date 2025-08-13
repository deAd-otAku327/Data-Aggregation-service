package service

import "errors"

var (
	ErrSubscriptionActivePeriodInvalid = errors.New("there is the same subscription with active period overlapping")
	ErrSubscriptionNotFound            = errors.New("subscription with provided id not found")
)
