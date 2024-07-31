package services

import (
	dto "github.com/Kazzzz195/GoProject/DTO"
	"github.com/Kazzzz195/GoProject/models"
)

type TodoServicer interface {
	GetAllTodo() ([]models.Todo, error)
	GetTodo(todoID int) (models.Todo, error)
	CreateTodo(dto dto.CreateTodoDTO) (models.Todo, error)
	UpdateTodo(dto dto.UpdateTodoDTO) (models.Todo, error)
	DeleteTodo(todoID int) (models.Todo, error)
	CompleteTodo(todoId int) (models.Todo, error)
	SearchTodoByBody(value string) ([]models.Todo, error)
	SearchTodoByTitle(value string) ([]models.Todo, error)
	SearchCompletedTodo() ([]models.Todo, error)
	SearchOngoingTodo() ([]models.Todo, error)
}
