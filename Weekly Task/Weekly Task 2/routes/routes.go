package routes

import (
	"github.com/labstack/echo"
	"weekly-task2/controllers"
)

func New() *echo.Echo {
	e := echo.New()
	// User routes
	e.GET("/items", controllers.GetItemsController)
	u := e.Group("/items")
	u.GET("/:id", controllers.GetItemController)
	u.POST("", controllers.CreateItemController)
	u.DELETE("/:id", controllers.DeleteItemController)
	u.PUT("/:id", controllers.UpdateItemController)
	u.GET("/:keyword", controllers.SearchItemController)
	// by category
	u.GET("/category/:category_id", controllers.GetItemByCategoryController)
	return e
}
