package internal

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"net/http"
)

// EmployeeRepository is an interface that contains the methods that the employee repository should support
type EmployeeRepository interface {
	// FindAll returns all the employees
	FindAll() ([]mod.Employee, error)
	// FindByID returns the employee with the given ID
	FindByID(id int) (employee mod.Employee, err error)
	// Save saves the given employee
	Save(employee *mod.Employee) error
	// Update updates the given employee
	Update(id int, employee *mod.Employee) error
	// Delete deletes the employee with the given ID
	Delete(id int) error
}

// EmployeeService is an interface that contains the methods that the employee service should support
type EmployeeService interface {
	// FindAll returns all the employees
	FindAll() ([]mod.Employee, error)
	// FindByID returns the employee with the given ID
	FindByID(id int) (mod.Employee, error)
	// Save saves the given employee
	Save(employee *mod.Employee) error
	// Update updates the given employee
	Update(id int, employee *mod.Employee) error
	// Delete deletes the employee with the given ID
	Delete(id int) error
}

// EmployeeService is an interface that contains the methods that the buyer service should support
type EmployeeHandler interface {
	// FindAll returns all the buyers
	GetAll() http.HandlerFunc
	// FindByID returns the buyer with the given ID
	GetById() http.HandlerFunc
	// Save saves the given buyer
	Create() http.HandlerFunc
	// Update updates the given buyer
	Edit() http.HandlerFunc
	// Delete deletes the buyer with the given ID
	Delete() http.HandlerFunc
}
