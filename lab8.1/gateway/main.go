package main

import (
	"gateway/middleware"
	"gateway/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/api/auth", service.HandleAuth)
	forecast := app.Group("/api/forecast")
	forecast.Use(middleware.JWTMiddleware)
	forecast.Get("/now", service.ProxyForecastNow)
	forecast.Get("/history", service.ProxyForecastHistory)
	app.Listen(":3000")
}
