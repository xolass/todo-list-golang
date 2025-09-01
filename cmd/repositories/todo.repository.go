package repositories

import (
	"database/sql"
	"todolist/cmd/models"
)

type TodoStore interface {
	CreateTodo(models.Todo) (int64, error)
	GetAllTodos() ([]models.Todo, error)
	GetTodoById(id int) (*models.Todo, error)
	UpdateTodo(id int, todo models.Todo) (*models.Todo, error)
	DeleteTodo(id int) error
}

type TodoSql struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoStore {
	return &TodoSql{db: db}
}

func (s *TodoSql) CreateTodo(newItem models.Todo) (int64, error) {
	result, err := s.db.Exec("INSERT INTO todo (todo) VALUES ($1)", newItem.Todo)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *TodoSql) GetAllTodos() ([]models.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todo")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todoList []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Todo); err != nil {
			return nil, err
		}
		todoList = append(todoList, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todoList, err
}

func (s *TodoSql) GetTodoById(id int) (*models.Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todo WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todo models.Todo

	if rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Todo); err != nil {
			return nil, err
		}
	} else {
		return nil, sql.ErrNoRows
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (s *TodoSql) UpdateTodo(id int, todo models.Todo) (*models.Todo, error) {
	row, err := s.db.Query(`
		UPDATE todo SET 
			todo = $1
		WHERE id = $2
		RETURNING *
	`, todo.Todo, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	var returningTodo models.Todo

	if row.Next() {
		if err := row.Scan(&returningTodo.ID, &returningTodo.Todo); err != nil {
			return nil, err
		}
	} else {
		return nil, sql.ErrNoRows
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return &returningTodo, nil
}

func (s *TodoSql) DeleteTodo(id int) error {
	_, err := s.db.Exec("DELETE from todo WHERE id = $1", id)

	return err
}
