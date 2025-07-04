package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	"net/http"
	"strconv"
)

// NewEmployeeHandler creates a new instance of the employee handler
func NewEmployeeHandler(sv internal.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		sv: sv,
	}
}

// EmployeeHandler is the default implementation of the employee handler
type EmployeeHandler struct {
	// sv is the service used by the handler
	sv internal.EmployeeService
}

// GetAll returns all employees
func (h *EmployeeHandler) GetAllEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.FindAllEmployees()
		if err != nil {
			utils.BadResponse(w, 400, err.Error())
			return
		}
		utils.GoodResponse(w, 200, "succes", result)
	}
}

// GetByID returns a employee
func (h *EmployeeHandler) GetEmployeeById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idUrl := chi.URLParam(r, "id")
		if idUrl == "" {
			utils.BadResponse(w, http.StatusBadRequest, "id required")
			return
		}
		idNum, err := strconv.Atoi(idUrl)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestIdMustBeInt.Error())
			return
		}
		result, err := h.sv.FindEmployeeByID(idNum)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, "succes", result)
	}
}

// Create creates a new employee~
func (h *EmployeeHandler) CreateEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee mod.Employee
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if employee.FirstName == "" || employee.LastName == "" || employee.CardNumberID == "" || employee.WarehouseID == 0 {
			utils.BadResponse(w, http.StatusUnprocessableEntity, utils.ErrRequestWrongBody.Error())
			return
		}
		_, err = strconv.Atoi(employee.CardNumberID)
		if err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, "required only number in card Number")
			return
		}
		err = h.sv.SaveEmployee(&employee)
		if err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, utils.ErrEmployeeRepositoryDuplicated.Error())
			return
		}
		utils.GoodResponse(w, http.StatusCreated, "succes", employee)
	}
}

// Update updates a employee
func (h *EmployeeHandler) EditEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var model mod.Employee
		idObj := chi.URLParam(r, "id")
		if idObj == "" {
			utils.BadResponse(w, http.StatusBadRequest, "id required")
			return
		}
		idNum, err := strconv.Atoi(idObj)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestIdMustBeInt.Error())
			return
		}
		err = json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, utils.ErrRequestWrongBody.Error())
			return
		}
		emplo, err := h.sv.FindEmployeeByID(idNum)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		employee := common.PatchEmployees(model, emplo)
		employee.ID = idNum
		err = h.sv.UpdateEmployee(idNum, &employee)
		if err != nil {
			utils.BadResponse(w, http.StatusConflict, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusOK, "succes", employee)
	}
}

// Delete deletes a employee
func (h *EmployeeHandler) DeleteEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idObj := chi.URLParam(r, "id")
		if idObj == "" {
			utils.BadResponse(w, http.StatusBadRequest, "id required")
			return
		}
		idNum, err := strconv.Atoi(idObj)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, utils.ErrRequestIdMustBeInt.Error())
			return
		}
		err = h.sv.DeleteEmployee(idNum)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, err.Error())
			return
		}
		utils.GoodResponse(w, http.StatusNoContent, "succes", nil)

	}
}
