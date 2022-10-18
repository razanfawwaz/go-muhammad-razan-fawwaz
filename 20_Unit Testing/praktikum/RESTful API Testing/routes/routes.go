package routes

import (
	"github.com/labstack/echo/v4"
	echomid "github.com/labstack/echo/v4/middleware"
	"part3-orm/constants"
	"part3-orm/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controllers.CreateUserController)
	e.GET("/books", controllers.GetBooksControllers)
	e.GET("/books/:id", controllers.GetBookController)

	//eAuth := e.Group("")
	//eAuth.Use(echomid.BasicAuth(middlewares.BasicAuthDB))
	//eAuth.DELETE("/:id", controllers.DeleteUserController)
	//eAuth.PUT("/:id", controllers.UpdateUserController)
	//eAuth.DELETE("/books/:id", controllers.DeleteBookController)
	//eAuth.PUT("/books/:id", controllers.UpdateBookController)

	jwtAuth := e.Group("")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("/users", controllers.GetUsersControllers)
	jwtAuth.GET("/users/:id", controllers.GetUserController)
	jwtAuth.DELETE("/users/:id", controllers.DeleteUserController)
	jwtAuth.PUT("/users/:id", controllers.UpdateUserController)
	jwtAuth.POST("/books", controllers.CreateBookController)
	jwtAuth.DELETE("/books/:id", controllers.DeleteBookController)
	jwtAuth.PUT("/books/:id", controllers.UpdateBookController)
	return e
}
