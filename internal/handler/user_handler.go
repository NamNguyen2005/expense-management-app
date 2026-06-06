package handler

import (
	"expense-management/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	// Handler logic to get all users would go here
	ctx.JSON(200, gin.H{"message": "Get all users"})
}