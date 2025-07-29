package handler_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	hd "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/handler"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	tests2 "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func addChiURLParam(r *http.Request, key, value string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, value)
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func TestSellerHandler_GetAll(t *testing.T) {
	tests := []struct {
		name           string
		mockReturnData []mod.Seller
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "#1 Success - Get All Sellers",
			mockReturnData: []mod.Seller{
				{ID: 1, CID: 101, CompanyName: "Test Corp", Address: "123 Test St", Telephone: "555-1234", Locality: 1},
				{ID: 2, CID: 102, CompanyName: "Sample Inc", Address: "456 Sample Ave", Telephone: "555-5678", Locality: 2},
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"succes","data":[{"id":1,"cid":101,"company_name":"Test Corp","address":"123 Test St","telephone":"555-1234","locality_id":1},{"id":2,"cid":102,"company_name":"Sample Inc","address":"456 Sample Ave","telephone":"555-5678","locality_id":2}]}`,
		},
		{
			name:           "#2 Error - Service failure",
			mockReturnData: nil,
			mockReturnErr:  e.ErrRepositoryDatabase,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"repository: database operation failed","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockSellerService)
			handler := hd.NewSellerHandler(mockService)

			mockService.On("FindAll").Return(tt.mockReturnData, tt.mockReturnErr).Once()

			req := httptest.NewRequest(http.MethodGet, "/sellers", nil)
			rr := httptest.NewRecorder()

			handler.GetAll().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}

func TestSellerHandler_GetByID(t *testing.T) {
	tests := []struct {
		name           string
		sellerID       string
		mockReturnData mod.Seller
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "#1 Success - Seller Found",
			sellerID: "1",
			mockReturnData: mod.Seller{
				ID: 1, CID: 101, CompanyName: "Test Corp", Address: "123 Test St", Telephone: "555-1234", Locality: 1,
			},
			mockReturnErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"success":true,"message":"success","data":{"id":1,"cid":101,"company_name":"Test Corp","address":"123 Test St","telephone":"555-1234","locality_id":1}}`,
		},
		{
			name:           "#2 Error - Seller Not Found",
			sellerID:       "99",
			mockReturnData: mod.Seller{},
			mockReturnErr:  e.ErrSellerRepositoryNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"success":false,"message":"repository: seller not found","data":null}`,
		},
		{
			name:           "#3 Error - Invalid ID",
			sellerID:       "abc",
			mockReturnData: mod.Seller{},
			mockReturnErr:  nil, // Service is not called
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"handler: id must be an integer","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockSellerService)
			handler := hd.NewSellerHandler(mockService)

			if tt.mockReturnErr != nil || (tt.mockReturnErr == nil && tt.expectedStatus == http.StatusOK) {
				mockService.On("FindByID", mock.AnythingOfType("int")).Return(tt.mockReturnData, tt.mockReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodGet, "/sellers/"+tt.sellerID, nil)
			req = addChiURLParam(req, "id", tt.sellerID)
			rr := httptest.NewRecorder()

			handler.GetByID().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}

func TestSellerHandler_Create(t *testing.T) {
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
			name:              "#1 Success - Seller Created",
			requestBody:       `{"cid": 101, "company_name": "New Company", "address": "123 Main St", "telephone": "555-0101", "locality_id": 1}`,
			mockReturnID:      1,
			mockReturnErr:     nil,
			expectServiceCall: true,
			expectedStatus:    http.StatusCreated,
			expectedBody:      `{"success":true,"message":"success","data":1}`,
		},
		{
			name:              "#2 Error - Invalid JSON",
			requestBody:       `{"cid": "charlie", "company_name": "New Company", "address": "123 Main St", "telephone": "555-0101", "locality_id": 1}`,
			expectServiceCall: false,
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      `{"success":false,"message":"handler: body does not meet requirements","data":null}`,
		},
		{
			name:              "#3 Error - Validation Failed",
			requestBody:       `{"company_name": "New Company", "address": "123 Main St", "telephone": "555-0101", "locality_id": 1}`, // Missing required fields
			expectServiceCall: false,
			expectedStatus:    http.StatusUnprocessableEntity,
			expectedBody:      `{"success":false,"message":"CID is required, ","data":null}`,
		},
		{
			name:              "#4 Error - Service Conflict",
			requestBody:       `{"cid": 101, "company_name": "New Company", "address": "123 Main St", "telephone": "555-0101", "locality_id": 1}`,
			mockReturnID:      0,
			mockReturnErr:     e.ErrSellerRepositoryDuplicated,
			expectServiceCall: true,
			expectedStatus:    http.StatusConflict,
			expectedBody:      `{"success":false,"message":"repository: seller already exists","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockSellerService)
			handler := hd.NewSellerHandler(mockService)

			if tt.expectServiceCall {
				mockService.On("Save", mock.AnythingOfType("*models.Seller")).Return(tt.mockReturnID, tt.mockReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodPost, "/sellers", bytes.NewBufferString(tt.requestBody))
			rr := httptest.NewRecorder()

			handler.Create().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			// For validation error, we check for contains because the order of fields is not guaranteed.
			if tt.expectedStatus == http.StatusUnprocessableEntity {
				assert.Contains(t, rr.Body.String(), "is required")
			} else {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			}
			mockService.AssertExpectations(t)
		})
	}
}

