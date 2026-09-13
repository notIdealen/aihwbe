package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/notIdealen/aihwbe.git/internal/logger"
	"github.com/notIdealen/aihwbe.git/internal/middleware"
	"github.com/notIdealen/aihwbe.git/internal/repo"
	"github.com/notIdealen/aihwbe.git/internal/server"
	"github.com/notIdealen/aihwbe.git/internal/service"
	"github.com/notIdealen/aihwbe.git/internal/transport"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Not found .env file")
	}

	logger := logger.NewBlank(logger.NewConfigMust())
	logger.WithCTextHandlerConsoleOut()
	// logger.WithCTextHandlerFileOut()
	logger.Construct()

	logger.Info("Logger check")

	stop, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer cancel()

	db := repo.NewDB()
	db.Fill()
	s := service.NewTaskService(db)
	taskHandler := transport.NewTaskHandler(s)

	router := server.NewRouter()
	router.SetHandlers(taskHandler.GetRoutes())
	wrappedRouter := middleware.WrapRouter(router,
		middleware.Logger(logger),
		middleware.RequestID(),
	)

	server := server.NewCServer(server.NewConfigMust())
	server.AttachHandler(wrappedRouter)
	server.Run(stop, logger)
}
