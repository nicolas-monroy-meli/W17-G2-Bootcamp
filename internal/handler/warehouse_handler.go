package handler

import (
	"net/http"

	internal "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/internal/interfaces"
)

// NewWarehouseHandler creates a new instance of the warehouse handler
func NewWarehouseHandler(sv internal.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		sv: sv,
	}
}

// WarehouseHandler is the default implementation of the warehouse handler
type WarehouseHandler struct {
	// sv is the service used by the handler
	sv internal.WarehouseService
}

// GetAll returns all warehouses
func (h *WarehouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// GetByID returns a warehouse
func (h *WarehouseHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Create creates a new warehouse
func (h *WarehouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Update updates a warehouse
func (h *WarehouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Delete deletes a warehouse
func (h *WarehouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
