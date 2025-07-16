package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// SellerRepository is an interface that contains the methods that the seller repository should support
type SellerRepository interface {
	// FindAll returns all the sellers
	FindAll() (sellers []mod.Seller, err error)
	// FindByID returns the seller with the given ID
	FindByID(id int) (mod.Seller, error)
	// Save saves the given seller
	Save(seller *mod.Seller) (id int, err error)
	// Update updates the given seller
	Update(seller *mod.Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
}

// SellerService is an interface that contains the methods that the seller service should support
type SellerService interface {
	// FindAll returns all the sellers
	FindAll() (sellers []mod.Seller, err error)
	// FindByID returns the seller with the given ID
	FindByID(id int) (mod.Seller, error)
	// Save saves the given seller
	Save(seller *mod.Seller) (id int, err error)
	// Update updates the given seller
	Update(seller *mod.Seller) error
	// Delete deletes the seller with the given ID
	Delete(id int) error
}

// SellerService is an interface that contains the methods that the seller service should support
type SellerHandler interface {
	// FindAll returns all the sellers
	GetAll() http.HandlerFunc
	// FindByID returns the seller with the given ID
	GetByID(id int) http.HandlerFunc
	// Save saves the given se
	Create(seller *mod.Seller) http.HandlerFunc
	// Update updates the given seller
	Update(seller *mod.Seller) http.HandlerFunc
	// Delete deletes the seller with the given ID
	Delete(id int) http.HandlerFunc
}
