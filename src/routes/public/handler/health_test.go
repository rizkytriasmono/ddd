package public_handler

import (
	"testing"

	httptest "net/http/httptest"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PublicHealthHandlerSuite struct {
	suite.Suite
	handler IPublicHealthHandler
	app     *fiber.App
}

func (suite *PublicHealthHandlerSuite) SetupTest() {
	// Use an in-memory SQLite database for testing.
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	}

	// Create the handler with the in-memory database.
	suite.handler = NewPublicHealthHandler(db)
	suite.app = fiber.New()
	suite.app.Get("/health", suite.handler.Health)
}

func (suite *PublicHealthHandlerSuite) TestHealth_Success() {
	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := suite.app.Test(req, 1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)
}

type MockDB struct {
	mock.Mock
}

func (m *MockDB) GormDB() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Ping() error {
	args := m.Called()
	return args.Error(0)
}

func (suite *PublicHealthHandlerSuite) TestHealth_DBError() {
	// Create a handler with a mocked database connection error.
	gormDB := &gorm.DB{
		Config:       &gorm.Config{},
		Error:        nil,
		RowsAffected: 0,
		Statement:    &gorm.Statement{},
	} // Customize this based on your needs

	suite.handler = NewPublicHealthHandler(gormDB)
	suite.app = fiber.New()
	suite.app.Get("/health", suite.handler.Health)

	// Create the Fiber app with the mocked handler's route.
	app := fiber.New()
	app.Get("/health", suite.handler.Health)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := suite.app.Test(req, 1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), fiber.StatusInternalServerError, resp.StatusCode)

	// Add more assertions as needed for the response body.
}

func TestPublicHealthHandlerSuite(t *testing.T) {
	suite.Run(t, new(PublicHealthHandlerSuite))
}
