package apperrors

import "errors"

type AppError struct {
	apierr error
	svcerr error
}

func New(apierr, svcerr error) error {
	return &AppError{
		apierr: apierr,
		svcerr: svcerr,
	}
}

func (se AppError) Error() string {
	return errors.Join(se.apierr, se.svcerr).Error()
}

func (se AppError) GetAPIErr() error {
	return se.apierr
}

func (se AppError) GetSvcErr() error {
	return se.svcerr
}
