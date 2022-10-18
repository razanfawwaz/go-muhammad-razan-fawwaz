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

func GetBooksControllers(c echo.Context) error {
	book, err := database.GetBooks()
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

func GetBookController(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	book, err := database.GetBook(idInt)
	fmt.Println(err)
	if err != nil {
		return c.JSON(500, map[string]string{
			"status": err.Error(),
		})

	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		// return name from models book
		"book": book,
	})
}

func CreateBookController(c echo.Context) error {
	// create book
	book := models.Book{}
	user := models.User{}
	c.Bind(&book)

	err := config.DB.Save(&book).Error
	fmt.Println(err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middlewares.CreateToken(strconv.Itoa(int(book.ID)), user.Email)

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
		"message": "success create book",
		"book":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	book := models.Book{}
	c.Bind(&book)

	err := config.DB.Model(&book).Where("id = ?", idInt).Updates(book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	book := models.Book{}

	err := config.DB.Delete(&book, idInt).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"book":    book,
	})
}
