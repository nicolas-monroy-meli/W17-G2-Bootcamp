package repository

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
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
	if len(r.db) == 0 {
		return nil, utils.ErrEmptySectionDB
	}
	return r.db, nil
}

// SectionExists returns a boolean that verifies if a section is in the db through its id
func (r *SectionDB) SectionExists(id int) bool {
	_, exists := r.db[id]
	return exists
}

// FindByID returns a section from the database by its id
func (r *SectionDB) FindByID(id int) (section mod.Section, err error) {
	if !r.SectionExists(id) {
		return mod.Section{}, utils.ErrSectionRepositoryNotFound
	}
	return r.db[id], nil
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
	if !r.SectionExists(id) {
		return utils.ErrSectionRepositoryNotFound
	}
	delete(r.db, id)
	return nil
}
