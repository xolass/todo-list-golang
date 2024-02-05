package services

import (
	"todolist/cmd/models"
	"todolist/cmd/repositories"
)

func GetTodos() ([]models.Todo, error) {
	rows, err := repositories.GetTodos()
	if err != nil {
		return nil, err
	}
	var todoList []models.Todo

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Todo)
		if err != nil {
			return nil, err
		}
		todoList = append(todoList, todo)
	}

	return todoList, err
}

func CreateTodo(newItem string) (int64, error) {

	result, err := repositories.CreateTodo(newItem)
	id, _ := result.LastInsertId()
	return id, err
}
