package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigMethods(t *testing.T) {
	c := &config{
		db: dbConfig{
			user:     "testuser",
			password: "testpassword",
			host:     "localhost",
			port:     "5432",
			name:     "testdb",
		},
	}

	assert.Equal(t, "testuser", c.DBUser())
	assert.Equal(t, "testpassword", c.DBPassword())
	assert.Equal(t, "localhost", c.DBHost())
	assert.Equal(t, "5432", c.DBPort())
	assert.Equal(t, "testdb", c.DBName())
}

func TestGetConfig(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("DB_USERNAME", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")

	// Ensure environment variables are cleared after the test
	defer func() {
		os.Unsetenv("DB_USERNAME")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_NAME")
	}()

	// Test GetConfig function
	conf, err := GetConfig("../../.env")

	// Assert that no error occurred during configuration loading
	assert.NoError(t, err)

	// Assert that the returned configuration matches the expected values
	assert.Equal(t, "testuser", conf.DBUser())
	assert.Equal(t, "testpassword", conf.DBPassword())
	assert.Equal(t, "localhost", conf.DBHost())
	assert.Equal(t, "5432", conf.DBPort())
	assert.Equal(t, "testdb", conf.DBName())
}

func TestGetConfigError(t *testing.T) {

	// Unset environment variables to simulate an error loading the .env file
	os.Unsetenv("DB_USERNAME")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_NAME")

	// Test GetConfig function
	conf, err := GetConfig("not found")

	// Assert that an error occurred during configuration loading
	assert.Error(t, err)
	assert.Nil(t, conf)
}
