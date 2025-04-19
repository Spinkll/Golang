package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"weather/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Отримати поточну погоду за координатами
	app.Get("/api/forecast/now", func(c *fiber.Ctx) error {
		lat := c.Query("lat")
		lon := c.Query("lon")

		if lat == "" || lon == "" {
			return c.Status(400).JSON(fiber.Map{"error": "lat and lon are required"})
		}

		apiKey := "babe5359cad9e2a96757184d8e28c631"
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, apiKey)

		resp, err := http.Get(url)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch weather"})
		}
		defer resp.Body.Close()

		var weather map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to parse weather"})
		}

		return c.JSON(weather)
	})

	app.Get("/api/forecast/history", func(c *fiber.Ctx) error {
		city := c.Query("city")
		start := c.Query("start")
		end := c.Query("end")

		if city == "" {
			return c.Status(400).JSON(fiber.Map{"error": "city is required"})
		}

		query := db.Where("city ILIKE ?", city) // нечутливий до регістру

		if start != "" {
			if t, err := time.Parse(time.RFC3339, start); err == nil {
				query = query.Where("timestamp >= ?", t)
			}
		}
		if end != "" {
			if t, err := time.Parse(time.RFC3339, end); err == nil {
				query = query.Where("timestamp <= ?", t)
			}
		}

		var forecasts []models.Forecast
		if err := query.Find(&forecasts).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch data"})
		}

		return c.JSON(forecasts)
	})

}
