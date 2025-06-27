package service

import (
	internal "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/pkg/models"
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
func (s *SectionService) FindAll() (sections map[int]mod.Section, err error) {
	return
}

// FindByID returns a section
func (s *SectionService) FindByID(id int) (section mod.Section, err error) {
	return
}

// Save creates a new section
func (s *SectionService) Save(section *mod.Section) (err error) {
	return
}

// Update updates a section
func (s *SectionService) Update(section *mod.Section) (err error) {
	return
}

// Delete deletes a section
func (s *SectionService) Delete(id int) (err error) {
	return
}
