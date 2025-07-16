package internal

import (
	"net/http"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type CarryHandler interface {
	GetAll() http.HandlerFunc
	GetByID() http.HandlerFunc
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
	GetReportByLocality() http.HandlerFunc
}

type CarryService interface {
	FindAll() ([]models.Carry, error)
	FindByID(id int) (models.Carry, error)
	Create(c *models.Carry) error
	Update(c *models.Carry) error
	Delete(id int) error
	ReportByLocality(localityID int) ([]models.LocalityCarryReport, error)
	ReportByLocalityAll() ([]models.LocalityCarryReport, error)
}

type CarryRepository interface {
	GetAll() ([]models.Carry, error)
	GetByID(id int) (models.Carry, error)
	Save(c *models.Carry) error
	Update(c *models.Carry) error
	Delete(id int) error
	ExistsLocality(localityID int) (bool, error)
	ExistsCID(cid string) (bool, error)
	GetReportByLocality(localityID int) ([]models.LocalityCarryReport, error)
	GetByCID(cid string) (models.Carry, error)
	GetReportByLocalityAll() ([]models.LocalityCarryReport, error)
}
