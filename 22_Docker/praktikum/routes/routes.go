package routes

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.POST("/users", userController.Create(config.DB))
	e.POST("/login", userController.Auth(config.DB))
	e.GET("/users", userController.Find(config.DB))

	a := e.Group("/users/jwt")
	data := config.LoadENV()
	a.Use(middleware.JWT([]byte(data["JWT_SECRET"])))
	a.GET("", userController.Find(config.DB))
	return e
}
