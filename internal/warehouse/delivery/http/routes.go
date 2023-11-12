package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teatou/lamoda/internal/warehouse"
)

func MapWarehouseRoutes(router fiber.Router, h warehouse.Handlers) {
	router.Post("/items/reserve", h.CreateItemsReserve())
	router.Post("/items/release", h.ReleaseItemsReserve())
	router.Post("/items/left", h.GetItemsLeft())
	router.Post("/item", h.GetItemFromAllWarehouses())
}
