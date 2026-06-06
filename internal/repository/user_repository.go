package repository

import "expense-management/internal/model"

type InMemoryUserRepository struct {
	users []model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []model.User{},
	}
}