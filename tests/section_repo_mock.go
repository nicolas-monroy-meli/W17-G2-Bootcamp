package tests

import (
	"database/sql/driver"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

var SectionTableStruct = []string{`id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`}

var SectionDataValuesSelect = [][]driver.Value{
	{
		1, 1, 0, -5, 50, 20, 100, 1, 1,
	},
	{
		2, 2, -2, -6, 60, 30, 110, 2, 2,
	},
}

var SectionDataValuesSelectByID = []driver.Value{
	2, 2, -2, -6, 60, 30, 110, 2, 2,
}

var SectionSelectExpectedQuery = "SELECT `id`, `section_number`,`current_temperature`,`minimum_temperature`,`current_capacity`, `minimum_capacity`,`maximum_capacity`,`warehouse_id`,`product_type_id` FROM `sections`"

var SectionSelectWhereExpectedQuery = "SELECT `id`, `section_number`,`current_temperature`,`minimum_temperature`,`current_capacity`, `minimum_capacity`,`maximum_capacity`,`warehouse_id`,`product_type_id`  FROM `sections` WHERE `id`=?"

type MockSectionService struct {
	MockFindAll        func() ([]mod.Section, error)
	MockFindByID       func(id int) (mod.Section, error)
	MockSave           func(section *mod.Section) error
	MockDelete         func(id int) error
	MockUpdate         func(id int, fields map[string]interface{}) (*mod.Section, error)
	MockReportProducts func(ids []int) ([]mod.ReportProductsResponse, error)
}

func (m *MockSectionService) FindAll() ([]mod.Section, error) {
	return m.MockFindAll()
}

func (m *MockSectionService) FindByID(id int) (mod.Section, error) {
	return m.MockFindByID(id)
}
func (m *MockSectionService) Delete(id int) error {
	return m.MockDelete(id)
}
func (m *MockSectionService) Update(id int, fields map[string]interface{}) (*mod.Section, error) {
	return m.MockUpdate(id, fields)
}
func (m *MockSectionService) ReportProducts(ids []int) ([]mod.ReportProductsResponse, error) {
	return m.MockReportProducts(ids)
}

func (m *MockSectionService) Save(section *mod.Section) error {
	return m.MockSave(section)
}
