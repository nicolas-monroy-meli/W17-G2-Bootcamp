package handler

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/tests"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSectionHandler_GetAll(t *testing.T) {
	sectionSlice := []mod.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 2,
			MinimumTemperature: 0,
			CurrentCapacity:    12,
			MinimumCapacity:    3,
			MaximumCapacity:    15,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      2,
			CurrentTemperature: 3,
			MinimumTemperature: 0,
			CurrentCapacity:    14,
			MinimumCapacity:    13,
			MaximumCapacity:    15,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
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
			expectedContent: ``,
		},
	}

	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &tests.MockSectionService{
				MockFindAll: tc.mockFindAll,
			}
			handler := NewSectionHandler(svc)

			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			handler.GetAll().ServeHTTP(rr, req)

			require.Equal(t, tc.expectedStatus, rr.Code)
			require.JSONEq(t, tc.expectedContent, rr.Body.String())
		})
	}
}

func TestSectionHandler_GetByID(t *testing.T) {
	/*testsSlice := []struct {
		name            string
		mockFindAll     func() ([]mod.ProductBatch, error)
		expectedStatus  int
		expectedContent string
	}{{}}*/
}

func TestSectionHandler_Create(t *testing.T) {
	/*testsSlice := []struct {
		name            string
		mockFindAll     func() (mod.ProductBatch, error)
		expectedStatus  int
		expectedContent string
	}{{}}*/
}

func TestBuyerHandler_Delete(t *testing.T) {
	/*testsSlice := []struct {
		name            string
		mockFindAll     func() error
		expectedStatus  int
		expectedContent string
	}{{}}*/
}

func TestSectionHandler_Update(t *testing.T) {
	/*	testsSlice := []struct {
		name            string
		mockFindAll     func() ([]mod.ProductBatch, error)
		expectedStatus  int
		expectedContent string
	}{{}}*/
}

func TestSectionHandler_ReportProducts(t *testing.T) {
	/*testsSlice := []struct {
		name            string
		mockFindAll     func() ([]mod.ProductBatch, error)
		expectedStatus  int
		expectedContent string
	}{{}}*/
}
