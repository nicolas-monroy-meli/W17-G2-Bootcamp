package models

import (
	"time"
)

type ProductBatch struct {
	ID                 int       `json:"id,omitempty"`
	BatchNumber        int       `json:"batchNumber" validate:"required,gt=0"`
	CurrentQuantity    int       `json:"currentQuantity" validate:"required,gtefield=InitialQuantity"`
	InitialQuantity    int       `json:"initialQuantity" validate:"required,gt=0"`
	CurrentTemperature int       `json:"currentTemperature" validate:"required,gtefield=MinimumTemperature"`
	MinimumTemperature int       `json:"minimumTemperature" validate:"required"`
	DueDate            time.Time `json:"dueDate" validate:"required"`
	ManufacturingDate  time.Time `json:"manufacturingDate" validate:"required"`
	ManufacturingHour  string    `json:"manufacturingHour" validate:"required,hhmmss"`
	ProductId          int       `json:"productId" validate:"required,gt=0"`
	SectionId          int       `json:"sectionId" validate:"required,gt=0"`
}
