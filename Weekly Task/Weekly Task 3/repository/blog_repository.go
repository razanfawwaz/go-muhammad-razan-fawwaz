package repository

import (
	"belajar-go-echo/model"
	"gorm.io/gorm"
)

type BlogRepository interface {
	CreateBlog(blog model.Blog) (model.Blog, error)
	FindBlog() (model.Blog, error)
	FindByID(id int) (model.Blog, error)
	Update(id int, blog model.Blog) (model.Blog, error)
	Delete(id int) error
	FindByTitle(title string) (model.Blog, error)
}

func NewBlogRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBlog(blog model.Blog) (model.Blog, error) {
	err := r.db.Create(&blog).Error
	return blog, err
}

func (r *repository) FindBlog() (model.Blog, error) {
	var blogs model.Blog
	err := r.db.Find(&blogs).Error
	return blogs, err
}

func (r *repository) FindByID(id int) (model.Blog, error) {
	var blog model.Blog
	err := r.db.Where("id = ?", id).First(&blog).Error
	return blog, err
}

func (r *repository) Update(id int, blog model.Blog) (model.Blog, error) {
	err := r.db.Where("id = ?", id).Updates(&blog).Error
	return blog, err
}

func (r *repository) Delete(id int) error {
	err := r.db.Where("id = ?", id).Delete(&model.Blog{}).Error
	return err
}

func (r *repository) FindByTitle(title string) (model.Blog, error) {
	var blogs model.Blog
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&blogs).Error
	return blogs, err
}
