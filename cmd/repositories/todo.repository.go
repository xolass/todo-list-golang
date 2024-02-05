package repositories

import (
	"database/sql"
	"todolist/cmd/providers"
)

func GetTodos() (*sql.Rows, error) {
	todo_list, err := providers.DB.Query("SELECT * FROM todo")
	return todo_list, err
}

func CreateTodo(newItem string) (sql.Result, error) {
	result, err := providers.DB.Exec("INSERT INTO todo (todo) VALUES ($1)", newItem)
	return result, err
}
