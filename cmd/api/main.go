package main

import (
	"expense-management/internal/app"
	"expense-management/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// API initialization code would go here
	// Intitialization configuration, repository, service, handler, and routes
	app.LoadEnv()
	cfg := config.NewConfig()
	r := gin.Default()
	application := app.NewApplication(cfg, r)
	if err := application.Run(); err != nil {
		panic(err)
	}
}