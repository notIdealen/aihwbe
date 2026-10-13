package service_task

import (
	"context"

	"github.com/notIdealen/aihwbe.git/internal/domain"
)

type Repo interface {
	GetTask(context.Context, string) (domain.Task, error)
	CreateTask(context.Context, domain.Task) (string, error)
	DeleteTask(context.Context, string) error
	UpdateTask(context.Context, domain.Task) error
	GetList(context.Context) ([]domain.Task, error)
}

type TaskService struct {
	DB Repo
}

func NewTaskService(db Repo) *TaskService {
	return &TaskService{
		DB: db,
	}
}
