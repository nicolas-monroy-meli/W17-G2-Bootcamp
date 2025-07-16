package common

import (
	"errors"
	"strings"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// ValidateWarehouseUpdate valida que los campos requeridos estén presentes para el put
func ValidateWarehouseUpdate(w models.Warehouse) error {
	if strings.TrimSpace(w.Address) == "" {
		return errors.New("el campo 'address' es requerido")
	}
	if strings.TrimSpace(w.Telephone) == "" {
		return errors.New("el campo 'telephone' es requerido")
	}
	if strings.TrimSpace(w.WarehouseCode) == "" {
		return errors.New("el campo 'warehouse_code' es requerido")
	}
	if w.MinimumCapacity <= 0 {

		return errors.New("el campo 'minimum_capacity' debe ser mayor a 0")
	}
	// Validamos temperatura real, no solo que no sea cero
	if w.MinimumTemperature < -50 || w.MinimumTemperature > 50 {
		return errors.New("el campo 'minimum_temperature' debe estar entre -50 y 50")
	}
	return nil
}
