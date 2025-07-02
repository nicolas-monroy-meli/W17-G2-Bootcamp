package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"net/http"
	"strconv"
	"strings"
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
		var newBuyer mod.Buyer

		if err := json.NewDecoder(r.Body).Decode(&newBuyer); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestFailedBody.Error())
			return
		}

		errValidation := utils.ValidateStruct(newBuyer)

		//validate := validator.New()
		//err := validate.Struct(newBuyer)

		if errValidation != nil {
			//err = err.(validator.ValidationErrors)
			str := make([]string, 0, len(errValidation))
			for _, err := range errValidation {
				str = append(str, err)
			}

			err := fmt.Errorf("%w: %v", utils.ErrRequestWrongBody, strings.Join(str, ", "))
			utils.BadResponse(w, http.StatusUnprocessableEntity, err.Error())
			return

		}

		err := h.sv.Save(&newBuyer)

		if err != nil {
			switch {
			case errors.Is(err, utils.ErrBuyerRepositoryCardDuplicated):
				utils.BadResponse(w, http.StatusConflict, err.Error())
			default:
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		utils.GoodResponse(w, http.StatusCreated, "Buyer creado exitosamente", newBuyer)
		return
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
