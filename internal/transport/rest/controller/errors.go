package controller

import "errors"

var (
	ErrParsingRequest     = errors.New("invalid request")
	ErrSomethingWentWrong = errors.New("sorry, something went wrong")
)
