package service

import (
	internal "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/pkg/models"
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
func (s *EmployeeService) FindAll() (employees map[int]mod.Employee, err error) {
	return
}

// FindByID returns a employee
func (s *EmployeeService) FindByID(id int) (employee mod.Employee, err error) {
	return
}

// Save creates a new employee
func (s *EmployeeService) Save(employee *mod.Employee) (err error) {
	return
}

// Update updates a employee
func (s *EmployeeService) Update(employee *mod.Employee) (err error) {
	return
}

// Delete deletes a employee
func (s *EmployeeService) Delete(id int) (err error) {
	return
}
