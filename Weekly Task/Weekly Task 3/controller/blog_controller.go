package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// blog controller

type BlogController interface{}

type blogController struct {
	useBlogusecase usecase.BlogUsecase
}

func NewBlogController(blogUsecase usecase.BlogUsecase) *blogController {
	return &blogController{
		blogUsecase,
	}
}

func (u *blogController) FindById(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var blog model.Blog
		if err := c.Bind(&blog); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		blog, err := u.useBlogusecase.FindByID(int(blog.ID))
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    blog,
		})
	}
}

func (u *blogController) FindByTitle(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var blog model.Blog
		if err := c.Bind(&blog); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		blog, err := u.useBlogusecase.FindByTitle(string(blog.Title))
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    blog,
		})
	}
}

func (u *blogController) FindAll(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var blog model.Blog
		if err := c.Bind(&blog); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		blog, err := u.useBlogusecase.FindBlog()
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    blog,
		})
	}
}

func (u *blogController) CreateBlog(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var blog model.Blog
		if err := c.Bind(&blog); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		blog, err := u.useBlogusecase.CreateBlog(blog)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    blog,
		})
	}
}

func (u *blogController) UpdateBlog(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var blog model.Blog
		if err := c.Bind(&blog); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		blog, err := u.useBlogusecase.Update(int(blog.ID), blog)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    blog,
		})
	}
}

func (u *blogController) DeleteBlog(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var blog model.Blog
		if err := c.Bind(&blog); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		err := u.useBlogusecase.Delete(int(blog.ID))
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
		})
	}
}
