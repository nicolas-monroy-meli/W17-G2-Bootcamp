package repository

import (
	"errors"
	"fmt"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"gorm.io/gorm"
)

type ProductBatchDB struct {
	db *gorm.DB
}

func NewProductBatchRepo(db *gorm.DB) *ProductBatchDB {
	return &ProductBatchDB{
		db: db,
	}
}

func (r *ProductBatchDB) FindAll() (batches []mod.ProductBatch, err error) {
	result := r.db.Find(&batches)
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %w", e.ErrRepositoryDatabase, result.Error)
	}
	if len(batches) == 0 {
		return nil, e.ErrEmptySectionDB
	}
	return batches, nil
}

func (r *ProductBatchDB) Save(batch *mod.ProductBatch) (err error) {
	result := r.db.Create(batch)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return e.ErrSectionRepositoryDuplicated
	}
	if result.Error != nil {
		return fmt.Errorf("%w: %w", e.ErrRepositoryDatabase, result.Error)
	}
	return
}
