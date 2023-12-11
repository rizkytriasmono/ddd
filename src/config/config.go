package config

import (
	"os"

	"github.com/joho/godotenv"
)

type IConfig interface {
	DBUser() string
	DBPassword() string
	DBHost() string
	DBPort() string
	DBName() string
}

type dbConfig struct {
	user     string
	password string
	host     string
	port     string
	name     string
}

type config struct {
	db dbConfig
}

// DBName implements IConfig.
func (c *config) DBName() string {
	return c.db.name
}

// DBHost implements IConfig.
func (c *config) DBHost() string {
	return c.db.host
}

// DBPassword implements IConfig.
func (c *config) DBPassword() string {
	return c.db.password
}

// DBPort implements IConfig.
func (c *config) DBPort() string {
	return c.db.port
}

// DBUser implements IConfig.
func (c *config) DBUser() string {
	return c.db.user
}

func GetConfig(file string) (IConfig, error) {
	// Load environment variables from the .env file
	if err := godotenv.Load(file); err != nil {
		return nil, err
	}

	db := dbConfig{
		user:     os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("DB_NAME"),
	}
	return &config{
		db: db,
	}, nil
}
