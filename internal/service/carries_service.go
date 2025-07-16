package service

import (
	"errors"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
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
		return e.ErrCarryRepositoryLocalityNotFound
	}

	existsCID, err := s.repo.ExistsCID(c.CID)
	if err != nil {
		return errors.New("error verificando CID: " + err.Error())
	}
	if existsCID {
		return e.ErrCarryRepositoryDuplicated
	}

	// Guardamos el nuevo "carry"
	if err := s.repo.Save(c); err != nil {
		return err
	}

	return nil
}

func (s *carryService) FindAll() ([]mod.Carry, error) {
	carries, err := s.repo.GetAll()
	if err != nil || len(carries) == 0 {
		return nil, e.ErrCarryRepositoryNotFound

	}
	return carries, nil
}

func (s *carryService) FindByID(id int) (mod.Carry, error) {
	carry, err := s.repo.GetByID(id)
	if err != nil {
		return mod.Carry{}, e.ErrCarryRepositoryNotFound
	}
	return carry, nil
}

func (s *carryService) Update(c *mod.Carry) error {

	if err := s.repo.Update(c); err != nil {
		return err
	}
	return nil
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
