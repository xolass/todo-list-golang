package main

import (
	"fmt"
	"log"
	"net/http"
	"todolist/cmd/controllers"
	middlewares "todolist/cmd/middlewares"
	providers "todolist/cmd/providers"
	"todolist/cmd/repositories"
	routes "todolist/cmd/router"
	"todolist/cmd/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Bootstrap the application
	router := bootstrap()

	// Start the server
	startServer(router)
}

func bootstrap() *gin.Engine {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	providers.ConnectDB()

	// Create Gin router
	router := gin.Default()
	router.Use(middlewares.EnableCORS)

	// Wire up dependencies
	wireDependencies(router)

	return router
}

func wireDependencies(router *gin.Engine) {
	todoRepository := repositories.NewTodoRepository(providers.DB)

	todoService := services.NewTodoService(todoRepository)

	todoController := controllers.NewTodoController(todoService)

	todoRoutes := routes.NewTodoRoutes(todoController)

	todoRoutes.SetupRoutes(router)
}

func startServer(router *gin.Engine) {
	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
