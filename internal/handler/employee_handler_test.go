package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmployeeHandler_GetAll(t *testing.T) {
	mockService := new(MockEmployeeService)
	handler := NewEmployeeHandler(mockService)

	tests := []struct {
		name           string
		mockReturnEmp  []mod.Employee
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success - Retrieve All Employees",
			mockReturnEmp: []mod.Employee{
				{ID: 1, FirstName: "Juan", LastName: "Perez", CardNumberID: "123", WarehouseID: 1},
				{ID: 2, FirstName: "Maria", LastName: "Gomez", CardNumberID: "456", WarehouseID: 2},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"handler: data retrieved successfully","data":[{"id":1,"first_name":"Juan","last_name":"Perez","card_number_id":"123","warehouse_id":1},{"id":2,"first_name":"Maria","last_name":"Gomez","card_number_id":"456","warehouse_id":2}]}`,
		},
		{
			name:           "Success - No Employees Found (Empty List)",
			mockReturnEmp:  []mod.Employee{}, // Servicio devuelve una lista vacía
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"handler: data retrieved successfully","data":[]}`,
		},
		{
			name:           "Error - Service Returns Error",
			mockReturnEmp:  nil,
			mockReturnErr:  errors.New("database connection failed"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"database connection failed","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar el comportamiento esperado del mock para FindAll
			// Aquí siempre esperamos que FindAll sea llamado
			mockService.On("FindAll").Return(tt.mockReturnEmp, tt.mockReturnErr).Once()

			// Crear una petición HTTP simulada para GET /employees (sin ID en la URL)
			req := httptest.NewRequest(http.MethodGet, "/employees", nil)
			rr := httptest.NewRecorder()

			// Llamar al handler GetAll
			handler.GetAll().ServeHTTP(rr, req)

			// Realizar aserciones
			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status code mismatch")
			assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected response body mismatch")

			// Verificar que el mock fue llamado como se esperaba
			mockService.AssertExpectations(t)
		})
	}
}

