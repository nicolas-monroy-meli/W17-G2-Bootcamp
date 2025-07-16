package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewLocalityHandler creates a new instance of the seller handler
func NewLocalityHandler(sv internal.LocalityService) *LocalityHandler {
	return &LocalityHandler{
		sv: sv,
	}
}

// LocalityHandler is the default implementation of the seller handler
type LocalityHandler struct {
	// sv is the service used by the handler
	sv internal.LocalityService
}

// GetByID returns a seller
func (h *LocalityHandler) GetSelByLoc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.FindAllLocalities()
		if err != nil {
			utils.BadResponse(w, 404, err.Error())
			return
		}
		utils.GoodResponse(w, 200, "success", result)
	}
}

func (h *LocalityHandler) GetSelByLocID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		var err error
		req := r.URL.Query().Get("id")
		switch req {
		case "":
			id = -1
		default:
			id, err = strconv.Atoi(req)
			if id < 0 {
				utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
				return
			}
		}
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeGte0.Error())
			return
		}
		result, err := h.sv.FindSellersByLocID(id)
		if err != nil {
			utils.BadResponse(w, 404, err.Error())
			return
		}
		utils.GoodResponse(w, 200, "success", result)
	}
}

// Create creates a new locality
func (h *LocalityHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.Locality
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
