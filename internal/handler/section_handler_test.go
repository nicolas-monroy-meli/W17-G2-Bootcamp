package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/tests/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper for JSON responses
func toJSON(t *testing.T, v interface{}) string {
	d, err := json.Marshal(v)
	require.NoError(t, err)
	return string(d)
}

func TestSectionHandler_GetAll(t *testing.T) {
	sectionSlice := []mod.Section{
		{ID: 1, SectionNumber: 1, CurrentTemperature: 2, MinimumTemperature: 0, CurrentCapacity: 12, MinimumCapacity: 3, MaximumCapacity: 15, WarehouseID: 1, ProductTypeID: 1},
		{ID: 2, SectionNumber: 2, CurrentTemperature: 3, MinimumTemperature: 0, CurrentCapacity: 14, MinimumCapacity: 13, MaximumCapacity: 15, WarehouseID: 2, ProductTypeID: 2},
	}
	testsSlice := []struct {
		name            string
		mockFindAll     func() ([]mod.Section, error)
		expectedStatus  int
		expectedContent string
	}{
		{
			name: "Get all items",
			mockFindAll: func() ([]mod.Section, error) {
				return sectionSlice, nil
			},
			expectedStatus:  http.StatusOK,
			expectedContent: toJSON(t, sectionSlice),
		},
		{
			name: "repo error",
			mockFindAll: func() ([]mod.Section, error) {
				return nil, e.ErrEmptyDB
			},
			expectedStatus:  http.StatusNotFound,
			expectedContent: `{"success":false,"message":"repository: empty DB","data":null}`,
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockSectionService{
				MockFindAll: tc.mockFindAll,
			}
			handler := NewSectionHandler(svc)
			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			handler.GetAll().ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedContent)
		})
	}
}

func TestSectionHandler_GetByID(t *testing.T) {
	item := mod.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 2,
		MinimumTemperature: 1,
		CurrentCapacity:    33,
		MinimumCapacity:    1,
		MaximumCapacity:    333,
		WarehouseID:        1,
		ProductTypeID:      1,
	}
	testsSlice := []struct {
		name            string
		id              string
		mockFindByID    func(int) (mod.Section, error)
		expectedStatus  int
		expectedContent string
	}{
		{
			name: "found",
			id:   "1",
			mockFindByID: func(id int) (mod.Section, error) {
				return item, nil
			},
			expectedStatus:  http.StatusOK,
			expectedContent: toJSON(t, item),
		},
		{
			name: "not found",
			id:   "99",
			mockFindByID: func(id int) (mod.Section, error) {
				return mod.Section{}, e.ErrSectionRepositoryNotFound
			},
			expectedStatus:  http.StatusNotFound,
			expectedContent: "not found",
		},
		{
			name:            "bad id",
			id:              "foo",
			mockFindByID:    func(id int) (mod.Section, error) { return mod.Section{}, nil },
			expectedStatus:  http.StatusBadRequest,
			expectedContent: "invalid",
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockSectionService{
				MockFindByID: tc.mockFindByID,
			}
			handler := NewSectionHandler(svc)
			req := httptest.NewRequest("GET", "/"+tc.id, nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()
			handler.GetByID().ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedContent)
		})
	}
}

func TestSectionHandler_Create(t *testing.T) {
	valid := mod.Section{SectionNumber: 5, CurrentTemperature: 2, MinimumTemperature: 0, CurrentCapacity: 6, MinimumCapacity: 3, MaximumCapacity: 10, WarehouseID: 1, ProductTypeID: 1}
	invalid := mod.Section{SectionNumber: 5, CurrentTemperature: -2, MinimumTemperature: 0, CurrentCapacity: -6, MinimumCapacity: 3, MaximumCapacity: 10, WarehouseID: 1, ProductTypeID: 1}
	badJSON := `{"section_number":5,` // Malformed

	testsSlice := []struct {
		name           string
		body           string
		mockSave       func(*mod.Section) error
		expectedStatus int
		expectedString string
	}{
		{
			name:           "create success",
			body:           toJSON(t, valid),
			mockSave:       func(s *mod.Section) error { return nil },
			expectedStatus: http.StatusCreated,
			expectedString: "created",
		},
		{
			name: "invalid update body problems",
			body: toJSON(t, invalid),
			mockSave: func(s *mod.Section) error {
				return errors.New("not reached")
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedString: "must be greater than or equal to",
		},
		{
			name:           "bad json input",
			body:           badJSON,
			mockSave:       func(s *mod.Section) error { return nil },
			expectedStatus: http.StatusBadRequest,
			expectedString: "failed to read body",
		},
		{
			name:           "repo conflict",
			body:           toJSON(t, valid),
			mockSave:       func(s *mod.Section) error { return e.ErrSectionRepositoryDuplicated },
			expectedStatus: http.StatusConflict,
			expectedString: "section already exists",
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockSectionService{MockSave: tc.mockSave}
			handler := NewSectionHandler(svc)
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(tc.body))
			rr := httptest.NewRecorder()
			handler.Create().ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedString)
		})
	}
}

