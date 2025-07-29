package models

// Estructura
type Warehouse struct {
	ID                 int    `json:"ID"`
	WarehouseCode      string `json:"Warehouse_Code" validate:"required"`
	Address            string `json:"Address" validate:"required"`
	Telephone          string `json:"Telephone" validate:"required"`
	MinimumCapacity    int    `json:"Minimum_Capacity" validate:"required,min=1"`
	MinimumTemperature int    `json:"Minimum_Temperature" validate:"required,min=0"`
}
