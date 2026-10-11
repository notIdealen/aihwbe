package service_task

import (
	"context"

	"github.com/google/uuid"

	// "uuid"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	transport_task "github.com/notIdealen/aihwbe.git/internal/transport/task"
)

func (s *TaskService) UpdateTask(ctx context.Context, id string, dto transport_task.CreateTaskRequestDTO) (*domain.Task, error) {
	// ctxwt, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(2900))
	// defer cancel()

	task := domain.Task{
		ID: uuid.MustParse(id),
	}
	err := s.DB.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
