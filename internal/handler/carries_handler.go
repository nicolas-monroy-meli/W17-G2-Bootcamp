package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
)

type carryHandler struct {
	sv       internal.CarryService
	validate *validator.Validate
}

func NewCarryHandler(sv internal.CarryService) *carryHandler {
	validate := validator.New() // Inicializar primero el validador

	return &carryHandler{
		sv:       sv,
		validate: validate,
	}
}

// POST /carries
func (h *carryHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var carry models.Carry

		if err := json.NewDecoder(r.Body).Decode(&carry); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestWrongBody.Error())
			return
		}

		if err := h.validate.Struct(carry); err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, "Campos inválidos: "+err.Error())
			return
		}

		if err := h.sv.Create(&carry); err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, "No se pudo crear el carry: "+err.Error())
			return
		}

		utils.GoodResponse(w, http.StatusCreated, "", carry)
	}
}

// GET /localities/reportCarries?id={id}
func (h *carryHandler) GetReportByLocality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			utils.BadResponse(w, http.StatusBadRequest, "Debe especificar un ID")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestIdMustBeInt.Error())
			return
		}

		report, err := h.sv.ReportByLocality(id)
		if err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, "Error al generar el reporte: "+err.Error())
			return
		}

		utils.GoodResponse(w, http.StatusOK, "", report)
	}
}
