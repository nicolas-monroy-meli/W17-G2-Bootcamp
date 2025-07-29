package tests

import (
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/mock"
)

type CarryMock struct {
	mock.Mock
}

func NewCarryMock() *CarryMock {
	return &CarryMock{}
}

func (m *CarryMock) GetAll() ([]models.Carry, error) {
	args := m.Called()
	return args.Get(0).([]models.Carry), args.Error(1)
}

func (m *CarryMock) GetByID(id int) (models.Carry, error) {
	args := m.Called(id)
	return args.Get(0).(models.Carry), args.Error(1)
}

func (m *CarryMock) Save(c *models.Carry) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *CarryMock) Update(c *models.Carry) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *CarryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *CarryMock) ExistsLocality(localityID int) (bool, error) {
	args := m.Called(localityID)
	return args.Bool(0), args.Error(1)
}

func (m *CarryMock) ExistsCID(cid string) (bool, error) {
	args := m.Called(cid)
	return args.Bool(0), args.Error(1)
}

func (m *CarryMock) GetReportByLocality(localityID int) ([]models.LocalityCarryReport, error) {
	args := m.Called(localityID)
	return args.Get(0).([]models.LocalityCarryReport), args.Error(1)
}

func (m *CarryMock) GetByCID(cid string) (models.Carry, error) {
	args := m.Called(cid)
	return args.Get(0).(models.Carry), args.Error(1)
}

func (m *CarryMock) GetReportByLocalityAll() ([]models.LocalityCarryReport, error) {
	args := m.Called()
	return args.Get(0).([]models.LocalityCarryReport), args.Error(1)
}
