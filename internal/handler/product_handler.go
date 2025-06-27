package handler

import (
	"net/http"

	internal "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/internal/interfaces"
)

// NewProductHandler creates a new instance of the product handler
func NewProductHandler(sv internal.ProductService) *ProductHandler {
	return &ProductHandler{
		sv: sv,
	}
}

// ProductHandler is the default implementation of the product handler
type ProductHandler struct {
	// sv is the service used by the handler
	sv internal.ProductService
}

// GetAll returns all products
func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a product
func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new product
func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a product
func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a product
func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
