package handler

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
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

type THandlerBuyerSuite struct {
	suite.Suite
	mockService *mocks.MockBuyerService
	handler     *BuyerHandler
}

func (s *THandlerBuyerSuite) SetupTest() {
	s.mockService = new(mocks.MockBuyerService)
	s.handler = NewBuyerHandler(s.mockService)
}

func (s *THandlerBuyerSuite) TestGetAll() {
	t := s.T()

	type TestTableBuyerGet struct {
		name           string
		mockReturn     []mod.Buyer
		mockError      error
		isError        bool
		expectedStatus int
		expectedBody   string
	}

	buyers := []mod.Buyer{
		{
			ID:           1,
			CardNumberID: "1234567890123456",
			FirstName:    "Juan",
			LastName:     "Pérez",
		},
		{
			ID:           2,
			CardNumberID: "9876543210987654",
			FirstName:    "Ana",
			LastName:     "González",
		},
	}

	testsTable := []TestTableBuyerGet{
		{
			name:           "Case 1: Success",
			mockReturn:     buyers,
			mockError:      nil,
			isError:        false,
			expectedBody:   `{"success":true,"message":"","data":[{"id":1,"card_number_id":"1234567890123456","first_name":"Juan","last_name":"Pérez"},{"id":2,"card_number_id":"9876543210987654","first_name":"Ana","last_name":"González"}]}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Case 2: Fail Case",
			mockReturn:     nil,
			mockError:      errors.New("bad request error"),
			isError:        false,
			expectedBody:   `{"data": null, "message":"bad request error", "success":false}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			//Given
			s.SetupTest()
			s.mockService.On("FindAll").Return(test.mockReturn, test.mockError)

			req := httptest.NewRequest(http.MethodGet, "/buyers", nil)
			req.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()

			s.handler.GetAll()(recorder, req)

			expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
			require.Equal(t, test.expectedStatus, recorder.Code)
			require.Equal(t, expectedHeaders, recorder.Header())
			require.JSONEq(t, test.expectedBody, recorder.Body.String())
		})
	}

}

