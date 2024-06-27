package middleware

import (
	"fmt"
	"kenshilabs_demo/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("kenshilabskey")

func AuthMiddleware(c *fiber.Ctx) error {
	var tokenString string

	if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Result{
			Status:  false,
			Message: "unauthorized",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return sampleSecretKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Result{
			Status:  false,
			Message: "unauthorized",
		})
	}
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Result{
			Status:  false,
			Message: "unauthorized",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Result{
			Status:  false,
			Message: "unauthorized",
		})
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Result{
			Status:  false,
			Message: "unauthorized",
		})
	}

	return c.Next()
}
