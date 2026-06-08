package repository

type UserRepository interface {
	FindAll()
	Create()
	Find()
	Update()
	Delete()
}