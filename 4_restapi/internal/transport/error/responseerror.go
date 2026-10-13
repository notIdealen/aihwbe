package transport_error

import (
	"errors"
	"net/http"

	service_error "github.com/notIdealen/aihwbe.git/internal/service/error"
)

type ResponseError struct {
	StatusCode int
	Message    string
	Log        string
}

func (re *ResponseError) Error() string {
	return re.Log
}

func ConcertToResponseError(err error) *ResponseError {
	var se *service_error.ServiceError
	if errors.As(err, &se) {
		switch se.C {
		case service_error.ErrValidation:
			return &ResponseError{
				StatusCode: http.StatusBadRequest,
				Message:    se.M,
				Log:        se.L,
			}
		}
	}
	return &ResponseError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Server error",
		Log:        err.Error(),
	}
}
