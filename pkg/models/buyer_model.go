package models

// Buyer is a struct that contains the buyer's information
type Buyer struct {
	// ID is the unique identifier of the buyer
	ID int `json:"id"`
	// CardNumberID is the unique identifier of the card number
	CardNumberID int `json:"card_number_id" validate:"required,gt=0"`
	// FirstName is the first name of the buyer
	FirstName string `json:"first_name" validate:"required"`
	// LastName is the last name of the buyer
	LastName string `json:"last_name" validate:"required"`
}

type BuyerPatch struct {
	CardNumberID *int    `json:"card_number_id,omitempty" validate:"omitempty,gt=0"`
	FirstName    *string `json:"first_name,omitempty" validate:"omitempty,min=1"`
	LastName     *string `json:"last_name,omitempty" validate:"omitempty,min=1"`
}

func (b *BuyerPatch) MapPatchToEntity(buyer *Buyer) {
	if b.CardNumberID != nil {
		buyer.CardNumberID = *b.CardNumberID
	}
	if b.FirstName != nil {
		buyer.FirstName = *b.FirstName
	}
	if b.LastName != nil {
		buyer.LastName = *b.LastName
	}
}
