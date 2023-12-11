package public

import (
	public_handler "ddd/src/routes/public/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app fiber.Router, db *gorm.DB) {
	public := app.Group("/public") // add middleware here if needed
	health := public_handler.NewPublicHealthHandler(db)
	public.Get("/health", health.Health)
}
