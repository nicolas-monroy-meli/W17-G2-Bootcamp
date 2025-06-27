package internal

import (
	"net/http"

	mod "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/pkg/models"
)

// WarehouseRepository is an interface that contains the methods that the warehouse repository should support
type WarehouseRepository interface {
	// FindAll returns all the warehouses
	FindAll() (map[int]mod.Warehouse, error)
	// FindByID returns the warehouse with the given ID
	FindByID(id int) (mod.Warehouse, error)
	// Save saves the given warehouse
	Save(warehouse *mod.Warehouse) error
	// Update updates the given warehouse
	Update(warehouse *mod.Warehouse) error
	// Delete deletes the warehouse with the given ID
	Delete(id int) error
}

// WarehouseService is an interface that contains the methods that the warehouse service should support
type WarehouseService interface {
	// FindAll returns all the warehouses
	FindAll() (map[int]mod.Warehouse, error)
	// FindByID returns the warehouse with the given ID
	FindByID(id int) (mod.Warehouse, error)
	// Save saves the given warehouse
	Save(warehouse *mod.Warehouse) error
	// Update updates the given warehouse
	Update(warehouse *mod.Warehouse) error
	// Delete deletes the warehouse with the given ID
	Delete(id int) error
}

// WarehouseService is an interface that contains the methods that the buyer service should support
type WarehouseHandler interface {
	// FindAll returns all the buyers
	GetAll() http.HandlerFunc
	// FindByID returns the buyer with the given ID
	GetByID(id int) http.HandlerFunc
	// Save saves the given buyer
	Create(buyer *mod.Warehouse) http.HandlerFunc
	// Update updates the given buyer
	Update(buyer *mod.Warehouse) http.HandlerFunc
	// Delete deletes the buyer with the given ID
	Delete(id int) http.HandlerFunc
}
