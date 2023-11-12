package warehouse

type Usecase interface {
	CreateItemsReserve(codes []string) error
	ReleaseItemsReserve(codes []string) error
	GetItemsLeft(warehouseId int) ([]GetItemsFromOneResponse, error)
	GetItemFromAllWarehouses(itemId int) ([]GetItemFromAllWarehousesResponse, error)
}
