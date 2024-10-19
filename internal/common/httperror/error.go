package httperror

import (
	"errors"
	"net/http"
)

var HTTPErrorMap = make(map[error]int)

type HTTPError struct {
	statusCode int
	err        error
}

func New(statusCode int, err error) *HTTPError {
	return &HTTPError{
		statusCode: statusCode,
		err:        err,
	}
}

func (e *HTTPError) Error() string {
	return e.err.Error()
}

func GetStatusCode(err error) int {
	if val, ok := err.(*HTTPError); ok {
		return val.statusCode
	}

	if statusCode, ok := HTTPErrorMap[err]; ok {
		return statusCode
	}

	if unwrap := errors.Unwrap(err); unwrap != nil {
		return GetStatusCode(unwrap)
	}

	return http.StatusInternalServerError
}
