package models

// Estructura
type Warehouse struct {
	ID                 int     `json:"ID"`
	WarehouseCode      string  `json:"WarehouseCode"`
	Address            string  `json:"Address"`
	Telephone          string  `json:"Telephone"`
	MinimumCapacity    int     `json:"MinimumCapacity"`
	MinimumTemperature float64 `json:"MinimumTemperature"`
}
type Carry struct {
	ID         int    `json:"id"`
	CID        string `json:"cid" validate:"required"`
	LocalityID int    `json:"locality_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Address    string `json:"address" validate:"required"`
}

type LocalityCarryReport struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int    `json:"carries_count"`
}
