package mock

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockLocalityService struct {
	mock.Mock
}

func (m *MockLocalityService) FindAllLocalities() ([]mod.Locality, error) {
	args := m.Called()
	return args.Get(0).([]mod.Locality), args.Error(1)
}

func (m *MockLocalityService) FindSellersByLocID(id int) (result []mod.SelByLoc, err error) {
	args := m.Called(id)
	return args.Get(0).([]mod.SelByLoc), args.Error(1)
}

func (m *MockLocalityService) Save(locality *mod.Locality) (int, error) {
	args := m.Called(locality)
	return args.Get(0).(int), args.Error(1)
}
