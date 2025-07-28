package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/internal/service"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// Mock del repositorio

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) FindAll() ([]mod.Product, error) {
	args := m.Called()
	return args.Get(0).([]mod.Product), args.Error(1)
}

func (m *MockProductRepo) FindByID(id int) (mod.Product, error) {
	args := m.Called(id)
	return args.Get(0).(mod.Product), args.Error(1)
}

func (m *MockProductRepo) Save(product *mod.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepo) Update(product *mod.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestProductService_FindAll(t *testing.T) {
	mockRepo := new(MockProductRepo)
	expected := []mod.Product{{ID: 1, ProductCode: "P001", Description: "Test product"}}
	mockRepo.On("FindAll").Return(expected, nil)

	svc := service.NewProductService(mockRepo)
	result, err := svc.FindAll()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestProductService_FindByID(t *testing.T) {
	mockRepo := new(MockProductRepo)
	expected := mod.Product{ID: 1, ProductCode: "P001", Description: "Test product"}
	mockRepo.On("FindByID", 1).Return(expected, nil)

	svc := service.NewProductService(mockRepo)
	result, err := svc.FindByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestProductService_Save(t *testing.T) {
	mockRepo := new(MockProductRepo)
	product := &mod.Product{ID: 1, ProductCode: "P001", Description: "Test product"}
	mockRepo.On("Save", product).Return(nil)

	svc := service.NewProductService(mockRepo)
	err := svc.Save(product)

	assert.NoError(t, err)
}

func TestProductService_Update(t *testing.T) {
	mockRepo := new(MockProductRepo)
	product := &mod.Product{ID: 1, ProductCode: "P001", Description: "Test product"}
	mockRepo.On("Update", product).Return(nil)

	svc := service.NewProductService(mockRepo)
	err := svc.Update(product)

	assert.NoError(t, err)
}

func TestProductService_Delete(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockRepo.On("Delete", 1).Return(nil)

	svc := service.NewProductService(mockRepo)
	err := svc.Delete(1)

	assert.NoError(t, err)
}

// Tests de error
func TestProductService_FindAll_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockRepo.On("FindAll").Return([]mod.Product{}, errors.New("db error"))

	svc := service.NewProductService(mockRepo)
	result, err := svc.FindAll()

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestProductService_FindByID_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockRepo.On("FindByID", 1).Return(mod.Product{}, errors.New("not found"))

	svc := service.NewProductService(mockRepo)
	result, err := svc.FindByID(1)

	assert.Error(t, err)
	assert.Equal(t, mod.Product{}, result)
}

func TestProductService_Save_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	product := &mod.Product{ID: 1, ProductCode: "P001", Description: "Test product"}
	mockRepo.On("Save", product).Return(errors.New("save error"))

	svc := service.NewProductService(mockRepo)
	err := svc.Save(product)

	assert.Error(t, err)
}

func TestProductService_Update_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	product := &mod.Product{ID: 1, ProductCode: "P001", Description: "Test product"}
	mockRepo.On("Update", product).Return(errors.New("update error"))

	svc := service.NewProductService(mockRepo)
	err := svc.Update(product)

	assert.Error(t, err)
}

func TestProductService_Delete_Error(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockRepo.On("Delete", 1).Return(errors.New("delete error"))

	svc := service.NewProductService(mockRepo)
	err := svc.Delete(1)

	assert.Error(t, err)
}
