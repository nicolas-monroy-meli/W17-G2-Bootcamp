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
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
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

func (h *warehouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehousesMap, err := h.sv.FindAll()
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, "Error al obtener los almacenes")
			return
		}

		var warehouses []models.Warehouse
		warehouses = append(warehouses, warehousesMap...)

		sort.Slice(warehouses, func(i, j int) bool {
			return warehouses[i].ID < warehouses[j].ID
		})

		utils.GoodResponse(w, http.StatusOK, "", warehouses)
	}
}

func (h *warehouseHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		wh, err := h.sv.FindByID(id)
		if err != nil {
			utils.BadResponse(w, http.StatusNotFound, e.ErrWarehouseRepositoryNotFound.Error())
			return
		}

		utils.GoodResponse(w, http.StatusOK, "success", wh)
	}
}

func (h *warehouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouse models.Warehouse

		if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestWrongBody.Error())
			return
		}

		if err := h.validate.Struct(warehouse); err != nil {
			utils.BadResponse(w, http.StatusUnprocessableEntity, "Campos inválidos: "+err.Error())
			return
		}

		if err := h.sv.Save(&warehouse); err != nil {
			utils.BadResponse(w, http.StatusConflict, e.ErrWarehouseRepositoryDuplicated.Error())
			return
		}

		utils.GoodResponse(w, http.StatusCreated, "success", warehouse)
	}
}

func (h *warehouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		var warehouse models.Warehouse
		if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestWrongBody.Error())
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
			utils.BadResponse(w, http.StatusNotFound, e.ErrWarehouseRepositoryNotFound.Error())
			return
		}

		utils.GoodResponse(w, http.StatusOK, "success", warehouse)
	}
}

func (h *warehouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			utils.BadResponse(w, http.StatusBadRequest, e.ErrRequestIdMustBeInt.Error())
			return
		}

		if err := h.sv.Delete(id); err != nil {
			utils.BadResponse(w, http.StatusNotFound, e.ErrWarehouseRepositoryNotFound.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
