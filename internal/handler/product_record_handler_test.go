package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/assert"
)

// MockProductRecordService implementa ProductRecordService para tests
// Puedes agregar campos para controlar el comportamiento en cada test

type MockProductRecordService struct {
	FindProductByIDFunc      func(id int) (models.Product, error)
	FindAllByProductIDPRFunc func(id int) (map[int]models.ProductRecord, error)
	FindAllPRFunc            func() (map[int]models.ProductRecord, error)
	SavePRFunc               func(pr *models.ProductRecord) error
}

func (m *MockProductRecordService) FindProductByID(id int) (models.Product, error) {
	return m.FindProductByIDFunc(id)
}
func (m *MockProductRecordService) FindAllByProductIDPR(id int) (map[int]models.ProductRecord, error) {
	return m.FindAllByProductIDPRFunc(id)
}
func (m *MockProductRecordService) FindAllPR() (map[int]models.ProductRecord, error) {
	return m.FindAllPRFunc()
}
func (m *MockProductRecordService) SavePR(pr *models.ProductRecord) error {
	return m.SavePRFunc(pr)
}

// Tests para GetRecords
func TestProductRecordHandler_GetRecords_All(t *testing.T) {
	mock := &MockProductRecordService{
		FindAllPRFunc: func() (map[int]models.ProductRecord, error) {
			return map[int]models.ProductRecord{1: {ID: 1}}, nil
		},
	}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductRecordHandler_GetRecords_AllNotFound(t *testing.T) {
	mock := &MockProductRecordService{
		FindAllPRFunc: func() (map[int]models.ProductRecord, error) {
			return nil, e.ErrProductRecordRepositoryNotFound
		},
	}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestProductRecordHandler_GetRecords_AllOtherError(t *testing.T) {
	mock := &MockProductRecordService{
		FindAllPRFunc: func() (map[int]models.ProductRecord, error) {
			return nil, errors.New("error interno")
		},
	}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error interno")
}

func TestProductRecordHandler_GetRecords_ByProductID_Success(t *testing.T) {
	mock := &MockProductRecordService{
		FindProductByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		FindAllByProductIDPRFunc: func(id int) (map[int]models.ProductRecord, error) {
			return map[int]models.ProductRecord{1: {ID: 1}}, nil
		},
	}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records?id=1", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductRecordHandler_GetRecords_ByProductID_InvalidID(t *testing.T) {
	mock := &MockProductRecordService{}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records?id=abc", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestIdMustBeInt.Error())
}

func TestProductRecordHandler_GetRecords_ByProductID_NotFound(t *testing.T) {
	mock := &MockProductRecordService{
		FindProductByIDFunc: func(id int) (models.Product, error) {
			return models.Product{}, e.ErrProductRepositoryNotFound
		},
	}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records?id=99", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrProductRepositoryNotFound.Error())
}

func TestProductRecordHandler_GetRecords_ByProductID_RecordsNotFound(t *testing.T) {
	mock := &MockProductRecordService{
		FindProductByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		FindAllByProductIDPRFunc: func(id int) (map[int]models.ProductRecord, error) {
			return map[int]models.ProductRecord{}, nil
		},
	}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records?id=1", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrProductRecordRepositoryNotFound.Error())
}

func TestProductRecordHandler_GetRecords_ByProductID_InternalError(t *testing.T) {
	mock := &MockProductRecordService{
		FindProductByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		FindAllByProductIDPRFunc: func(id int) (map[int]models.ProductRecord, error) {
			return nil, errors.New("error interno")
		},
	}
	h := NewProductRecordHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/product-records?id=1", nil)
	w := httptest.NewRecorder()

	h.GetRecords()(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error interno")
}

// Tests para CreateRecord
func TestProductRecordHandler_CreateRecord_Success(t *testing.T) {
	mock := &MockProductRecordService{
		SavePRFunc: func(pr *models.ProductRecord) error { return nil },
	}
	h := NewProductRecordHandler(mock)
	record := models.ProductRecord{ID: 1, ProductID: 1, LastUpdateDate: "2025-07-28", PurchasePrice: 100.0, SalePrice: 120.0}
	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/product-records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.CreateRecord()(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProductRecordHandler_CreateRecord_BadJSON(t *testing.T) {
	mock := &MockProductRecordService{
		SavePRFunc: func(pr *models.ProductRecord) error { return nil },
	}
	h := NewProductRecordHandler(mock)
	body := []byte(`{mal json}`)
	req := httptest.NewRequest(http.MethodPost, "/product-records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.CreateRecord()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestNoBody.Error())
}

func TestProductRecordHandler_CreateRecord_ValidationError(t *testing.T) {
	mock := &MockProductRecordService{
		SavePRFunc: func(pr *models.ProductRecord) error { return nil },
	}
	h := NewProductRecordHandler(mock)
	record := models.ProductRecord{ID: 1} // Faltan campos requeridos
	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/product-records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.CreateRecord()(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestWrongBody.Error())
}

func TestProductRecordHandler_CreateRecord_ProductNotFound(t *testing.T) {
	mock := &MockProductRecordService{
		SavePRFunc: func(pr *models.ProductRecord) error { return e.ErrProductRepositoryNotFound },
	}
	h := NewProductRecordHandler(mock)
	record := models.ProductRecord{ID: 1, ProductID: 1, LastUpdateDate: "2025-07-28", PurchasePrice: 100.0, SalePrice: 120.0}
	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/product-records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.CreateRecord()(w, req)
	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrProductRepositoryNotFound.Error())
}

func TestProductRecordHandler_CreateRecord_InternalError(t *testing.T) {
	mock := &MockProductRecordService{
		SavePRFunc: func(pr *models.ProductRecord) error { return errors.New("error interno") },
	}
	h := NewProductRecordHandler(mock)
	record := models.ProductRecord{ID: 1, ProductID: 1, LastUpdateDate: "2025-07-28", PurchasePrice: 100.0, SalePrice: 120.0}
	body, _ := json.Marshal(record)
	req := httptest.NewRequest(http.MethodPost, "/product-records", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.CreateRecord()(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error interno")
}
