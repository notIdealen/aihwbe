package middleware_router

import (
	"context"
	"net/http"

	"github.com/notIdealen/aihwbe.git/internal/logger"
	"github.com/notIdealen/aihwbe.git/internal/middleware"
)

func Logger(logger *logger.Logger) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "logger", logger)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}
