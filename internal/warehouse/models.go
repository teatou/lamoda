package warehouse

type Warehouse struct {
	Name         string `json:"name"`
	IsAccessible bool   `json:"accessable"`
}

type Item struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity"`
}
