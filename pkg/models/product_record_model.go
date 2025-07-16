package models

// ProductRecords is a struct that connrects a product with its records
type ProductRecord struct {
	// ID is the unique identifier of the product
	ID int `json:"id"`
	// LastUpdateDate is the last update date of the product
	LastUpdateDate string `json:"last_update_date" validate:"required"`
	// PurchasePrice is the purchase price of the product
	PurchasePrice float64 `json:"purchase_price" validate:"required"`
	// SalePrice is the sale price of the product
	SalePrice float64 `json:"sale_price" validate:"required"`
	// ProductID is the unique identifier of the product
	ProductID int `json:"product_id" validate:"required"`
}
