package utils

import (
	"context"

	"github.com/notIdealen/aihwbe.git/internal/logger"
)

func GetLoggerFromRequestContext(ctx context.Context) *logger.Logger {
	logger, ok := ctx.Value("logger").(*logger.Logger)
	if !ok {
		panic("Logger in request context not found")
	}
	return logger
}
