package repository

import (
	"github.com/smartineztri_meli/W17-G2-Bootcamp/docs"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
)

// NewProductRepo creates a new instance of the Product repository
func NewProductRepo(products map[int]mod.Product) *ProductDB {
	return &ProductDB{
		db: products,
	}
}

// ProductDB is the implementation of the Product database
type ProductDB struct {
	db map[int]mod.Product
}

// FindAll returns all products from the database
func (r *ProductDB) FindAll() (products map[int]mod.Product, err error) {
	result := r.db
	if len(r.db) == 0 {
		return nil, utils.ErrProductRepositoryNotFound
	}
	return result, nil
}

// FindByID returns a product from the database by its id
func (r *ProductDB) FindByID(id int) (product mod.Product, err error) {
	val, ok := r.db[id]
	if !ok {
		return mod.Product{}, utils.ErrProductRepositoryNotFound
	}
	return val, nil
}

// Save saves a product into the database
func (r *ProductDB) Save(product *mod.Product) (err error) {
	for _, v := range r.db {
		if v.ProductCode == product.ProductCode {
			return utils.ErrProductRepositoryDuplicated
		}
	}
	product.ID = len(r.db) + 1
	r.db[product.ID] = *product
	docs.WriterFile("products.json", r.db)
	return nil
}

// Update updates a product in the database
func (r *ProductDB) Update(product *mod.Product) (err error) {
	// Check if the product exists
	if _, exists := r.db[product.ID]; !exists {
		return utils.ErrProductRepositoryNotFound
	}
	// Update the product in the database
	r.db[product.ID] = *product
	docs.WriterFile("products.json", r.db)
	// Return nil to indicate success
	return nil
}

// Delete deletes a product from the database by its id
func (r *ProductDB) Delete(id int) (err error) {
	// Check if the product exists
	if _, exists := r.db[id]; !exists {
		return utils.ErrProductRepositoryNotFound
	}
	// Delete the product from the database
	delete(r.db, id)
	docs.WriterFile("products.json", r.db)
	// Return nil to indicate success
	return nil
}
