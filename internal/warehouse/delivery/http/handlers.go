package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/teatou/lamoda/internal/config"
	"github.com/teatou/lamoda/internal/warehouse"
	"github.com/teatou/lamoda/pkg/logger"
)

type WarehouseHandler struct {
	cfg         *config.Config
	warehouseUC warehouse.Usecase
	logger      logger.Logger
}

func NewWarehouseHandlers(cfg *config.Config, warehouseUC warehouse.Usecase, logger logger.Logger) warehouse.Handlers {
	return &WarehouseHandler{
		cfg:         cfg,
		logger:      logger,
		warehouseUC: warehouseUC,
	}
}

func (w WarehouseHandler) CreateItemsReserve() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params warehouse.ReserveItemsParams
		if err := c.BodyParser(&params); err != nil {
			w.logger.Error(err)
			return c.Status(http.StatusBadRequest).SendString("bad body to parse")
		}

		err := w.warehouseUC.CreateItemsReserve(params.Codes)
		if err != nil {
			w.logger.Error(err)
			return c.Status(http.StatusBadRequest).SendString("error reserving items")
		}

		return c.Status(http.StatusOK).SendString("OK")
	}
}

func (w WarehouseHandler) ReleaseItemsReserve() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params warehouse.ReleaseItemsParams
		if err := c.BodyParser(&params); err != nil {
			w.logger.Error(err)
			return c.Status(http.StatusBadRequest).SendString("bad body to parse")
		}

		err := w.warehouseUC.ReleaseItemsReserve(params.Codes)
		if err != nil {
			w.logger.Error(err)
			return c.Status(http.StatusBadRequest).SendString("error releasing reserve")
		}

		return c.Status(http.StatusOK).SendString("OK")
	}
}

func (w WarehouseHandler) GetItemsLeft() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params warehouse.GetItemsLeftParams
		if err := c.BodyParser(&params); err != nil {
			w.logger.Error(err)
			return c.Status(http.StatusBadRequest).SendString("bad body to parse")
		}

		items, err := w.warehouseUC.GetItemsLeft(params.WarehouseId)
		if err != nil {
			w.logger.Error(err)
			return c.Status(http.StatusBadRequest).SendString("error getting items left")
		}

		return c.Status(http.StatusOK).JSON(items)
	}
}
