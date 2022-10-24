package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	Find() ([]model.User, error)
	Create(user model.User) (model.User, error)
	Auth(user model.User) (model.User, error)
}

type userHandler struct {
	Repository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userHandler {
	return &userHandler{userRepo}
}

func (s *userHandler) Create(user model.User) (model.User, error) {
	return s.Repository.Create(user)
}

func (s *userHandler) Find() ([]model.User, error) {
	return s.Repository.Find()
}

func (s *userHandler) Auth(user model.User) (model.User, error) {
	return s.Repository.Auth(&user)
}
