package main

import (
	"expense-management/internal/config"
	"expense-management/internal/handler"
	"expense-management/internal/repository"
	"expense-management/internal/routes"
	"expense-management/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// API initialization code would go here
	// Intitialization configuration, repository, service, handler, and routes
	cfg := config.NewConfig()
	repo := repository.NewInMemoryUserRepository()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	route := routes.NewUserRoutes(handler)

	r := gin.Default()
	// Register routes here using routes object
	routes.RegisterRoutes(r, route)
	r.Run(cfg.ServerPort)
}