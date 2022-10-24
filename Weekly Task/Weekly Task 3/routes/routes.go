package routes

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserValidator struct {
	validator *validator.Validate
}

func (v *UserValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &UserValidator{validator: validator.New()}

	// user repo
	userRepo := repository.NewUserRepository(config.DB)
	userService := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userService)

	blogRepo := repository.NewBlogRepository(config.DB)
	blogService := usecase.NewBlogUsecase(blogRepo)
	blogController := controller.NewBlogController(blogService)

	e.POST("/signup", userController.Create(config.DB))
	e.POST("/login", userController.Auth(config.DB))
	e.GET("/blogs", blogController.FindAll(config.DB))
	e.GET("/blogs/:id", blogController.FindById(config.DB))
	e.POST("/blogs", blogController.CreateBlog(config.DB))
	e.PUT("/blogs/:id", blogController.UpdateBlog(config.DB))
	e.DELETE("/blogs/:id", blogController.DeleteBlog(config.DB))
	
	return e
}
