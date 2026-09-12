package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/notIdealen/aihwbe.git/internal/logger"
)

type Middleware func(http.Handler) http.Handler

func Logger(logger *logger.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "logger", logger)
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

type MyResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *MyResponseWriter) SetStatusCode(sc int) {
	rw.StatusCode = sc
}

func RequestID() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			XRequestID := r.Header.Get("X-Request-ID")
			if XRequestID == "" {
				id, err := uuid.NewV7()
				if err != nil {

				}
				XRequestID = id.String()
				r.Header.Set("X-Request-ID", XRequestID)
			}

			rw := &MyResponseWriter{ResponseWriter: w}

			logger := GetLoggerFromRequest(r).With(
				slog.String("Request_ID", XRequestID),
				slog.String("Method", r.Method),
				slog.String("Host", r.Host),
				slog.String("Path", r.URL.Path),
			)

			start := time.Now()
			logger.Debug(">>>")

			h.ServeHTTP(rw, r)

			logger.Debug("<<<",
				// slog.Int64("Duration", time.Since(start).Milliseconds()),
				slog.Float64("Duration", float64(time.Since(start).Milliseconds())/1000),
				slog.Int("StatusCode", rw.StatusCode),
			)
		})
	}
}

func WrapRouter(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; 0 <= i; i-- {
		h = mws[i](h)
	}
	return h
}

func GetLoggerFromRequest(r *http.Request) *logger.Logger {
	ctx := r.Context()
	logger, ok := ctx.Value("logger").(*logger.Logger)
	if !ok {
		//вернуть ошибку, но куда?
		return nil
	}
	return logger
}
