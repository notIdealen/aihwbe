package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/notIdealen/aihwbe.git/internal/logger"
	"github.com/notIdealen/aihwbe.git/internal/middleware"
	middleware_router "github.com/notIdealen/aihwbe.git/internal/middleware/router"
	"github.com/notIdealen/aihwbe.git/internal/repo/psql"
	server "github.com/notIdealen/aihwbe.git/internal/server/http"
	service_task "github.com/notIdealen/aihwbe.git/internal/service/task"
	transport_task "github.com/notIdealen/aihwbe.git/internal/transport/task"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Not found .env file")
	}

	// h := transport_task.NewTaskHandler(&service_task.TaskService{}) //   {s: service_task.TaskService{}}
	// server_mock.ServerForTestHandler("/tasks", http.HandlerFunc(h.CreateTask))

	logger := logger.NewBlank(logger.NewConfigMust())
	logger.WithCTextHandlerConsoleOut()
	// // logger.WithCTextHandlerFileOut()
	logger.Construct()
	logger.Info("Logger run")

	ctx := context.Background()
	config := psql.NewConfig()
	db := psql.NewPostgresPool(ctx, config)
	defer db.Close()

	stop, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer cancel()

	s := service_task.NewTaskService(db)
	taskHandler := transport_task.NewTaskHandler(s)

	router := server.NewRouter()
	router.SetHandlers(taskHandler.GetRoutes())
	wrappedRouter := middleware.WrapRouter(router,
		middleware_router.Logger(logger),
		middleware_router.RequestID(),
		middleware_router.SetWriterAndWriteLog(),
	)

	server := server.NewCServer(server.NewConfigMust())
	server.AttachHandler(wrappedRouter)
	server.Run(stop, logger)
}
