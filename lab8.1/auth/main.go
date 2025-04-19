package main

import (
	"auth/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/api/auth", service.Login)
	app.Listen(":3001")
}
