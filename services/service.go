package services

import (
	"database/sql"
	"time"
)

type TodoService struct {
	db *sql.DB
	currentTime time.Time
}

func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		db:         db,
		currentTime: time.Now(),
	}
}
