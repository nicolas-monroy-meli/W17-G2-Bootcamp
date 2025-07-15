package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// ProductRepository is an interface that contains the methods that the product repository should support
type ProductRepository interface {
	// FindAll returns all the products
	FindAll() (map[int]mod.Product, error)
	// FindByID returns the product with the given ID
	FindByID(id int) (mod.Product, error)
	// Save saves the given product
	Save(product *mod.Product) error
	// Update updates the given product
	Update(product *mod.Product) error
	// Delete deletes the product with the given ID
	Delete(id int) error
	// FindAllPR returns all product records from the database
	FindAllPR() (map[int]mod.ProductRecord, error)
	// FindAllByProductIDPR returns all product records for a given product ID
	FindAllByProductIDPR(productID int) (map[int]mod.ProductRecord, error)
	// SavePR saves the given product record
	SavePR(productRecord *mod.ProductRecord) error
}

// ProductService is an interface that contains the methods that the product service should support
type ProductService interface {
	// FindAll returns all the products
	FindAll() (map[int]mod.Product, error)
	// FindByID returns the product with the given ID
	FindByID(id int) (mod.Product, error)
	// Save saves the given product
	Save(product *mod.Product) error
	// Update updates the given product
	Update(product *mod.Product) error
	// Delete deletes the product with the given ID
	Delete(id int) error
	// FindAllPR returns all product records from the database
	FindAllPR() (map[int]mod.ProductRecord, error)
	// FindAllByProductIDPR returns all product records for a given product ID
	FindAllByProductIDPR(productID int) (map[int]mod.ProductRecord, error)
	// SavePR saves the given product record
	SavePR(productRecord *mod.ProductRecord) error
}

// ProductService is an interface that contains the methods that the buyer service should support
type ProductHandler interface {
	// FindAll returns all the buyers
	GetAll() http.HandlerFunc
	// FindByID returns the buyer with the given ID
	GetByID() http.HandlerFunc
	// Save saves the given buyer
	Create() http.HandlerFunc
	// Update updates the given buyer
	Update() http.HandlerFunc
	// Delete deletes the buyer with the given ID
	Delete() http.HandlerFunc
	// GetRecords returns all product records, if productId is provided, it returns the records for that product, if not, it returns all records
	GetRecords() http.HandlerFunc
	// CreateRecord creates a new product record
	CreateRecord() http.HandlerFunc
}
