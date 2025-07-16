package handler

import (
	"encoding/json"
	"errors"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"net/http"
	"strconv"
)

// NewInboundHandler creates a new instance of the inbound handler
func NewInboundHandler(sv internal.InboundService) *InboundHandler {
	return &InboundHandler{
		sv: sv,
	}
}

// EmployeeHandler is the default implementation of the employee handler
type InboundHandler struct {
	// sv is the service used by the handler
	sv internal.InboundService
}

func (h *InboundHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inboundOrder mod.InboundOrders
		if err := json.NewDecoder(r.Body).Decode(&inboundOrder); err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, "Invalid JSON format")
			return
		}

		createdOrder, err := h.sv.Save(&inboundOrder)
		if err != nil {
			// Manejo de errores específicos
			if errors.Is(err, e.ErrInboundOrderInvalidData) {
				utils.BadResponse(w, http.StatusUnprocessableEntity, err.Error())
				return
			}
			if err.Error() == "order number already exists" || err.Error() == "employee not found" {
				utils.BadResponse(w, http.StatusConflict, err.Error())
				return
			}

			// Error interno
			utils.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusCreated, "succes", createdOrder)

	}
}

func (h *InboundHandler) GetOrdersByEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeIDStr := r.URL.Query().Get("id")

		var employeeID int
		if employeeIDStr != "" {
			id, err := strconv.Atoi(employeeIDStr)
			if err != nil {
				utils.BadResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			employeeID = id
		}

		report, err := h.sv.FindOrdersByEmployee(employeeID)
		if err != nil {
			if errors.Is(err, e.ErrEmployeeNotFound) {
				utils.BadResponse(w, http.StatusNotFound, err.Error())
				return
			}
			utils.BadResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusCreated, "succes", report)
	}
}
