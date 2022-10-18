package controllers

import (
	"fmt"
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
	user, err := database.GetUsers()
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

func GetUserController(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	user, err := database.GetUser(idInt)
	fmt.Println(err)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})

	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		// return name from models user
		"user": user,
	})
}

func CreateUserController(c echo.Context) error {
	// create user
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	fmt.Println(err)
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
	} else {
		fmt.Println(err)
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Model(&user).Where("id = ?", idInt).Updates(user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	user := models.User{}

	err := config.DB.Delete(&user, idInt).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}
