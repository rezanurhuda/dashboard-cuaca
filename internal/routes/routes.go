package routes

import (
	"dashboard-cuaca/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Dashboard Cuaca API",
			"version": "1.0.0",
			"endpoints": fiber.Map{
				"weather": "/weather",
				"docs":    "/docs",
			},
		})
	})

	// Endpoint Mock
	app.Get("/weather", handler.MockWeatherHandler)

	// Endpoint Dokumentasi
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/docs.html")
	})
}
