package service

import (
	"auth/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret")

func Login(c *fiber.Ctx) error {
	var req model.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	if req.Login != "admin" || req.Password != "admin" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   req.Login,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return c.Status(500).SendString("Token generation failed")
	}

	return c.JSON(fiber.Map{"token": tokenStr})
}
