package handler

import (
	"bytes"
	"errors"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/tests/buyers/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type THandlerPurchaseOrderSuite struct {
	suite.Suite
	mockService *mocks.MockPurchaseOrderService
	handler     *PurchaseOrderHandler
}

func (s *THandlerPurchaseOrderSuite) SetupTest() {
	s.mockService = new(mocks.MockPurchaseOrderService)
	s.handler = NewPurchaseOrderHandler(s.mockService)
}

func (s *THandlerPurchaseOrderSuite) TestPurchaseOrderHandler_Create() {
	t := s.T()

	type TestTablePOCreate struct {
		name           string
		requestBody    string
		mockError      error
		requireMock    bool
		expectedStatus int
		expectedBody   string
		funcRun        func(args mock.Arguments)
	}

	testsTable := []TestTablePOCreate{
		{
			name:      "Case 1: Success",
			mockError: nil,
			funcRun: func(args mock.Arguments) {
				b := args.Get(0).(*mod.PurchaseOrder)
				b.ID = 1
			},
			requestBody: `
				{
					"order_number": "ORD-2024-101",
					"order_date": "2024-12-01",
					"tracking_code": "abscf123",
					"buyer_id": 6,
					"products_details": [
						{
							"quantity": 1,
							"product_record_id": 1,
							"clean_liness_status": "Ready",
							"temperature": 10.6
						},
						{
							"quantity": 4,
							"product_record_id": 2,
							"clean_liness_status": "Ready",
							"temperature": 9.6
						}
					]
				}`,
			requireMock:    true,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"data":{"buyer_id":6, "id":1, "order_date":"2024-12-01", "order_number":"ORD-2024-101", "products_details":[{"clean_liness_status":"Ready", "id":0, "product_record_id":1, "purchase_order_id":0, "quantity":1, "temperature":10.6}, {"clean_liness_status":"Ready", "id":0, "product_record_id":2, "purchase_order_id":0, "quantity":4, "temperature":9.6}], "tracking_code":"abscf123"}, "message":"Purchase order creado exitosamente", "success":true}`,
		},
		{
			name:           "Case 2: Fail - Failed body",
			mockError:      e.ErrRequestFailedBody,
			expectedBody:   `{"success":false,"message":"handler: failed to read body","data":null}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:      "Case 3: Fail - Failed validation Purchase Order",
			mockError: e.ErrRequestWrongBody,
			requestBody: `
				{
					"order_number": "ORD-2024-101",
					"order_date": "2024-12-01",
					"tracking_code": "",
					"buyer_id": 6,
					"products_details": [
						{
							"quantity": 1,
							"product_record_id": 1,
							"clean_liness_status": "Ready",
							"temperature": 10.6
						}
					]
				}`,
			expectedBody:   `{"success":false,"message":"handler: body does not meet requirements: TrackingCode is required","data":null}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:      "Case 4: Fail - Failed validation Order Details",
			mockError: e.ErrRequestWrongBody,
			requestBody: `
				{
					"order_number": "ORD-2024-101",
					"order_date": "2024-12-01",
					"tracking_code": "120",
					"buyer_id": 6,
					"products_details": [
						{
							"quantity": 1,
							"product_record_id": 1,
							"temperature": 10.6
						}
					]
				}`,
			expectedBody:   `{"success":false,"message":"handler: body does not meet requirements: ProductDetails[0]: CleanLinessStatus is required","data":null}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:      "Case 5: Fail - Card duplicated",
			mockError: e.ErrPORepositoryOrderNumberDuplicated,
			requestBody: `
				{
					"order_number": "ORD-2024-101",
					"order_date": "2024-12-01",
					"tracking_code": "120",
					"buyer_id": 6,
					"products_details": [
						{
							"quantity": 1,
							"product_record_id": 1,
							"temperature": 10.6,
							"clean_liness_status": "Ready"
						}
					]
				}`,
			expectedBody:   `{"success":false,"message":"repository: Order number duplicated","data":null}`,
			expectedStatus: http.StatusConflict,
			requireMock:    true,
			funcRun:        func(args mock.Arguments) {},
		},
		{
			name:      "Case 5: Fail - Card duplicated",
			mockError: e.ErrForeignKeyError,
			requestBody: `
				{
					"order_number": "ORD-2024-101",
					"order_date": "2024-12-01",
					"tracking_code": "120",
					"buyer_id": 6,
					"products_details": [
						{
							"quantity": 1,
							"product_record_id": 1,
							"temperature": 10.6,
							"clean_liness_status": "Ready"
						}
					]
				}`,
			expectedBody:   `{"success":false,"message":"repository: unable to execute query due to foreign key error","data":null}`,
			expectedStatus: http.StatusConflict,
			requireMock:    true,
			funcRun:        func(args mock.Arguments) {},
		},
		{
			name:      "Case 6: Fail - Buyer not found",
			mockError: e.ErrBuyerRepositoryNotFound,
			requestBody: `
				{
					"order_number": "ORD-2024-101",
					"order_date": "2024-12-01",
					"tracking_code": "120",
					"buyer_id": 6,
					"products_details": [
						{
							"quantity": 1,
							"product_record_id": 1,
							"temperature": 10.6,
							"clean_liness_status": "Ready"
						}
					]
				}`,
			expectedBody:   `{"success":false,"message":"repository: buyer not found","data":null}`,
			expectedStatus: http.StatusConflict,
			requireMock:    true,
			funcRun:        func(args mock.Arguments) {},
		},
		{
			name:      "Case 6: Fail - Buyer not found",
			mockError: errors.New("bad request"),
			requestBody: `
				{
					"order_number": "ORD-2024-101",
					"order_date": "2024-12-01",
					"tracking_code": "120",
					"buyer_id": 6,
					"products_details": [
						{
							"quantity": 1,
							"product_record_id": 1,
							"temperature": 10.6,
							"clean_liness_status": "Ready"
						}
					]
				}`,
			expectedBody:   `{"success":false,"message":"bad request","data":null}`,
			expectedStatus: http.StatusBadRequest,
			requireMock:    true,
			funcRun:        func(args mock.Arguments) {},
		},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			//Given
			s.SetupTest()
			if test.requireMock {
				s.mockService.On("Save", mock.AnythingOfType("*models.PurchaseOrder")).
					Run(func(args mock.Arguments) {
						test.funcRun(args)
					}).Return(test.mockError)
			}

			//When
			req := httptest.NewRequest(http.MethodPost, "/purchaseOrders", bytes.NewBufferString(test.requestBody))
			req.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()

			s.handler.Create()(recorder, req)

			//Then
			expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
			require.Equal(t, test.expectedStatus, recorder.Code)
			require.Equal(t, expectedHeaders, recorder.Header())
			require.JSONEq(t, test.expectedBody, recorder.Body.String())

			s.mockService.AssertExpectations(s.T())
		})
	}
}

func TestHanlderSuite(t *testing.T) {
	suite.Run(t, new(THandlerPurchaseOrderSuite))
}
