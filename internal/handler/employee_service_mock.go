package handler

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeService struct {
	mock.Mock
}

func (m *MockEmployeeService) FindAll() ([]mod.Employee, error) {
	args := m.Called()
	return args.Get(0).([]mod.Employee), args.Error(1)
}

func (m *MockEmployeeService) FindByID(id int) (*mod.Employee, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mod.Employee), args.Error(1)
}

func (m *MockEmployeeService) Save(employee *mod.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeService) Update(id int, employee *mod.Employee) error {
	args := m.Called(id, employee)
	return args.Error(0)
}

func (m *MockEmployeeService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
