package repository

import (
	"github.com/smartineztri_meli/W17-G2-Bootcamp/docs"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
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

// FindAll returns all employees
func (r *EmployeeDB) FindAllEmployees() (employees map[int]mod.Employee, err error) {
	employees = r.db
	if len(r.db) == 0 {
		return nil, utils.ErrSellerRepositoryNotFound
	}
	return employees, nil
}

// FindByID returns a employee
func (r *EmployeeDB) FindEmployeeByID(id int) (employee mod.Employee, err error) {
	for _, e := range r.db {
		if e.ID == id {
			employee = e
			break
		}
	}
	if employee.ID == 0 {
		return employee, utils.ErrProductRepositoryNotFound
	}
	return employee, nil
}

// Save creates a new employee
func (r *EmployeeDB) SaveEmployee(employee *mod.Employee) (err error) {
	for _, e := range r.db {
		if e.ID == employee.ID {
			return utils.ErrEmployeeRepositoryDuplicated
		}
	}
	employee.ID = len(r.db) + 1
	r.db[employee.ID] = *employee
	docs.WriterFile("employees.json", r.db)
	return nil
}

// Update updates a employee
func (r *EmployeeDB) UpdateEmployee(id int, employee *mod.Employee) (err error) {
	for _, e := range r.db {
		if e.ID == id {
			r.db[employee.ID] = *employee
			docs.WriterFile("employees.json", r.db)
			return nil
		}
	}
	return utils.ErrEmployeeRepositoryNotFound
}

// Delete deletes a employee
func (r *EmployeeDB) DeleteEmployee(id int) (err error) {
	return
}