func TestEmployeeHandler_GetByID(t *testing.T) {
	mockService := new(MockEmployeeService)
	handler := NewEmployeeHandler(mockService)
	tests := []struct {
		name           string
		employeeID     string
		mockReturnEmp  *mod.Employee
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success - Employee Found",
			employeeID:     "1",
			mockReturnEmp:  &mod.Employee{ID: 1, FirstName: "Juan", LastName: "Perez", CardNumberID: "123", WarehouseID: 1},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"handler: data retrieved successfully","data":{"id":1,"first_name":"Juan","last_name":"Perez","card_number_id":"123","warehouse_id":1}}`,
		},
		{
			name:           "Not Found - Employee Does Not Exist",
			employeeID:     "99",
			mockReturnEmp:  nil,
			mockReturnErr:  e.ErrEmployeeNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"success":false,"message":"employee not found","data":null}`,
		},
		{
			name:           "Bad Request - Invalid ID Format",
			employeeID:     "abc",
			mockReturnEmp:  nil,
			mockReturnErr:  nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"handler: id must be an integer","data":null}`,
		},
		{
			name:           "Bad Request - Missing ID",
			employeeID:     "", // Simula que el parámetro ID no está en la URL
			mockReturnEmp:  nil,
			mockReturnErr:  nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"id required","data":null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockReturnEmp != nil || tt.mockReturnErr != nil {
				mockService.On("FindByID", mock.AnythingOfType("int")).Return(tt.mockReturnEmp, tt.mockReturnErr).Once()
			}

			//Crear una petición HTTP simulada
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/employees/%s", tt.employeeID), nil)
			rr := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.employeeID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			// Llamar al handler
			handler.GetById().ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status code mismatch")

			// Verificar el cuerpo de la respuesta
			assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected response body mismatch")

			if tt.mockReturnEmp != nil || tt.mockReturnErr != nil {
				mockService.AssertExpectations(t)
			} else {
				mockService.AssertNotCalled(t, "FindByID")
			}
		})
	}
}

func TestEmployeeHandler_Create(t *testing.T) {

	tests := []struct {
		name           string
		requestBody    string
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success - Employee Created",
			requestBody:    `{"first_name":"Test","last_name":"User","card_number_id":"12345","warehouse_id":10}`,
			mockReturnErr:  nil, // Service will save successfully
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"success":true,"message":"handler: data retrieved successfully","data":{"id":0,"first_name":"Test","last_name":"User","card_number_id":"12345","warehouse_id":10}}`,
		},
		{
			name:           "Bad Request - Invalid JSON",
			requestBody:    `{"first_name":"Test","last_name":"User",`, // Malformed JSON
			mockReturnErr:  nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"unexpected EOF","data":null}`, // json.NewDecoder error message (often "EOF" for incomplete JSON)
		},
		{
			name:           "Unprocessable Entity - Missing Required Fields",
			requestBody:    `{"first_name":"","last_name":"User","card_number_id":"12345","warehouse_id":10}`, // Missing FirstName
			mockReturnErr:  nil,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"success":false,"message":"` + e.ErrRequestWrongBody.Error() + `","data":null}`,
		},
		{
			name:           "Unprocessable Entity - Non-numeric CardNumberID",
			requestBody:    `{"first_name":"","last_name":"User","card_number_id":"abc","warehouse_id":10}`, // CardNumberID is "abc"
			mockReturnErr:  nil,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   `{"success":false,"message":"handler: body does not meet requirements","data":null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockEmployeeService)
			handler := NewEmployeeHandler(mockService)

			if tt.mockReturnErr != nil || (tt.mockReturnErr == nil && tt.expectedStatus == http.StatusCreated) {
				mockService.On("Save", mock.AnythingOfType("*models.Employee")).Return(tt.mockReturnErr).Once()
			}
			req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json") // Important: set Content-Type header
			rr := httptest.NewRecorder()

			// Call the handler
			handler.Create().ServeHTTP(rr, req)

			// Assertions
			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status code mismatch")
			assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected response body mismatch")

			// Verify mock interactions
			if tt.mockReturnErr != nil || tt.expectedStatus == http.StatusCreated {
				// If service.Save was expected to be called (either for success or service error)
				mockService.AssertExpectations(t)
			} else {
				// If service.Save was NOT expected to be called (due to handler-level validation errors)
				mockService.AssertNotCalled(t, "Save", mock.Anything)
			}
		})
	}
}

func TestEmployeeHandler_Edit(t *testing.T) {
	tests := []struct {
		name                  string
		employeeID            string
		requestBody           string
		mockFindByIDReturnEmp *mod.Employee
		mockFindByIDReturnErr error
		mockUpdateReturnErr   error
		expectedStatus        int
		expectedBody          string
	}{
		{
			name:                  "Success - Employee Updated",
			employeeID:            "1",
			requestBody:           `{"first_name":"UpdatedName","Warehouse_id":20}`, // Partial update
			mockFindByIDReturnEmp: &mod.Employee{ID: 1, FirstName: "OriginalName", LastName: "User", CardNumberID: "123", WarehouseID: 10},
			mockFindByIDReturnErr: nil,
			mockUpdateReturnErr:   nil,
			expectedStatus:        http.StatusOK,
			expectedBody:          `{"success":true,"message":"handler: data retrieved successfully","data":{"id":1,"first_name":"UpdatedName","last_name":"User","card_number_id":"123","warehouse_id":20}}`,
			// The `ID` in data should be the updated ID (from URL param)
		},
		{
			name:                  "Bad Request - Missing ID in URL",
			employeeID:            "",
			requestBody:           `{"FirstName":"UpdatedName"}`,
			mockFindByIDReturnEmp: nil,
			mockFindByIDReturnErr: nil,
			mockUpdateReturnErr:   nil,
			expectedStatus:        http.StatusBadRequest,
			expectedBody:          `{"success":false,"message":"id required","data":null}`,
		},
		{
			name:                  "Bad Request - Invalid ID Format in URL",
			employeeID:            "abc",
			requestBody:           `{"FirstName":"UpdatedName"}`,
			mockFindByIDReturnEmp: nil,
			mockFindByIDReturnErr: nil,
			mockUpdateReturnErr:   nil,
			expectedStatus:        http.StatusBadRequest,
			expectedBody:          `{"success":false,"message":"` + e.ErrRequestIdMustBeInt.Error() + `","data":null}`,
		},
		{
			name:                  "Unprocessable Entity - Invalid JSON Body",
			employeeID:            "1",
			requestBody:           `{"FirstName":"UpdatedName`, // Malformed JSON
			mockFindByIDReturnEmp: nil,                         // FindByID is not called if JSON decode fails
			mockFindByIDReturnErr: nil,
			mockUpdateReturnErr:   nil,
			expectedStatus:        http.StatusUnprocessableEntity,                                                       // Or 400, depending on utils.BadResponse
			expectedBody:          `{"success":false,"message":"handler: body does not meet requirements","data":null}`, // Common json.NewDecoder error
		},
		{
			name:                  "Not Found - Employee Does Not Exist for Update",
			employeeID:            "99", // Valid ID, but not found
			requestBody:           `{"FirstName":"UpdatedName"}`,
			mockFindByIDReturnEmp: nil,
			mockFindByIDReturnErr: errors.New("employee not found"), // Service returns Not Found
			mockUpdateReturnErr:   nil,
			expectedStatus:        http.StatusNotFound,
			expectedBody:          `{"success":false,"message":"employee not found","data":null}`,
		},
		{
			name:                  "Conflict - Service Returns Update Error",
			employeeID:            "1",
			requestBody:           `{"FirstName":"UpdatedName"}`,
			mockFindByIDReturnEmp: &mod.Employee{ID: 1, FirstName: "OriginalName", LastName: "User", CardNumberID: "123", WarehouseID: 10},
			mockFindByIDReturnErr: nil,
			mockUpdateReturnErr:   errors.New("update conflict: card number already exists"), // Simulate service conflict
			expectedStatus:        http.StatusConflict,                                       // Or 422 if your service returns that
			expectedBody:          `{"success":false,"message":"update conflict: card number already exists","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockEmployeeService)
			handler := NewEmployeeHandler(mockService)

			if tt.mockFindByIDReturnEmp != nil || tt.mockFindByIDReturnErr != nil {
				mockService.On("FindByID", mock.AnythingOfType("int")).Return(tt.mockFindByIDReturnEmp, tt.mockFindByIDReturnErr).Once()
			}

			if tt.mockUpdateReturnErr != nil || (tt.mockFindByIDReturnEmp != nil && tt.mockUpdateReturnErr == nil && tt.expectedStatus == http.StatusOK) {
				mockService.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("*models.Employee")).Return(tt.mockUpdateReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/employees/%s", tt.employeeID), bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.employeeID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			handler.Edit().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status code mismatch")
			assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected response body mismatch")

			if tt.name == "Bad Request - Missing ID in URL" ||
				tt.name == "Bad Request - Invalid ID Format in URL" ||
				tt.name == "Unprocessable Entity - Invalid JSON Body" {
				mockService.AssertNotCalled(t, "FindByID", mock.Anything)
				mockService.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
			} else if tt.name == "Not Found - Employee Does Not Exist for Update" {
				mockService.AssertExpectations(t)
				mockService.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
			} else {
				mockService.AssertExpectations(t)
			}
		})
	}
}

func TestEmployeeHandler_Delete(t *testing.T) {
	tests := []struct {
		name                string
		employeeID          string
		mockDeleteReturnErr error
		expectedStatus      int
		expectedBody        string // 204 No Content should have no body
	}{
		{
			name:                "Success - Employee Deleted",
			employeeID:          "1",
			mockDeleteReturnErr: nil,
			expectedStatus:      http.StatusNoContent, // 204 No Content
			expectedBody:        "",
		},
		{
			name:                "Bad Request - Missing ID in URL",
			employeeID:          "",
			mockDeleteReturnErr: nil,
			expectedStatus:      http.StatusBadRequest,
			expectedBody:        `{"success":false,"message":"id required","data":null}`,
		},
		{
			name:                "Bad Request - Invalid ID Format in URL",
			employeeID:          "xyz",
			mockDeleteReturnErr: nil,
			expectedStatus:      http.StatusBadRequest,
			expectedBody:        `{"success":false,"message":"` + e.ErrRequestIdMustBeInt.Error() + `","data":null}`,
		},
		{
			name:                "Not Found - Employee Does Not Exist for Deletion",
			employeeID:          "99",
			mockDeleteReturnErr: errors.New("employee not found"), // Service returns Not Found
			expectedStatus:      http.StatusNotFound,
			expectedBody:        `{"success":false,"message":"employee not found","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockEmployeeService)
			handler := NewEmployeeHandler(mockService)

			if tt.mockDeleteReturnErr != nil || tt.expectedStatus == http.StatusNoContent {
				mockService.On("Delete", mock.AnythingOfType("int")).Return(tt.mockDeleteReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/employees/%s", tt.employeeID), nil)
			rr := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.employeeID)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			handler.Delete().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status code mismatch")
			if tt.expectedStatus == http.StatusNoContent {
				assert.Empty(t, rr.Body.String(), "Expected empty body for 204 No Content")
			} else {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected response body mismatch")
			}

			if tt.name == "Bad Request - Missing ID in URL" ||
				tt.name == "Bad Request - Invalid ID Format in URL" {
				mockService.AssertNotCalled(t, "Delete", mock.Anything)
			} else {
				mockService.AssertExpectations(t)
			}
		})
	}
}
