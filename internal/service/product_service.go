package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewProductService creates a new instance of the product service
func NewProductService(products internal.ProductRepository) *ProductService {
	return &ProductService{
		rp: products,
	}
}

// ProductService is the default implementation of the product service
type ProductService struct {
	// rp is the repository used by the service
	rp internal.ProductRepository
}

// FindAll returns all products
func (s *ProductService) FindAll() (products []mod.Product, err error) {
	return s.rp.FindAll()
}

// FindByID returns a product
func (s *ProductService) FindByID(id int) (product mod.Product, err error) {
	return s.rp.FindByID(id)
}

// Save creates a new product
func (s *ProductService) Save(product *mod.Product) (err error) {
	return s.rp.Save(product)
}

// Update updates a product
func (s *ProductService) Update(product *mod.Product) (err error) {
	return s.rp.Update(product)
}

// Delete deletes a product
func (s *ProductService) Delete(id int) (err error) {
	return s.rp.Delete(id)
}