func TestSectionHandler_Update(t *testing.T) {
	validUpdate := mod.Section{ID: 1, SectionNumber: 2, CurrentTemperature: 2, MinimumTemperature: 1, CurrentCapacity: 5, MinimumCapacity: 1, MaximumCapacity: 10, WarehouseID: 1, ProductTypeID: 1}
	invalidUpdate := mod.Section{SectionNumber: 1, CurrentTemperature: -1}
	badUpdate := `{"section_number":`

	testsSlice := []struct {
		name           string
		id             string
		body           string
		mockUpdate     func(int, map[string]interface{}) (*mod.Section, error)
		expectedStatus int
		expectedString string
	}{
		{
			name: "update success",
			id:   "1",
			body: toJSON(t, validUpdate),
			mockUpdate: func(id int, fields map[string]interface{}) (*mod.Section, error) {
				return &validUpdate, nil
			},
			expectedStatus: http.StatusOK,
			expectedString: "2", // SectionNumber after update
		},
		{
			name: "invalid update body problems",
			id:   "1",
			body: toJSON(t, invalidUpdate),
			mockUpdate: func(i int, m map[string]interface{}) (*mod.Section, error) {
				return nil, errors.New("not reached")
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedString: "must be greater than or equal to",
		},
		{
			name: "invalid update nothing to update",
			id:   "1",
			body: toJSON(t, validUpdate),
			mockUpdate: func(i int, m map[string]interface{}) (*mod.Section, error) {
				return nil, e.ErrNoRowsAffected
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedString: "handler: nothing to update",
		},
		{
			name: "bad id",
			id:   "foo",
			body: toJSON(t, validUpdate),
			mockUpdate: func(id int, fields map[string]interface{}) (*mod.Section, error) {
				return nil, e.ErrQueryError
			},
			expectedStatus: http.StatusBadRequest,
			expectedString: "id must be an integer",
		},
		{
			name: "bad json",
			id:   "1",
			body: badUpdate,
			mockUpdate: func(id int, fields map[string]interface{}) (*mod.Section, error) {
				return nil, errors.New("fail")
			},
			expectedStatus: http.StatusBadRequest,
			expectedString: "body",
		},
		{
			name: "update not found",
			id:   "100",
			body: toJSON(t, validUpdate),
			mockUpdate: func(id int, fields map[string]interface{}) (*mod.Section, error) {
				return nil, e.ErrSectionRepositoryNotFound
			},
			expectedStatus: http.StatusNotFound,
			expectedString: "not found",
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockSectionService{MockUpdate: tc.mockUpdate}
			handler := NewSectionHandler(svc)
			req := httptest.NewRequest("PATCH", "/"+tc.id, bytes.NewBufferString(tc.body))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()
			handler.Update().ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedString)
		})
	}
}

func TestSectionHandler_ReportProducts(t *testing.T) {
	oneRes := []mod.ReportProductsResponse{{SectionId: 1, ProductsCount: 3}}
	multipleRes := []mod.ReportProductsResponse{{
		SectionId:     1,
		ProductsCount: 5,
	},
		{
			SectionId:     2,
			ProductsCount: 10,
		}}
	testsSlice := []struct {
		name           string
		ids            string
		mockReport     func([]int) ([]mod.ReportProductsResponse, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "success-one id",
			ids:  "1",
			mockReport: func(ids []int) ([]mod.ReportProductsResponse, error) {
				return oneRes, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   toJSON(t, oneRes),
		},
		{
			name: "success-multiple-ids",
			ids:  "1,2",
			mockReport: func(ints []int) ([]mod.ReportProductsResponse, error) {
				return multipleRes, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   toJSON(t, multipleRes),
		},
		{
			name: "success-no-args",
			mockReport: func(ints []int) ([]mod.ReportProductsResponse, error) {
				return multipleRes, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   toJSON(t, multipleRes),
		},
		{
			name: "error in ID parsing",
			ids:  "test",
			mockReport: func(ints []int) ([]mod.ReportProductsResponse, error) {
				return nil, e.ErrRequestIdMustBeInt
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "id must be an integer",
		},
		{
			name: "section not found",
			ids:  "500",
			mockReport: func(ids []int) ([]mod.ReportProductsResponse, error) {
				return nil, errors.New("not found")
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "not found",
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockSectionService{MockReportProducts: tc.mockReport}
			handler := NewSectionHandler(svc)
			req := httptest.NewRequest("GET", "/?ids="+tc.ids, nil)
			rr := httptest.NewRecorder()
			handler.ReportProducts().ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedBody)
		})
	}
}

func TestSectionHandler_Delete(t *testing.T) {
	testsSlice := []struct {
		name           string
		id             string
		mockDelete     func(int) error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "success",
			id:             "1",
			mockDelete:     func(id int) error { return nil },
			expectedStatus: http.StatusNoContent,
			expectedBody:   "",
		},
		{
			name:           "not found",
			id:             "4",
			mockDelete:     func(id int) error { return errors.New("not found") },
			expectedStatus: http.StatusNotFound,
			expectedBody:   "not found",
		},
		{
			name:           "bad id",
			id:             "foo",
			mockDelete:     func(id int) error { return e.ErrRequestIdMustBeInt },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "handler: id must be an integer",
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockSectionService{MockDelete: tc.mockDelete}
			handler := NewSectionHandler(svc)
			req := httptest.NewRequest("DELETE", "/"+tc.id, nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()
			handler.Delete().ServeHTTP(rr, req)
			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedBody)
		})
	}
}
