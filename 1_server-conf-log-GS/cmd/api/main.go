package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/notIdealen/simple_server-config-logger/internal/handler"
	"github.com/notIdealen/simple_server-config-logger/internal/httpserver"
	"github.com/notIdealen/simple_server-config-logger/internal/logger"
	"github.com/notIdealen/simple_server-config-logger/internal/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}

	interruptctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	l := logger.NewLogger().
		WithLevelFromConfig().
		// WithFileHandlerFromConfig().
		WithSTDOutHandler()
	l.Exec()

	l.Debug("app start")

	routes := handler.GetRoutes()
	router := httpserver.NewRouter()
	router.RegistRoutes(routes)

	handler := middleware.Enrich(router,
		middleware.Logger(l),
		middleware.RequestID(),
		middleware.Recover(),
	)

	server := httpserver.NewHTTPServer(handler, l)
	server.Run(interruptctx)
	l.Info("app end")
}
