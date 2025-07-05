package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

// MockWeatherHandler mengembalikan data cuaca hard-coded
func MockWeatherHandler(c *fiber.Ctx) error {
	// Simulasikan penundaan seperti request real
	time.Sleep(300 * time.Millisecond)

	// Hard-coded response
	response := fiber.Map{
		"location": fiber.Map{
			"name":            "Bekasi",
			"region":          "West Java",
			"country":         "Indonesia",
			"lat":             -6.2349,
			"lon":             106.9896,
			"tz_id":           "Asia/Jakarta",
			"localtime_epoch": 1751497482,
			"localtime":       "2025-07-03 06:04",
		},
		"current": fiber.Map{
			"last_updated_epoch": 1751497200,
			"last_updated":       "2025-07-03 06:00",
			"temp_c":             25.1,
			"temp_f":             77.2,
			"is_day":             1,
			"condition": fiber.Map{
				"text": "Mist",
				"icon": "//cdn.weatherapi.com/weather/64x64/day/143.png",
				"code": 1030,
			},
			"wind_mph":    2.9,
			"wind_kph":    4.7,
			"wind_degree": 136,
			"wind_dir":    "SE",
			"pressure_mb": 1011,
			"pressure_in": 29.85,
			"precip_mm":   0,
			"precip_in":   0,
			"humidity":    100,
			"cloud":       75,
			"feelslike_c": 28,
			"feelslike_f": 82.4,
			"windchill_c": 24.8,
			"windchill_f": 76.6,
			"heatindex_c": 27.5,
			"heatindex_f": 81.5,
			"dewpoint_c":  23.2,
			"dewpoint_f":  73.7,
			"vis_km":      5,
			"vis_miles":   3,
			"uv":          0,
			"gust_mph":    4.8,
			"gust_kph":    7.7,
		},
	}

	return c.JSON(response)
}
