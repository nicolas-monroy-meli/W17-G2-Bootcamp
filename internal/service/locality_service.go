package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewLocalityService creates a new instance of the seller service
func NewLocalityService(localities internal.LocalityRepository) *LocalityService {
	return &LocalityService{
		rp: localities,
	}
}

// LocalityService is the default implementation of the seller service
type LocalityService struct {
	// rp is the repository used by the service
	rp internal.LocalityRepository
}

// FindAll returns all sellers
func (s *LocalityService) FindAllLocalities() (result []mod.Locality, err error) {
	return s.rp.FindAllLocalities()
}

func (s *LocalityService) FindSellersByLocID(id int) (result []mod.SelByLoc, err error) {
	return s.rp.FindSellersByLocID(id)
}

// Save creates a new locality
func (s *LocalityService) Save(locality *mod.Locality) (id int, err error) {
	return s.rp.Save(locality)
}
