package middleware_router

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/notIdealen/aihwbe.git/internal/middleware"
)

func RequestID() middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			XRequestID := r.Header.Get("X-Request-ID")
			if XRequestID == "" {
				id, err := uuid.NewV7()
				if err != nil {
					// ошибка 5хх
					return
				}
				XRequestID = id.String()
				r.Header.Set("X-Request-ID", XRequestID)
			}
			h.ServeHTTP(w, r)
		})
	}
}
