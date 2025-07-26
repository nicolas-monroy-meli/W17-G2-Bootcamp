package tests

import "database/sql/driver"

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
