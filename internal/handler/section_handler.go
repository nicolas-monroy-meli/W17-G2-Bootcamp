package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"net/http"
	"strconv"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
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

	}
}

// Update updates a section
func (h *SectionHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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
