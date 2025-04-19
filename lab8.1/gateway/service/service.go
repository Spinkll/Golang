package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret")

func HandleAuth(c *fiber.Ctx) error {
	resp, err := http.Post(
		"http://localhost:3001/api/auth",
		"application/json",
		bytes.NewReader(c.Body()),
	)
	if err != nil {
		return c.Status(500).SendString("Auth service error")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return c.Status(resp.StatusCode).Send(body)
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
	}

	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}

	return c.Next()
}

func ProxyForecastNow(c *fiber.Ctx) error {
	lat := c.Query("lat")
	lon := c.Query("lon")

	url := fmt.Sprintf("http://localhost:3002/api/forecast/now?lat=%s&lon=%s", lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(500).SendString("Forecast service unavailable")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return c.Status(resp.StatusCode).Send(body)
}

func ProxyForecastHistory(c *fiber.Ctx) error {
	city := c.Query("city")
	start := c.Query("start")
	end := c.Query("end")

	url := fmt.Sprintf("http://localhost:3002/api/forecast/history?city=%s", city)
	if start != "" {
		url += fmt.Sprintf("&start=%s", start)
	}
	if end != "" {
		url += fmt.Sprintf("&end=%s", end)
	}

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(500).SendString("Forecast service unavailable")
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return c.Status(resp.StatusCode).Send(body)
}
