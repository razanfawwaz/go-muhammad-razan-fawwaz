package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	BookName string `json:"book_name" form:"book_name"`
	Desc     string `json:"desc" form:"desc"`
	Author   string `json:"author_name" form:"author_name"`
}
