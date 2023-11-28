package handler

import (
	"ddd/src/public/response"

	"github.com/gofiber/fiber/v2"
)

func HealthHandler(c *fiber.Ctx) error {

	// logic to check health

	resp := response.HealthMessage{
		Status: "Ok!",
	}

	return c.JSON(resp)
}
