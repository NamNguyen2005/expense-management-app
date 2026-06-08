package handler

import (
	"expense-management/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	// Handler logic to get all users would go here
	ctx.JSON(200, gin.H{"message": "Get all users"})
	h.service.GetAllUsers()
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	// Handler logic to create a new user would go here
	ctx.JSON(200, gin.H{"message": "Create a new user"})
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	// Handler logic to get a user by UUID would go here
	ctx.JSON(200, gin.H{"message": "Get user by UUID"})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Update user by UUID"})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Delete user by UUID"})
}