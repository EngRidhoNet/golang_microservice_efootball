package error

import "errors"

const (
	Success = "Success"
	Error   = "Error"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrSQLError		   = errors.New("database server failed to excetute the query")
	ErrTooManyRequests = errors.New("Too many requests")
	ErrUnauthorized = errors.New("Unauthorized")
	ErrInvalidToken = errors.New("Invalid token")
	ErrForbidden = errors.New("Forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}