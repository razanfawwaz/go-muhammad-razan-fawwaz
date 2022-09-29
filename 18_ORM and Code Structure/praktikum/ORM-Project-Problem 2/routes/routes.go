package routes

import (
	"github.com/labstack/echo/v4"
	"part3-orm/controllers"
)

func New() *echo.Echo {
	e := echo.New()

	// User routes
	e.GET("/users", controllers.GetUsersControllers)
	u := e.Group("/user")
	u.GET("/:id", controllers.GetUserController)
	u.POST("", controllers.CreateUserController)
	u.DELETE("/:id", controllers.DeleteUserController)
	u.PUT("/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/book/:id", controllers.GetBookController)
	e.POST("/book", controllers.CreateBookController)
	e.PUT("/book/:id", controllers.UpdateBookController)
	e.DELETE("book/:id", controllers.DeleteBookController)

	return e
}
