package service

import (
	"errors"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type warehouseService struct {
	repo internal.WarehouseRepository
}

// NewWarehouseService creates a new instance of the service
func NewWarehouseService(repo internal.WarehouseRepository) *warehouseService {
	return &warehouseService{
		repo: repo,
	}
}

// FindAll returns all warehouses
func (s *warehouseService) FindAll() (map[int]mod.Warehouse, error) {
	return s.repo.FindAll()
}

// FindByID returns a warehouse by ID
func (s *warehouseService) FindByID(id int) (mod.Warehouse, error) {
	return s.repo.FindByID(id)
}

// Save adds a new warehouse after checking for duplicate WarehouseCode
func (s *warehouseService) Save(w *mod.Warehouse) error {
	all, err := s.repo.FindAll()
	if err != nil {
		return err
	}

	// Verificar código duplicado
	for _, existing := range all {
		if existing.WarehouseCode == w.WarehouseCode {
			return errors.New("ya existe un almacén con ese código")
		}
	}

	return s.repo.Save(w)
}

// Update modifies an existing warehouse
func (s *warehouseService) Update(w *mod.Warehouse) error {
	_, err := s.repo.FindByID(w.ID)
	if err != nil {
		return err
	}

	return s.repo.Update(w)
}

// Delete removes a warehouse by ID
func (s *warehouseService) Delete(id int) error {
	return s.repo.Delete(id)
}
