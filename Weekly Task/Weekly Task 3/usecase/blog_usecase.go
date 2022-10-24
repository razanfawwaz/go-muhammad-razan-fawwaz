package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type BlogUsecase interface {
	CreateBlog(blog model.Blog) (model.Blog, error)
	FindBlog() (model.Blog, error)
	FindByID(id int) (model.Blog, error)
	Update(id int, blog model.Blog) (model.Blog, error)
	Delete(id int) error
	FindByTitle(title string) (model.Blog, error)
}

type blogHandler struct {
	Repository repository.BlogRepository
}

func NewBlogUsecase(blogRepo repository.BlogRepository) *blogHandler {
	return &blogHandler{blogRepo}
}

func (s *blogHandler) CreateBlog(blog model.Blog) (model.Blog, error) {
	return s.Repository.CreateBlog(blog)
}

func (s *blogHandler) FindBlog() (model.Blog, error) {
	return s.Repository.FindBlog()
}

func (s *blogHandler) FindByID(id int) (model.Blog, error) {
	return s.Repository.FindByID(id)
}

func (s *blogHandler) FindByTitle(title string) (model.Blog, error) {
	return s.Repository.FindByTitle(title)
}

func (s *blogHandler) Update(id int, blog model.Blog) (model.Blog, error) {
	return s.Repository.Update(id, blog)
}

func (s *blogHandler) Delete(id int) error {
	return s.Repository.Delete(id)
}
