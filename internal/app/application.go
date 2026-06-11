package app

import (
	"expense-management/internal/config"
	"expense-management/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

type Application struct {
	config *config.Config
	route  *gin.Engine
	modules []Module
}

type Module interface {
	GetRoute() routes.Routes
}

func NewApplication(config *config.Config, r *gin.Engine) *Application {
	modules := []Module{
		NewUserModule(),	
	}

	routes.RegisterRoutes(r, GetModulesRoutes(modules)...)
	return &Application{
		config: config,
		route:  r,
		modules: modules,
	}
}

func (app *Application) Run() error {
	return app.route.Run(app.config.ServerPort)
}

func GetModulesRoutes(modules []Module) []routes.Routes {
	routeList := make([]routes.Routes, len(modules))
	for i, module := range modules {
		routeList[i] = module.GetRoute()
	}
	return routeList
}

func LoadEnv() {
	err := godotenv.Load(".env")
  	if err != nil {
   		log.Fatal("Error loading .env file")
 	 }
}
