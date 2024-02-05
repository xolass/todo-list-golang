package controllers

import (
	"encoding/json"
	"net/http"
	todo_service "todolist/cmd/services"
)

func GetTodosController(w http.ResponseWriter, r *http.Request) {
	todoList, err := todo_service.GetTodos()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonTodoList, err := json.Marshal(todoList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonTodoList))
}

func CreateTodoController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo string
	err := decoder.Decode(&todo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	todoList, err := todo_service.CreateTodo(todo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	jsonTodoList, err := json.Marshal(todoList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonTodoList))
}
