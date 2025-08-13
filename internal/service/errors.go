package service

import "errors"

var (
	ErrSubscriptionActivePeriodInvalid = errors.New("there is the same subscription with active period overlapping")
)
