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
	SectionExists(id int) (bool, error)
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

// SectionHandler is an interface that contains the methods that the buyer service should support
type SectionHandler interface {
	// GetAll returns all the buyers
	GetAll() http.HandlerFunc
	// GetByID returns the buyer with the given ID
	GetByID(id int) http.HandlerFunc
	// Create saves the given buyer
	Create(buyer *mod.Section) http.HandlerFunc
	// Update updates the given buyer
	Update(buyer *mod.Section) http.HandlerFunc
	// Delete deletes the buyer with the given ID
	Delete(id int) http.HandlerFunc
}
