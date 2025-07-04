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
		"city":    "Jakarta",
		"country": "Indonesia",
		"current": fiber.Map{
			"temp":      28.5,
			"condition": "Partly Cloudy",
			"icon":      "//cdn.weatherapi.com/weather/64x64/day/116.png",
		},
		"forecast": []fiber.Map{
			{
				"date": time.Now().Format("2006-01-02"),
				"day": fiber.Map{
					"maxtemp_c": 31.2,
					"mintemp_c": 24.5,
					"condition": fiber.Map{
						"text": "Moderate rain",
						"icon": "//cdn.weatherapi.com/weather/64x64/day/302.png",
					},
				},
			},
			{
				"date": time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
				"day": fiber.Map{
					"maxtemp_c": 32.1,
					"mintemp_c": 25.3,
					"condition": fiber.Map{
						"text": "Sunny",
						"icon": "//cdn.weatherapi.com/weather/64x64/day/113.png",
					},
				},
			},
		},
	}

	return c.JSON(response)
}
