package controllers

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"part3-orm/lib/database"
	"part3-orm/models"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookController(c echo.Context) error {
	id := c.Param("id")
	books, err := database.GetBook(id)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func LoginBookController(c echo.Context) error {
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

func GetBookDetailController(c echo.Context) error {
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

func CreateBookController(c echo.Context) error {
	book := models.Book{}

	err := c.Bind(&book)
	if err != nil {
		return err
	}
	_, err = database.CreateBook(&book)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}

	err := c.Bind(&book)
	if err != nil {
		return err
	}
	_, err = database.UpdateBook(id, &book)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	_, err := database.DeleteBook(id)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
	})
}
