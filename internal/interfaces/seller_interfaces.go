package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// SellerRepository is an interface that contains the methods that the seller repository should support
type SellerRepository interface {
	// FindAll returns all the sellers
	FindAll() (map[int]mod.Seller, error)
	// FindByID returns the seller with the given ID
	FindByID(id int) (mod.Seller, error)
	// Save saves the given seller
	Save(seller *mod.Seller) error
	// Update updates the given seller
	Update(seller *mod.Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
}

// SellerService is an interface that contains the methods that the seller service should support
type SellerService interface {
	// FindAll returns all the sellers
	FindAll() (map[int]mod.Seller, error)
	// FindByID returns the seller with the given ID
	FindByID(id int) (mod.Seller, error)
	// Save saves the given seller
	Save(seller *mod.Seller) error
	// Update updates the given seller
	Update(seller *mod.Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
	//Valdiates Patch Request
	PatchValidator(seller, newSeller mod.Seller) mod.Seller
}

// SellerService is an interface that contains the methods that the buyer service should support
type SellerHandler interface {
	// FindAll returns all the buyers
	GetAll() http.HandlerFunc
	// FindByID returns the buyer with the given ID
	GetByID(id int) http.HandlerFunc
	// Save saves the given buyer
	Create(buyer *mod.Seller) http.HandlerFunc
	// Update updates the given buyer
	Update(buyer *mod.Seller) http.HandlerFunc
	// Delete deletes the buyer with the given ID
	Delete(id int) http.HandlerFunc
	//Valdiates Patch Request
	PatchValidator(seller, newSeller mod.Seller) mod.Seller
}
