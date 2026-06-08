package routes

import (
	"expense-management/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		handler: handler,
	}
}


func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", ur.handler.CreateUser)
		userGroup.GET("/", ur.handler.GetAllUsers)
		userGroup.GET("/:uuid", ur.handler.GetUser)
		userGroup.PUT("/:uuid", ur.handler.UpdateUser)
		userGroup.DELETE("/:uuid", ur.handler.DeleteUser)
	}
}