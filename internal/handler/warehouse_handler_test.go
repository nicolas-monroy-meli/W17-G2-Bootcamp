package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/internal/service"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/tests"
	"github.com/stretchr/testify/require"
)

func TestWarehouseController_Create(t *testing.T) {
	mock := tests.NewWarehouseMock()
	serv := service.NewWarehouseService(mock)
	handler := NewWarehouseHandler(serv)

	warehouseOk := models.Warehouse{ID: 0,
		WarehouseCode:      "holiis como prueba",
		Address:            "a",
		Telephone:          "1234",
		MinimumCapacity:    1,
		MinimumTemperature: 1}

	warehouseEmpty := models.Warehouse{
		ID:                 0,
		WarehouseCode:      "",
		Address:            "",
		Telephone:          "",
		MinimumCapacity:    0,
		MinimumTemperature: 0,
	}
	mock.On("Save", &warehouseOk).Return(nil)
	mock.On("Save", &warehouseEmpty).Return(nil)
	mock.On("ExistsWarehouseCode", "holiis como prueba").Return(false, nil)
	mock.On("ExistsWarehouseCode", "").Return(false, nil)
	mock.On("ExistsWarehouseCode", "holiis").Return(true, nil)
	t.Run("create_ok", func(t *testing.T) {
		body := strings.NewReader(`{
  "address": "a",
  "telephone": "1234",
  "warehouse_code": "holiis como prueba",
  "minimum_capacity": 1,
  "minimum_temperature": 1

		}`)
		req := httptest.NewRequest(http.MethodPost, "/warehouses", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{
			
    "success": true,
    "message": "success",
    "data": 
        {
            "ID": 0,
            "Warehouse_Code": "holiis como prueba",
            "Address": "a",
            "Telephone": "1234",
            "Minimum_Capacity": 1,
            "Minimum_Temperature": 1
    }
		}`

		handler.Create().ServeHTTP(w, req)
		require.Equal(t, http.StatusCreated, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("create_fail_missing_fields", func(t *testing.T) {
		body := strings.NewReader(`{"address": "W1"}`) // Falta address
		req := httptest.NewRequest(http.MethodPost, "/warehouses", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{"data":null, "message":"Campos inválidos: Key: 'Warehouse.WarehouseCode' Error:Field validation for 'WarehouseCode' failed on the 'required' tag\nKey: 'Warehouse.Telephone' Error:Field validation for 'Telephone' failed on the 'required' tag\nKey: 'Warehouse.MinimumCapacity' Error:Field validation for 'MinimumCapacity' failed on the 'required' tag\nKey: 'Warehouse.MinimumTemperature' Error:Field validation for 'MinimumTemperature' failed on the 'required' tag", "success":false}`

		handler.Create().ServeHTTP(w, req)
		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("create_conflict", func(t *testing.T) {
		body := strings.NewReader(`{
			
  "Address": "dfghhgfds",
  "Telephone": "1277779",
  "Warehouse_Code": "holiis",
  "Minimum_Capacity": 300,
  "Minimum_Temperature": 8

		}`)
		req := httptest.NewRequest(http.MethodPost, "/warehouses", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{
	"success": false,
    "message": "repository: warehouse already exists",
    "data": null
		}`

		handler.Create().ServeHTTP(w, req)
		require.Equal(t, http.StatusConflict, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("create_conflictNill", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/warehouses", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{"data":null, "message":"handler: body does not meet requirements", "success":false}`

		handler.Create().ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

}

func TestWarehouseController_Read(t *testing.T) {
	mock := tests.NewWarehouseMock()
	serv := service.NewWarehouseService(mock)
	handler := NewWarehouseHandler(serv)

	warehouse := models.Warehouse{ID: 1,
		WarehouseCode:      "a",
		Address:            "a",
		Telephone:          "1234",
		MinimumCapacity:    1,
		MinimumTemperature: 1}
	mock.On("GetByID", 1).Return(warehouse, nil)
	mock.On("GetByID", 3).Return(models.Warehouse{}, e.ErrWarehouseRepositoryNotFound)
	t.Run("find_by_id_ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/warehouses/1", nil)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		w := httptest.NewRecorder()

		expected := `{
			
    "success": true,
    "message": "success",
    "data": {
        "ID": 1,
        "Warehouse_Code": "a",
        "Address": "a",
        "Telephone": "1234",
        "Minimum_Capacity": 1,
        "Minimum_Temperature": 1
    
}
		}`

		handler.GetByID().ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("find_by_id_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/warehouses/3", nil)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "3")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		w := httptest.NewRecorder()

		expected := `{
			"data":null, "message":"repository: warehouse not found", "success":false
		}`

		handler.GetByID().ServeHTTP(w, req)
		require.Equal(t, http.StatusNotFound, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("find_by_id_bad_request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/warehouses/invalid", nil)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "invalid")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		w := httptest.NewRecorder()

		expected := `{
			"data":null, "message":"handler: id must be an integer", "success":false
		}`

		handler.GetByID().ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
}

func TestWarehouseController_ReadAll(t *testing.T) {
	t.Run("find_all_ok", func(t *testing.T) {
		mock := tests.NewWarehouseMock()
		serv := service.NewWarehouseService(mock)
		handler := NewWarehouseHandler(serv)

		warehouses := []models.Warehouse{
			{ID: 0,
				WarehouseCode:      "a",
				Address:            "a",
				Telephone:          "1234",
				MinimumCapacity:    1,
				MinimumTemperature: 1},
			{ID: 1,
				WarehouseCode:      "a",
				Address:            "a",
				Telephone:          "1234",
				MinimumCapacity:    1,
				MinimumTemperature: 1},
		}
		mock.On("GetAll").Return(warehouses, nil)
		req := httptest.NewRequest(http.MethodGet, "/warehouses", nil)
		w := httptest.NewRecorder()

		expected := `{
			
    "success": true,
    "message": "",
    "data": [
        {
            "ID": 0,
            "Warehouse_Code": "a",
            "Address": "a",
            "Telephone": "1234",
            "Minimum_Capacity": 1,
            "Minimum_Temperature": 1
        },
		 {
            "ID": 1,
            "Warehouse_Code": "a",
            "Address": "a",
            "Telephone": "1234",
            "Minimum_Capacity": 1,
            "Minimum_Temperature": 1
	} ]
		}`

		handler.GetAll().ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("find_all_empty", func(t *testing.T) {
		mock := tests.NewWarehouseMock()
		serv := service.NewWarehouseService(mock)
		handler := NewWarehouseHandler(serv)

		warehouses := []models.Warehouse{}
		mock.On("GetAll").Return(warehouses, e.ErrWarehouseRepositoryNotFound)
		req := httptest.NewRequest(http.MethodGet, "/warehouses", nil)
		w := httptest.NewRecorder()

		expected := `{
			"success":false,
			"message": "Error al obtener los almacenes",
			"data":null
		}`

		handler.GetAll().ServeHTTP(w, req)
		require.Equal(t, http.StatusNotFound, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
}

func TestWarehouseController_Update(t *testing.T) {
	mock := tests.NewWarehouseMock()
	serv := service.NewWarehouseService(mock)
	handler := NewWarehouseHandler(serv)

	updatedWarehouse := models.Warehouse{ID: 1,
		WarehouseCode:      "a",
		Address:            "a",
		Telephone:          "1234",
		MinimumCapacity:    1,
		MinimumTemperature: 1}
	mock.On("GetByID", 1).Return(updatedWarehouse, nil)
	mock.On("GetByID", 2).Return(models.Warehouse{}, e.ErrWarehouseRepositoryNotFound)

	mock.On("Update", &updatedWarehouse).Return(nil)
	mock.On("GetByWarehouseCode", "a").Return(updatedWarehouse, nil)

	t.Run("update_ok", func(t *testing.T) {
		body := strings.NewReader(`{
			
  "Address": "a",
  "Telephone": "1234",
  "Warehouse_Code": "a",
  "Minimum_Capacity": 1,
  "Minimum_Temperature": 1

		}`)
		req := httptest.NewRequest(http.MethodPatch, "/warehouses/1", body)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{
			"data":{"Address":"a", "ID":1, "Minimum_Capacity":1, "Minimum_Temperature":1, "Telephone":"1234", "Warehouse_Code":"a"}, "message":"success", "success":true
			
		}`

		handler.Update().ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("update_not_found", func(t *testing.T) {
		body := strings.NewReader(`{"Address": "a",
  "Telephone": "1234",
  "Warehouse_Code": "a",
  "Minimum_Capacity": 1,
  "Minimum_Temperature": 1}`)
		req := httptest.NewRequest(http.MethodPatch, "/warehouses/2", body)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "2")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{
		"data":null, "message":"repository: warehouse not found", "success":false
		}`

		handler.Update().ServeHTTP(w, req)
		require.Equal(t, http.StatusNotFound, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("update_invalidedID", func(t *testing.T) {
		body := strings.NewReader(`{"Address": "a",
  "Telephone": "1234",
  "Warehouse_Code": "a",
  "Minimum_Capacity": 1,
  "Minimum_Temperature": 1}`)
		req := httptest.NewRequest(http.MethodPatch, "/warehouses/a", body)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "a")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{
		"data":null, "message":"handler: id must be an integer", "success":false
		}`

		handler.Update().ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("update_invalidedIDNill", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPatch, "/warehouses/3", nil)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "3")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{"data":null, "message":"handler: body does not meet requirements", "success":false}`

		handler.Update().ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("update_invalidedIDNot", func(t *testing.T) {
		body := strings.NewReader(`{"Address": "a",
  "Telephone": "1234",
  "Warehouse_Code": "a",
  "Minimum_Capacity": -1,
  "Minimum_Temperature": 1}`)
		req := httptest.NewRequest(http.MethodPatch, "/warehouses/4", body)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "4")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{"data":null, "message":"Campos inválidos: Key: 'Warehouse.MinimumCapacity' Error:Field validation for 'MinimumCapacity' failed on the 'min' tag", "success":false}`

		handler.Update().ServeHTTP(w, req)
		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

}

func TestWarehouseController_Delete(t *testing.T) {
	mock := tests.NewWarehouseMock()
	serv := service.NewWarehouseService(mock)
	handler := NewWarehouseHandler(serv)

	mock.On("Delete", 1).Return(nil)
	mock.On("Delete", 2).Return(e.ErrWarehouseRepositoryNotFound)

	t.Run("delete_ok", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/warehouses/1", nil)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		w := httptest.NewRecorder()

		handler.Delete().ServeHTTP(w, req)
		require.Equal(t, http.StatusNoContent, w.Code)
		require.Empty(t, w.Body.String())
	})

	t.Run("delete_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/warehouses/2", nil)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "2")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		w := httptest.NewRecorder()

		expected := `{
			"data":null, "message":"repository: warehouse not found", "success":false
		}`

		handler.Delete().ServeHTTP(w, req)
		require.Equal(t, http.StatusNotFound, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})

	t.Run("delete_invalidedID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/warehouses/a", nil)
		routeCtx := chi.NewRouteContext()
		routeCtx.URLParams.Add("id", "a")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))
		w := httptest.NewRecorder()

		expected := `{
			"data":null, "message":"handler: id must be an integer", "success":false
		}`

		handler.Delete().ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
}
