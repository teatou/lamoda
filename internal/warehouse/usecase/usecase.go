package usecase

import "github.com/teatou/lamoda/internal/warehouse"

type WarehouseUsecase struct {
	warehouseRepo warehouse.Repository
}

func NewWarehouseUsecase(warehouseRepo warehouse.Repository) warehouse.Usecase {
	return &WarehouseUsecase{
		warehouseRepo: warehouseRepo,
	}
}
