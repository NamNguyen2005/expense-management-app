package main

import (
	"expense-management/internal/app"
	"expense-management/internal/config"
	// "expense-management/internal/handler"
	// "expense-management/internal/repository"
	// "expense-management/internal/routes"
	// "expense-management/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// API initialization code would go here
	// Intitialization configuration, repository, service, handler, and routes
	cfg := config.NewConfig()
	r := gin.Default()
	application := app.NewApplication(cfg, r)
	if err := application.Run(); err != nil {
		panic(err)
	}
}