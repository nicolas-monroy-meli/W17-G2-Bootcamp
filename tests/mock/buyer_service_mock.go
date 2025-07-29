package mock

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockBuyerService struct {
	mock.Mock
}

func (m *MockBuyerService) FindAll() ([]mod.Buyer, error) {
	args := m.Called()
	if buyers := args.Get(0); buyers != nil {
		return buyers.([]mod.Buyer), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBuyerService) FindByID(id int) (buyer mod.Buyer, err error) {
	args := m.Called(id)
	if buyers := args.Get(0); buyers != nil {
		return buyers.(mod.Buyer), args.Error(1)
	}
	return mod.Buyer{}, args.Error(1)
}

func (m *MockBuyerService) Save(buyer *mod.Buyer) (err error) {
	args := m.Called(buyer)
	return args.Error(0)
}

func (m *MockBuyerService) Update(buyer *mod.Buyer) (err error) {
	args := m.Called(buyer)
	return args.Error(0)
}

func (m *MockBuyerService) Delete(id int) (err error) {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBuyerService) GetPurchaseOrderReport(id *int) ([]mod.BuyerReportPO, error) {
	args := m.Called(id)
	if buyers := args.Get(0); buyers != nil {
		return buyers.([]mod.BuyerReportPO), args.Error(1)
	}
	return nil, args.Error(1)
}
