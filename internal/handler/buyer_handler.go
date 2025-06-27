package handler

import (
	"net/http"

	internal "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/internal/interfaces"
)

// NewBuyerHandler creates a new instance of the buyer handler
func NewBuyerHandler(sv internal.BuyerService) *BuyerHandler {
	return &BuyerHandler{
		sv: sv,
	}
}

// BuyerHandler is the default implementation of the buyer handler
type BuyerHandler struct {
	// sv is the service used by the handler
	sv internal.BuyerService
}

// GetAll returns all buyers
func (h *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a buyer
func (h *BuyerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new buyer
func (h *BuyerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a buyer
func (h *BuyerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a buyer
func (h *BuyerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
