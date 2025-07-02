package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

type WarehouseHandler struct {
	sv       internal.WarehouseService
	validate *validator.Validate
}

func NewWarehouseHandler(sv internal.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		sv:       sv,
		validate: validator.New(),
	}
}

// GET /warehouses
func (h *WarehouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehousesMap, err := h.sv.FindAll()
		if err != nil {
			http.Error(w, "Error al obtener los almacenes", http.StatusInternalServerError)
			return
		}

		var warehouses []mod.Warehouse
		for _, wh := range warehousesMap {
			warehouses = append(warehouses, wh)
		}

		// ✅ Ordenar por ID ascendente
		sort.Slice(warehouses, func(i, j int) bool {
			return warehouses[i].ID < warehouses[j].ID
		})

		json.NewEncoder(w).Encode(warehouses)
	}
}

// GET /warehouses/{id}
func (h *WarehouseHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		wh, err := h.sv.FindByID(id)
		if err != nil {
			http.Error(w, "Almacén no encontrado", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(wh)
	}
}

// POST /warehouses
func (h *WarehouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Address            string  `json:"address" validate:"required"`
			Telephone          string  `json:"telephone" validate:"required"`
			WarehouseCode      string  `json:"warehouse_code" validate:"required"`
			MinimumCapacity    int     `json:"minimun_capacity" validate:"required,min=1"`
			MinimumTemperature float64 `json:"minimun_temperature" validate:"required"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		if err := h.validate.Struct(input); err != nil {
			http.Error(w, "Campos inválidos: "+err.Error(), http.StatusUnprocessableEntity)
			return
		}

		warehouse := &mod.Warehouse{
			Address:            input.Address,
			Telephone:          input.Telephone,
			WarehouseCode:      input.WarehouseCode,
			MinimumCapacity:    input.MinimumCapacity,
			MinimumTemperature: input.MinimumTemperature,
		}

		err := h.sv.Save(warehouse)
		if err != nil {
			http.Error(w, "Error al crear almacén: "+err.Error(), http.StatusConflict)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(warehouse)
	}
}

// PUT /warehouses/{id}
func (h *WarehouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		var input struct {
			Address            string  `json:"address" validate:"required"`
			Telephone          string  `json:"telephone" validate:"required"`
			WarehouseCode      string  `json:"warehouse_code" validate:"required"`
			MinimumCapacity    int     `json:"minimun_capacity" validate:"required,min=1"`
			MinimumTemperature float64 `json:"minimun_temperature" validate:"required"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		if err := h.validate.Struct(input); err != nil {
			http.Error(w, "Campos inválidos: "+err.Error(), http.StatusUnprocessableEntity)
			return
		}

		warehouse := &mod.Warehouse{
			ID:                 id,
			Address:            input.Address,
			Telephone:          input.Telephone,
			WarehouseCode:      input.WarehouseCode,
			MinimumCapacity:    input.MinimumCapacity,
			MinimumTemperature: input.MinimumTemperature,
		}

		err = h.sv.Update(warehouse)
		if err != nil {
			http.Error(w, "Error al actualizar: "+err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(warehouse)
	}
}

// DELETE /warehouses/{id}
func (h *WarehouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			http.Error(w, "Error al eliminar: "+err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
