package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
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
			utils.BadResponse(w, 400, err.Error())
			return
		}
		utils.GoodResponse(w, 200, "succes", result)
	}
}

// GetByID returns a seller
func (h *SellerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		req, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestIdMustBeInt.Error())
			return
		}
		result, err := h.sv.FindByID(req)
		if err != nil {
			utils.BadResponse(w, 404, err.Error())
			return
		}
		utils.GoodResponse(w, 200, "success", result)
	}
}

// Create creates a new seller
func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.Seller
		body, err := io.ReadAll(r.Body)
		if err != nil || body == nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestFailedBody.Error())
			return
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestNoBody.Error())
			return
		}

		var validate = validator.New(validator.WithRequiredStructEnabled())
		errValidate := validate.Struct(req)
		if errValidate != nil {
			utils.BadResponse(w, 422,
				utils.ErrRequestWrongBody.Error()+"\n"+errValidate.Error())
			return
		}

		err = h.sv.Save(&req)
		if err != nil {
			utils.BadResponse(w, http.StatusConflict, err.Error())
			return
		}
		utils.GoodResponse(w, 201, "success", nil)
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
