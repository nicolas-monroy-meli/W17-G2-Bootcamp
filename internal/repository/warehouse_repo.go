package repository

import (
	"errors"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type WarehouseDB struct {
	db         map[int]mod.Warehouse
	carries    map[int]mod.Carry
	localities map[int]string
}

func NewWarehouseRepo(warehouses map[int]mod.Warehouse) *WarehouseDB {
	return &WarehouseDB{
		db:      warehouses,
		carries: make(map[int]mod.Carry),
		localities: map[int]string{ // ejemplo de localidades precargadas
			1: "Palermo",
			2: "Belgrano",
			3: "Caballito",
		},
	}
}

// Métodos existentes
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
	for _, wh := range r.db {
		if wh.WarehouseCode == warehouse.WarehouseCode {
			return errors.New("ya existe un warehouse con ese código")
		}
	}
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

///////////////////////////////////////////
//  Métodos nuevos para Carry
///////////////////////////////////////////

func (r *WarehouseDB) CreateCarry(c *mod.Carry) error {
	// Validación simple de duplicado por ID
	if _, exists := r.carries[c.ID]; exists {
		return errors.New("carry con ese ID ya existe")
	}
	r.carries[c.ID] = *c
	return nil
}

func (r *WarehouseDB) ExistsCID(cid int) bool {
	for _, c := range r.carries {
		if c.CID == cid {
			return true
		}
	}
	return false
}

func (r *WarehouseDB) ExistsLocality(id int) bool {
	_, ok := r.localities[id]
	return ok
}

func (r *WarehouseDB) ReportByLocality(id int) ([]mod.LocalityCarryReport, error) {
	count := 0
	for _, c := range r.carries {
		if c.LocalityID == id {
			count++
		}
	}

	name, ok := r.localities[id]
	if !ok {
		return nil, errors.New("localidad no encontrada")
	}

	report := mod.LocalityCarryReport{
		LocalityID:   id,
		LocalityName: name,
		CarriesCount: count,
	}

	return []mod.LocalityCarryReport{report}, nil
}
