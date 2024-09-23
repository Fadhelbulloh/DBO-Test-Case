package main

import (
	"log"

	"github.com/Fadhelbulloh/DBO-Test-Case/models/db"
	"github.com/Fadhelbulloh/DBO-Test-Case/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize DB connection
	db.ConnectDB()
	// Close the database connection on application shutdown
	defer func() {
		log.Println("Closing database connection...")
		db.CloseDB()
	}()

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
