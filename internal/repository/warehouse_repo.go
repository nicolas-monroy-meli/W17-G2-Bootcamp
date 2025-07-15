package repository

import (
	"errors"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type WarehouseDB struct {
	db map[int]mod.Warehouse
}

func NewWarehouseRepo(warehouses map[int]mod.Warehouse) *WarehouseDB {
	return &WarehouseDB{
		db: warehouses,
	}
}

func (r *WarehouseDB) FindAll() (map[int]mod.Warehouse, error) {
	return r.db, nil
}

func (r *WarehouseDB) FindByID(id int) (mod.Warehouse, error) {
	wh, exists := r.db[id]
	if !exists {
		return mod.Warehouse{}, errors.New("almacén no encontrado")
	}
	return wh, nil
}

func (r *WarehouseDB) Save(warehouse *mod.Warehouse) error {
	// Verificamos si ya existe un warehouse con el mismo código
	for _, wh := range r.db {
		if wh.WarehouseCode == warehouse.WarehouseCode {
			return errors.New("ya existe un warehouse con ese código")
		}
	}
	// Creamos un ID nuevo (máximo ID + 1)
	maxID := 0
	for id := range r.db {
		if id > maxID {
			maxID = id
		}
	}
	warehouse.ID = maxID + 1
	r.db[warehouse.ID] = *warehouse
	return nil
}

func (r *WarehouseDB) Update(warehouse *mod.Warehouse) error {
	_, exists := r.db[warehouse.ID]
	if !exists {
		return errors.New("warehouse no encontrado")
	}
	r.db[warehouse.ID] = *warehouse
	return nil
}

func (r *WarehouseDB) Delete(id int) error {
	_, exists := r.db[id]
	if !exists {
		return errors.New("warehouse no encontrado")
	}
	delete(r.db, id)
	return nil
}
