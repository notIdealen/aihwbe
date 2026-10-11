package service_task

import (
	"context"
	"fmt"
	"strings"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	transport_task "github.com/notIdealen/aihwbe.git/internal/transport/task"
)

func (s *TaskService) CreateTask(ctx context.Context, dto transport_task.CreateTaskRequestDTO) (string, error) {
	strings.TrimSpace(dto.Title)
	if dto.Title == "" {
		return "", domain.NewBadRequestError("Empty title",
			"TaskService: CreateTask: dto.Title == ''",
		)
	}
	status, err := domain.GetStatus(dto.Status)
	if err != nil {
		return "", domain.NewBadRequestError("invalid status",
			fmt.Sprintf("TaskService: domain.GetStatus: %s", err),
		)
	}

	task := domain.Task{
		Title:  dto.Title,
		Desc:   dto.Desc,
		Status: status,
	}
	id, err := s.DB.CreateTask(ctx, task)
	if err != nil {
		return "", err
	}

	return id, nil
}
