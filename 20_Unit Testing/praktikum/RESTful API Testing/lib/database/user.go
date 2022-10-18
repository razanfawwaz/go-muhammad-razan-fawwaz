package database

import (
	"part3-orm/config"
	"part3-orm/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	// select * from users
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	var userName = []string{}
	for _, user := range users {
		userName = append(userName, user.Name)
	}

	return userName, nil
}

func GetUser(id int) (interface{}, error) {
	// only return name from models user
	var user models.User
	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}
	return user.Name, nil
}
