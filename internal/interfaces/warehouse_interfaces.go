package internal

import (
	"net/http"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type WarehouseHandler interface {
	GetAll() http.HandlerFunc
	GetByID() http.HandlerFunc
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	CreateCarry() http.HandlerFunc
	GetReportByLocalityID() http.HandlerFunc
}

type WarehouseService interface {
	FindAll() (map[int]models.Warehouse, error)
	FindByID(id int) (models.Warehouse, error)
	Save(warehouse *models.Warehouse) error
	Update(warehouse *models.Warehouse) error
	Delete(id int) error

	// Métodos relacionados a "carry"
	SaveCarry(carry *models.Carry) error
	CIDExists(cid string) bool
	LocalityExists(localityID int) bool
	GetCarriesReportByLocality(localityID int) ([]models.LocalityCarryReport, error)
}

type WarehouseRepository interface {
	FindAll() (map[int]models.Warehouse, error)
	FindByID(id int) (models.Warehouse, error)
	Save(warehouse *models.Warehouse) error
	Update(warehouse *models.Warehouse) error
	Delete(id int) error
}
