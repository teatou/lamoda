package warehouse

type Repository interface {
	ReserveItems(codes []string) error
	ReleaseItems(codes []string) error
	GetItemsFromOne(warehouseId int) ([]GetItemsFromOneResponse, error)
	GetItemRemainder(itemId int) error
}
