package repository

import (
	"errors"
	"fmt"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"gorm.io/gorm"
)

type SectionDB struct {
	db *gorm.DB
}

// NewSectionRepo creates a new instance of the Section repository
func NewSectionRepo(db *gorm.DB) *SectionDB {
	return &SectionDB{
		db: db,
	}
}

// SectionDB is the implementation of the Section database

// FindAll returns all sections from the database
func (r *SectionDB) FindAll() (sections []mod.Section, err error) {
	//this query returns a gorm.DB obj which holds the errors
	result := r.db.Find(&sections)

	if len(sections) == 0 {
		return nil, e.ErrEmptySectionDB
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%w: %w", e.ErrRepositoryDatabase, result.Error)
	}
	return sections, nil
}

// FindByID returns a section from the database by its id
func (r *SectionDB) FindByID(id int) (section mod.Section, err error) {
	result := r.db.First(&section, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return mod.Section{}, e.ErrSectionRepositoryNotFound
	} else if result.Error != nil {
		return mod.Section{}, fmt.Errorf("%w: %w", e.ErrRepositoryDatabase, result.Error)
	}
	return section, nil
}

// Save saves a section into the database
func (r *SectionDB) Save(section *mod.Section) (err error) {
	result := r.db.Create(&section)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return e.ErrSectionRepositoryDuplicated
	} else if result.Error != nil {
		return fmt.Errorf("%w: %w", e.ErrRepositoryDatabase, result.Error)
	}
	return
}

// Update updates a section in the database
func (r *SectionDB) Update(id int, fields map[string]interface{}) (res *mod.Section, err error) {
	result := r.db.Model(&mod.Section{}).Where("id = ?", id).Updates(fields)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, e.ErrSectionRepositoryNotFound
	} else if result.Error != nil {
		return nil, fmt.Errorf("%w: %w", e.ErrRepositoryDatabase, result.Error)
	}
	sec, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &sec, nil
}

// Delete deletes a section from the database by its id
func (r *SectionDB) Delete(id int) (err error) {
	result := r.db.Delete(&mod.Section{}, id)

	fmt.Println(result)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return e.ErrSectionRepositoryNotFound
	} else if result.Error != nil {
		return fmt.Errorf("%w: %w", e.ErrRepositoryDatabase, result.Error)
	}

	return nil
}

func (r *SectionDB) ReportProducts(ids []int) ([]mod.ReportProductsResponse, error) {
	rows, err := r.db.Table("sections s").Select("s.id,s.section_number").Joins("LEFT JOIN products p ON p.id = s.product_type_id").Rows()
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
