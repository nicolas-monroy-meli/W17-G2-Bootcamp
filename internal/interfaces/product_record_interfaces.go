package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// ProductRecordRepository is an interface that contains the methods that the product record repository should support
type ProductRecordRepository interface {
	// FindAllPR returns all product records from the database
	FindAllPR() (map[int]mod.ProductRecord, error)
	// FindAllByProductIDPR returns all product records for a given product ID
	FindAllByProductIDPR(productID int) (map[int]mod.ProductRecord, error)
	// SavePR saves the given product record
	SavePR(productRecord *mod.ProductRecord) error
}

// ProductRecordService is an interface that contains the methods that the product record service should support
type ProductRecordService interface {
	// FindAllPR returns all product records from the database
	FindAllPR() (map[int]mod.ProductRecord, error)
	// FindAllByProductIDPR returns all product records for a given product ID
	FindAllByProductIDPR(productID int) (map[int]mod.ProductRecord, error)
	// SavePR saves the given product record
	SavePR(productRecord *mod.ProductRecord) error
	// FindProductByID retrieves a product by its ID
	FindProductByID(id int) (mod.Product, error)
}

// ProductRecordHandler is an interface that contains the methods that the product record service should support
type ProductRecordHandler interface {
	// GetRecords returns all product records, if productId is provided, it returns the records for that product, if not, it returns all records
	GetRecords() http.HandlerFunc
	// CreateRecord creates a new product record
	CreateRecord() http.HandlerFunc
}
