package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/assert"
)

// MockProductService implementa ProductService para tests
// Puedes agregar campos para controlar el comportamiento en cada test

type MockProductService struct {
	FindAllFunc  func() ([]models.Product, error)
	FindByIDFunc func(id int) (models.Product, error)
	SaveFunc     func(p *models.Product) error
	UpdateFunc   func(p *models.Product) error
	DeleteFunc   func(id int) error
}

func (m *MockProductService) FindAll() ([]models.Product, error) {
	return m.FindAllFunc()
}
func (m *MockProductService) FindByID(id int) (models.Product, error) {
	return m.FindByIDFunc(id)
}
func (m *MockProductService) Save(p *models.Product) error {
	return m.SaveFunc(p)
}
func (m *MockProductService) Update(p *models.Product) error {
	return m.UpdateFunc(p)
}
func (m *MockProductService) Delete(id int) error {
	return m.DeleteFunc(id)
}

func TestProductHandler_GetAll(t *testing.T) {
	mock := &MockProductService{
		FindAllFunc: func() ([]models.Product, error) {
			return []models.Product{{ID: 1, Description: "Test"}}, nil
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	h.GetAll()(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_GetAll_Error(t *testing.T) {
	mock := &MockProductService{
		FindAllFunc: func() ([]models.Product, error) {
			return nil, errors.New("error de prueba")
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()

	h.GetAll()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error de prueba")
}

func TestProductHandler_GetByID_Success(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id, Description: "Test"}, nil
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.GetByID()(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_GetByID_NotFound(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{}, e.ErrProductRepositoryNotFound
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/products/99", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "99")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.GetByID()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestProductHandler_GetByID_InvalidID(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{}, nil
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/products/abc", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.GetByID()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestIdMustBeInt.Error())
}

func TestProductHandler_GetByID_OtherError(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{}, errors.New("otro error")
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodGet, "/products/2", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "2")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.GetByID()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "otro error")
}

func TestProductHandler_Create_Success(t *testing.T) {
	mock := &MockProductService{
		SaveFunc: func(p *models.Product) error {
			return nil
		},
	}
	h := NewProductHandler(mock)
	product := models.Product{
		ID:             1,
		ProductCode:    "P001",
		Description:    "Producto de prueba",
		Height:         10.5,
		Length:         20.0,
		Width:          5.0,
		Weight:         1.2,
		ExpirationRate: 0.5,
		FreezingRate:   1.0,
		RecomFreezTemp: -18.0,
		ProductTypeID:  2,
		SellerID:       3,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Create()(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProductHandler_Create_Duplicated(t *testing.T) {
	mock := &MockProductService{
		SaveFunc: func(p *models.Product) error {
			return e.ErrProductRepositoryDuplicated
		},
	}
	h := NewProductHandler(mock)
	product := models.Product{
		ID:             1,
		ProductCode:    "P001",
		Description:    "Producto duplicado",
		Height:         10.5,
		Length:         20.0,
		Width:          5.0,
		Weight:         1.2,
		ExpirationRate: 0.5,
		FreezingRate:   1.0,
		RecomFreezTemp: -18.0,
		ProductTypeID:  2,
		SellerID:       3,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Create()(w, req)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestProductHandler_Create_BadJSON(t *testing.T) {
	mock := &MockProductService{
		SaveFunc: func(p *models.Product) error { return nil },
	}
	h := NewProductHandler(mock)
	body := []byte(`{mal json}`)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Create()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestNoBody.Error())
}

func TestProductHandler_Create_ValidationError(t *testing.T) {
	mock := &MockProductService{
		SaveFunc: func(p *models.Product) error { return nil },
	}
	h := NewProductHandler(mock)
	// Falta ProductCode y otros campos requeridos
	product := models.Product{ID: 1}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Create()(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestWrongBody.Error())
}

func TestProductHandler_Create_SellerNotFound(t *testing.T) {
	mock := &MockProductService{
		SaveFunc: func(p *models.Product) error { return e.ErrSellerRepositoryNotFound },
	}
	h := NewProductHandler(mock)
	product := models.Product{
		ID:             1,
		ProductCode:    "P001",
		Description:    "Producto",
		Height:         10.5,
		Length:         20.0,
		Width:          5.0,
		Weight:         1.2,
		ExpirationRate: 0.5,
		FreezingRate:   1.0,
		RecomFreezTemp: -18.0,
		ProductTypeID:  2,
		SellerID:       3,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Create()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrSellerRepositoryNotFound.Error())
}

func TestProductHandler_Create_InternalError(t *testing.T) {
	mock := &MockProductService{
		SaveFunc: func(p *models.Product) error { return errors.New("error interno") },
	}
	h := NewProductHandler(mock)
	product := models.Product{
		ID:             1,
		ProductCode:    "P001",
		Description:    "Producto",
		Height:         10.5,
		Length:         20.0,
		Width:          5.0,
		Weight:         1.2,
		ExpirationRate: 0.5,
		FreezingRate:   1.0,
		RecomFreezTemp: -18.0,
		ProductTypeID:  2,
		SellerID:       3,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	w := httptest.NewRecorder()

	h.Create()(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error interno")
}

func TestProductHandler_Update_Success(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{
				ID:             id,
				ProductCode:    "P001",
				Description:    "Producto original",
				Height:         10.5,
				Length:         20.0,
				Width:          5.0,
				Weight:         1.2,
				ExpirationRate: 0.5,
				FreezingRate:   1.0,
				RecomFreezTemp: -18.0,
				ProductTypeID:  2,
				SellerID:       3,
			}, nil
		},
		UpdateFunc: func(p *models.Product) error {
			return nil
		},
	}
	h := NewProductHandler(mock)
	actualizado := "Producto actualizado"
	code := "P002"
	height := 11.0
	patch := models.ProductPatch{
		Description: &actualizado,
		ProductCode: &code,
		Height:      &height,
	}
	body, _ := json.Marshal(patch)
	req := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_Update_InvalidID(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) { return models.Product{}, nil },
	}
	h := NewProductHandler(mock)
	body := []byte(`{}`)
	req := httptest.NewRequest(http.MethodPatch, "/products/abc", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestIdMustBeInt.Error())
}

func TestProductHandler_Update_NotFound(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) { return models.Product{}, e.ErrProductRepositoryNotFound },
	}
	h := NewProductHandler(mock)
	body := []byte(`{}`)
	req := httptest.NewRequest(http.MethodPatch, "/products/99", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "99")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrProductRepositoryNotFound.Error())
}

func TestProductHandler_Update_FindByIDError(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) { return models.Product{}, errors.New("error find") },
	}
	h := NewProductHandler(mock)
	body := []byte(`{}`)
	req := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error find")
}

func TestProductHandler_Update_BadJSON(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		UpdateFunc: func(p *models.Product) error { return nil },
	}
	h := NewProductHandler(mock)
	body := []byte(`{mal json}`)
	req := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestNoBody.Error())
}

func TestProductHandler_Update_ValidationError(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		UpdateFunc: func(p *models.Product) error { return nil },
	}
	h := NewProductHandler(mock)
	negativo := -1.0
	patch := models.ProductPatch{Height: &negativo} // Valor inválido para Height
	body, _ := json.Marshal(patch)
	req := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestProductHandler_Update_UpdateNotFound(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		UpdateFunc: func(p *models.Product) error { return e.ErrProductRepositoryNotFound },
	}
	h := NewProductHandler(mock)
	patch := models.ProductPatch{Description: strPtr("desc")}
	body, _ := json.Marshal(patch)
	req := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrProductRepositoryNotFound.Error())
}

