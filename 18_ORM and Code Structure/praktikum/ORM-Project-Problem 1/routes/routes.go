package routes

import (
	"github.com/labstack/echo/v4"
	"part3-orm/controllers"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUsersControllers)
	e.GET("/user/:id", controllers.GetUserController)
	e.POST("/user", controllers.CreateUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)
	return e
}
