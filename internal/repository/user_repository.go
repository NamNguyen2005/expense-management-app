package repository

import (
	"expense-management/internal/model"
	"log"
)

type InMemoryUserRepository struct {
	users []model.User
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: []model.User{},
	}
}



func (s *InMemoryUserRepository) FindAll() {
	log.Println("Get all users into repo layer")
}

func (s *InMemoryUserRepository) Create() {
	
}

func (s *InMemoryUserRepository) Find() {
	
}

func (s *InMemoryUserRepository) Update() {
	
}

func (s *InMemoryUserRepository) Delete() {
	
}