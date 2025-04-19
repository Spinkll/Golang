package main

import (
	"log"
	"weather/api"
	"weather/config"
	"weather/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.InitDB()
	go services.StartWeatherFetching(db)

	app := fiber.New()
	api.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3002"))
}
