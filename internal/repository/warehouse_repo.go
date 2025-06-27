package repository

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewWarehouseRepo creates a new instance of the Warehouse repository
func NewWarehouseRepo(warehouses map[int]mod.Warehouse) *WarehouseDB {
	return &WarehouseDB{
		db: warehouses,
	}
}

// WarehouseDB is the implementation of the Warehouse database
type WarehouseDB struct {
	db map[int]mod.Warehouse
}

// FindAll returns all warehouses from the database
func (r *WarehouseDB) FindAll() (warehouses map[int]mod.Warehouse, err error) {

	return
}

// FindByID returns a warehouse from the database by its id
func (r *WarehouseDB) FindByID(id int) (warehouse mod.Warehouse, err error) {

	return
}

// Save saves a warehouse into the database
func (r *WarehouseDB) Save(warehouse *mod.Warehouse) (err error) {

	return
}

// Update updates a warehouse in the database
func (r *WarehouseDB) Update(warehouse *mod.Warehouse) (err error) {

	return
}

// Delete deletes a warehouse from the database
func (r *WarehouseDB) Delete(id int) (err error) {

	return
}
