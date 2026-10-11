package main

import (
	"net/http"

	"github.com/joho/godotenv"
	server_mock "github.com/notIdealen/aihwbe.git/internal/server/mock"
	service_task "github.com/notIdealen/aihwbe.git/internal/service/task"
	transport_task "github.com/notIdealen/aihwbe.git/internal/transport/task"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Not found .env file")
	}

	h := transport_task.NewTaskHandler(&service_task.TaskService{}) //   {s: service_task.TaskService{}}
	server_mock.ServerForTestHandler("/tasks", http.HandlerFunc(h.CreateTask))

	// logger := logger.NewBlank(logger.NewConfigMust())
	// logger.WithCTextHandlerConsoleOut()
	// // // logger.WithCTextHandlerFileOut()
	// logger.Construct()
	// logger.Info("Logger run")

	// ctx := context.Background()
	// config := psql.NewConfig()
	// db := psql.NewPostgresPool(ctx, config)
	// db.Connect(ctx)
	// defer db.Close()

	// task1, _ := db.GetTask(ctx, "cd6f839f-dfb3-4b93-8f7b-6be25a7917e8")
	// fmt.Println(task1)
	// stop, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	// defer cancel()

	// // db := repo.NewDB()
	// // db.Fill()
	// s := service.NewTaskService(db)
	// taskHandler := transport_task.NewTaskHandler(s)

	// router := server.NewRouter()
	// router.SetHandlers(taskHandler.GetRoutes())
	// wrappedRouter := middleware.WrapRouter(router,
	// 	middleware.Logger(logger),
	// 	middleware.RequestID(),
	// middleware.ChangeWriter(),
	// )

	// server := server.NewCServer(server.NewConfigMust())
	// server.AttachHandler(wrappedRouter)
	// server.Run(stop, logger)
}
