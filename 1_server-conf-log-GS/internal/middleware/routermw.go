package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/notIdealen/simple_server-config-logger/internal/logger"
)

type Middleware func(http.Handler) http.Handler

func Logger(logger *logger.Logger) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			ctx = context.WithValue(ctx, "logger", logger)
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
				slog.Int64("Duration", time.Since(start).Microseconds()),
				slog.Int("StatusCode", rw.StatusCode),
			)
			fmt.Println()
		})
	}
}

func Recover() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("RECOVER RUN")
					w.WriteHeader(http.StatusInternalServerError)
					// w.Write([]byte(fmt.Sprintf("some trouble with request: %v\n", err)))
					fmt.Fprintf(w, "some trouble with request: %v\n", err)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}
