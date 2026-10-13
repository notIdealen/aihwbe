package transport_task

import (
	"context"
	"net/http"

	server "github.com/notIdealen/aihwbe.git/internal/server/http"
)

type Servicer interface {
	// GetTask(context.Context, string) (*domain.Task, error)
	CreateTask(context.Context, CreateTaskRequestDTO) (string, error)
	// DeleteTask(context.Context, string) error
	// UpdateTask(context.Context, string, CreateTaskRequestDTO) (*domain.Task, error)
	// GetList(context.Context) ([]domain.Task, error)
}

type TaskHandler struct {
	s Servicer
}

func NewTaskHandler(s Servicer) *TaskHandler {
	return &TaskHandler{
		s: s,
	}
}

func (h *TaskHandler) GetRoutes() []server.Route {

	return []server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetList,
		},
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodPut,
			Path:    "/tasks/{id}",
			Handler: h.UpdateTask,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTask,
		},
	}
}
