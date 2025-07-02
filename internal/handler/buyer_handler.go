package handler

import (
	"net/http"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
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
		buyers, err := h.sv.FindAll()

		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, "", buyers)
		return
	}
}

// GetByID returns a buyer
func (h *BuyerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestIdMustBeInt.Error())
			return
		}

		buyer, err := h.sv.FindByID(id)

		if err != nil {
			switch {
			case errors.Is(err, utils.ErrBuyerRepositoryNotFound):
				utils.BadResponse(w, http.StatusNotFound, err.Error())
			default:
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		utils.GoodResponse(w, http.StatusCreated, "Buyer obtenido con exito", buyer)
		return
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
