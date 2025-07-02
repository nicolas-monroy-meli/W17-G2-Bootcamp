package models

// Buyer is a struct that contains the buyer's information
type Buyer struct {
	// ID is the unique identifier of the buyer
	ID int `json:"id"`
	// CardNumberID is the unique identifier of the card number
	CardNumberID int `json:"card_number_id" validate:"gt=0"`
	// FirstName is the first name of the buyer
	FirstName string `json:"first_name" validate:"min=1"`
	// LastName is the last name of the buyer
	LastName string `json:"last_name" validate:"min=1"`
}

type BuyerPatch struct {
	CardNumberID *int    `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}
