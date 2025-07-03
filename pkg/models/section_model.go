package models

// Section is a struct that contains the section's information
type Section struct {
	// ID is the unique identifier of the section
	ID int `json:"id" validate:"required"`
	// SectionNumber is the number of the section
	SectionNumber int `json:"sectionNumber" validate:"required,gt=0"`
	// CurrentTemperature is the current temperature of the section
	CurrentTemperature float64 `json:"currentTemperature" validate:"required,gtefield=MinimumTemperature"`
	// MinimumTemperature is the minimum temperature that can be maintained in the section
	MinimumTemperature float64 `json:"minimumTemperature" validate:"required"`
	// CurrentCapacity is the current capacity of the section
	CurrentCapacity int `json:"CurrentCapacity" validate:"required,gtefield=MinimumCapacity,ltefield=MaximumCapacity"`
	// MinimumCapacity is the minimum capacity of the section
	MinimumCapacity int `json:"minimumCapacity" validate:"required,gt=0"`
	// MaximumCapacity is the maximum capacity of the section
	MaximumCapacity int `json:"maximumCapacity" validate:"required,gtfield=MinimumCapacity"`
	// WarehouseID is the unique identifier of the warehouse to which the section belongs
	WarehouseID int `json:"warehouseID" validate:"required,gte=1"`
	// ProductTypeID is the unique identifier of the type of product stored in the section
	ProductTypeID int `json:"productTypeID" validate:"required,gte=1"`
}

type SectionPatch struct {
	SectionNumber *int
	// CurrentTemperature is the current temperature of the section
	CurrentTemperature *float64 `json:"currentTemperature"`
	// MinimumTemperature is the minimum temperature that can be maintained in the section
	MinimumTemperature *float64 `json:"minimumTemperature" `
	// CurrentCapacity is the current capacity of the section
	CurrentCapacity *int `json:"CurrentCapacity"`
	// MinimumCapacity is the minimum capacity of the section
	MinimumCapacity *int `json:"minimumCapacity"`
	// MaximumCapacity is the maximum capacity of the section
	MaximumCapacity *int `json:"maximumCapacity"`
	// WarehouseID is the unique identifier of the warehouse to which the section belongs
	WarehouseID *int `json:"warehouseID"`
	// ProductTypeID is the unique identifier of the type of product stored in the section
	ProductTypeID *int `json:"productTypeID"`
}
