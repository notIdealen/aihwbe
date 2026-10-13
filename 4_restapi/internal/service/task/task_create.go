package service_task

import (
	"context"
	"fmt"
	"strings"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	service_error "github.com/notIdealen/aihwbe.git/internal/service/error"
	transport_task "github.com/notIdealen/aihwbe.git/internal/transport/task"
)

func (s *TaskService) CreateTask(ctx context.Context, dto transport_task.CreateTaskRequestDTO) (string, error) {
	dto.Title = strings.TrimSpace(dto.Title)
	if dto.Title == "" {
		return "", service_error.NewValidationError("Empty title", "TaskService: CreateTask: dto.Title == ''")
	}

	status, err := domain.GetStatus(dto.Status)
	if err != nil {
		return "", service_error.NewValidationError("invalid status",
			fmt.Sprintf("TaskService: CreateTask: GetStatus: %s", err),
		)
	}

	task := domain.Task{
		Title:  dto.Title,
		Desc:   dto.Desc,
		Status: status,
	}
	id, err := s.DB.CreateTask(ctx, task)
	if err != nil {
		return "", fmt.Errorf("TaskService: CreateTask: %w", err)
	}

	return id, nil
}
