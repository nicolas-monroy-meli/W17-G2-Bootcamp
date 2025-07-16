package models

// Section is a struct that contains the section's information
type Section struct {
	// ID is the unique identifier of the section
	ID int `json:"id,omitempty"`
	// SectionNumber is the number of the section
	SectionNumber int `json:"section_number" validate:"required,gt=0"`
	// CurrentTemperature is the current temperature of the section
	CurrentTemperature float64 `json:"current_temperature" validate:"gtefield=MinimumTemperature"`
	// MinimumTemperature is the minimum temperature that can be maintained in the section
	MinimumTemperature float64 `json:"minimum_temperature" validate:"required,ltefield=CurrentTemperature"`
	// CurrentCapacity is the current capacity of the section
	CurrentCapacity int `json:"current_capacity" validate:"required,gtefield=MinimumCapacity,ltefield=MaximumCapacity"`
	// MinimumCapacity is the minimum capacity of the section
	MinimumCapacity int `json:"minimum_capacity" validate:"required,gt=0"`
	// MaximumCapacity is the maximum capacity of the section
	MaximumCapacity int `json:"maximum_capacity" validate:"required,gtfield=MinimumCapacity"`
	// WarehouseID is the unique identifier of the warehouse to which the section belongs
	WarehouseID int `json:"warehouse_id" validate:"required,gte=1"`
	// ProductTypeID is the unique identifier of the type of product stored in the section
	ProductTypeID int `json:"product_type_id" validate:"required,gte=1"`
}

type SectionPatch struct {
	SectionNumber *int `json:"section_number,omitempty" validate:"omitempty,gt=0"`
	// CurrentTemperature is the current temperature of the section
	CurrentTemperature *float64 `json:"current_temperature,omitempty" validate:"omitempty,gtefield=MinimumTemperature"`
	// MinimumTemperature is the minimum temperature that can be maintained in the section
	MinimumTemperature *float64 `json:"minimum_temperature,omitempty" validate:"omitempty,ltefield=CurrentTemperature"`
	// CurrentCapacity is the current capacity of the section
	CurrentCapacity *int `json:"current_capacity,omitempty" validate:"omitempty,gtefield=MinimumCapacity,ltefield=MaximumCapacity"`
	// MinimumCapacity is the minimum capacity of the section
	MinimumCapacity *int `json:"minimum_capacity,omitempty" validate:"omitempty,gt=0"`
	// MaximumCapacity is the maximum capacity of the section
	MaximumCapacity *int `json:"maximum_capacity,omitempty" validate:"omitempty,gtfield=MinimumCapacity"`
	// WarehouseID is the unique identifier of the warehouse to which the section belongs
	WarehouseID *int `json:"warehouse_id,omitempty" validate:"omitempty,gte=1"`
	// ProductTypeID is the unique identifier of the type of product stored in the section
	ProductTypeID *int `json:"product_type_id,omitempty" validate:"omitempty,gte=1"`
}

type ReportProductsResponse struct {
	SectionId     int
	SectionNumber int
	ProductsCount int
}
