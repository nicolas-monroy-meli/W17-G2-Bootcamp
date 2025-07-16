package repository

import (
	"database/sql"
	"errors"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

type EmployeeDB struct {
	db *sql.DB
}

func NewEmployeeRepo(database *sql.DB) *EmployeeDB {
	return &EmployeeDB{
		db: database,
	}
}

// FindAll returns all employees
func (r *EmployeeDB) FindAll() ([]mod.Employee, error) {
	var employees []mod.Employee
	rows, err := r.db.Query("SELECT id,id_card_number,first_name,last_name, wareHouse_id FROM employees") // Adjust columns
	if err != nil {
		return nil, errors.New("failed to query employees") // Use custom error type
	}
	defer rows.Close()

	for rows.Next() {

		var emp mod.Employee
		if err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.CardNumberID, &emp.WarehouseID); err != nil { // Adjust fields
			return nil, errors.New("failed to scan employee row")
		}
		employees = append(employees, emp)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.New("error during row iteration")
	}
	if len(employees) == 0 {
		return nil, e.ErrEmployeeRepositoryNotFound // Your custom error
	}
	return employees, nil
}

// FindById find 0ne employee by id
func (r *EmployeeDB) FindByID(id int) (employee mod.Employee, err error) {
	row := r.db.QueryRow("SELECT id,id_card_number,first_name,last_name, wareHouse_id  FROM employees WHERE id = ?", id) // Use appropriate placeholder for your DB
	err = row.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID) // Adjust fields
	if err != nil {
		if err == sql.ErrNoRows {
			return employee, e.ErrEmployeeRepositoryNotFound // Your custom error
		}
		return employee, errors.New("failed to scan employee by ID")
	}
	return employee, nil
}

// Save creates a new employee
func (r *EmployeeDB) Save(employee *mod.Employee) (err error) {
	res, err := r.db.Exec("INSERT INTO employees (id_card_number,first_name,last_name, wareHouse_id ) VALUES (?, ?,?,?)", employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID) // Adjust fields
	if err != nil {
		return errors.New("failed to insert employee")
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return errors.New("failed to get last insert ID")
	}
	employee.ID = int(lastID) // Update the employee object with the new ID
	return nil
}

// Update updates a employee
func (r *EmployeeDB) Update(id int, employee *mod.Employee) (err error) {
	res, err := r.db.Exec("UPDATE employees SET id_card_number = ?, first_name = ?, last_name = ?, wareHouse_id = ? WHERE id = ?", employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID, id) // Adjust fields
	if err != nil {
		return errors.New("failed to update employee")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("failed to get rows affected")
	}
	if rowsAffected == 0 {
		return e.ErrEmployeeRepositoryNotFound // Or a more specific "not found for update" error
	}
	return nil
}

// Delete deletes a employee
func (r *EmployeeDB) Delete(id int) (err error) {
	res, err := r.db.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		return errors.New("failed to delete employee")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("failed to get rows affected after delete")
	}
	if rowsAffected == 0 {
		return e.ErrEmployeeRepositoryNotFound
	}
	// docs.WriterFile("employees.json", r.db) // This line is for file-based storage, remove it
	return nil
}
