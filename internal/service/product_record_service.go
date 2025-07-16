package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewProductRecordService creates a new instance of the product service
func NewProductRecordService(productsRecords internal.ProductRecordRepository, products internal.ProductRepository) *ProductRecordService {
	return &ProductRecordService{
		rp:  productsRecords,
		prp: products,
	}
}

// ProductRecordService is the default implementation of the product service
type ProductRecordService struct {
	// rp is the repository used by the service
	rp  internal.ProductRecordRepository
	prp internal.ProductRepository
}

// FindAllPR returns all product records from the repository
func (s *ProductRecordService) FindAllPR() (productRecords map[int]mod.ProductRecord, err error) {
	return s.rp.FindAllPR()
}

// FindAllByProductIDPR returns all product records for a given product ID
func (s *ProductRecordService) FindAllByProductIDPR(productID int) (productRecords map[int]mod.ProductRecord, err error) {
	return s.rp.FindAllByProductIDPR(productID)
}

// SavePR creates a new product record
func (s *ProductRecordService) SavePR(productRecord *mod.ProductRecord) (err error) {
	_, err = s.prp.FindByID(productRecord.ProductID)
	if err != nil {
		return e.ErrProductRepositoryNotFound
	}
	return s.rp.SavePR(productRecord)
}

// FindProductByID retrieves a product by its ID
func (s *ProductRecordService) FindProductByID(id int) (mod.Product, error) {
	return s.prp.FindByID(id)
}
