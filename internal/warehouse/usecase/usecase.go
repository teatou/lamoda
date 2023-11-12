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

func (w WarehouseUsecase) CreateItemsReserve(codes []string) error {
	err := w.warehouseRepo.ReserveItems(codes)
	return err
}
func (w WarehouseUsecase) ReleaseItemsReserve(codes []string) error {
	err := w.warehouseRepo.ReleaseItems(codes)
	return err
}
func (w WarehouseUsecase) GetItemsLeft(warehouseId int) ([]warehouse.GetItemsFromOneResponse, error) {
	items, err := w.warehouseRepo.GetItemsFromOne(warehouseId)
	if err != nil {
		return nil, err
	}
	return items, err
}

func (w WarehouseUsecase) GetItemFromAllWarehouses(itemId int) ([]warehouse.GetItemFromAllWarehousesResponse, error) {
	items, err := w.warehouseRepo.GetItemRemainder(itemId)
	if err != nil {
		return nil, err
	}
	return items, err
}
