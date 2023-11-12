package warehouse

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	CreateItemsReserve(codes []string) fiber.Handler
	ReleaseItemsReserve(codes []string) fiber.Handler
	GetLeftItems(warehouseId int) fiber.Handler
}
