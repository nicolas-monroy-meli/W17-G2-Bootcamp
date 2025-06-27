package repository

import (
	mod "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/pkg/models"
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

	return
}

// FindByID returns a product from the database by its id
func (r *ProductDB) FindByID(id int) (product mod.Product, err error) {

	return
}

// Save saves a product into the database
func (r *ProductDB) Save(product *mod.Product) (err error) {

	return
}

// Update updates a product in the database
func (r *ProductDB) Update(product *mod.Product) (err error) {

	return
}

// Delete deletes a product from the database by its id
func (r *ProductDB) Delete(id int) (err error) {

	return
}
