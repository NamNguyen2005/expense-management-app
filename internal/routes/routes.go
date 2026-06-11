package routes

import (
	"expense-management/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(router *gin.Engine, routes ...Routes) {
	router.Use(
		middleware.LoggingMiddleware(),
		middleware.APIKeyMiddleware(),
		middleware.AuthMiddleware(),
		middleware.APIKeyMiddleware(),
	)
	router.Use(middleware.AuthMiddleware())
	api := router.Group("/api")
	for _, route := range routes {
		route.Register(api)
	}
}