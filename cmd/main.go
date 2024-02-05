package main

import (
	"fmt"
	"log"
	"net/http"
	middlewares "todolist/cmd/middlewares"
	providers "todolist/cmd/providers"
	custom_router "todolist/cmd/router"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	bootstrap()
	startServer()
}

func bootstrap() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	providers.ConnectDB()
}

func startServer() {
	router := mux.NewRouter()
	router.Use(middlewares.EnableCORS)

	custom_router.SetupRoutes(router)

	fmt.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
