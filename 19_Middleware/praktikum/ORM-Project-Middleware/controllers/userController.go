package controllers

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
	"part3-orm/config"
	"part3-orm/lib/database"
	"part3-orm/middlewares"
	"part3-orm/models"
	"strconv"
)

func GetUsersControllers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	id := c.Param("id")
	user, err := database.GetUser(id)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login user",
			"status":  err.Error(),
		})
	}

	token, err := middlewares.CreateToken(strconv.Itoa(int(user.ID)), user.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})
}

func CreateUserController(c echo.Context) error {
	// create user
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middlewares.CreateToken(strconv.Itoa(int(user.ID)), user.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func GetUserDetailController(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	_, err := database.DeleteUser(id)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
	})
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return err
	}

	_, err = database.UpdateUser(id, &user)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}
