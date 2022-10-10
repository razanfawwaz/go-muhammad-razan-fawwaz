package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoMid "github.com/labstack/echo/v4/middleware"
	"part3-orm/constants"
	"part3-orm/controllers"
	"part3-orm/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// User routes
	e.GET("/users", controllers.GetUsersControllers)
	e.POST("/users", controllers.LoginUsersController)
	u := e.Group("/user")
	u.GET("/:id", controllers.GetUserController)
	u.POST("", controllers.CreateUserController)

	eAuth := e.Group("")
	eAuth.Use(echoMid.BasicAuth(middlewares.BasicAuthDB))
	eAuth.DELETE("/:id", controllers.DeleteUserController)
	eAuth.PUT("/:id", controllers.UpdateUserController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/user/:id", controllers.GetUserDetailController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/book/:id", controllers.GetBookController)
	e.POST("/book", controllers.CreateBookController)
	e.PUT("/book/:id", controllers.UpdateBookController)
	e.DELETE("book/:id", controllers.DeleteBookController)

	return e
}
