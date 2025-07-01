package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// NewEmployeeService creates a new instance of the employee service
func NewEmployeeService(employees internal.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		rp: employees,
	}
}

// EmployeeService is the default implementation of the employee service
type EmployeeService struct {
	// rp is the repository used by the service
	rp internal.EmployeeRepository
}

// FindAll returns all employees
func (s *EmployeeService) FindAllEmployees() (employees map[int]mod.Employee, err error) {
	employees, err = s.rp.FindAllEmployees()
	return
}

// FindByID returns a employee
func (s *EmployeeService) FindEmployeeByID(id int) (employee mod.Employee, err error) {
	employee, err = s.rp.FindEmployeeByID(id)
	return
}

// Save creates a new employee
func (s *EmployeeService) SaveEmployee(employee *mod.Employee) (err error) {
	err = s.rp.SaveEmployee(employee)
	return
}

// Update updates a employee
func (s *EmployeeService) UpdateEmployee(id int, employee *mod.Employee) (err error) {
	err = s.rp.UpdateEmployee(id, employee)
	return
}

// Delete deletes a employee
func (s *EmployeeService) DeleteEmployee(id int) (err error) {
	return
}
