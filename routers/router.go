package routers

import (
	"net/http"

	"github.com/Kazzzz195/GoProject/controllers"
	"github.com/gorilla/mux"
)

func NewRouter(con *controllers.TodoController) *mux.Router {
	r := mux.NewRouter()

	// /api/v1/todo
	todoRouter := r.PathPrefix("/api/v1/todo").Subrouter()

	todoRouter.HandleFunc("/post", con.PostTodoHandler).Methods(http.MethodPost)
	todoRouter.HandleFunc("/list", con.TodoListHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/{id:[0-9]+}", con.TodoDetailHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/edit/{id:[0-9]+}", con.EditTodoHandler).Methods(http.MethodPatch)
	todoRouter.HandleFunc("/{id:[0-9]+}", con.DeleteTodoHandler).Methods(http.MethodDelete)
	todoRouter.HandleFunc("/complete/{id:[0-9]+}", con.CompleteTodoHandler).Methods(http.MethodPost)
	todoRouter.HandleFunc("/search/body", con.TodoSearchByBodyHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/search/title", con.TodoSearchByTitleHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/search/completed", con.TodoSearchByCompletedHandler).Methods(http.MethodGet)
	todoRouter.HandleFunc("/search/ongoing", con.TodoSearchByOngoingHandler).Methods(http.MethodGet)

	return r
}
