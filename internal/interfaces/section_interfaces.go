package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// SectionRepository is an interface that contains the methods that the section repository should support
type SectionRepository interface {
	// FindAll returns all the sections
	FindAll() ([]mod.Section, error)
	//SectionExists returns a boolean that verifies if a section is in the db through its id
	SectionExists(id int, sectionNumber *int) (res bool, err error)
	// FindByID returns the section with the given ID
	FindByID(id int) (mod.Section, error)
	// Save saves the given section
	Save(section *mod.Section) error
	// Update updates the given section
	Update(section *mod.Section) error
	// Delete deletes the section with the given ID
	Delete(id int) error
}

// SectionService is an interface that contains the methods that the section service should support
type SectionService interface {
	// FindAll returns all the sections
	FindAll() ([]mod.Section, error)
	// FindByID returns the section with the given ID
	FindByID(id int) (mod.Section, error)
	// Save saves the given section
	Save(section *mod.Section) error
	// Update updates the given section
	Update(section *mod.Section) error
	// Delete deletes the section with the given ID
	Delete(id int) error
}

// SectionHandler is an interface that contains the methods that the section service should support
type SectionHandler interface {
	// GetAll returns all the sections
	GetAll() http.HandlerFunc
	// GetByID returns the section with the given ID
	GetByID() http.HandlerFunc
	// Create saves the given section
	Create() http.HandlerFunc
	// Update updates the given section
	Update() http.HandlerFunc
	// Delete deletes the section with the given ID
	Delete() http.HandlerFunc
}
