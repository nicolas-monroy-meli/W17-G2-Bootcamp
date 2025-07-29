package tests

import "database/sql/driver"

var EmployeeTableStruct = []string{"id", "id_card_number", "first_name", "last_name", "wareHouse_id"}

var EmployeeDataValuesSelect = [][]driver.Value{
	{1, "123", "John", "Doe", 10},
	{2, "456", "Jane", "Smith", 20},
}

var EmployeeDataValuesSelectByID = []driver.Value{
	1, "123", "John", "Doe", 10,
}

var (
	EmployeeSelectExpectedQuery      = "SELECT id,id_card_number,first_name,last_name, wareHouse_id FROM employees"
	EmployeeSelectWhereExpectedQuery = "SELECT id,id_card_number,first_name,last_name, wareHouse_id FROM employees WHERE id = ?"
	EmployeeInsertExpectedQuery      = "INSERT INTO employees (id_card_number,first_name,last_name, wareHouse_id ) VALUES (?, ?,?,?)"
	EmployeeUpdateExpectedQuery      = "UPDATE employees SET id_card_number = ?, first_name = ?, last_name = ?, wareHouse_id = ? WHERE id = ?"
	EmployeeDeleteExpectedQuery      = "DELETE FROM employees WHERE id = ?"
)
