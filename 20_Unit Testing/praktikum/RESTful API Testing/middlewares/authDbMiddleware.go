package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"part3-orm/config"
	"part3-orm/models"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var user models.User
	fmt.Println(email, password)
	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
