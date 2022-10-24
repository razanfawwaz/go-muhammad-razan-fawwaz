package middlewares

import (
	"belajar-go-echo/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

func CreateToken(userEmail string, e echo.Context) (string, error) {
	claims := jwt.MapClaims{
		"authorization": true,
		"userEmail":     userEmail,
		"exp":           time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	data := config.LoadENV()
	secret := data["jwtSecret"]

	return token.SignedString([]byte(secret))
}
