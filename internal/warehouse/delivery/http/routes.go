package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teatou/lamoda/internal/warehouse"
)

func MapWarehouseRoutes(router fiber.Router, h warehouse.Handlers) {
	router.Post("/", nil)
}
