package controllers

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"part3-orm/lib/database"
	"part3-orm/models"
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
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	users, err := database.LoginUser(&user)
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

func GetUserDetailController(c echo.Context) error {
	id := c.Param("id")
	user, err := database.GetDetailUsers(id)
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

func CreateUserController(c echo.Context) error {
	// create user
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		return err
	}

	_, err = database.CreateUser(&user)
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
