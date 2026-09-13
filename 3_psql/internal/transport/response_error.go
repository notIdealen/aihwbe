package transport

import (
	"errors"
	"net/http"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	"github.com/notIdealen/aihwbe.git/internal/middleware"
)

// func writeResponseError(w *middleware.MyResponseWriter, err error) {
func writeResponseError(w *middleware.MyResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		w.SetStatusCode(http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
	case errors.Is(err, domain.ErrGatewayTimeout):
		w.SetStatusCode(http.StatusNotFound)
		w.WriteHeader(http.StatusGatewayTimeout)
		// w.Write([]byte(""))
	}
}
