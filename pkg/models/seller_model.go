package models

// Seller is a struct that contains the seller's information
type Seller struct {
	// ID is the unique identifier of the seller
	ID int `json:"id"`
	// CID is the unique identifier of the company
	CID int `json:"cid" validate:"required,gte=1"`
	// CompanyName is the name of the company
	CompanyName string `json:"company_name" validate:"required"`
	// Address is the address of the company
	Address string `json:"address" validate:"required"`
	// Telephone is the telephone number of the company
	Telephone string `json:"telephone" validate:"required"`
}

type SellerPatch struct {
	// CID is the unique identifier of the company
	CID *int `json:"cid"`
	// CompanyName is the name of the company
	CompanyName *string `json:"company_name"`
	// Address is the address of the company
	Address *string `json:"address"`
	// Telephone is the telephone number of the company
	Telephone *string `json:"telephone"`
}
