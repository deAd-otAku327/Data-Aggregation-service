package apperrors

import "errors"

var ErrNoErr = errors.New("no error found")

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
	if se.apierr != nil {
		return se.apierr
	}
	return ErrNoErr
}

func (se AppError) GetSvcErr() error {
	if se.svcerr != nil {
		return se.svcerr
	}
	return ErrNoErr
}
