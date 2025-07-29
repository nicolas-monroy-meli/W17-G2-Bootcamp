package tests

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockInboundService struct {
	mock.Mock
}

func (m *MockInboundService) Save(inbound *mod.InboundOrders) (*mod.InboundOrders, error) {
	args := m.Called(inbound)
	return args.Get(0).(*mod.InboundOrders), args.Error(1)
}

func (m *MockInboundService) FindOrdersByEmployee(id int) ([]mod.EmployeeReport, error) {
	args := m.Called(id)
	return args.Get(0).([]mod.EmployeeReport), args.Error(1)
}
