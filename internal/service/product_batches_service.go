package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type ProductBatchService struct {
	rp internal.ProductBatchRepository
}

func NewProductBatchRepository(batchesRepo internal.ProductBatchRepository) *ProductBatchService {
	return &ProductBatchService{rp: batchesRepo}
}

func (s *ProductBatchService) FindAll() (batches []mod.ProductBatch, err error) {
	return s.rp.FindAll()
}

func (s *ProductBatchService) Save(batch *mod.ProductBatch) error {
	return s.rp.Save(batch)
}
