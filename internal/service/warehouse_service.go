package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

type warehouseService struct {
	repo internal.WarehouseRepository
}

func NewWarehouseService(repo internal.WarehouseRepository) *warehouseService {
	return &warehouseService{repo: repo}
}

// FindAll devuelve un slice de warehouses (no un mapa)
func (s *warehouseService) FindAll() ([]mod.Warehouse, error) {
	return s.repo.GetAll() // Usamos el método GetAll del repositorio
}

func (s *warehouseService) FindByID(id int) (mod.Warehouse, error) {
	return s.repo.GetByID(id) // Usamos GetByID del repositorio
}

func (s *warehouseService) Save(w *mod.Warehouse) error {
	// Usamos el método ExistsWarehouseCode del repositorio
	exists, err := s.repo.ExistsWarehouseCode(w.WarehouseCode)
	if err != nil {
		return err
	}
	if exists {
		return e.ErrWarehouseRepositoryDuplicated // Usamos el error definido
	}

	return s.repo.Save(w)
}

func (s *warehouseService) Update(w *mod.Warehouse) error {
	// Verificamos si el warehouse existe antes de actualizar
	_, err := s.repo.GetByID(w.ID)
	if err != nil {
		return err
	}

	// Verificamos si el nuevo código ya existe en otro registro
	existingWarehouse, err := s.repo.GetByWarehouseCode(w.WarehouseCode)
	if err != nil && err != e.ErrWarehouseRepositoryNotFound {
		return err
	}

	if existingWarehouse.ID != 0 && existingWarehouse.ID != w.ID {
		return e.ErrWarehouseRepositoryDuplicated
	}

	return s.repo.Update(w)
}

func (s *warehouseService) Delete(id int) error {
	return s.repo.Delete(id)
}
