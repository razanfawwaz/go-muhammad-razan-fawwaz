package database

import (
	"part3-orm/config"
	"part3-orm/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
func GetBook(id string) (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books, id).Error; err != nil {
		return nil, err
	}
	return books, nil

}
func CreateBook(book *models.Book) (interface{}, error) {
	if err := config.DB.Create(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id string) (interface{}, error) {
	var book []models.Book
	if err := config.DB.Delete(&book, id).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(id string, book *models.Book) (interface{}, error) {
	if err := config.DB.Find(&book, id).Updates(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}
