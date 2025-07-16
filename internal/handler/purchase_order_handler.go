package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"net/http"
	"strings"
)

func NewPurchaseOrderHandler(sv internal.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{
		sv: sv,
	}
}

type PurchaseOrderHandler struct {
	// sv is the service used by the handler
	sv internal.PurchaseOrderService
}

func (h *PurchaseOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPurchaseOrder mod.PurchaseOrder

		if err := json.NewDecoder(r.Body).Decode(&newPurchaseOrder); err != nil {
			fmt.Println(err)
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestFailedBody.Error())
			return
		}

		errValidation := e.ValidateStruct(newPurchaseOrder)
		str := make([]string, 0, len(errValidation))
		if errValidation != nil {
			for _, err := range errValidation {
				str = append(str, err)
			}

			err := fmt.Errorf("%w: %v", e.ErrRequestWrongBody, strings.Join(str, ", "))
			utils.BadResponse(w, http.StatusUnprocessableEntity, err.Error())
			return

		}

		for idx, pd := range newPurchaseOrder.ProductsDetails {
			errValidationDetail := e.ValidateStruct(pd)
			if errValidationDetail != nil {

				for _, err := range errValidationDetail {
					str = append(str, err)
				}

				err := fmt.Errorf("%w: ProductDetails[%d]: %v", e.ErrRequestWrongBody, idx, strings.Join(str, ", "))
				utils.BadResponse(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
		}

		err := h.sv.Save(&newPurchaseOrder)

		if err != nil {
			switch {
			case errors.Is(err, e.ErrPORepositoryOrderNumberDuplicated):
				utils.BadResponse(w, http.StatusConflict, err.Error())
			case errors.Is(err, e.ErrBuyerRepositoryNotFound):
				utils.BadResponse(w, http.StatusConflict, err.Error())
			default:
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		utils.GoodResponse(w, http.StatusCreated, "Buyer creado exitosamente", newPurchaseOrder)
		return
	}
}
