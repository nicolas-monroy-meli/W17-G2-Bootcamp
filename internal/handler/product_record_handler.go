package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewProductHandler creates a new instance of the product handler
func NewProductRecordHandler(sv internal.ProductRecordService) *ProductRecordHandler {
	return &ProductRecordHandler{
		sv: sv,
	}
}

// ProductHandler is the default implementation of the product handler
type ProductRecordHandler struct {
	// sv is the service used by the handler
	sv internal.ProductRecordService
}

// GetRecords returns all product records, if productId is provided, it returns the records for that product, if not, it returns all records
func (h *ProductRecordHandler) GetRecords() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id != "" {
			// If an id is provided, we check if it is a valid integer
			idInt, err := strconv.Atoi(id)
			if err != nil {
				utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
				return
			}
			// If an id is provided, we return the records for that product
			// Check if the product exists
			_, err = h.sv.FindProductByID(idInt)
			if errors.Is(err, e.ErrProductRepositoryNotFound) {
				utils.BadResponse(w, http.StatusNotFound, err.Error())
				return
			}
			// Search for records by product ID
			producRecords, err := h.sv.FindAllByProductIDPR(idInt)
			if err != nil {
				utils.BadResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			if len(producRecords) == 0 {
				utils.BadResponse(w, http.StatusNotFound, e.ErrProductRecordRepositoryNotFound.Error())
				return
			}
			utils.GoodResponse(w, http.StatusOK, "success", producRecords)
			return
		}
		// If no id is provided, we return all records
		result, err := h.sv.FindAllPR()
		if errors.Is(err, e.ErrProductRecordRepositoryNotFound) {
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

// CreateRecord creates a new product record
func (h *ProductRecordHandler) CreateRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.ProductRecord

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

		err = h.sv.SavePR(&req)
		if errors.Is(err, e.ErrProductRepositoryNotFound) {
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
