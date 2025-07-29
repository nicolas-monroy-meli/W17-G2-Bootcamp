package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
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
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
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
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestWrongBody.Error())
			return
		}

		errValidate := e.ValidateStruct(req)
		if len(errValidate) > 0 {
			str := ""
			for _, err := range errValidate {
				str += err + ", "
			}
			utils.BadResponse(w, http.StatusUnprocessableEntity, str)
			return
		}

		id, err := h.sv.Save(&req)
		if err != nil {
			utils.BadResponse(w, http.StatusConflict, err.Error())
			return
		}
		utils.GoodResponse(w, 201, "success", id)
	}
}

// Update updates a seller
func (h *SellerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.SellerPatch
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		currentSeller, err := h.sv.FindByID(id)
		if err != nil {
			utils.BadResponse(w, 404, err.Error())
			return
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestWrongBody.Error())
			return
		}
		req.ID = id
		seller := common.PatchSeller(currentSeller, req)

		errValidate := e.ValidateStruct(seller)
		if len(errValidate) > 0 {
			str := ""
			for _, err := range errValidate {
				str += err + ", "
			}
			utils.BadResponse(w, http.StatusUnprocessableEntity, str)
			return
		}

		err = h.sv.Update(seller)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.GoodResponse(w, 200, "success", nil)

	}
}

// Delete deletes a seller
func (h *SellerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}
		err = h.sv.Delete(req)
		if err != nil {
			utils.BadResponse(w, 404, err.Error())
			return
		}
		utils.GoodResponse(w, 204, "success", nil)
	}
}
