package services

import (
	dto "github.com/Kazzzz195/GoProject/DTO"
	"github.com/Kazzzz195/GoProject/models"
	"github.com/Kazzzz195/GoProject/repositories"
)

func (s *TodoService) GetTodo(todoID int) (models.Todo, error) {

	todo, err := repositories.GetTodoByID(s.db, todoID)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) GetAllTodo() ([]models.Todo, error) {

	todos, err := repositories.GetAllTodos(s.db)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoService) CreateTodo(dto dto.CreateTodoDTO) (models.Todo, error) {

	todo := &models.Todo{
		Title:     dto.Title,
		Body:      dto.Body,
		DueDate:   dto.DueDate,
		CreatedAt: s.currentTime,
	}

	newTodo, err := repositories.InsertTodo(s.db, *todo)
	if err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil
}

func (s *TodoService) UpdateTodo(dto dto.UpdateTodoDTO) (models.Todo, error) {

	todo := &models.Todo{
		ID:         dto.ID,
		Title:      dto.Title,
		Body:       dto.Body,
		DueDate:    dto.DueDate,
		CompleteAt: dto.CompleteAt,
		UpdateAt:   s.currentTime,
	}

	newTodo, err := repositories.UpdateTodo(s.db, *todo)
	if err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil
}

func (s *TodoService) DeleteTodo(todoID int) (models.Todo, error) {

	// Todoを削除
	err := repositories.DeleteTodoById(s.db, todoID)
	if err != nil {
		return models.Todo{}, err
	}

	// 空のTodoを返す
	return models.Todo{}, nil
}

func (s *TodoService) CompleteTodo(todoId int) (models.Todo, error) {

	newTodo, err := repositories.CompleteTodo(s.db, todoId, s.currentTime)
	if err != nil {
		return models.Todo{}, err
	}

	return newTodo, nil
}

func (s *TodoService) SearchTodoByBody(value string) ([]models.Todo, error) {

	todos, err := repositories.SearchTodosByBody(s.db, value)
	if err != nil {
		return []models.Todo{}, nil
	}

	return todos, nil
}

func (s *TodoService) SearchTodoByTitle(value string) ([]models.Todo, error) {

	todos, err := repositories.SearchTodosByTitle(s.db, value)
	if err != nil {
		return []models.Todo{}, nil
	}

	return todos, nil
}

func (s *TodoService) SearchCompletedTodo() ([]models.Todo, error) {

	todos, err := repositories.SearchCompletedTodos(s.db, s.currentTime)
	if err != nil {
		return []models.Todo{}, nil
	}

	return todos, nil
}
func (s *TodoService) SearchOngoingTodo() ([]models.Todo, error) {

	todos, err := repositories.SearchOngoingTodos(s.db, s.currentTime)
	if err != nil {
		return []models.Todo{}, nil
	}

	return todos, nil
}
