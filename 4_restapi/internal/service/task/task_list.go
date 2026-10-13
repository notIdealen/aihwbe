package service_task

import (
	"context"

	"github.com/notIdealen/aihwbe.git/internal/domain"
)

func (s *TaskService) GetList(ctx context.Context) ([]domain.Task, error) {
	// ctxwt, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(5000))
	// defer cancel()
	return s.DB.GetList(ctx)
}
