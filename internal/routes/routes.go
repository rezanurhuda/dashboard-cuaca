package routes

import (
	"dashboard-cuaca/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Endpoint Mock
	app.Get("/weather", handler.MockWeatherHandler)

	// Endpoint Dokumentasi
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/docs.html")
	})

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}
