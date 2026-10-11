package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Desc      string     `json:"description"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
