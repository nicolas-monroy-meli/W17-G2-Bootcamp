package repository

import (
	"database/sql"
	"fmt"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewSectionRepo creates a new instance of the Section repository
func NewSectionRepo(db *sql.DB) *SectionDB {
	return &SectionDB{
		db: db,
	}
}

// SectionDB is the implementation of the Section database
type SectionDB struct {
	db *sql.DB
}

// FindAll returns all sections from the database
func (r *SectionDB) FindAll() (sections []mod.Section, err error) {
	rows, err := r.db.Query("SELECT `id`, `sectionNumber`,`currentTemperature`,`minimumTemperature`,`currentCapacity`, `minimumCapacity`,`maximumCapacity`,`warehouseID`,`productTypeID` FROM `sections`")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var section mod.Section
		err = rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		sections = append(sections, section)
	}
	return
}

// SectionExists returns a boolean that verifies if a section is in the db through its id
func (r *SectionDB) SectionExists(id int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM sections WHERE id = ?)`
	err := r.db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// FindByID returns a section from the database by its id
func (r *SectionDB) FindByID(id int) (section mod.Section, err error) {
	if exists, _ := r.SectionExists(id); !exists {
		return mod.Section{}, errors.ErrSectionRepositoryNotFound
	}

	row := r.db.QueryRow("SELECT `id`, `sectionNumber`,`currentTemperature`,`minimumTemperature`,`currentCapacity`,`minimumCapacity`,`maximumCapacity`,`warehouseID`,`productTypeID`  FROM `sections` WHERE `id`=?", id)

	err = row.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		fmt.Println(err.Error())
		return mod.Section{}, err
	}
	return section, nil
}

// Save saves a section into the database
func (r *SectionDB) Save(section *mod.Section) (err error) {
	if exists, _ := r.SectionExists(section.SectionNumber); exists {
		return errors.ErrSectionRepositoryDuplicated
	}
	result, err := r.db.Exec(
		"INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		(*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// get the id of the inserted section
	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	// set the id of the section
	(*section).ID = int(id)

	return
}

// Update updates a section in the database
func (r *SectionDB) Update(section *mod.Section) (err error) {
	if exists, _ := r.SectionExists(section.SectionNumber); !exists {
		return errors.ErrSectionRepositoryNotFound
	}
	// execute the query
	_, err = r.db.Exec(
		"UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?",
		(*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID, (*section).ID,
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

// Delete deletes a section from the database by its id
func (r *SectionDB) Delete(id int) (err error) {
	if exists, _ := r.SectionExists(id); !exists {
		return errors.ErrSectionRepositoryNotFound
	}
	// execute the query
	_, err = r.db.Exec("DELETE FROM `sections` WHERE `id` = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
