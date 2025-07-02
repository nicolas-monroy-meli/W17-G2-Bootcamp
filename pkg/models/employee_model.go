package models

// Employee is a struct that contains the employee's information
type Employee struct {
	// ID is the unique identifier of the employee
	ID int `json:"id" validate:"numeric,min=1"`
	// CardNumberID is the unique identifier of the card number
	CardNumberID string `json:"card_number_id" validate:"numeric"`
	// FirstName is the first name of the employee
	FirstName string `json:"first_name" `
	// LastName is the last name of the employee
	LastName string `json:"last_name"`
	// WarehouseID is the unique identifier of the warehouse to which the employee belongs
	WarehouseID int `json:"warehouse_id" validate:"numeric, required"`
}
