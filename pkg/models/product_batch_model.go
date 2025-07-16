package models

import (
	"time"
)

type ProductBatch struct {
	ID                 int       `json:"id,omitempty"`
	BatchNumber        int       `json:"batch_number" validate:"required,gt=0"`
	CurrentQuantity    int       `json:"current_quantity" validate:"required,gtefield=InitialQuantity"`
	InitialQuantity    int       `json:"initial_quantity" validate:"required,gt=0"`
	CurrentTemperature int       `json:"current_temperature" validate:"required,gtefield=MinimumTemperature"`
	MinimumTemperature int       `json:"minimum_temperature" validate:"required"`
	DueDate            time.Time `json:"due_date" validate:"required"`
	ManufacturingDate  time.Time `json:"manufacturing_date" validate:"required"`
	ManufacturingHour  string    `json:"manufacturing_hour" validate:"required,hhmmss"`
	ProductId          int       `json:"product_id" validate:"required,gt=0"`
	SectionId          int       `json:"section_id" validate:"required,gt=0"`
}
