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
func (s *EmployeeService) FindAll() (employees []mod.Employee, err error) {
	employees, err = s.rp.FindAll()
	return
}

// FindByID returns a employee
func (s *EmployeeService) FindByID(id int) (employee mod.Employee, err error) {
	employee, err = s.rp.FindByID(id)
	return
}

// Save creates a new employee
func (s *EmployeeService) Save(employee *mod.Employee) (err error) {
	err = s.rp.Save(employee)
	return
}

// Update updates a employee
func (s *EmployeeService) Update(id int, employee *mod.Employee) (err error) {
	err = s.rp.Update(id, employee)
	return
}

// Delete deletes a employee
func (s *EmployeeService) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}
