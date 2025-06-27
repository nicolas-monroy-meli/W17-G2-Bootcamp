package handler

import (
	"net/http"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
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
func (h *EmployeeHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a employee
func (h *EmployeeHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new employee
func (h *EmployeeHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a employee
func (h *EmployeeHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a employee
func (h *EmployeeHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
