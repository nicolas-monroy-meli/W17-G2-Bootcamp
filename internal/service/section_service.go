package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewSectionService creates a new instance of the section service
func NewSectionService(sections internal.SectionRepository) *SectionService {
	return &SectionService{
		rp: sections,
	}
}

// SectionService is the default implementation of the section service
type SectionService struct {
	// rp is the repository used by the service
	rp internal.SectionRepository
}

// FindAll returns all sections
func (s *SectionService) FindAll() (sections []mod.Section, err error) {
	return s.rp.FindAll()
}

// FindByID returns a section
func (s *SectionService) FindByID(id int) (section mod.Section, err error) {

	return s.rp.FindByID(id)
}

// Save creates a new section
func (s *SectionService) Save(section *mod.Section) (err error) {
	return s.rp.Save(section)
}

// Update updates a section
func (s *SectionService) Update(id int, fields map[string]interface{}) (section *mod.Section, err error) {
	return s.rp.Update(id, fields)
}

// Delete deletes a section
func (s *SectionService) Delete(id int) (err error) {
	return s.rp.Delete(id)
}

func (s *SectionService) ReportProducts(ids []int) ([]mod.ReportProductsResponse, error) {
	return s.rp.ReportProducts(ids)
}
