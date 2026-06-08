package service

type UserService interface {
	GetAllUsers()
	CreateUser()
	GetUser()
	UpdateUser()
	DeleteUser()
}