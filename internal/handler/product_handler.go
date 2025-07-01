package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
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
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestIdMustBeInt.Error())
			return
		}
		result, err := h.sv.FindByID(id)
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
			utils.BadResponse(w, http.StatusUnprocessableEntity, utils.ErrRequestWrongBody.Error()+"\n"+errValidate.Error())
			return
		}

		err = h.sv.Save(&req)
		if errors.Is(err, utils.ErrBuyerRepositoryDuplicated) {
			utils.BadResponse(w, http.StatusConflict, err.Error())
			return
		}
		if err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusCreated, "success", nil)
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
