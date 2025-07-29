package handler

import (
	"bytes"
	"errors"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	tests2 "github.com/smartineztri_meli/W17-G2-Bootcamp/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInboundHandler_Create(t *testing.T) {
	// Define test cases using a slice of structs
	tests := []struct {
		name               string
		requestBody        string             // Changed to string to directly provide JSON
		mockReturnOrder    *mod.InboundOrders // Expected return value from mock service
		mockReturnErr      error              // Error to be returned by mock service
		expectedStatusCode int
		expectedBody       string // Full JSON string for assertion
	}{
		{
			name:               "Success - Inbound Order Created",
			requestBody:        `{"order_number":"ORD001","order_date":"2023-01-01","employee_id":1,"product_batch_id":101,"warehouse_id":1001}`,
			mockReturnOrder:    &mod.InboundOrders{Id: 1, OrderNumber: "ORD001", OrderDate: "2023-01-01", EmployeeId: 1, ProductBatchId: 101, WarehouseId: 1001},
			mockReturnErr:      nil,
			expectedStatusCode: http.StatusCreated,
			expectedBody:       `{"success":true,"message":"handler: data retrieved successfully","data":{"id":1,"order_number":"ORD001","order_date":"2023-01-01","employee_id":1,"product_batch_id":101,"warehouse_id":1001}}`,
		},
		{
			name:               "Failure - Invalid JSON Format",
			requestBody:        `{"order_number": "ORD002", "order_date": "2023-01-02", "employee_id": "invalid"`,
			mockReturnOrder:    nil,
			mockReturnErr:      nil,
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedBody:       `{"success":false,"message":"Invalid JSON format","data":null}`,
		},
		{
			name:               "Failure - Invalid Data",
			requestBody:        `{"order_number":"","order_date":"2023-01-03","employee_id":2,"product_record_id":102,"warehouse_id":1002}`,
			mockReturnOrder:    &mod.InboundOrders{},
			mockReturnErr:      e.ErrInboundOrderInvalidData,
			expectedStatusCode: http.StatusUnprocessableEntity,
			expectedBody:       `{"success":false,"message":"invalid inbound order data","data":null}`,
		},
		{
			name:               "Failure - Order Number Already Exists",
			requestBody:        `{"order_number":"ORD001","order_date":"2023-01-04","employee_id":3,"product_record_id":103,"warehouse_id":1003}`,
			mockReturnOrder:    &mod.InboundOrders{},
			mockReturnErr:      errors.New("order number already exists"),
			expectedStatusCode: http.StatusConflict,
			expectedBody:       `{"success":false,"message":"order number already exists","data":null}`,
		},
		{
			name:               "Failure - Employee Not Found",
			requestBody:        `{"order_number":"ORD005","order_date":"2023-01-05","employee_id":999,"product_record_id":105,"warehouse_id":1005}`,
			mockReturnOrder:    &mod.InboundOrders{},
			mockReturnErr:      errors.New("employee not found"),
			expectedStatusCode: http.StatusConflict,
			expectedBody:       `{"success":false,"message":"employee not found","data":null}`,
		},
		{
			name:               "Failure - Internal Server Error",
			requestBody:        `{"order_number":"ORD006","order_date":"2023-01-06","employee_id":4,"product_record_id":106,"warehouse_id":1006}`,
			mockReturnOrder:    &mod.InboundOrders{},
			mockReturnErr:      errors.New("something went wrong"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       `{"success":false,"message":"something went wrong","data":null}`,
		},
	}

	// Iterate over test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(tests2.MockInboundService)

			if tc.mockReturnErr != nil || tc.expectedStatusCode == http.StatusCreated {
				mockService.On("Save", mock.AnythingOfType("*models.InboundOrders")).Return(tc.mockReturnOrder, tc.mockReturnErr).Once()
			}

			handler := NewInboundHandler(mockService)

			req := httptest.NewRequest(http.MethodPost, "/inbound-orders", bytes.NewBufferString(tc.requestBody))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			handler.Create().ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedStatusCode, rr.Code, "Expected status code mismatch")
			assert.JSONEq(t, tc.expectedBody, rr.Body.String(), "Expected response body mismatch")

			if tc.mockReturnErr != nil || tc.expectedStatusCode == http.StatusCreated {
				mockService.AssertExpectations(t)
			} else {
				mockService.AssertNotCalled(t, "Save", mock.Anything)
			}
		})
	}
}

