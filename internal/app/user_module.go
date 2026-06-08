package app

import (
	"expense-management/internal/handler"
	"expense-management/internal/repository"
	"expense-management/internal/routes"
	"expense-management/internal/service"
)

type UserModule struct {
	// This struct can hold references to the service, repository, and handler for the user module
	route routes.Routes
}

func NewUserModule() *UserModule {
	repo := repository.NewInMemoryUserRepository()
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	route := routes.NewUserRoutes(handler)
	return &UserModule{
		route: route,
	}
}

func (m *UserModule) GetRoute() routes.Routes {
	return m.route
}