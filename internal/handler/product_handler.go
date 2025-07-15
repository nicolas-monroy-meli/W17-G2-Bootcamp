package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
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
		result, err := h.sv.FindAll()
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, "success", result)
	}
}

// GetByID returns a product
func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}
		result, err := h.sv.FindByID(id)
		if errors.Is(err, e.ErrProductRepositoryNotFound) {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, "success", result)
	}
}

// Create creates a new product
func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.Product

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestNoBody.Error())
			return
		}

		var validate = validator.New(validator.WithRequiredStructEnabled())
		errValidate := validate.Struct(req)
		if errValidate != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, e.ErrRequestWrongBody.Error()+"\n"+errValidate.Error())
			return
		}

		err = h.sv.Save(&req)
		if errors.Is(err, e.ErrProductRepositoryDuplicated) {
			utils.BadResponse(w, http.StatusConflict, err.Error())
			return
		}
		if err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusCreated, "success", req)
	}
}

// Update updates a product
func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.ProductPatch

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		currentProduct, err := h.sv.FindByID(id)
		if errors.Is(err, e.ErrProductRepositoryNotFound) {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestNoBody.Error())
			return
		}

		common.PatchProduct(&currentProduct, req)
		var validate = validator.New(validator.WithRequiredStructEnabled())
		errValidate := validate.Struct(req)
		if errValidate != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, errValidate.Error())
			return
		}

		err = h.sv.Update(&currentProduct)
		if errors.Is(err, e.ErrProductRepositoryNotFound) {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		if err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, "success", currentProduct)
	}
}

// Delete deletes a product
func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}
		err = h.sv.Delete(id)
		if errors.Is(err, e.ErrProductRepositoryNotFound) {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		if err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
