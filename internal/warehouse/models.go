package warehouse

type Warehouse struct {
	Title        string `json:"title"`
	IsAccessible bool   `json:"accessable"`
}

type Item struct {
	Title    string `json:"title"`
	Size     int    `json:"size"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}

type ReserveItemsParams struct {
	Codes []string `json:"codes"`
}

type ReleaseItemsParams struct {
	Codes []string `json:"codes"`
}

type GetItemsLeftParams struct {
	WarehouseId int `json:"id"`
}

type GetItemsFromOneResponse struct {
	ItemTitle      string `json:"item_title" db:"item_title"`
	ItemCode       string `json:"item_code" db:"item_code"`
	Reserved       bool   `json:"reserved" db:"reserved"`
	Quantity       int    `json:"quantity" db:"quantity"`
	WarehouseTitle string `json:"warehouse_title" db:"warehouse_title"`
}

type GetItemFromAllWarehousesResponse struct {
	ItemTitle      string `json:"item_title" db:"item_title"`
	WarehouseTitle string `json:"warehouse_title" db:"warehouse_title"`
	Quantity       int    `json:"quantity" db:"quantity"`
}