func (s *THandlerBuyerSuite) TestBuyerHandler_GetById() {
	t := s.T()

	type TestTableBuyerGetById struct {
		name           string
		param          string
		mockReturn     mod.Buyer
		mockError      error
		requireMock    bool
		expectedStatus int
		expectedBody   string
	}

	buyer := mod.Buyer{
		ID:           1,
		CardNumberID: "1234567890123456",
		FirstName:    "Juan",
		LastName:     "Pérez",
	}

	testsTable := []TestTableBuyerGetById{
		{
			name:           "Case 1: Success - Get by id",
			mockReturn:     buyer,
			mockError:      nil,
			param:          "1",
			expectedBody:   `{"success":true,"message":"Buyer obtenido con exito","data":{"id":1,"card_number_id":"1234567890123456","first_name":"Juan","last_name":"Pérez"}}`,
			expectedStatus: http.StatusOK,
			requireMock:    true,
		},
		{
			name:           "Case 2: Fail - Get by id",
			mockReturn:     mod.Buyer{},
			mockError:      e.ErrRequestIdMustBeInt,
			param:          "NOT_INT",
			expectedBody:   `{"data": null, "message":"handler: id must be an integer", "success":false}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Case 3: Fail - Buyer not found",
			mockReturn:     mod.Buyer{},
			mockError:      e.ErrBuyerRepositoryNotFound,
			param:          "99",
			expectedBody:   `{"data": null, "message":"repository: buyer not found", "success":false}`,
			expectedStatus: http.StatusNotFound,
			requireMock:    true,
		},
		{
			name:           "Case 4: Fail - Bad request",
			mockReturn:     mod.Buyer{},
			mockError:      errors.New("bad request error"),
			param:          "1",
			expectedBody:   `{"success":false,"message":"bad request error","data":null}`,
			expectedStatus: http.StatusBadRequest,
			requireMock:    true,
		},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			//Given
			s.SetupTest()
			if test.requireMock {
				s.mockService.On("FindByID", mock.AnythingOfType("int")).Return(test.mockReturn, test.mockError)
			}

			req := httptest.NewRequest(http.MethodGet, "/buyers/"+test.param, nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))
			req.Header.Set("Content-Type", "application/json")
			chiCtx := chi.RouteContext(req.Context())
			chiCtx.URLParams.Add("id", test.param)

			recorder := httptest.NewRecorder()

			s.handler.GetByID()(recorder, req)

			//Then
			expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
			require.Equal(t, test.expectedStatus, recorder.Code)
			require.Equal(t, expectedHeaders, recorder.Header())
			require.JSONEq(t, test.expectedBody, recorder.Body.String())

			s.mockService.AssertExpectations(s.T())

		})
	}

}

func (s *THandlerBuyerSuite) TestBuyerHandler_GetReport() {
	t := s.T()

	type TestTableBuyerGetReport struct {
		name           string
		param          string
		mockReturn     []mod.BuyerReportPO
		mockError      error
		requireMock    bool
		expectedStatus int
		expectedBody   string
	}

	report := []mod.BuyerReportPO{
		{
			Buyer: mod.Buyer{
				ID:           1,
				CardNumberID: "1234567890123456",
				FirstName:    "Juan",
				LastName:     "Pérez",
			},
			PurchaseOrderCount: 10,
		},
		{
			Buyer: mod.Buyer{
				ID:           2,
				CardNumberID: "1234412",
				FirstName:    "Michael",
				LastName:     "Cordoba",
			},
			PurchaseOrderCount: 10,
		},
	}

	testsTable := []TestTableBuyerGetReport{
		{
			name:           "Case 1: Success - Get report",
			mockReturn:     report,
			mockError:      nil,
			expectedBody:   `{"success":true,"message":"Reporte generado exitosamente","data":[{"id":1,"card_number_id":"1234567890123456","first_name":"Juan","last_name":"Pérez","purchase_orders_count":10},{"id":2,"card_number_id":"1234412","first_name":"Michael","last_name":"Cordoba","purchase_orders_count":10}]}`,
			expectedStatus: http.StatusOK,
			requireMock:    true,
		},
		{
			name:           "Case 2: Fail - Get by id",
			mockReturn:     nil,
			mockError:      e.ErrRequestIdMustBeInt,
			param:          "l",
			expectedBody:   `{"success":false,"message":"handler: id must be an integer","data":null}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Case 3: Fail - Buyer not found",
			mockReturn:     nil,
			mockError:      e.ErrBuyerRepositoryNotFound,
			param:          "10",
			expectedBody:   `{"success":false,"message":"repository: buyer not found","data":null}`,
			expectedStatus: http.StatusNotFound,
			requireMock:    true,
		},
		{
			name:           "Case 4: Fail - Bad request",
			mockReturn:     nil,
			mockError:      e.ErrRequestInternalServer,
			param:          "10",
			expectedBody:   `{"success":false,"message":"handler: internal server error","data":null}`,
			expectedStatus: http.StatusInternalServerError,
			requireMock:    true,
		},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			//Given
			s.SetupTest()
			if test.requireMock {
				s.mockService.On("GetPurchaseOrderReport", mock.AnythingOfType("*int")).
					Return(test.mockReturn, test.mockError)
			}

			query := ""
			if test.param != "" {
				query = "?id=" + test.param
			}
			req := httptest.NewRequest(http.MethodGet, "/buyers/reportPurchaseOrders"+query, nil)
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()

			s.handler.GetReport()(recorder, req)

			//Then
			expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
			require.Equal(t, test.expectedStatus, recorder.Code)
			require.Equal(t, expectedHeaders, recorder.Header())
			require.JSONEq(t, test.expectedBody, recorder.Body.String())

			s.mockService.AssertExpectations(s.T())

		})
	}

}