func TestSellerHandler_Update(t *testing.T) {
	tests := []struct {
		name             string
		sellerID         string
		requestBody      string
		mockFindData     mod.Seller
		mockFindErr      error
		mockUpdateErr    error
		expectFindCall   bool
		expectUpdateCall bool
		expectedStatus   int
		expectedBody     string
	}{
		{
			name:             "#1 Success - Seller Updated",
			sellerID:         "1",
			requestBody:      `{"company_name": "Updated Corp", "telephone": "555-9999"}`,
			mockFindData:     mod.Seller{ID: 1, CID: 101, CompanyName: "Test Corp", Address: "123 Test St", Telephone: "555-1234", Locality: 1},
			mockFindErr:      nil,
			mockUpdateErr:    nil,
			expectFindCall:   true,
			expectUpdateCall: true,
			expectedStatus:   http.StatusOK,
			expectedBody:     `{"success":true,"message":"success","data":null}`,
		},
		{
			name:             "#2 Error - Seller Not Found on FindByID",
			sellerID:         "99",
			requestBody:      `{"company_name": "Updated Corp"}`,
			mockFindData:     mod.Seller{},
			mockFindErr:      e.ErrSellerRepositoryNotFound,
			expectFindCall:   true,
			expectUpdateCall: false,
			expectedStatus:   http.StatusNotFound,
			expectedBody:     `{"success":false,"message":"repository: seller not found","data":null}`,
		},
		{
			name:             "#3 Error - Invalid ID",
			sellerID:         "abc",
			requestBody:      `{}`,
			expectFindCall:   false,
			expectUpdateCall: false,
			expectedStatus:   http.StatusBadRequest,
			expectedBody:     `{"success":false,"message":"handler: id must be an integer","data":null}`,
		},
		{
			name:             "#4 Error - Invalid JSON Body",
			sellerID:         "1",
			requestBody:      `{"company_name":}`,
			mockFindData:     mod.Seller{ID: 1, CID: 101, CompanyName: "Test Corp", Address: "123 Test St", Telephone: "555-1234", Locality: 1},
			mockFindErr:      nil,
			expectFindCall:   true,
			expectUpdateCall: false,
			expectedStatus:   http.StatusBadRequest,
			expectedBody:     `{"success":false,"message":"handler: body does not meet requirements","data":null}`,
		},
		{
			name:             "#5 Error - Update Conflict",
			sellerID:         "1",
			requestBody:      `{"cid": 1001}`,
			mockFindData:     mod.Seller{ID: 1, CID: 101, CompanyName: "Test Corp", Address: "123 Test St", Telephone: "555-1234", Locality: 1},
			mockFindErr:      nil,
			mockUpdateErr:    e.ErrSellerRepositoryDuplicated,
			expectFindCall:   true,
			expectUpdateCall: true,
			expectedStatus:   http.StatusBadRequest,
			expectedBody:     `{"success":false,"message":"repository: seller already exists","data":null}`,
		},
		{
			name:             "#6 Error - Validation Failed After Patch",
			sellerID:         "1",
			requestBody:      `{"cid":0}`,
			mockFindData:     mod.Seller{ID: 1, CID: 101, CompanyName: "Alpha", Address: "123 Test St", Telephone: "555-1234", Locality: 1},
			mockFindErr:      nil,
			expectFindCall:   true,
			expectUpdateCall: false,
			expectedStatus:   http.StatusUnprocessableEntity,
			expectedBody:     `{"success":false,"message":"CID is required, ","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockSellerService)
			handler := hd.NewSellerHandler(mockService)

			if tt.expectFindCall {
				mockService.On("FindByID", mock.AnythingOfType("int")).Return(tt.mockFindData, tt.mockFindErr).Once()
			}

			if tt.expectUpdateCall {
				mockService.On("Update", mock.AnythingOfType("*models.Seller")).Return(tt.mockUpdateErr).Once()
			}

			req := httptest.NewRequest(http.MethodPatch, "/sellers/"+tt.sellerID, bytes.NewBufferString(tt.requestBody))
			req = addChiURLParam(req, "id", tt.sellerID)
			rr := httptest.NewRecorder()

			handler.Update().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.expectedStatus == http.StatusUnprocessableEntity {
				assert.Contains(t, rr.Body.String(), "is required, ")
			} else {
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			}
			mockService.AssertExpectations(t)
		})
	}
}

func TestSellerHandler_Delete(t *testing.T) {
	tests := []struct {
		name           string
		sellerID       string
		mockReturnErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "#1 Success - Seller Deleted",
			sellerID:       "1",
			mockReturnErr:  nil,
			expectedStatus: http.StatusNoContent,
			expectedBody:   `{"success":true,"message":"success","data":null}`,
		},
		{
			name:           "#2 Error - Seller Not Found",
			sellerID:       "99",
			mockReturnErr:  e.ErrSellerRepositoryNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"success":false,"message":"repository: seller not found","data":null}`,
		},
		{
			name:           "#3 Error - Invalid ID",
			sellerID:       "abc",
			mockReturnErr:  nil, // Service is not called
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"handler: id must be an integer","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockSellerService)
			handler := hd.NewSellerHandler(mockService)

			if tt.sellerID != "abc" {
				mockService.On("Delete", mock.AnythingOfType("int")).Return(tt.mockReturnErr).Once()
			}

			req := httptest.NewRequest(http.MethodDelete, "/sellers/"+tt.sellerID, nil)
			req = addChiURLParam(req, "id", tt.sellerID)
			rr := httptest.NewRecorder()

			handler.Delete().ServeHTTP(rr, req)

			// For 204 No Content, the body is often empty, but your handler sends a body.
			if tt.expectedStatus == http.StatusNoContent {
				assert.Equal(t, http.StatusNoContent, rr.Code)
			} else {
				assert.Equal(t, tt.expectedStatus, rr.Code)
				assert.JSONEq(t, tt.expectedBody, rr.Body.String())
			}
			mockService.AssertExpectations(t)
		})
	}
}
