package internal

import (
	"net/http"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

// EmployeeRepository is an interface that contains the methods that the employee repository should support
type EmployeeRepository interface {
	// FindAll returns all the employees
	FindAll() (map[int]mod.Employee, error)
	// FindByID returns the employee with the given ID
	FindByID(id int) (mod.Employee, error)
	// Save saves the given employee
	Save(employee *mod.Employee) error
	// Update updates the given employee
	Update(employee *mod.Employee) error
	// Delete deletes the employee with the given ID
	Delete(id int) error
}

// EmployeeService is an interface that contains the methods that the employee service should support
type EmployeeService interface {
	// FindAll returns all the employees
	FindAll() (map[int]mod.Employee, error)
	// FindByID returns the employee with the given ID
	FindByID(id int) (mod.Employee, error)
	// Save saves the given employee
	Save(employee *mod.Employee) error
	// Update updates the given employee
	Update(employee *mod.Employee) error
	// Delete deletes the employee with the given ID
	Delete(id int) error
}

// EmployeeService is an interface that contains the methods that the buyer service should support
type EmployeeHandler interface {
	// FindAll returns all the buyers
	GetAll() http.HandlerFunc
	// FindByID returns the buyer with the given ID
	GetByID(id int) http.HandlerFunc
	// Save saves the given buyer
	Create(buyer *mod.Employee) http.HandlerFunc
	// Update updates the given buyer
	Update(buyer *mod.Employee) http.HandlerFunc
	// Delete deletes the buyer with the given ID
	Delete(id int) http.HandlerFunc
}
