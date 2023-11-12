package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teatou/lamoda/internal/config"
	"github.com/teatou/lamoda/internal/warehouse"
	"github.com/teatou/lamoda/pkg/logger"
)

type WarehouseHandler struct {
	cfg         *config.Config
	warehouseUC warehouse.Usecase
	logger      *logger.Logger
}

func NewWarehouseHandlers(cfg *config.Config, warehouseUC warehouse.Usecase, logger *logger.Logger) warehouse.Handlers {
	return &WarehouseHandler{
		cfg:         cfg,
		logger:      logger,
		warehouseUC: warehouseUC,
	}
}

func (w WarehouseHandler) CreateItemsReserve(codes []string) fiber.Handler {
	return nil
}
func (w WarehouseHandler) ReleaseItemsReserve(codes []string) fiber.Handler {
	return nil
}

func (w WarehouseHandler) GetLeftItems(warehouseId int) fiber.Handler {
	return nil
}
