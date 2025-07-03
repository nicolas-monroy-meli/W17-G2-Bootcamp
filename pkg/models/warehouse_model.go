package models

type Warehouse struct {
	ID                 int     `json:"id"`
	Address            string  `json:"address" validate:"required"`
	Telephone          string  `json:"telephone" validate:"required"`
	WarehouseCode      string  `json:"warehouse_code" validate:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" validate:"required"`
	MinimumTemperature float64 `json:"minimum_temperature" validate:"required"`
}
