package repository

import (
	"belajar-go-echo/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	Find() ([]model.User, error)
	Auth(user *model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Find() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) Auth(user *model.User) (model.User, error) {
	// auth with jwt token
	err := r.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	return *user, err
}
