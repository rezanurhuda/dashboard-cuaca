package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"dashboard-cuaca/internal/routes"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Setup routes
	routes.SetupRoutes(app)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
