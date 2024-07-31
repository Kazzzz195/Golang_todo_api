package models

import "time"

type Todo struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	DueDate    time.Time `json:"due_date"`
	CompleteAt time.Time `json:"complete_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}