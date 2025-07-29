package mocks

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockPurchaseOrderService struct {
	mock.Mock
}

func (m *MockPurchaseOrderService) Save(buyer *mod.PurchaseOrder) (err error) {
	args := m.Called(buyer)
	return args.Error(0)
}
