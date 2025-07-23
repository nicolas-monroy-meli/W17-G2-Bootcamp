package internal

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"net/http"
)

type ProductBatchRepository interface {
	FindAll() (batches []mod.ProductBatch, err error)
	Save(batch *mod.ProductBatch) error
}

type ProductBatchService interface {
	FindAll() (batches []mod.ProductBatch, err error)
	Save(batch *mod.ProductBatch) error
}

type ProductBatchHandler interface {
	GetAll() http.HandlerFunc
	Create() http.HandlerFunc
}
