package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	dto "github.com/Kazzzz195/GoProject/DTO"
	"github.com/Kazzzz195/GoProject/controllers/services"
	"github.com/Kazzzz195/GoProject/models"

	"github.com/gorilla/mux"
)

type TodoController struct {
	service services.TodoServicer
}

func NewTodoController(s services.TodoServicer) *TodoController {
	return &TodoController{service: s}
}

// /todo/post
func (c *TodoController) PostTodoHandler(w http.ResponseWriter, req *http.Request) {
	// JSONデータを構造体にデコードする
	var createDTO dto.CreateTodoDTO
	if err := json.NewDecoder(req.Body).Decode(&createDTO); err != nil {
		http.Error(w, "Failed to decode JSON\n", http.StatusBadRequest)
		return
	}

	// バリデーションの実行
	if err := createDTO.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// デコードされた構造体を再度JSONにエンコードしてレスポンスとして返す
	todo, err := c.service.CreateTodo(createDTO)
	if err != nil {
		http.Error(w, "Internal server errors at repo\n", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// /todo-list のハンドラ
func (c *TodoController) TodoListHandler(w http.ResponseWriter, req *http.Request) {

	todos, err := c.service.GetAllTodo()
	if err != nil {
		http.Error(w, "Fail internal exec: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// /todo/id のハンドラ
func (c *TodoController) TodoDetailHandler(w http.ResponseWriter, req *http.Request) {

	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	todo, err := c.service.GetTodo(todoID)

	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	if todo.ID == 0 {
		// Todo item not found, return a 404 error
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// post/todo
func (c *TodoController) EditTodoHandler(w http.ResponseWriter, req *http.Request) {

	// JSONデータを構造体にデコードする
	var updateDTO dto.UpdateTodoDTO
	if err := json.NewDecoder(req.Body).Decode(&updateDTO); err != nil {
		http.Error(w, "Failed to decode JSON\n", http.StatusBadRequest)
		return
	}

	// バリデーションの実行
	if err := updateDTO.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// デコードされた構造体を再度JSONにエンコードしてレスポンスとして返す
	todo, err := c.service.UpdateTodo(updateDTO)
	if err != nil {
		http.Error(w, "fail internal excec\n", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}

// /todo/delete/
func (c *TodoController) DeleteTodoHandler(w http.ResponseWriter, req *http.Request) {
	// クエリパラメータからIDを取得
	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	// Todoを削除
	todo, err := c.service.DeleteTodo(todoID)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
}

// /todo/complete
func (c *TodoController) CompleteTodoHandler(w http.ResponseWriter, req *http.Request) {

	todoID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// 構造体をJSONにエンコードしてレスポンスとして返す
	todo, err := c.service.CompleteTodo(todoID)
	if err != nil {
		http.Error(w, "fail internal error\n", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}

// todo/search/body
func (c *TodoController) TodoSearchByBodyHandler(w http.ResponseWriter, req *http.Request) {

	searchValue := req.URL.Query().Get("body")
	if searchValue == "" {
		http.Error(w, "Query parameter 'body' is required", http.StatusBadRequest)
		return
	}
	//var todos []models.Todo
	var err error
	todos, err := c.service.SearchTodoByBody(searchValue)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// /todo/search/title
func (c *TodoController) TodoSearchByTitleHandler(w http.ResponseWriter, req *http.Request) {

	searchValue := req.URL.Query().Get("title")
	if searchValue == "" {
		http.Error(w, "Query parameter 'title' is required", http.StatusBadRequest)
		return
	}
	var todos []models.Todo
	var err error
	todos, err = c.service.SearchTodoByTitle(searchValue)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}

// /todo/search/complete
func (c *TodoController) TodoSearchByCompletedHandler(w http.ResponseWriter, req *http.Request) {

	var todos []models.Todo
	var err error
	todos, err = c.service.SearchCompletedTodo()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}

// /todo/search/ongoing
func (c *TodoController) TodoSearchByOngoingHandler(w http.ResponseWriter, req *http.Request) {
	var todos []models.Todo
	var err error
	todos, err = c.service.SearchOngoingTodo()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
