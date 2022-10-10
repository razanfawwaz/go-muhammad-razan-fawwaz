package routes

import (
	"github.com/labstack/echo/v4"
	echomid "github.com/labstack/echo/v4/middleware"
	"part3-orm/constants"
	"part3-orm/controllers"
	"part3-orm/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)
	e.GET("/users", controllers.GetUserDetailController)

	eAuth := e.Group("")
	eAuth.Use(echomid.BasicAuth(middlewares.BasicAuthDB))
	eAuth.DELETE("/:id", controllers.DeleteUserController)
	eAuth.PUT("/:id", controllers.UpdateUserController)

	jwtAuth := e.Group("")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/users", controllers.GetUsersControllers)
	return e
}
