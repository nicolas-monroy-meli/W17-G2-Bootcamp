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
}

type WarehouseService interface {
	FindAll() ([]models.Warehouse, error) // Cambiado a slice
	FindByID(id int) (models.Warehouse, error)
	Save(warehouse *models.Warehouse) error
	Update(warehouse *models.Warehouse) error
	Delete(id int) error
}

type WarehouseRepository interface {
	GetAll() ([]models.Warehouse, error)
	GetByID(id int) (models.Warehouse, error)
	GetByWarehouseCode(code string) (models.Warehouse, error)
	Save(wh *models.Warehouse) error
	Update(wh *models.Warehouse) error
	Delete(id int) error
	ExistsWarehouseCode(code string) (bool, error)
}
