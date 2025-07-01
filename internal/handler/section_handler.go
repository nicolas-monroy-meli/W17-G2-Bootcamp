package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"net/http"
	"strconv"
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

// idRequests encapsulates the process of getting the id parameter and returns an int number and an error if necessary
func idRequests(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return 0, errors.New(utils.EmptyParams)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return id, nil
}

// GetByID returns a section
func (h *SectionHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := idRequests(r)
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

func patchSection(model models.Section, section models.Section) models.Section {
	switch {
	case model.SectionNumber != 0 && model.SectionNumber != section.SectionNumber:
		section.SectionNumber = model.SectionNumber
	case model.CurrentTemperature != 0 && model.CurrentTemperature != section.CurrentTemperature:
		section.CurrentTemperature = model.CurrentTemperature
	case model.MinimumTemperature != 0 && model.MinimumTemperature != section.MinimumTemperature:
		section.MinimumTemperature = model.MinimumTemperature
	case model.CurrentCapacity != 0 && model.CurrentCapacity != section.CurrentCapacity:
		section.CurrentCapacity = model.CurrentCapacity
	case model.MinimumCapacity != 0 && model.MinimumCapacity != section.MinimumCapacity:
		section.MinimumCapacity = model.MinimumCapacity
	case model.MaximumCapacity != 0 && model.MaximumCapacity != section.MaximumCapacity:
		section.MaximumCapacity = model.MaximumCapacity
	case model.WarehouseID != 0 && model.WarehouseID != section.WarehouseID:
		section.WarehouseID = model.WarehouseID
	case model.ProductTypeID != 0 && model.ProductTypeID != section.ProductTypeID:
		section.ProductTypeID = model.ProductTypeID
	}
	return section
}

// Update updates a section
func (h *SectionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model models.Section
		id, err := idRequests(r)
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
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		section = patchSection(model, section)
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
		id, err := idRequests(r)
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
