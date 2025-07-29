package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/internal/service"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/tests"
	"github.com/stretchr/testify/require"
)

func TestCarryHandler_Create(t *testing.T) {
	mock := tests.NewCarryMock()
	serv := service.NewCarryService(mock)
	handler := NewCarryHandler(serv)

	validCarry := models.Carry{
		ID:          0,
		CID:         "CID#1",
		CompanyName: "some name",
		Address:     "corrientes 800",
		Telephone:   "45674567",
		LocalityID:  6700,
	}

	t.Run("create_ok", func(t *testing.T) {
		mock.On("Save", &validCarry).Return(nil)
		mock.On("ExistsLocality", 6700).Return(true, nil)
		mock.On("ExistsCID", "CID#1").Return(false, nil)

		body := strings.NewReader(`{

  "cid": "CID#1",
  "company_name": "some name",
  "address": "corrientes 800",
  "telephone": "45674567",
  "locality_id": 6700

		}`)

		req := httptest.NewRequest(http.MethodPost, "/carries", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expected := `{
			
  
    "success": true,
    "message": "Created",
    "data": {
        "id": 0,
        "cid": "CID#1",
        "company_name": "some name",
        "address": "corrientes 800",
        "telephone": "45674567",
        "locality_id": 6700
    }

		}`

		handler.Create().ServeHTTP(w, req)

		require.Equal(t, http.StatusCreated, w.Code)
		require.JSONEq(t, expected, w.Body.String())
	})
	t.Run("create_missing_required_field", func(t *testing.T) {
		body := strings.NewReader(`{
        "company_name": "some name",
        "address": "corrientes 800",
        "telephone": "45674567",
        "locality_id": 6700
    }`)

		req := httptest.NewRequest(http.MethodPost, "/carries", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Create().ServeHTTP(w, req)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
		require.Contains(t, w.Body.String(), "Campos inválidos")
	})
	t.Run("create_invalid_json", func(t *testing.T) {
		body := strings.NewReader(`{invalid_json}`) // JSON malformado
		req := httptest.NewRequest(http.MethodPost, "/carries", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Create().ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Contains(t, w.Body.String(), e.ErrRequestWrongBody.Error())
	})

	t.Run("create_invalid_telephone", func(t *testing.T) {
		body := strings.NewReader(`{
			"cid": "CID123",
			"company_name": "Test",
			"address": "Test",
			"telephone": "abc",
			"locality_id": 1
		}`)

		req := httptest.NewRequest(http.MethodPost, "/carries", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.Create().ServeHTTP(w, req)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
		require.Contains(t, w.Body.String(), "Telephone' failed on the 'numeric' tag")
	})

	t.Run("create_invalid_locality_id", func(t *testing.T) {
		body := strings.NewReader(`{ "cid": "CID#1",
  "company_name": "some name",
  "address": "corrientes 800",
  "telephone": "45674567",
  "locality_id": 0
		}`)

		req := httptest.NewRequest(http.MethodPost, "/carries", body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler.Create().ServeHTTP(w, req)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
		require.Contains(t, w.Body.String(), "Campos inválidos: Key: 'Carry.LocalityID'")
	})
}

func TestCarryHandler_GetReportByLocality(t *testing.T) {
	mock := tests.NewCarryMock()
	serv := service.NewCarryService(mock)
	handler := NewCarryHandler(serv)

	report := []models.LocalityCarryReport{
		{
			LocalityID:   1,
			LocalityName: "Test Locality",
			CarriesCount: 5,
		},
	}

	t.Run("get_report_with_id", func(t *testing.T) {
		mock.On("GetReportByLocality", 1).Return(report, nil)

		req := httptest.NewRequest(http.MethodGet, "/localities/reportCarries?id=1", nil)
		w := httptest.NewRecorder()

		handler.GetReportByLocality().ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("get_report_all", func(t *testing.T) {
		mock.On("GetReportByLocalityAll").Return(report, nil)

		req := httptest.NewRequest(http.MethodGet, "/localities/reportCarries", nil)
		w := httptest.NewRecorder()

		handler.GetReportByLocality().ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("invalid_locality_id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/localities/reportCarries?id=invalid", nil)
		w := httptest.NewRecorder()

		handler.GetReportByLocality().ServeHTTP(w, req)

		require.Equal(t, http.StatusBadRequest, w.Code)
		require.Contains(t, w.Body.String(), "id must be an integer")
	})

	t.Run("locality_not_found", func(t *testing.T) {
		mock.On("GetReportByLocality", 999).Return([]models.LocalityCarryReport{}, e.ErrLocalityRepositoryNotFound)

		req := httptest.NewRequest(http.MethodGet, "/localities/reportCarries?id=999", nil)
		w := httptest.NewRecorder()

		handler.GetReportByLocality().ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
		require.Contains(t, w.Body.String(), "locality not found")
	})

	t.Run("get_report_all_success", func(t *testing.T) {
		// Mock del servicio: Retorna datos y ningún error
		mock.On("GetReportByLocalityAll").Return(report, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/localities/reportCarries", nil)
		w := httptest.NewRecorder()

		handler.GetReportByLocality().ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		require.Contains(t, w.Body.String(), "success")
		require.Contains(t, w.Body.String(), "Test Locality") // Verifica datos del reporte
	})

}
