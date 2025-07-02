package handler

import (
	"encoding/json"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	"net/http"
)

// NewSectionHandler creates a new instance of the section handler
func NewSectionHandler(sv internal.SectionService) *SectionHandler {
	return &SectionHandler{
		sv: sv,
	}
}

// SectionHandler is the default implementation of the section handler
type SectionHandler struct {
	// sv is the service used by the handler
	sv internal.SectionService
}

// GetAll returns all sections
func (h *SectionHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.FindAll()
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, "data retrieved successfully", result)
	}
}

// GetByID returns a section
func (h *SectionHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := common.IdRequests(r)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		result, err := h.sv.FindByID(id)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, utils.DataRetrievedSuccess, result)
	}
}

// Create creates a new section
func (h *SectionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.Section

		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		validationErrors := utils.ValidateStruct(model)
		if len(validationErrors) > 0 {
			str := ""
			for _, err := range validationErrors {
				str += err + ", "
			}
			utils.BadResponse(w, http.StatusUnprocessableEntity, str)
			return
		}
		err = h.sv.Save(&model)
		if err != nil {
			utils.BadResponse(w, http.StatusConflict, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusCreated, utils.SectionCreated, model)
	}
}

// Update updates a section
func (h *SectionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.SectionPatch
		id, err := common.IdRequests(r)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		section, err := h.sv.FindByID(id)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}

		section = common.PatchSection(model, section)

		validationErrors := utils.ValidateStruct(section)
		if len(validationErrors) > 0 {
			str := ""
			for _, err := range validationErrors {
				str += err + ", "
			}
			utils.BadResponse(w, http.StatusUnprocessableEntity, str)
			return
		}

		err = h.sv.Update(&section)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, utils.SectionUpdated, section)
	}
}

// Delete deletes a section
func (h *SectionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := common.IdRequests(r)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		err = h.sv.Delete(id)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusNoContent, utils.SectionDeleted, nil)
	}
}
