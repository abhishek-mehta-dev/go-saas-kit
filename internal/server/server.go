package server

import (
	"fmt"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/db"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/handlers"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/repositories"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/routes"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/services"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/pkg/email"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() (*Server, error) {
	// Use the global DB
	if db.DB == nil {
		return nil, fmt.Errorf("database is not initialized, call db.ConnectDB() first")
	}

	// Initialize modules
	userRepo := repositories.NewUserRepository(db.DB)
	emailSender := email.NewSMTPSender()
	userService := services.NewUserService(userRepo, emailSender)
	authHandler := handlers.NewAuthHandler(userService)

	// Initialize router
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Register health check route
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to the world of GoLang!"})
	})

	// Register API routes
	routes.RegisterRoutes(r, authHandler)

	return &Server{Router: r}, nil
}

func (s *Server) Run(port string) {
	s.Router.Run(":" + port)
}
