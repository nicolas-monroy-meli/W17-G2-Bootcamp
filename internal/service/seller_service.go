package service

import (
	internal "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/pkg/models"
)

// NewSellerService creates a new instance of the seller service
func NewSellerService(sellers internal.SellerRepository) *SellerService {
	return &SellerService{
		rp: sellers,
	}
}

// SellerService is the default implementation of the seller service
type SellerService struct {
	// rp is the repository used by the service
	rp internal.SellerRepository
}

// FindAll returns all sellers
func (s *SellerService) FindAll() (sellers map[int]mod.Seller, err error) {
	return s.rp.FindAll()
}

// FindByID returns a seller
func (s *SellerService) FindByID(id int) (seller mod.Seller, err error) {
	return
}

// Save creates a new seller
func (s *SellerService) Save(seller *mod.Seller) (err error) {
	return
}

// Update updates a seller
func (s *SellerService) Update(seller *mod.Seller) (err error) {
	return
}

// Delete deletes a seller
func (s *SellerService) Delete(id int) (err error) {
	return
}
