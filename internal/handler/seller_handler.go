package handler

import (
	"net/http"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
)

// NewSellerHandler creates a new instance of the seller handler
func NewSellerHandler(sv internal.SellerService) *SellerHandler {
	return &SellerHandler{
		sv: sv,
	}
}

// SellerHandler is the default implementation of the seller handler
type SellerHandler struct {
	// sv is the service used by the handler
	sv internal.SellerService
}

// GetAll returns all sellers
func (h *SellerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.FindAll()
		if err != nil {
			utils.BadResponse(w, 400, "unable to get sellers")
			return
		}
		utils.GoodResponse(w, 200, "succes", result)
	}
}

// GetByID returns a seller
func (h *SellerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new seller
func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a seller
func (h *SellerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a seller
func (h *SellerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
