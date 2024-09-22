package main

import (
	"log"

	"github.com/Fadhelbulloh/DBO-Test-Case/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	r.Run() // Start server on default port 8080
}
