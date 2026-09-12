package repo

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/notIdealen/aihwbe.git/internal/domain"
	"github.com/notIdealen/aihwbe.git/internal/utils"
)

type MemoryMapDB struct {
	mu sync.RWMutex
	db map[string]domain.Task
}

func NewDB() *MemoryMapDB {
	return &MemoryMapDB{
		mu: sync.RWMutex{},
		db: make(map[string]domain.Task, 1000),
	}
}

func (r *MemoryMapDB) GetTask(ctx context.Context, id string) (domain.Task, error) {
	waiting := rand.IntN(4) + 1
	timer := time.NewTimer(time.Second * time.Duration(waiting))

	logger := utils.GetLoggerFromRequestContext(ctx)
	var task domain.Task

	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-timer.C:
		task, ok := r.db[id]
		if !ok {
			return task, domain.ErrNotFound
		}
		return task, nil
	case <-ctx.Done():
		logger.Error("Long waiting! Time out!!!!")
		return task, domain.ErrGatewayTimeout
	}
}

func (r *MemoryMapDB) CreateTask(ctx context.Context, task domain.Task) (string, error) {
	var id string
	for {
		id = uuid.New().String()
		if id == "" {
			return "", fmt.Errorf("Invalid ID generated")
		}
		_, ok := r.db[id]
		if !ok {
			break
		}
	}

	logger := utils.GetLoggerFromRequestContext(ctx)

	waiting := rand.IntN(4) + 1
	timer := time.NewTimer(time.Second * time.Duration(waiting))

	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-timer.C:
		task.ID = id
		task.CreatedAt = time.Now()
		task.UpdatedAt = task.CreatedAt
		r.db[id] = task
		return id, nil
	case <-ctx.Done():
		logger.Error("Long waiting! Time out!!!!")
		return "", domain.ErrGatewayTimeout
	}
}

func (r *MemoryMapDB) DeleteTask(ctx context.Context, id string) error {
	waiting := rand.IntN(4) + 1
	timer := time.NewTimer(time.Second * time.Duration(waiting))

	logger := utils.GetLoggerFromRequestContext(ctx)

	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-timer.C:
		if _, ok := r.db[id]; ok {
			delete(r.db, id)
			return nil
		}
		return domain.ErrNotFound
	case <-ctx.Done():
		logger.Error("Long waiting! Time out!!!!")
		return domain.ErrGatewayTimeout
	}
}

func (r *MemoryMapDB) UpdateTask(ctx context.Context, data domain.Task) error {
	waiting := rand.IntN(4) + 1
	timer := time.NewTimer(time.Second * time.Duration(waiting))

	logger := utils.GetLoggerFromRequestContext(ctx)

	r.mu.RLock()
	defer r.mu.RUnlock()

	select {
	case <-timer.C:
		task, ok := r.db[data.ID]
		if !ok {
			return domain.ErrNotFound
		}
		task.Title = data.Title
		task.Desc = data.Desc
		task.Status = data.Status
		task.UpdatedAt = time.Now()
		r.db[data.ID] = task
		return nil
	case <-ctx.Done():
		logger.Error("UpdateTask->Long waiting! Time out!!!!")
		return domain.ErrGatewayTimeout
	}
}

func (r *MemoryMapDB) GetList(ctx context.Context) ([]domain.Task, error) {
	tasks := make([]domain.Task, 0, len(r.db))

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, task := range r.db {
		tasks = append(tasks, task)
	}
	return tasks, nil
}
