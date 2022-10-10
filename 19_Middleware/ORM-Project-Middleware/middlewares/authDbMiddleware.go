package middlewares

import (
	"github.com/labstack/echo/v4"
	"part3-orm/config"
	"part3-orm/models"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	db.Where("username = ? AND password = ?", username, password).First(&user)
	if user.ID != 0 {
		return true, nil
	}
	return false, nil
}
