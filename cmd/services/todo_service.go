package services

import (
	"todolist/cmd/models"
	"todolist/cmd/repositories"
)

type TodoService interface {
	CreateTodo(todo models.Todo) (int64, error)
	GetTodoById(id int) (*models.Todo, error)
	GetAllTodos() ([]models.Todo, error)
	UpdateTodo(id int, todo models.Todo) (*models.Todo, error)
	DeleteTodo(id int) error
}

type TodoServiceImpl struct {
	store repositories.TodoStore
}

// NewTodoService creates a new TodoService instance
func NewTodoService(store repositories.TodoStore) TodoService {
	return &TodoServiceImpl{store: store}
}

func (s *TodoServiceImpl) CreateTodo(todo models.Todo) (int64, error) {
	return s.store.CreateTodo(todo)
}

func (s *TodoServiceImpl) GetAllTodos() ([]models.Todo, error) {
	return s.store.GetAllTodos()
}

func (s *TodoServiceImpl) GetTodoById(id int) (*models.Todo, error) {
	return s.store.GetTodoById(id)
}

func (s *TodoServiceImpl) UpdateTodo(id int, todo models.Todo) (*models.Todo, error) {
	return s.store.UpdateTodo(id, todo)
}

func (s *TodoServiceImpl) DeleteTodo(id int) error {
	return s.store.DeleteTodo(id)
}
