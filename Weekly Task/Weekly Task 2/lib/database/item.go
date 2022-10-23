package database

import (
	"weekly-task2/config"
	"weekly-task2/models"
)

func GetItems() (interface{}, error) {
	var items []models.Item

	if err := config.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
func GetItem(id string) (interface{}, error) {
	var item []models.Item

	if err := config.DB.Find(&item, id).Error; err != nil {
		return nil, err
	}
	return item, nil

}
func CreateItem(item *models.Item) (interface{}, error) {
	if err := config.DB.Create(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func DeleteItem(id string) (interface{}, error) {
	var item []models.Item
	if err := config.DB.Delete(&item, id).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func UpdateItem(id string, item *models.Item) (interface{}, error) {
	if err := config.DB.Find(&item, id).Updates(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func SearchItem(keyword string) (interface{}, error) {
	var items []models.Item

	if err := config.DB.Where("item_name LIKE ?", "%"+keyword+"%").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func GetItemByCategory(categoryId string) (interface{}, error) {
	var item []models.Item

	if err := config.DB.Find(&item, categoryId).Error; err != nil {
		return nil, err
	}
	return item, nil
}
