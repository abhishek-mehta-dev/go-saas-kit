package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}, message string, statusCode int, pagination interface{}) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	c.JSON(statusCode, ApiResponse{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
		Pagination: pagination,
	})
}

func ErrorResponse(c *gin.Context, message string, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}
	c.JSON(statusCode, ApiResponse{
		Message:    message,
		StatusCode: statusCode,
	})
}
