package middleware_router

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	"github.com/notIdealen/aihwbe.git/internal/logger"
	"github.com/notIdealen/aihwbe.git/internal/middleware"
)

func SetWriterAndWriteLog() middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger, ok := ctx.Value("logger").(*logger.Logger)
			if !ok {
				// ошибка 5хх
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			wlogger := logger.With(
				slog.String("Request_ID", r.Header.Get("X-Request-ID")),
				slog.String("Method", r.Method),
				slog.String("Host", r.Host),
				slog.String("Path", r.URL.Path),
			)
			wlogger.Info(">>>")

			mrw := &domain.MyResponseWriter{ResponseWriter: w, StatusCode: -1, Log: ""}
			start := time.Now()
			h.ServeHTTP(mrw, r)

			logger.Info("<<<",
				slog.Float64("Duration", float64(time.Since(start).Milliseconds())/1000),
				slog.Int("StatusCode", mrw.StatusCode),
				slog.String("Text error", mrw.Log),
			)
		})
	}
}
