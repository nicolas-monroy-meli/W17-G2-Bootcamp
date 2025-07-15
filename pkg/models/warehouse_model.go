package models

// Estructura
type Warehouse struct {
	ID                 int     `json:"ID"`
	WarehouseCode      string  `json:"Warehouse_Code"`
	Address            string  `json:"Address"`
	Telephone          string  `json:"Telephone"`
	MinimumCapacity    int     `json:"Minimum_Capacity"`
	MinimumTemperature float64 `json:"Minimum_Temperature"`
}
