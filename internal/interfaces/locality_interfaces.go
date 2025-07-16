package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// LocalityRepository is an interface that contains the methods that the seller repository should support
type LocalityRepository interface {
	// FindByID returns the seller with the given ID
	FindAllLocalities() (result []mod.Locality, err error)

	FindSellersByLocID(id int) (result []mod.SelByLoc, err error)
	// Save saves the given locality
	Save(locality *mod.Locality) (id int, err error)
}

// LocalityService is an interface that contains the methods that the seller service should support
type LocalityService interface {
	// FindByID returns the seller with the given ID
	FindAllLocalities() (result []mod.Locality, err error)

	FindSellersByLocID(id int) (result []mod.SelByLoc, err error)
	// Save saves the given locality
	Save(locality *mod.Locality) (id int, err error)
}

// LocalityService is an interface that contains the methods that the seller service should support
type LocalityHandler interface {
	// FindAll returns all the sellers
	GetLocalities() http.HandlerFunc

	GetSelByLocID(id int) http.HandlerFunc
	// Save saves the given locality
	Create(locality *mod.Locality) http.HandlerFunc
}
