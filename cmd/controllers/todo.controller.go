package controllers

import (
	"net/http"
	"strconv"
	"todolist/cmd/models"
	"todolist/cmd/services"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	CreateTodo(c *gin.Context)
	GetTodoById(c *gin.Context)
	GetAllTodos(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type TodoControllerImpl struct {
	todoService services.TodoService
}

// NewTodoController creates a new TodoController instance
func NewTodoController(todoService services.TodoService) TodoController {
	return &TodoControllerImpl{todoService: todoService}
}

func (s *TodoControllerImpl) CreateTodo(c *gin.Context) {
	var body models.Todo
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	todoId, err := s.todoService.CreateTodo(body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"id": todoId})
}

func (s *TodoControllerImpl) GetAllTodos(c *gin.Context) {
	todoList, err := s.todoService.GetAllTodos()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, todoList)
}

func (s *TodoControllerImpl) GetTodoById(c *gin.Context) {
	todoId, found := c.Params.Get("id")

	if !found {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing id info"})
		return
	}

	todoIdAsNumber, err := strconv.Atoi(todoId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id doesn't have the correct format"})
		return
	}

	todo, err := s.todoService.GetTodoById(todoIdAsNumber)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, todo)
}

func (s *TodoControllerImpl) UpdateTodo(c *gin.Context) {
	todoId, found := c.Params.Get("id")

	if !found {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing id info"})
		return
	}

	todoIdAsNumber, err := strconv.Atoi(todoId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id doesn't have the correct format"})
		return
	}

	var todoBody models.Todo
	err = c.ShouldBindJSON(&todoBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	updatedTodo, err := s.todoService.UpdateTodo(todoIdAsNumber, todoBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, updatedTodo)
}

func (s *TodoControllerImpl) DeleteTodo(c *gin.Context) {
	todoId, found := c.Params.Get("id")

	if !found {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing id info"})
		return
	}

	todoIdAsNumber, err := strconv.Atoi(todoId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id doesn't have the correct format"})
		return
	}

	err = s.todoService.DeleteTodo(todoIdAsNumber)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
