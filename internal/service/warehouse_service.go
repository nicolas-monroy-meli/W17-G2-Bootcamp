package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewWarehouseService creates a new instance of the warehouse service
func NewWarehouseService(warehouses internal.WarehouseRepository) *WarehouseService {
	return &WarehouseService{
		rp: warehouses,
	}
}

// WarehouseService is the default implementation of the warehouse service
type WarehouseService struct {
	// rp is the repository used by the service
	rp internal.WarehouseRepository
}

// FindAll returns all warehouses
func (s *WarehouseService) FindAll() (warehouses map[int]mod.Warehouse, err error) {
	return
}

// FindByID returns a warehouse
func (s *WarehouseService) FindByID(id int) (warehouse mod.Warehouse, err error) {
	return
}

// Save creates a new warehouse
func (s *WarehouseService) Save(warehouse *mod.Warehouse) (err error) {
	return
}

// Update updates a warehouse
func (s *WarehouseService) Update(warehouse *mod.Warehouse) (err error) {
	return
}

// Delete deletes a warehouse
func (s *WarehouseService) Delete(id int) (err error) {
	return
}
