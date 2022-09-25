package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type User struct {
	Id       int    `json:id form:id`
	Name     string `json:name form:name`
	Email    string `json:email form:email`
	Password string `json:password form:password`
}

var users []User

// Controller

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, user := range users {
		if user.Id == id {
			users = append(users[:index], users[index+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success delete user",
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, user := range users {
		if user.Id == id {
			newUser := User{}
			c.Bind(&newUser)
			newUser.Id = user.Id
			users[index] = newUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success update user",
				"user":    newUser,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// get user by id
func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success get user by id",
				"user":    user,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "user not found",
	})
}

// create user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		user.Id = users[len(users)-1].Id + 1
	}
	users = append(users, user)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}
func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserByIdController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.Logger.Fatal(e.Start(":8000"))
}
