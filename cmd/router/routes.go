package routes

import (
	todo_controller "todolist/cmd/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/todo", todo_controller.GetTodosController).Methods("GET")
	apiRouter.HandleFunc("/todo", todo_controller.CreateTodoController).Methods("POST")

}
