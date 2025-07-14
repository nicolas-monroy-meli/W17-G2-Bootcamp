package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	myerrors "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors" // 👈 esta es la corrección
)

type warehouseHandler struct {
	sv       internal.WarehouseService
	validate *validator.Validate
}

func NewWarehouseHandler(sv internal.WarehouseService) internal.WarehouseHandler {
	return &warehouseHandler{
		sv:       sv,
		validate: validator.New(),
	}
}

// GET /warehouses
func (h *warehouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehousesMap, err := h.sv.FindAll()
		if err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, "Error al obtener los almacenes")
			return
		}

		var warehouses []models.Warehouse
		for _, wh := range warehousesMap {
			warehouses = append(warehouses, wh)
		}

		sort.Slice(warehouses, func(i, j int) bool {
			return warehouses[i].ID < warehouses[j].ID
		})

		utils.GoodResponse(w, http.StatusOK, warehouses)
	}
}

// GET /warehouses/{id}
func (h *warehouseHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, myerrors.ErrRequestIdMustBeInt.Error())
			return
		}

		wh, err := h.sv.FindByID(id)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, myerrors.ErrWarehouseRepositoryNotFound.Error())
			return
		}

		utils.GoodResponse(w, http.StatusOK, wh)
	}
}

// POST /warehouses
func (h *warehouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouse models.Warehouse

		if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, myerrors.ErrRequestWrongBody.Error())
			return
		}

		if err := h.validate.Struct(warehouse); err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, "Campos inválidos: "+err.Error())
			return
		}

		if err := h.sv.Save(&warehouse); err != nil {
			utils.BadResponse(w, http.StatusConflict, myerrors.ErrWarehouseRepositoryDuplicated.Error())
			return
		}

		utils.GoodResponse(w, http.StatusCreated, warehouse)
	}
}

// PUT /warehouses/{id}
func (h *warehouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, myerrors.ErrRequestIdMustBeInt.Error())
			return
		}

		var warehouse models.Warehouse
		if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, myerrors.ErrRequestWrongBody.Error())
			return
		}

		if err := h.validate.Struct(warehouse); err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, "Campos inválidos: "+err.Error())
			return
		}

		if err := common.ValidateWarehouseUpdate(warehouse); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		warehouse.ID = id
		if err := h.sv.Update(&warehouse); err != nil {
			utils.BadResponse(w, http.StatusNotFound, myerrors.ErrWarehouseRepositoryNotFound.Error())
			return
		}

		utils.GoodResponse(w, http.StatusOK, warehouse)
	}
}

// DELETE /warehouses/{id}
func (h *warehouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, myerrors.ErrRequestIdMustBeInt.Error())
			return
		}

		if err := h.sv.Delete(id); err != nil {
			utils.BadResponse(w, http.StatusNotFound, myerrors.ErrWarehouseRepositoryNotFound.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// POST /carries
func (h *warehouseHandler) CreateCarry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var carry models.Carry

		if err := json.NewDecoder(r.Body).Decode(&carry); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, myerrors.ErrRequestWrongBody.Error())
			return
		}

		if err := h.validate.Struct(carry); err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, "Campos inválidos: "+err.Error())
			return
		}

		if h.sv.CIDExists(carry.CID) {
			utils.BadResponse(w, http.StatusConflict, "El CID ya existe")
			return
		}

		if !h.sv.LocalityExists(carry.LocalityID) {
			utils.BadResponse(w, http.StatusConflict, "La localidad no existe")
			return
		}

		if err := h.sv.SaveCarry(&carry); err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, "No se pudo crear el carry")
			return
		}

		utils.GoodResponse(w, http.StatusCreated, carry)
	}
}

// GET /localities/reportCarries?id={id}
func (h *warehouseHandler) GetReportByLocalityID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			utils.BadResponse(w, http.StatusBadRequest, "Debe especificar un ID")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, myerrors.ErrRequestIdMustBeInt.Error())
			return
		}

		if !h.sv.LocalityExists(id) {
			utils.BadResponse(w, http.StatusNotFound, "Localidad no encontrada")
			return
		}

		report, err := h.sv.GetCarriesReportByLocality(id)
		if err != nil {
			utils.BadResponse(w, http.StatusInternalServerError, "Error al generar el reporte")
			return
		}

		utils.GoodResponse(w, http.StatusOK, report)
	}
}
