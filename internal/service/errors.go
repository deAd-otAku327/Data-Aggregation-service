package service

import "errors"

var (
	ErrInvalidPrice                    = errors.New("price value must be >= 0")
	ErrInvalidEndDate                  = errors.New("end date must be after start date")
	ErrSubscriptionActivePeriodInvalid = errors.New("there is the same subscription with active period overlapping")
)
