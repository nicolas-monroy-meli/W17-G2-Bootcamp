package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// BuyerRepository is an interface that contains the methods that the buyer repository should support
type BuyerRepository interface {
	// FindAll returns all the buyers
	FindAll() ([]mod.Buyer, error)
	// FindByID returns the buyer with the given ID
	FindByID(id int) (mod.Buyer, error)
	// Save saves the given buyer
	Save(buyer *mod.Buyer) error
	// Update updates the given buyer
	Update(buyer *mod.Buyer) error
	// Delete deletes the buyer with the given ID
	Delete(id int) error
}

// BuyerService is an interface that contains the methods that the buyer service should support
type BuyerService interface {
	// FindAll returns all the buyers
	FindAll() ([]mod.Buyer, error)
	// FindByID returns the buyer with the given ID
	FindByID(id int) (mod.Buyer, error)
	// Save saves the given buyer
	Save(buyer *mod.Buyer) error
	// Update updates the given buyer
	Update(buyer *mod.Buyer) error
	// Delete deletes the buyer with the given ID
	Delete(id int) error
}

// BuyerService is an interface that contains the methods that the buyer service should support
type BuyerHandler interface {
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
}
