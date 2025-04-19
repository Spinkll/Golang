package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"weather/models"

	"gorm.io/gorm"
)

type OpenWeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Visibility int `json:"visibility"`
}

var locations = map[string][2]float64{
	"Kyiv":   {50.45, 30.52},
	"Lviv":   {49.83, 24.02},
	"Odesa":  {46.48, 30.72},
	"Dnipro": {48.46, 35.04},
}

func StartWeatherFetching(db *gorm.DB) {
	for {
		for name, coords := range locations {
			saveWeather(db, name, coords[0], coords[1])
		}
		time.Sleep(1 * time.Hour)
	}
}

func saveWeather(db *gorm.DB, name string, lat, lon float64) {
	apiKey := "babe5359cad9e2a96757184d8e28c631"
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching weather:", err)
		return
	}
	defer resp.Body.Close()

	var weather OpenWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Println("Error decoding weather response:", err)
		return
	}

	forecast := models.Forecast{
		Timestamp:  time.Now(),
		City:       weather.Name,
		Temp:       weather.Main.Temp,
		FeelsLike:  weather.Main.FeelsLike,
		Pressure:   weather.Main.Pressure,
		Humidity:   weather.Main.Humidity,
		WindSpeed:  weather.Wind.Speed,
		WindDeg:    weather.Wind.Deg,
		Clouds:     weather.Clouds.All,
		Visibility: weather.Visibility,
		Weather:    weather.Weather[0].Main,
		Lat:        lat,
		Lon:        lon,
	}
	db.Create(&forecast)
}
