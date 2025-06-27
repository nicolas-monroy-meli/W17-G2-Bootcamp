package repository

import (
	mod "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/pkg/models"
)

// NewSectionRepo creates a new instance of the Section repository
func NewSectionRepo(sections map[int]mod.Section) *SectionDB {
	return &SectionDB{
		db: sections,
	}
}

// SectionDB is the implementation of the Section database
type SectionDB struct {
	db map[int]mod.Section
}

// FindAll returns all sections from the database
func (r *SectionDB) FindAll() (sections map[int]mod.Section, err error) {

	return
}

// FindByID returns a section from the database by its id
func (r *SectionDB) FindByID(id int) (section mod.Section, err error) {

	return
}

// Save saves a section into the database
func (r *SectionDB) Save(section *mod.Section) (err error) {

	return
}

// Update updates a section in the database
func (r *SectionDB) Update(section *mod.Section) (err error) {

	return
}

// Delete deletes a section from the database by its id
func (r *SectionDB) Delete(id int) (err error) {

	return
}
