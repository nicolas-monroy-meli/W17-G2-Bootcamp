package repository

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewEmployeeRepo creates a new instance of the Employee repository
func NewEmployeeRepo(employees map[int]mod.Employee) *EmployeeDB {
	return &EmployeeDB{
		db: employees,
	}
}

// EmployeeDB is the implementation of the Employee database
type EmployeeDB struct {
	db map[int]mod.Employee
}

// FindAll returns all employees from the database
func (r *EmployeeDB) FindAll() (employees map[int]mod.Employee, err error) {

	return
}

// FindByID returns a employee from the database by its id
func (r *EmployeeDB) FindByID(id int) (employee mod.Employee, err error) {

	return
}

// Save saves the given employee in the database
func (r *EmployeeDB) Save(employee *mod.Employee) (err error) {

	return
}

// Update updates the given employee in the database
func (r *EmployeeDB) Update(employee *mod.Employee) (err error) {

	return
}

// Delete deletes the given employee from the database
func (r *EmployeeDB) Delete(id int) (err error) {

	return
}
