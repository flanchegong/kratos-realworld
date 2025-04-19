package errors

import (
	"errors"
	"fmt"
)

func NewHTTPError(code int, field, detail string) *HTTPError {

	return &HTTPError{
		Code: code,
		Errors: map[string][]string{
			field: {detail},
		},
	}
}

type HTTPError struct {
	Errors map[string][]string `json:"errors"`
	Code   int                 `json:"-"` // http status code
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError: %d", e.Code)
}

// FromError try to convert an error to *HTTPError.
func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	return &HTTPError{}
}
