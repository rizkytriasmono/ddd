// main.go
package main

import (
	"ddd/src/public"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// setup routes
	public.SetupRoutes(app)

	// Start the Fiber application on port 3000
	if err := app.Listen(":3000"); err != nil {
		// Print any errors that occur during startup
		panic(err)
	}
}
