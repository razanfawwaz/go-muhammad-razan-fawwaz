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
	var booksData = []string{}
	for _, book := range books {
		booksData = append(booksData, book.BookName)
	}
	return booksData, nil
}
func GetBook(id int) (interface{}, error) {
	var books models.Book

	if err := config.DB.Find(&books, id).Error; err != nil {
		return nil, err
	}
	return books.BookName, nil

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
