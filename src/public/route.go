package public

import (
	"ddd/src/public/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	app.Get("/health", handler.HealthHandler)
}