func (s *THandlerBuyerSuite) TestBuyerHandler_Create() {
	t := s.T()

	type TestTableBuyerCreate struct {
		name           string
		requestBody    string
		mockError      error
		requireMock    bool
		expectedStatus int
		expectedBody   string
		funcRun        func(args mock.Arguments)
	}

	testsTable := []TestTableBuyerCreate{
		{
			name:      "Case 1: Success - Create",
			mockError: nil,
			requestBody: `{
				"card_number_id": "1234567890123456",
				"first_name": "Juan",
				"last_name": "Perez"
				}`,
			expectedBody:   `{"success":true,"message":"Buyer creado exitosamente","data":{"id":1,"card_number_id":"1234567890123456","first_name":"Juan","last_name":"Perez"}}`,
			expectedStatus: http.StatusCreated,
			requireMock:    true,
			funcRun: func(args mock.Arguments) {
				b := args.Get(0).(*mod.Buyer)
				b.ID = 1
			},
		},
		{
			name:           "Case 2: Fail - Bad request",
			mockError:      e.ErrRequestFailedBody,
			expectedBody:   `{"success":false,"message":"handler: failed to read body","data":null}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:      "Case 3: Fail - Failed validation",
			mockError: e.ErrRequestWrongBody,
			requestBody: `{
				"card_number_id": "",
				"first_name": "Juan",
				"last_name": "Carlos"
				}`,
			expectedBody:   `{"success":false,"message":"handler: body does not meet requirements: CardNumberID failed on min validation","data":null}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:      "Case 4: Fail - Card duplicated",
			mockError: e.ErrBuyerRepositoryCardDuplicated,
			requestBody: `{
				"card_number_id": "12",
				"first_name": "Juan",
				"last_name": "Perez"
				}`,
			expectedBody:   `{"success":false,"message":"repository: Card id duplicated","data":null}`,
			expectedStatus: http.StatusConflict,
			requireMock:    true,
			funcRun:        func(args mock.Arguments) {},
		},
		{
			name:      "Case 4: Fail - Card duplicated",
			mockError: errors.New("bad request"),
			requestBody: `{
				"card_number_id": "12",
				"first_name": "Juan",
				"last_name": "Perez"
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
				s.mockService.On("Save", mock.AnythingOfType("*models.Buyer")).Run(func(args mock.Arguments) {
					test.funcRun(args)
				}).Return(test.mockError)
			}

			req := httptest.NewRequest(http.MethodPost, "/buyers", bytes.NewBufferString(test.requestBody))
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

func (s *THandlerBuyerSuite) TestBuyerHandler_Update() {
	t := s.T()

	type TestTableBuyerUpdate struct {
		name              string
		requestBody       string
		param             string
		mockGet           mod.Buyer
		mockGetError      error
		mockUpdateError   error
		requireMockGet    bool
		requireMockUpdate bool
		expectedStatus    int
		expectedBody      string
		funcRun           func(args mock.Arguments)
	}

	buyer := mod.Buyer{
		ID:           1,
		CardNumberID: "1234567890123456",
		FirstName:    "Juan",
		LastName:     "Pérez",
	}

	testTables := []TestTableBuyerUpdate{
		{
			name:              "Case 1: Success - Update",
			requireMockGet:    true,
			requireMockUpdate: true,
			mockGet:           buyer,
			funcRun: func(args mock.Arguments) {
				b := args.Get(0).(*mod.Buyer)
				b.ID = 1
			},
			requestBody: `{
				"card_number_id": "1234567890123456",
				"first_name": "Juan",
				"last_name": "Perez"
				}`,
			param:          "1",
			expectedBody:   `{"success":true,"message":"Buyer actualizado exitosamente","data":{"id":1,"card_number_id":"1234567890123456","first_name":"Juan","last_name":"Perez"}}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Case 2: Fail - Id integer",
			param:          "NOT_INT",
			mockGetError:   e.ErrRequestIdMustBeInt,
			expectedBody:   `{"success":false,"message":"handler: id must be an integer","data":null}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:              "Case 3: Fail - Buyer not found",
			requireMockGet:    true,
			requireMockUpdate: false,
			mockGet:           buyer,
			param:             "1",
			mockGetError:      e.ErrBuyerRepositoryNotFound,
			expectedBody:      `{"success":false,"message":"repository: buyer not found","data":null}`,
			expectedStatus:    http.StatusNotFound,
		},
		{
			name:              "Case 4: Fail - Bad request",
			requireMockGet:    true,
			requireMockUpdate: false,
			mockGet:           buyer,
			requestBody: `{
				"card_number_id": "1234567890123456",
				"first_name": "Juan",
				"last_name": "Perez"
				}`,
			param:          "1",
			mockGetError:   errors.New("bad request"),
			expectedBody:   `{"success":false,"message":"bad request","data":null}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:              "Case 5: Fail - Failed body",
			requireMockGet:    true,
			requireMockUpdate: false,
			mockGet:           buyer,
			requestBody: `
				"card_number_id": "1234567890123456",
				"first_name": "Juan",
				"last_name": "Perez",
				`,
			param:          "1",
			expectedBody:   `{"success":false,"message":"handler: failed to read body","data":null}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:              "Case 6: Fail - Validation failed",
			requireMockGet:    true,
			requireMockUpdate: false,
			mockGet:           buyer,
			requestBody: `{
				"last_name": ""
				}`,
			param:          "1",
			expectedBody:   `{"success":false,"message":"handler: body does not meet requirements: LastName failed on min validation","data":null}`,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:              "Case 7: Fail - Card duplicated",
			requireMockGet:    true,
			requireMockUpdate: true,
			funcRun:           func(args mock.Arguments) {},
			mockGet:           buyer,
			requestBody: `{
				"card_number_id": "1231231",
				"first_name": "Juan",
				"last_name": "Perez"
				}`,
			param:           "1",
			mockUpdateError: e.ErrBuyerRepositoryCardDuplicated,
			expectedBody:    `{"success":false,"message":"repository: Card id duplicated","data":null}`,
			expectedStatus:  http.StatusConflict,
		},
		{
			name:              "Case 8: Fail - Final update Bad request",
			requireMockGet:    true,
			requireMockUpdate: true,
			funcRun:           func(args mock.Arguments) {},
			mockGet:           buyer,
			requestBody: `{
				"card_number_id": "1231231",
				"first_name": "Juan",
				"last_name": "Perez"
				}`,
			param:           "1",
			mockUpdateError: errors.New("bad request"),
			expectedBody:    `{"success":false,"message":"bad request","data":null}`,
			expectedStatus:  http.StatusBadRequest,
		},
	}

	for _, test := range testTables {
		t.Run(test.name, func(t *testing.T) {
			//Given

			s.SetupTest()
			if test.requireMockGet {
				s.mockService.On("FindByID", mock.AnythingOfType("int")).Return(test.mockGet, test.mockGetError)
			}
			if test.requireMockUpdate {
				s.mockService.On("Update", mock.AnythingOfType("*models.Buyer")).Run(func(args mock.Arguments) {
					test.funcRun(args)
				}).Return(test.mockUpdateError)
			}

			req := httptest.NewRequest(http.MethodPatch, "/buyers/"+test.param, bytes.NewBufferString(test.requestBody))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))
			chiCtx := chi.RouteContext(req.Context())
			chiCtx.URLParams.Add("id", test.param)
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()

			s.handler.Update()(recorder, req)

			//Then
			expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
			require.Equal(t, test.expectedStatus, recorder.Code)
			require.Equal(t, expectedHeaders, recorder.Header())
			require.JSONEq(t, test.expectedBody, recorder.Body.String())

			s.mockService.AssertExpectations(s.T())
		})
	}

}

func (s *THandlerBuyerSuite) TestBuyerHandler_Delete() {
	t := s.T()

	type TestTableBuyerDelete struct {
		name           string
		param          string
		mockReturn     mod.Buyer
		mockError      error
		expectedStatus int
		expectedBody   string
		requireMock    bool
	}

	testsTable := []TestTableBuyerDelete{
		{
			name:           "Case 1: Success - Delete",
			param:          "1",
			expectedStatus: http.StatusNoContent,
			expectedBody:   `{"success":true,"message":"","data":null}`,
			requireMock:    true,
		},
		{
			name:           "Case 2: Fail - Id integer",
			param:          "NOT_INT",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"handler: id must be an integer","data": null}`,
		},
		{
			name:           "Case 3: Fail - Buyer not found",
			param:          "1",
			expectedStatus: http.StatusNotFound,
			mockError:      e.ErrBuyerRepositoryNotFound,
			expectedBody:   `{"success":false,"message":"repository: buyer not found","data": null}`,
			requireMock:    true,
		},
		{
			name:           "Case 4: Fail - bad request",
			param:          "1",
			mockError:      errors.New("bad request"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"success":false,"message":"bad request","data": null}`,
			requireMock:    true,
		},
	}

	for _, test := range testsTable {
		t.Run(test.name, func(t *testing.T) {
			s.SetupTest()
			if test.requireMock {
				s.mockService.On("Delete", mock.AnythingOfType("int")).Return(test.mockError)
			}

			req := httptest.NewRequest(http.MethodDelete, "/buyers/"+test.param, nil)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))
			chiCtx := chi.RouteContext(req.Context())
			chiCtx.URLParams.Add("id", test.param)
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()

			s.handler.Delete()(recorder, req)

			//Then
			expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
			require.Equal(t, test.expectedStatus, recorder.Code)
			require.Equal(t, expectedHeaders, recorder.Header())
			require.JSONEq(t, test.expectedBody, recorder.Body.String())

			s.mockService.AssertExpectations(s.T())
		})
	}

}

func TestBuyerHanlderSuite(t *testing.T) {
	suite.Run(t, new(THandlerBuyerSuite))
}
