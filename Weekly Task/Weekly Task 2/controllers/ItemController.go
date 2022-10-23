package controllers

import (
	"github.com/labstack/echo"
	"weekly-task2/lib/database"
	"weekly-task2/models"
)

func GetItemsController(c echo.Context) error {
	items, err := database.GetItems()
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   items,
	})
}

func GetItemController(c echo.Context) error {
	id := c.Param("id")
	items, err := database.GetItem(id)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   items,
	})
}

func CreateItemController(c echo.Context) error {
	item := models.Item{}

	err := c.Bind(&item)
	if err != nil {
		return err
	}
	_, err = database.CreateItem(&item)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   item,
	})
}

func UpdateItemController(c echo.Context) error {
	id := c.Param("id")
	item := models.Item{}

	err := c.Bind(&item)
	if err != nil {
		return err
	}
	_, err = database.UpdateItem(id, &item)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   item,
	})
}

func DeleteItemController(c echo.Context) error {
	id := c.Param("id")
	_, err := database.DeleteItem(id)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
	})
}

func SearchItemController(c echo.Context) error {
	keyword := c.Param("keyword")
	items, err := database.SearchItem(keyword)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   items,
	})
}

// get by category id
func GetItemByCategoryController(c echo.Context) error {
	id := c.Param("id")
	items, err := database.GetItemByCategory(id)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   items,
	})
}
