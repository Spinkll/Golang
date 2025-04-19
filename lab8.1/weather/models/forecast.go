package models

import "time"

type Forecast struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	City       string    `json:"city"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	Weather    string    `json:"weather"`
	Lat        float64   `json:"lat"`
	Lon        float64   `json:"lon"`
}
