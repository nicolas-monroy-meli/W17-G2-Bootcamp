package internal

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"net/http"
)

// EmployeeRepository is an interface that contains the methods that the employee repository should support
type EmployeeRepository interface {
	// FindAll returns all the employees
	FindAllEmployees() (map[int]mod.Employee, error)
	// FindByID returns the employee with the given ID
	FindEmployeeByID(id int) (mod.Employee, error)
	// Save saves the given employee
	SaveEmployee(employee *mod.Employee) error
	// Update updates the given employee
	UpdateEmployee(id int, employee *mod.Employee) error
	// Delete deletes the employee with the given ID
	DeleteEmployee(id int) error
}

// EmployeeService is an interface that contains the methods that the employee service should support
type EmployeeService interface {
	// FindAll returns all the employees
	FindAllEmployees() (map[int]mod.Employee, error)
	// FindByID returns the employee with the given ID
	FindEmployeeByID(id int) (mod.Employee, error)
	// Save saves the given employee
	SaveEmployee(employee *mod.Employee) error
	// Update updates the given employee
	UpdateEmployee(id int, employee *mod.Employee) error
	// Delete deletes the employee with the given ID
	DeleteEmployee(id int) error
}

// EmployeeService is an interface that contains the methods that the buyer service should support
type EmployeeHandler interface {
	// FindAll returns all the buyers
	GetAllEmployees() http.HandlerFunc
	// FindByID returns the buyer with the given ID
	GetEmployeeById() http.HandlerFunc
	// Save saves the given buyer
	CreateEmployee() http.HandlerFunc
	// Update updates the given buyer
	EditEmployee() http.HandlerFunc
	// Delete deletes the buyer with the given ID
	DeleteEmployee() http.HandlerFunc
}
