package mock

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockSellerService struct {
	mock.Mock
}

func (m *MockSellerService) FindAll() ([]mod.Seller, error) {
	args := m.Called()
	return args.Get(0).([]mod.Seller), args.Error(1)
}

func (m *MockSellerService) FindByID(id int) (*mod.Seller, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mod.Seller), args.Error(1)
}

func (m *MockSellerService) Save(seller *mod.Seller) error {
	args := m.Called(seller)
	return args.Error(0)
}

func (m *MockSellerService) Update(id int, seller *mod.Seller) error {
	args := m.Called(id, seller)
	return args.Error(0)
}

func (m *MockSellerService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
