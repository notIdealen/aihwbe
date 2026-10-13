package psql

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/notIdealen/aihwbe.git/internal/domain"
)

func (db *PostgresDB) CreateTask(ctx context.Context, task domain.Task) (string, error) {
	ctxwt, cancel := context.WithTimeout(ctx, db.timerForContext)
	defer cancel()

	var id uuid.UUID

	q := fmt.Sprintf("INSERT INTO %s(title, description, status) VALUES($1, $2, $3) RETURNING id;", db.table)
	err := db.QueryRow(ctxwt, q, task.Title, task.Desc, task.Status).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("PostgresDB: CreateTask: QueryRow: %w", err)
	}

	return id.String(), nil
}

func (db *PostgresDB) GetTask(ctx context.Context, id string) (domain.Task, error) {
	// ctxwt, cancel := context.WithTimeout(ctx, db.timerForContext)
	// defer cancel()
	// q := fmt.Sprintf("SELECT id, title, description, status, created_at, updated_at FROM %s WHERE id=$1;", db.table)
	var task domain.Task
	// err := db.QueryRow(ctxwt, q, id).Scan(&task.ID, &task.Title, &task.Desc, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	// if err != nil {
	// 	logger.Error("Bad scan when get task:", slog.String("err", err.Error()))
	// }
	return task, nil
}

func (db *PostgresDB) GetList(context.Context) ([]domain.Task, error) {
	return []domain.Task{}, nil
}

func (db *PostgresDB) DeleteTask(context.Context, string) error { return nil }

func (db *PostgresDB) UpdateTask(context.Context, domain.Task) error { return nil }
