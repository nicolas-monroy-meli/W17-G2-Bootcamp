package repository

import (
	"database/sql"
	"fmt"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"strconv"
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
	rows, err := r.db.Query("SELECT `id`, `section_number`,`current_temperature`,`minimum_temperature`,`current_capacity`, `minimum_capacity`,`maximum_capacity`,`warehouse_id`,`product_type_id` FROM `sections`")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var section mod.Section
		err = rows.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
		if err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}

	if len(sections) == 0 {
		return nil, errors.ErrEmptySectionDB
	}
	return sections, nil
}

// SectionExists returns a boolean that verifies if a section is in the db through its id
func (r *SectionDB) SectionExists(id int, sectionNumber *int) (res bool, err error) {
	var exists bool
	if sectionNumber != nil {
		query := `SELECT EXISTS (SELECT 1 FROM sections WHERE id = ? or section_number=?)`
		err = r.db.QueryRow(query, id, *sectionNumber).Scan(&exists)
	} else {
		query := `SELECT EXISTS (SELECT 1 FROM sections WHERE id = ?)`
		err = r.db.QueryRow(query, id).Scan(&exists)
	}
	if err != nil {
		return false, err
	}
	return exists, nil
}

// FindByID returns a section from the database by its id
func (r *SectionDB) FindByID(id int) (section mod.Section, err error) {
	row := r.db.QueryRow("SELECT `id`, `section_number`,`current_temperature`,`minimum_temperature`,`current_capacity`, `minimum_capacity`,`maximum_capacity`,`warehouse_id`,`product_type_id`  FROM `sections` WHERE `id`=?", id)

	err = row.Scan(&section.ID, &section.SectionNumber, &section.CurrentTemperature, &section.MinimumTemperature, &section.CurrentCapacity, &section.MinimumCapacity, &section.MaximumCapacity, &section.WarehouseID, &section.ProductTypeID)
	if err != nil {
		return mod.Section{}, errors.ErrSectionRepositoryNotFound
	}
	return section, nil
}

// Save saves a section into the database
func (r *SectionDB) Save(section *mod.Section) (err error) {
	if exists, _ := r.SectionExists(section.ID, &section.SectionNumber); exists {
		return errors.ErrSectionRepositoryDuplicated
	}
	result, err := r.db.Exec(
		"INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		(*section).SectionNumber, (*section).CurrentTemperature, (*section).MinimumTemperature, (*section).CurrentCapacity, (*section).MinimumCapacity, (*section).MaximumCapacity, (*section).WarehouseID, (*section).ProductTypeID,
	)
	if err != nil {
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
func (r *SectionDB) Update(id int, fields map[string]interface{}) (result *mod.Section, err error) {
	//Build query
	query, args := common.BuildPatchQuery("sections", fields, strconv.Itoa(id))
	// execute the query
	res, err := r.db.Exec(query, args...)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sec, err := r.FindByID(id)
	if err != nil {
		return nil, errors.ErrSectionRepositoryNotFound
	}
	rowsAffected, _ := res.RowsAffected()
	if int(rowsAffected) == 0 {
		return nil, errors.NoRowsAffected
	}
	return &sec, nil
}

// Delete deletes a section from the database by its id
func (r *SectionDB) Delete(id int) (err error) {
	if exists, _ := r.SectionExists(id, nil); !exists {
		return errors.ErrSectionRepositoryNotFound
	}
	// execute the query
	_, err = r.db.Exec("DELETE FROM `sections` WHERE `id` = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return nil
}

func (r *SectionDB) ReportProducts(ids []int) ([]mod.ReportProductsResponse, error) {
	query, args := common.GetQueryReport(ids)
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]mod.ReportProductsResponse, 0)
	foundIDs := make(map[int]bool)

	for rows.Next() {
		var result mod.ReportProductsResponse
		err := rows.Scan(&result.SectionId, &result.SectionNumber, &result.ProductsCount)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
		foundIDs[result.SectionId] = true
	}

	// If filtering by IDs, ensure missing sections are added
	if len(ids) > 0 {
		for _, id := range ids {
			if !foundIDs[id] {
				results = append(results, mod.ReportProductsResponse{
					SectionId:     id,
					SectionNumber: 0,
					ProductsCount: 0,
				})
			}
		}
	}

	return results, nil
}
