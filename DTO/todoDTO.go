package dto

import (
	"errors"
	"time"
)

type TaskDTO struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	DueDate    time.Time `json:"due_date"`
	CompleteAt time.Time `json:"complete_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}

type CreateTodoDTO struct {
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	DueDate time.Time `json:"due_date"`
}

type UpdateTodoDTO struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	DueDate    time.Time `json:"due_date"`
	CompleteAt time.Time `json:"complete_at"`
}

func (dto *CreateTodoDTO) Validate() error {
	if dto.Title == "" {
		return errors.New("title is required")
	}
	if dto.Body == "" {
		return errors.New("body is required")
	}
	if dto.DueDate.IsZero() {
		return errors.New("due date is required")
	}
	return nil
}

func (dto *UpdateTodoDTO) Validate() error {
	if dto.ID == 0 {
		return errors.New("id is required")
	}
	if dto.Title == "" {
		return errors.New("title is required")
	}
	if dto.Body == "" {
		return errors.New("body is required")
	}
	if dto.DueDate.IsZero() {
		return errors.New("due date is required")
	}
	return nil
}
