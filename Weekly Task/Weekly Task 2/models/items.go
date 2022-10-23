package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (item *Item) BeforeCreate(scope *gorm.Scope) (err error) {
	return scope.SetColumn("ID", uuid.New())
}

type Item struct {
	ID         uuid.UUID `gorm:"type:char(36);primarykey"`
	ItemName   string    `json:"item_name" form:"item_name"`
	Desc       string    `json:"desc" form:"desc"`
	CategoryID string    `json:"category_id" form:"category_id"`
}
