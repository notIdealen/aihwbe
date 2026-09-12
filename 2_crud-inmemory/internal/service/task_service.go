package service

import (
	"context"
	"time"

	"github.com/notIdealen/aihwbe.git/internal/domain"
	"github.com/notIdealen/aihwbe.git/internal/transport"
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

func (s *TaskService) GetTask(ctx context.Context, id string) (*domain.Task, error) {
	ctxwt, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(2100))
	defer cancel()

	task, err := s.DB.GetTask(ctxwt, id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) CreateTask(ctx context.Context, dto transport.CreateTaskRequestDTO) (string, error) {
	ctxwt, cancel := context.WithTimeout(ctx, time.Second*time.Duration(2))
	defer cancel()

	task := domain.Task{
		Title:  dto.Title,
		Desc:   dto.Desc,
		Status: dto.Status,
	}
	id, err := s.DB.CreateTask(ctxwt, task)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	ctxwt, cancel := context.WithTimeout(ctx, time.Second*time.Duration(3))
	defer cancel()

	err := s.DB.DeleteTask(ctxwt, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) UpdateTask(ctx context.Context, id string, dto transport.CreateTaskRequestDTO) (*domain.Task, error) {
	ctxwt, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(2900))
	defer cancel()

	task := domain.Task{
		ID: id,
	}
	err := s.DB.UpdateTask(ctxwt, task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) GetList(ctx context.Context) ([]domain.Task, error) {
	ctxwt, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(5000))
	defer cancel()
	return s.DB.GetList(ctxwt)
}
