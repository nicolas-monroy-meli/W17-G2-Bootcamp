package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/internal/service"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// Mock del repositorio de ProductRecord

type MockProductRecordRepo struct {
	mock.Mock
}

func (m *MockProductRecordRepo) FindAllPR() (map[int]mod.ProductRecord, error) {
	args := m.Called()
	return args.Get(0).(map[int]mod.ProductRecord), args.Error(1)
}

func (m *MockProductRecordRepo) FindAllByProductIDPR(productID int) (map[int]mod.ProductRecord, error) {
	args := m.Called(productID)
	return args.Get(0).(map[int]mod.ProductRecord), args.Error(1)
}

func (m *MockProductRecordRepo) SavePR(productRecord *mod.ProductRecord) error {
	args := m.Called(productRecord)
	return args.Error(0)
}

// Mock del repositorio de Product

type MockProductRepoPRService struct {
	mock.Mock
}

func (m *MockProductRepoPRService) FindByID(id int) (mod.Product, error) {
	args := m.Called(id)
	return args.Get(0).(mod.Product), args.Error(1)
}

func (m *MockProductRepoPRService) FindAll() ([]mod.Product, error) {
	args := m.Called()
	return args.Get(0).([]mod.Product), args.Error(1)
}

func (m *MockProductRepoPRService) Save(product *mod.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepoPRService) Update(product *mod.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepoPRService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestProductRecordService_FindAllPR(t *testing.T) {
	mockPRRepo := new(MockProductRecordRepo)
	mockProdRepo := new(MockProductRepoPRService)
	expected := map[int]mod.ProductRecord{1: {ID: 1, ProductID: 10}}
	mockPRRepo.On("FindAllPR").Return(expected, nil)

	svc := service.NewProductRecordService(mockPRRepo, mockProdRepo)
	result, err := svc.FindAllPR()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestProductRecordService_FindAllByProductIDPR(t *testing.T) {
	mockPRRepo := new(MockProductRecordRepo)
	mockProdRepo := new(MockProductRepoPRService)
	expected := map[int]mod.ProductRecord{1: {ID: 1, ProductID: 10}}
	mockPRRepo.On("FindAllByProductIDPR", 10).Return(expected, nil)

	svc := service.NewProductRecordService(mockPRRepo, mockProdRepo)
	result, err := svc.FindAllByProductIDPR(10)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestProductRecordService_SavePR(t *testing.T) {
	mockPRRepo := new(MockProductRecordRepo)
	mockProdRepo := new(MockProductRepoPRService)
	pr := &mod.ProductRecord{ID: 1, ProductID: 10}
	mockProdRepo.On("FindByID", 10).Return(mod.Product{ID: 10}, nil)
	mockPRRepo.On("SavePR", pr).Return(nil)

	svc := service.NewProductRecordService(mockPRRepo, mockProdRepo)
	err := svc.SavePR(pr)

	assert.NoError(t, err)
}

func TestProductRecordService_SavePR_ProductNotFound(t *testing.T) {
	mockPRRepo := new(MockProductRecordRepo)
	mockProdRepo := new(MockProductRepoPRService)
	pr := &mod.ProductRecord{ID: 1, ProductID: 10}
	mockProdRepo.On("FindByID", 10).Return(mod.Product{}, errors.New("not found"))

	svc := service.NewProductRecordService(mockPRRepo, mockProdRepo)
	err := svc.SavePR(pr)

	assert.Error(t, err)
}

func TestProductRecordService_FindProductByID(t *testing.T) {
	mockPRRepo := new(MockProductRecordRepo)
	mockProdRepo := new(MockProductRepoPRService)
	mockProdRepo.On("FindByID", 10).Return(mod.Product{ID: 10}, nil)

	svc := service.NewProductRecordService(mockPRRepo, mockProdRepo)
	result, err := svc.FindProductByID(10)

	assert.NoError(t, err)
	assert.Equal(t, 10, result.ID)
}

func TestProductRecordService_FindAllPR_Error(t *testing.T) {
	mockPRRepo := new(MockProductRecordRepo)
	mockProdRepo := new(MockProductRepoPRService)
	mockPRRepo.On("FindAllPR").Return(map[int]mod.ProductRecord{}, errors.New("db error"))

	svc := service.NewProductRecordService(mockPRRepo, mockProdRepo)
	result, err := svc.FindAllPR()

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestProductRecordService_FindAllByProductIDPR_Error(t *testing.T) {
	mockPRRepo := new(MockProductRecordRepo)
	mockProdRepo := new(MockProductRepoPRService)
	mockPRRepo.On("FindAllByProductIDPR", 10).Return(map[int]mod.ProductRecord{}, errors.New("db error"))

	svc := service.NewProductRecordService(mockPRRepo, mockProdRepo)
	result, err := svc.FindAllByProductIDPR(10)

	assert.Error(t, err)
	assert.Empty(t, result)
}
