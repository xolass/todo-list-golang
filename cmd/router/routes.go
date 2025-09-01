package routes

import (
	"todolist/cmd/controllers"

	"github.com/gin-gonic/gin"
)

type TodoRoutes interface {
	SetupRoutes(router *gin.Engine)
}

type TodoRoutesImpl struct {
	todoController controllers.TodoController
}

// NewTodoRoutes creates a new TodoRoutes instance
func NewTodoRoutes(todoController controllers.TodoController) TodoRoutes {
	return &TodoRoutesImpl{todoController: todoController}
}

func (c *TodoRoutesImpl) SetupRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")

	apiGroup.POST("/todo", c.todoController.CreateTodo)
	apiGroup.GET("/todo", c.todoController.GetAllTodos)
	apiGroup.GET("/todo/:id", c.todoController.GetTodoById)
	apiGroup.PUT("/todo/:id", c.todoController.UpdateTodo)
	apiGroup.DELETE("/todo/:id", c.todoController.DeleteTodo)
}
