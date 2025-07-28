package tests

import (
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type WarehouseMock struct {
	mock.Mock
}

func NewWarehouseMock() *WarehouseMock {
	return &WarehouseMock{}
}

func (m *WarehouseMock) GetAll() ([]models.Warehouse, error) {
	args := m.Called()
	return args.Get(0).([]models.Warehouse), args.Error(1)
}

func (m *WarehouseMock) GetByID(id int) (models.Warehouse, error) {
	args := m.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehouseMock) GetByWarehouseCode(code string) (models.Warehouse, error) {
	args := m.Called(code)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *WarehouseMock) Save(wh *models.Warehouse) error {
	args := m.Called(wh)
	return args.Error(0)
}

func (m *WarehouseMock) Update(wh *models.Warehouse) error {
	args := m.Called(wh)
	return args.Error(0)
}

func (m *WarehouseMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *WarehouseMock) ExistsWarehouseCode(code string) (bool, error) {
	args := m.Called(code)
	return args.Bool(0), args.Error(1)
}