// TestInboundHandler_GetOrdersByEmployee tests the GetOrdersByEmployee method of InboundHandler
func TestInboundHandler_GetOrdersByEmployee(t *testing.T) {
	tests := []struct {
		name               string
		employeeIDParam    string // Query parameter string
		mockServiceSetup   func(*tests2.MockInboundService)
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:            "Success - Orders Found for Employee ID",
			employeeIDParam: "id=1",
			mockServiceSetup: func(m *tests2.MockInboundService) {
				// Expect FindOrdersByEmployee to be called with ID 1 and return data
				m.On("FindOrdersByEmployee", 1).Return([]mod.EmployeeReport{
					{ID: 1, CardNumberID: "EMP001", FirstName: "John", LastName: "Doe", WarehouseID: 100, InboundOrdersCount: 5},
				}, nil).Once()
			},
			expectedStatusCode: http.StatusCreated, // Note: The handler returns StatusCreated, which might be unusual for GET but matches the original code.
			expectedBody:       `{"success":true,"message":"handler: data retrieved successfully","data":[{"id":1,"card_number_id":"EMP001","first_name":"John","last_name":"Doe","warehouse_id":100,"inbound_orders_count":5}]}`,
		},
		{
			name:            "Success - No Employee ID Provided (All Orders)",
			employeeIDParam: "", // No ID parameter
			mockServiceSetup: func(m *tests2.MockInboundService) {
				// Expect FindOrdersByEmployee to be called with ID 0 (default for Atoi) and return data
				m.On("FindOrdersByEmployee", 0).Return([]mod.EmployeeReport{
					{ID: 1, CardNumberID: "EMP001", InboundOrdersCount: 5},
					{ID: 2, CardNumberID: "EMP002", InboundOrdersCount: 3},
				}, nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
			expectedBody:       `{"success":true,"message":"handler: data retrieved successfully","data":[{"id":1,"card_number_id":"EMP001","first_name":"","last_name":"","warehouse_id":0,"inbound_orders_count":5},{"id":2,"card_number_id":"EMP002","first_name":"","last_name":"","warehouse_id":0,"inbound_orders_count":3}]}`,
		},
		{
			name:               "Failure - Invalid Employee ID Format",
			employeeIDParam:    "id=abc", // Non-integer ID
			mockServiceSetup:   func(m *tests2.MockInboundService) {},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"success":false,"message":"strconv.Atoi: parsing \"abc\": invalid syntax","data":null}`,
		},
		{
			name:            "Failure - Employee Not Found",
			employeeIDParam: "id=999",
			mockServiceSetup: func(m *tests2.MockInboundService) {
				// Expect FindOrdersByEmployee to be called with ID 999 and return ErrEmployeeNotFound
				m.On("FindOrdersByEmployee", 999).Return([]mod.EmployeeReport{}, e.ErrEmployeeNotFound).Once()
			},
			expectedStatusCode: http.StatusNotFound,
			expectedBody:       `{"success":false,"message":"employee not found","data":null}`,
		},
		{
			name:            "Failure - Internal Server Error",
			employeeIDParam: "id=2",
			mockServiceSetup: func(m *tests2.MockInboundService) {
				// Expect FindOrdersByEmployee to be called with ID 2 and return a generic error
				m.On("FindOrdersByEmployee", 2).Return([]mod.EmployeeReport{}, errors.New("database connection error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       `{"success":false,"message":"database connection error","data":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(tests2.MockInboundService)
			tt.mockServiceSetup(mockService)

			handler := NewInboundHandler(mockService)

			// Construct the request URL with the query parameter
			reqURL := "/inbound-orders/report"
			if tt.employeeIDParam != "" {
				reqURL += "?" + tt.employeeIDParam
			}

			req := httptest.NewRequest(http.MethodGet, reqURL, nil)
			rr := httptest.NewRecorder()

			handler.GetOrdersByEmployee().ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatusCode, rr.Code, "Expected status code mismatch")
			assert.JSONEq(t, tt.expectedBody, rr.Body.String(), "Expected body content mismatch")

			mockService.AssertExpectations(t)
		})
	}
}
