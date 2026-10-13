package service_task

import (
	"context"

	"github.com/notIdealen/aihwbe.git/internal/domain"
)

func (s *TaskService) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	// ctxwt, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(2100))
	// defer cancel()

	task, err := s.DB.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
