package transport

import (
	"errors"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	middleware_router "github.com/notIdealen/aihwbe.git/internal/middleware/router"
)

func ParseResponseError(err error, w *middleware_router.MyResponseWriter) string {
	// respErr := err.(*domain.ResponseError)
	if re, ok := errors.AsType[*domain.ResponseError](err); ok {
		switch re.Code {
		case domain.ErrBadRequest:
			w.StatusCode = 400
			w.Log = re.Log
			return re.Message
		case domain.ErrNotFound:
			w.StatusCode = 404
			w.Log = re.Log
			return re.Message
		case domain.ErrInternalServer:
			w.StatusCode = 500
			w.Log = re.Log
			return re.Message
		}
		w.StatusCode = 400
		w.Log = re.Log
		return re.Message
	}
	return ""
}
