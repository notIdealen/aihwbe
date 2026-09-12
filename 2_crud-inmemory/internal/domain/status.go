package domain

import "errors"

type TaskStatus string

const (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in_progress"
	Done       TaskStatus = "done"
)

func (s TaskStatus) IsValidStatus() error {
	switch s {
	case Todo:
	case InProgress:
	case Done:
	default:
		return errors.New("Invalid task status")
	}
	return nil
}
