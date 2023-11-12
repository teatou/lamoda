package warehouse

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	CreateItemsReserve() fiber.Handler
	ReleaseItemsReserve() fiber.Handler
	GetItemsLeft() fiber.Handler
}
