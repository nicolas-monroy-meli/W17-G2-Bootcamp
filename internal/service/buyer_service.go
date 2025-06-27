package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewBuyerService creates a new instance of the buyer service
func NewBuyerService(buyers internal.BuyerRepository) *BuyerService {
	return &BuyerService{
		rp: buyers,
	}
}

// BuyerService is the default implementation of the buyer service
type BuyerService struct {
	// rp is the repository used by the service
	rp internal.BuyerRepository
}

// FindAll returns all buyers
func (s *BuyerService) FindAll() (buyers map[int]mod.Buyer, err error) {
	return
}

// FindByID returns a buyer
func (s *BuyerService) FindByID(id int) (buyer mod.Buyer, err error) {
	return
}

// Save creates a new buyer
func (s *BuyerService) Save(buyer *mod.Buyer) (err error) {
	return
}

// Update updates a buyer
func (s *BuyerService) Update(buyer *mod.Buyer) (err error) {
	return
}

// Delete deletes a buyer
func (s *BuyerService) Delete(id int) (err error) {
	return
}
