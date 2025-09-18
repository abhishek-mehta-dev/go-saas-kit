package routes

import (
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
	}
}
