package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/db"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment")
	}

	// Connect to database (must be done before server)
	db.ConnectDB()

	// Run migrations (optional, after DB connection)
	//migrations.RunMigrations()

	// Initialize server (this wires all modules)
	srv, err := server.NewServer()
	if err != nil {
		log.Fatal("Failed to initialize server:", err)
	}

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Get PORT from env, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("[[Server is Successfully Running on Port: %s]]\n", port)
	srv.Run(port)
}