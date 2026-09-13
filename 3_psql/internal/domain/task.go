package domain

import "time"

type Task struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Desc      string     `json:"description"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
