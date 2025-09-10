package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println(".env file not found, using system environment")
    }

    // Get PORT from env, default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    db.ConnectDB()

    // Set Gin mode from env
    ginMode := os.Getenv("GIN_MODE")
    if ginMode == "" {
        ginMode = gin.DebugMode
    }
    gin.SetMode(ginMode)

    router := gin.Default()
    router.SetTrustedProxies([]string{"127.0.0.1"})

    // Routes
    router.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{"message": "Welcome to the world of GoLang!"})
    })

    fmt.Printf("[[Server is Successfully Running on Port: %s]]\n", port)
    router.Run(":" + port)
}
