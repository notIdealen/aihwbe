package middleware

import (
	"net/http"

	"github.com/notIdealen/simple_server-config-logger/internal/logger"
)

func Enrich(h http.Handler, mws ...Middleware) http.Handler {
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
