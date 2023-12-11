// main.go
package main

import (
	"ddd/src/config"
	"ddd/src/routes/public"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ENV_FILE = ".env"

func main() {
	// get env config
	cfg, err := config.GetConfig(ENV_FILE)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}

	// Create a new Fiber instance
	app := fiber.New()
	api := app.Group("/api")

	// setup routes
	public.SetupRoutes(api, db)

	// Start the Fiber application on port 3000
	if err := app.Listen(":3000"); err != nil {
		// Print any errors that occur during startup
		panic(err)
	}
}

func connectDB(cfg config.IConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost(),
		cfg.DBPort(),
		cfg.DBUser(),
		cfg.DBPassword(),
		cfg.DBName(),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