func TestProductHandler_Update_SellerNotFound(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		UpdateFunc: func(p *models.Product) error { return e.ErrSellerRepositoryNotFound },
	}
	h := NewProductHandler(mock)
	patch := models.ProductPatch{Description: strPtr("desc")}
	body, _ := json.Marshal(patch)
	req := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrSellerRepositoryNotFound.Error())
}

func TestProductHandler_Update_InternalError(t *testing.T) {
	mock := &MockProductService{
		FindByIDFunc: func(id int) (models.Product, error) {
			return models.Product{ID: id}, nil
		},
		UpdateFunc: func(p *models.Product) error { return errors.New("error interno") },
	}
	h := NewProductHandler(mock)
	patch := models.ProductPatch{Description: strPtr("desc")}
	body, _ := json.Marshal(patch)
	req := httptest.NewRequest(http.MethodPatch, "/products/1", bytes.NewReader(body))
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Update()(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error interno")
}

func strPtr(s string) *string { return &s }

func TestProductHandler_Delete_Success(t *testing.T) {
	mock := &MockProductService{
		DeleteFunc: func(id int) error {
			return nil
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Delete()(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestProductHandler_Delete_NotFound(t *testing.T) {
	mock := &MockProductService{
		DeleteFunc: func(id int) error {
			return e.ErrProductRepositoryNotFound
		},
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodDelete, "/products/99", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "99")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Delete()(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestProductHandler_Delete_InvalidID(t *testing.T) {
	mock := &MockProductService{
		DeleteFunc: func(id int) error { return nil },
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodDelete, "/products/abc", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Delete()(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), e.ErrRequestIdMustBeInt.Error())
}

func TestProductHandler_Delete_InternalError(t *testing.T) {
	mock := &MockProductService{
		DeleteFunc: func(id int) error { return errors.New("error interno") },
	}
	h := NewProductHandler(mock)
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	w := httptest.NewRecorder()

	h.Delete()(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "error interno")
}
