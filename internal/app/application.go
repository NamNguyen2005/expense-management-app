package app

import (
	"expense-management/internal/config"
	"expense-management/internal/routes"

	"github.com/gin-gonic/gin"
)

type Application struct {
	config *config.Config
	route  *gin.Engine
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


