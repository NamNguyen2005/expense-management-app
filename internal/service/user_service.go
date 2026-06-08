package service

import (
	"expense-management/internal/repository"
	"log"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo : repo,
	}
}

func (s *userService) GetAllUsers() {
	log.Println("Get all users into service layer")
	s.repo.FindAll()
}

func (s *userService) CreateUser() {
	
}

func (s *userService) GetUser() {
	
}

func (s *userService) UpdateUser() {
	
}

func (s *userService) DeleteUser() {
	
}