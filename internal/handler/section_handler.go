package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"net/http"
	"strconv"
)

var model models.Section

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

// ValidateStruct returns a string map of formatted errors
func ValidateStruct(s interface{}) map[string]string {
	v := validator.New()
	errorsList := make(map[string]string)

	err := v.Struct(s)
	if err == nil {
		return nil
	}
	for _, err := range err.(validator.ValidationErrors) {
		customMsg := "unexpected error"
		field := err.Field()
		switch err.Tag() {
		case "required":
			customMsg = fmt.Sprintf("%s is required", field)
		case "gte":
			customMsg = fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
		case "gt":
			customMsg = fmt.Sprintf("%s must be greater than %s", field, err.Param())
		case "gtefield":
			customMsg = fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
		case "gtfield":
			customMsg = fmt.Sprintf("%s must be greater than %s", field, err.Param())
		case "lte":
			customMsg = fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
		case "lt":
			customMsg = fmt.Sprintf("%s must be less than %s", field, err.Param())
		case "ltefield":
			customMsg = fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
		case "ltfield":
			customMsg = fmt.Sprintf("%s must be less than %s", field, err.Param())
		default:
			customMsg = fmt.Sprintf("%s failed on %s validation", field, err.Tag())
		}
		errorsList[field] = customMsg
	}
	return errorsList
}

// Create creates a new section
func (h *SectionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		validationErrors := ValidateStruct(model)
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
