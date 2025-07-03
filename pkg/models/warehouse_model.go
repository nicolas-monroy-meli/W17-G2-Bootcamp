package models

type Warehouse struct {
	ID                 int     `json:"ID"`
	WarehouseCode      string  `json:"WarehouseCode"`
	Address            string  `json:"Address"`
	Telephone          string  `json:"Telephone"`
	MinimumCapacity    int     `json:"MinimumCapacity"`
	MinimumTemperature float64 `json:"MinimumTemperature"`
}
