package domain

import (
	"errors"
	"strings"
)

type TaskStatus string

var ErrInvalidTaskStatus = errors.New("Invalid task status")

const (
	Todo       TaskStatus = "todo"
	InProgress TaskStatus = "in_progress"
	Done       TaskStatus = "done"
)

func GetStatus(status string) (TaskStatus, error) {
	strings.TrimSpace(status)
	s := TaskStatus(status)
	switch s {
	case Todo:
		fallthrough
	case InProgress:
		fallthrough
	case Done:
		return s, nil
	}
	return "", ErrInvalidTaskStatus
}
