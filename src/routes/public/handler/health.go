package public_handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HealthMessage struct {
	DB     string `json:"db"`
	Status int    `json:"status"`
}

type IPublicHealthHandler interface {
	Health(c *fiber.Ctx) error
}

type publicHealthHandler struct {
	db *gorm.DB
}

// Health implements IPublicHealthHandler.
func (h *publicHealthHandler) Health(c *fiber.Ctx) error {

	resp := HealthMessage{
		Status: fiber.StatusOK,
		DB:     "Ok!",
	}

	h.checkDBConnection(&resp)

	return c.Status(resp.Status).JSON(resp)
}

func (h *publicHealthHandler) checkDBConnection(resp *HealthMessage) {
	db, _ := h.db.DB()
	if db == nil || db.Ping() != nil {
		resp.Status = fiber.StatusInternalServerError
		resp.DB = "Can not connect"
	}
}

func NewPublicHealthHandler(db *gorm.DB) IPublicHealthHandler {
	return &publicHealthHandler{
		db: db,
	}
}
