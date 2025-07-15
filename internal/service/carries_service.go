package service

import (
	"errors"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type carryService struct {
	repo internal.CarryRepository
}

func NewCarryService(repo internal.CarryRepository) internal.CarryService {
	return &carryService{repo: repo}
}

func (s *carryService) Create(c *mod.Carry) error {
	existsLocality, err := s.repo.ExistsLocality(c.LocalityID)
	if err != nil {
		return errors.New("error verificando localidad: " + err.Error())
	}
	if !existsLocality {
		return errors.New("localidad no existe")
	}

	existsCID, err := s.repo.ExistsCID(c.CID)
	if err != nil {
		return errors.New("error verificando CID: " + err.Error())
	}
	if existsCID {
		return errors.New("el CID ya está registrado")
	}

	return s.repo.Save(c)
}

func (s *carryService) FindAll() ([]mod.Carry, error) {
	return s.repo.GetAll()
}

func (s *carryService) FindByID(id int) (mod.Carry, error) {
	return s.repo.GetByID(id)
}

func (s *carryService) Update(c *mod.Carry) error {
	return s.repo.Update(c)
}

func (s *carryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *carryService) ReportByLocality(localityID int) ([]mod.LocalityCarryReport, error) {
	return s.repo.GetReportByLocality(localityID)
}

func (s *carryService) ReportByLocalityAll() ([]mod.LocalityCarryReport, error) {
	return s.repo.GetReportByLocalityAll()
}
