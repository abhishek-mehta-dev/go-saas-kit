package handlers

import (
	"net/http"

	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/models"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/internal/services"
	"github.com/abhishek-mehta-dev/go-saas-kit.git/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService services.UserService
}

func NewAuthHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{userService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		FirstName string `json:"firstName" binding:"required"`
		LastName  string `json:"lastName"`
		UserName  string `json:"userName" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserName:  req.UserName,
		Email:     req.Email,
		Password:  req.Password,
	}

	if err := h.userService.Register(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessResponse(c, user, "User registered successfully", http.StatusCreated, nil)
}
