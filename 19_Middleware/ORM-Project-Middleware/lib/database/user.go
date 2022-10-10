package database

import (
	"part3-orm/config"
	"part3-orm/middlewares"
	"part3-orm/models"
	"strconv"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id string) (interface{}, error) {
	var user []models.User
	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func DeleteUser(id string) (interface{}, error) {
	var user []models.User
	if err := config.DB.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id string, user *models.User) (interface{}, error) {
	var userUpdate []models.User
	if err := config.DB.Find(&userUpdate, id).Updates(&user).Error; err != nil {
		return nil, err
	}
	return userUpdate, nil
}

func GetDetailUsers(id string) (interface{}, error) {
	var user []models.User
	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(user *models.User) (interface{}, error) {
	var users []models.User
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).Find(&users).Error; err != nil {
		return nil, err
	}

	var err error

	user.Token, err = middlewares.CreateToken(strconv.Itoa(int(user.ID)), user.Email)
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return users, nil
}
