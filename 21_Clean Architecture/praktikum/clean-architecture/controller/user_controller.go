package controller

import (
	"belajar-go-echo/middlewares"
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController interface{}

type userController struct {
	useUserusecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) Create(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		user, err := u.useUserusecase.Create(user)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    user,
		})
	}
}

func (u *userController) Find(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := u.useUserusecase.Find()
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(200, echo.Map{
			"message": "success",
			"data":    users,
		})
	}
}

func (u *userController) Auth(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"message": err.Error(),
			})
		}

		user, err := u.useUserusecase.Auth(user)
		{
			if err != nil {
				return c.JSON(500, echo.Map{
					"message": err.Error(),
				})
			}
		}

		// create token jwt
		token, err := middlewares.CreateToken(user.Email, c)
		if err != nil {
			return c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}

		c.Set("userEmail", user.Email)

		return c.JSON(200, echo.Map{
			"message": "success",
			"token":   token,
			"data":    user,
		})
	}
}
