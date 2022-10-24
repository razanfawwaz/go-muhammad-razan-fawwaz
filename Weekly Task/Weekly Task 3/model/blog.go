package model

import "gorm.io/gorm"

type Blog struct {
	*gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	// user id is the foreign key
	Author User `gorm:"foreignKey:Name"`
	UserID User `gorm:"foreignKey:ID"`
}
