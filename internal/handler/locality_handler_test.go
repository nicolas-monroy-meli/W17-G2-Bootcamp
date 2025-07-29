package handler_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	hd "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/handler"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	tests2 "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLocalityHandler_GetAll(t *testing.T) {
	mockService := new(tests2.MockLocalityService)
	handler := hd.NewLocalityHandler(mockService)

	tests := []struct {
		name           string
		mockReturnEmp  []mod.Locality
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "#1 Success - Retrieve All Localities",
			mockReturnEmp: []mod.Locality{
				{ID: 1, Name: "Manhattan", Province: "New York", Country: "USA"},
				{ID: 2, Name: "Downtown", Province: "California", Country: "USA"},
				{ID: 3, Name: "Lakeview", Province: "Illinois", Country: "USA"},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"success","data":[{"id":1,"locality_name":"Manhattan","province_name":"New York","country_name":"USA"},{"id":2,"locality_name":"Downtown","province_name":"California","country_name":"USA"},{"id":3,"locality_name":"Lakeview","province_name":"Illinois","country_name":"USA"}]}`},
		{
			name:           "#2 Error - Service Returns Error",
			mockReturnEmp:  nil,
			mockReturnErr:  e.ErrQueryIsEmpty,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"success":false,"message":"repository: query returned no info","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configurar el comportamiento esperado del mock para FindAll
			// Aquí siempre esperamos que FindAll sea llamado
			mockService.On("FindAllLocalities").Return(tt.mockReturnEmp, tt.mockReturnErr).Once()

			// Crear una petición HTTP simulada para GET /localities (sin ID en la URL)
			req := httptest.NewRequest(http.MethodGet, "/localities", nil)
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

func TestLocalityHandler_GetSelByLocID(t *testing.T) {
	tests := []struct {
		name              string
		localityID        string
		mockReturnData    []mod.SelByLoc
		mockReturnErr     error
		expectServiceCall bool
		expectedStatus    int
		expectedBody      string
	}{
		{
			name:              "#1 Success - Sellers by localities Found",
			localityID:        "",
			expectServiceCall: true,
			mockReturnData: []mod.SelByLoc{
				{ID: 1, Name: "Brooklyn", Count: 2},
				{ID: 2, Name: "Santa Monica", Count: 2},
				{ID: 3, Name: "Cambridge", Count: 4},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"success","data":[{"locality_id":1,"locality_name":"Brooklyn","sellers_count":2},{"locality_id":2,"locality_name":"Santa Monica","sellers_count":2},{"locality_id":3,"locality_name":"Cambridge","sellers_count":4}]}`,
		},
		{
			name:              "#2 Success - Locality Found",
			localityID:        "1",
			expectServiceCall: true,
			mockReturnData: []mod.SelByLoc{
				{ID: 1, Name: "Brooklyn", Count: 2},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"success","data":[{"locality_id":1,"locality_name":"Brooklyn","sellers_count":2}]}`,
		},
		{
			name:              "#3 Error - Bad Request ID Must be Int",
			localityID:        "abc",
			expectServiceCall: false,
			mockReturnData:    nil,
			mockReturnErr:     nil,
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      `{"success":false,"message":"handler: id must be an integer","data":null}`,
		},
		{
			name:              "#4 Error - Bad Request ID Must be Greater than 0",
			localityID:        "-1",
			expectServiceCall: false,
			mockReturnData:    nil,
			mockReturnErr:     nil,
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      `{"success":false,"message":"handler: id must be greater than 0","data":null}`,
		},
		{
			name:              "#5 Error - Service Returns Error",
			localityID:        "99",
			expectServiceCall: true,
			mockReturnData:    nil,
			mockReturnErr:     e.ErrLocalityRepositoryNotFound,
			expectedStatus:    http.StatusNotFound,
			expectedBody:      `{"success":false,"message":"repository: locality not found","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockLocalityService)
			handler := hd.NewLocalityHandler(mockService)

			if tt.expectServiceCall {
				var expectedID int
				if tt.localityID == "" {
					expectedID = -1
				} else {
					expectedID, _ = strconv.Atoi(tt.localityID)
				}
				mockService.On("FindSellersByLocID", expectedID).Return(tt.mockReturnData, tt.mockReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/localities?id=%s", tt.localityID), nil)
			rr := httptest.NewRecorder()

			rctx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			handler.GetSelByLocID().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status code mismatch")

			assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected response body mismatch")

			mockService.AssertExpectations(t)
		})
	}
}

func TestLocalityHandler_Create(t *testing.T) {
	tests := []struct {
		name              string
		requestBody       string
		mockReturnID      int
		mockReturnErr     error
		expectServiceCall bool
		expectedStatus    int
		expectedBody      string
	}{
		{
			name:              "#1 Success - Locality Created",
			requestBody:       `{"locality_name":"Bogota","province_name":"Cundinamarca","country_name":"Colombia"}`,
			mockReturnID:      1,
			mockReturnErr:     nil,
			expectServiceCall: true,
			expectedStatus:    http.StatusCreated,
			expectedBody:      `{"success":true,"message":"success","data":1}`,
		},
		{
			name:              "#2 Error - Bad Request - Invalid JSON Body",
			requestBody:       `{"locality_name":1,"province_name":"Cundinamarca","country_name":"Colombia"}`,
			expectServiceCall: false,
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      `{"success":false,"message":"handler: body does not meet requirements","data":null}`,
		},
		{
			name:              "#3 Error - Unprocessable Entity - Missing Required Fields",
			requestBody:       `{"province_name":"Cundinamarca","country_name":"Colombia"}`,
			expectServiceCall: false,
			expectedStatus:    http.StatusUnprocessableEntity,
			expectedBody:      `{"success":false,"message":"Name is required, ","data":null}`,
		},
		{
			name:              "#4 Error - Conflict - Locality duplicated",
			requestBody:       `{"locality_name":"Bogota","province_name":"Cundinamarca","country_name":"Colombia"}`,
			mockReturnID:      0,
			mockReturnErr:     e.ErrLocalityRepositoryDuplicated,
			expectServiceCall: true,
			expectedStatus:    http.StatusConflict,
			expectedBody:      `{"success":false,"message":"repository: locality already exists","data":null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockLocalityService)
			handler := hd.NewLocalityHandler(mockService)

			if tt.expectServiceCall {
				mockService.On("Save", mock.AnythingOfType("*models.Locality")).Return(tt.mockReturnID, tt.mockReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodPost, "/localities", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.Create().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code, "Expected status code mismatch")
			assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected response body mismatch")

			mockService.AssertExpectations(t)
		})
	}
}
