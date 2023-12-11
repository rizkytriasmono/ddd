package public

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	return db
}

func TestSetupRoutes(t *testing.T) {
	app := fiber.New()
	db := setupTestDB()

	SetupRoutes(app, db)

	// Create a test request to the /public/health endpoint
	req := httptest.NewRequest(http.MethodGet, "/public/health", nil)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Check if the response status code is 200 OK
	assert.Equal(t, http.StatusOK, res.StatusCode)

}
