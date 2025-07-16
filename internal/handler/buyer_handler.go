package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
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
		fmt.Print("Hola")
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
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		buyer, err := h.sv.FindByID(id)

		if err != nil {
			switch {
			case errors.Is(err, e.ErrBuyerRepositoryNotFound):
				utils.BadResponse(w, http.StatusNotFound, err.Error())
			default:
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		utils.GoodResponse(w, http.StatusOK, "Buyer obtenido con exito", buyer)
		return
	}
}

func (h *BuyerHandler) GetReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var id *int
		idQuery := r.URL.Query().Get("id")

		if idQuery != "" {
			idParsed, err := strconv.Atoi(idQuery)
			if err != nil {
				utils.BadResponse(w, http.StatusNotFound, e.ErrRequestIdMustBeInt.Error())
			}
			id = &idParsed
		}

		reports, err := h.sv.GetPurchaseOrderReport(id)
		if err != nil {
			switch {
			case errors.Is(err, e.ErrBuyerRepositoryNotFound):
				utils.BadResponse(w, http.StatusNotFound, err.Error())
			default:
				utils.BadResponse(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize product to JSON

		utils.GoodResponse(
			w,
			http.StatusOK,
			"Reporte generado exitosamente",
			reports,
		)
	}
}

// Create creates a new buyer
func (h *BuyerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBuyer mod.Buyer

		if err := json.NewDecoder(r.Body).Decode(&newBuyer); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestFailedBody.Error())
			return
		}

		errValidation := e.ValidateStruct(newBuyer)

		//validate := validator.New()
		//err := validate.Struct(newBuyer)

		if errValidation != nil {
			//err = err.(validator.ValidationErrors)
			str := make([]string, 0, len(errValidation))
			for _, err := range errValidation {
				str = append(str, err)
			}

			err := fmt.Errorf("%w: %v", e.ErrRequestWrongBody, strings.Join(str, ", "))
			utils.BadResponse(w, http.StatusUnprocessableEntity, err.Error())
			return

		}

		err := h.sv.Save(&newBuyer)

		if err != nil {
			switch {
			case errors.Is(err, e.ErrBuyerRepositoryCardDuplicated):
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
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		buyer, err := h.sv.FindByID(id)

		if err != nil {
			switch {
			case errors.Is(err, e.ErrBuyerRepositoryNotFound):
				utils.BadResponse(w, http.StatusNotFound, err.Error())
			default:
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		var buyerPatch mod.BuyerPatch

		if err := json.NewDecoder(r.Body).Decode(&buyerPatch); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestFailedBody.Error())
			return
		}

		buyerMapped := common.ValidatePatchRequest(buyer, buyerPatch)

		errValidation := e.ValidateStruct(buyerMapped)

		if errValidation != nil {
			str := make([]string, 0, len(errValidation))
			for _, err := range errValidation {
				str = append(str, err)
			}

			err := fmt.Errorf("%w: %v", e.ErrRequestWrongBody, strings.Join(str, ", "))
			utils.BadResponse(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		err = h.sv.Update(&buyer)

		if err != nil {
			switch {
			case errors.Is(err, e.ErrBuyerRepositoryCardDuplicated):
				utils.BadResponse(w, http.StatusConflict, err.Error())
			default:
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		utils.GoodResponse(w, http.StatusCreated, "Buyer actualizado exitosamente", buyer)
		return
	}
}

// Delete deletes a buyer
func (h *BuyerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		if err = h.sv.Delete(id); err != nil {

			switch {
			case errors.Is(err, e.ErrBuyerRepositoryNotFound):
				utils.BadResponse(w, http.StatusNotFound, err.Error())
			default:
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
			}
			return

		}

		utils.GoodResponse(w, http.StatusNoContent, "", nil)
		return
	}
}
