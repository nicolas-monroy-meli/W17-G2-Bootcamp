package handler

import (
	"bytes"
	"encoding/json"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/tests/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestProductBatchHandler_GetAll(t *testing.T) {
	testsSlice := []struct {
		name            string
		mockFindAll     func() ([]mod.ProductBatch, error)
		expectedStatus  int
		expectedContent string
	}{
		{
			name: "success",
			mockFindAll: func() ([]mod.ProductBatch, error) {
				return []mod.ProductBatch{
					{
						ID:                 1,
						BatchNumber:        1,
						CurrentQuantity:    200,
						InitialQuantity:    200,
						CurrentTemperature: 2,
						MinimumTemperature: -5,
						DueDate:            time.Date(2024, 07, 05, 17, 00, 00, 0, time.UTC),
						ManufacturingDate:  time.Date(2024, 06, 1, 0, 00, 00, 0, time.UTC),
						ManufacturingHour:  "08:00:00",
						ProductId:          1,
						SectionId:          1,
					},
					{
						ID:                 2,
						BatchNumber:        2,
						CurrentQuantity:    310,
						InitialQuantity:    310,
						CurrentTemperature: -2,
						MinimumTemperature: -6,
						DueDate:            time.Date(2024, 8, 01, 12, 00, 00, 0, time.UTC),
						ManufacturingDate:  time.Date(2024, 7, 1, 0, 00, 00, 0, time.UTC),
						ManufacturingHour:  "09:30:00",
						ProductId:          2,
						SectionId:          2,
					},
				}, nil
			},
			expectedStatus:  http.StatusOK,
			expectedContent: `{"success":true,"message":"handler: data retrieved successfully","data":[{"id":1,"batch_number":1,"current_quantity":200,"initial_quantity":200,"current_temperature":2,"minimum_temperature":-5,"due_date":"2024-07-05T17:00:00Z","manufacturing_date":"2024-06-01T00:00:00Z","manufacturing_hour":"08:00:00","product_id":1,"section_id":1},{"id":2,"batch_number":2,"current_quantity":310,"initial_quantity":310,"current_temperature":-2,"minimum_temperature":-6,"due_date":"2024-08-01T12:00:00Z","manufacturing_date":"2024-07-01T00:00:00Z","manufacturing_hour":"09:30:00","product_id":2,"section_id":2}]}`,
		},
		{
			name: "repo error",
			mockFindAll: func() ([]mod.ProductBatch, error) {
				return nil, e.ErrEmptyDB
			},
			expectedStatus:  http.StatusNotFound,
			expectedContent: `{"success":false,"message":"repository: empty DB","data":null}`,
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockProductBatchService{
				MockFindAll: tc.mockFindAll,
			}
			handler := NewProductBatchHandler(svc)

			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			handler.GetAll().ServeHTTP(rr, req)

			require.Equal(t, tc.expectedStatus, rr.Code)
			require.JSONEq(t, tc.expectedContent, rr.Body.String())
		})
	}
}

func TestProductBatchHandler_Create(t *testing.T) {
	validBatch := mod.ProductBatch{
		ID:                 1,
		BatchNumber:        1,
		CurrentQuantity:    200,
		InitialQuantity:    200,
		CurrentTemperature: 2,
		MinimumTemperature: -5,
		DueDate:            time.Date(2024, 07, 05, 17, 00, 00, 0, time.UTC),
		ManufacturingDate:  time.Date(2024, 06, 1, 0, 00, 00, 0, time.UTC),
		ManufacturingHour:  "08:00:00",
		ProductId:          1,
		SectionId:          1,
	}
	invalidBatch := mod.ProductBatch{
		ID:                 1,
		BatchNumber:        1,
		CurrentQuantity:    200,
		InitialQuantity:    300,
		CurrentTemperature: -6,
		MinimumTemperature: -5,
		DueDate:            time.Date(2024, 07, 05, 17, 00, 00, 0, time.UTC),
		ManufacturingDate:  time.Date(2024, 06, 1, 0, 00, 00, 0, time.UTC),
		ManufacturingHour:  "08:00:00",
		ProductId:          1,
		SectionId:          1,
	}
	invalidJSON := `{"BatchNumber":`
	toJSON := func(v interface{}) string {
		data, _ := json.Marshal(v)
		return string(data)
	}

	type testCase struct {
		name           string
		body           string
		mockSave       func(*mod.ProductBatch) error
		expectedStatus int
		expectedText   string
	}
	testsSlice := []testCase{
		{
			name: "success",
			body: toJSON(validBatch),
			mockSave: func(pb *mod.ProductBatch) error {
				return nil
			},
			expectedStatus: http.StatusCreated,
			expectedText:   `{"success":true,"message":"handler: section successfully created","data":{"id":1,"batch_number":1,"current_quantity":200,"initial_quantity":200,"current_temperature":2,"minimum_temperature":-5,"due_date":"2024-07-05T17:00:00Z","manufacturing_date":"2024-06-01T00:00:00Z","manufacturing_hour":"08:00:00","product_id":1,"section_id":1}}`,
		},
		{
			name:           "bad json",
			body:           invalidJSON,
			mockSave:       func(pb *mod.ProductBatch) error { return nil },
			expectedStatus: http.StatusBadRequest,
			expectedText:   `{"success":false,"message":"handler: failed to read body","data":null}`,
		},
		{
			name: "bad section",
			body: toJSON(invalidBatch),
			mockSave: func(batch *mod.ProductBatch) error {
				return nil
			},
			expectedStatus: http.StatusUnprocessableEntity,
			expectedText:   `{"success":false,"message":"CurrentQuantity must be greater than or equal to InitialQuantity, CurrentTemperature must be greater than or equal to MinimumTemperature, ","data":null}`,
		},
		{
			name:           "save error",
			body:           toJSON(validBatch),
			mockSave:       func(pb *mod.ProductBatch) error { return e.ErrSectionRepositoryDuplicated },
			expectedStatus: http.StatusConflict,
			expectedText:   `{"success":false,"message":"repository: section already exists","data":null}`,
		},
	}
	for _, tc := range testsSlice {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mock.MockProductBatchService{
				MockSave: tc.mockSave,
			}
			handler := NewProductBatchHandler(svc)

			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(tc.body))
			rr := httptest.NewRecorder()
			handler.Create().ServeHTTP(rr, req)

			require.Equal(t, tc.expectedStatus, rr.Code)
			require.Contains(t, rr.Body.String(), tc.expectedText)
		})
	}
}
