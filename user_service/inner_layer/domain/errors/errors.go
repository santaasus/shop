package errors

import "errors"

const (
	Unauthorized        = "Unauthorized"
	UnauthorizedMessage = "unauthorized"

	ValidationError        = "ValidationError"
	ValidationErrorMessage = "validation error"

	InternalServerError        = "InternalServerError"
	InternalServerErrorMessage = "internal server error"

	NotFound        = "NotFound"
	NotFoundMessage = "object not found"

	NotAuthenticated             = "NotAuthenticated"
	NotAuthenticatedErrorMessage = "no authenticated"

	UnknownError        = "UnknownError"
	unknownErrorMessage = "something went wrong"
)

type AppError struct {
	Err  error
	Type string
}

func ThrowAppErrorWith(errorType string) *AppError {
	var err error

	switch errorType {
	case Unauthorized:
		err = errors.New(UnauthorizedMessage)
	case ValidationError:
		err = errors.New(ValidationErrorMessage)
	case InternalServerError:
		err = errors.New(InternalServerErrorMessage)
	case NotFound:
		err = errors.New(NotFoundMessage)
	case UnknownError:
		err = errors.New(unknownErrorMessage)
	}

	return &AppError{err, errorType}
}

func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}
