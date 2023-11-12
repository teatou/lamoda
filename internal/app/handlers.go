package app

import (
	"github.com/gofiber/fiber/v2"
	warehouseHttp "github.com/teatou/lamoda/internal/warehouse/delivery/http"
	warehouseRepository "github.com/teatou/lamoda/internal/warehouse/repository"
	warehouseUsecase "github.com/teatou/lamoda/internal/warehouse/usecase"
	"github.com/teatou/lamoda/pkg/logger"
	storage_postgres "github.com/teatou/lamoda/pkg/storage/postgres"
)

func (a *App) MapHandlers(router *fiber.App, logger logger.Logger) error {
	db, err := storage_postgres.InitPsqlDB(a.cfg.Postgres)
	if err != nil {
		return nil
	}

	warehouseRepo := warehouseRepository.NewPostgresRepository(db, logger)
	warehouseUC := warehouseUsecase.NewWarehouseUsecase(warehouseRepo)
	warehouseHandlers := warehouseHttp.NewWarehouseHandlers(a.cfg, warehouseUC, logger)
	warehouseRouter := router.Group("/warehouse")
	warehouseHttp.MapWarehouseRoutes(warehouseRouter, warehouseHandlers)

	return nil
}
