package models

// Product is a struct that contains the product's information
type Product struct {
	// ID is the unique identifier of the product
	ID int `json:"id"`
	// ProductCode is the unique code of the product
	ProductCode string `json:"product_code" validate:"required,gte=0"`
	// Description is the description of the product
	Description string `json:"description" validate:"required"`
	// Height is the height of the product
	Height float64 `json:"height" validate:"required,gte=0"`
	// Length is the length of the product
	Length float64 `json:"lenght" validate:"required,gte=0"`
	// Width is the width of the product
	Width float64 `json:"width" validate:"required,gte=0"`
	// Weight is the weight of the product
	Weight float64 `json:"net_weight" validate:"required,gte=0"`
	// ExpirationRate is the rate at which the product expires
	ExpirationRate float64 `json:"expiration_rate" validate:"required,gte=0"`
	// FreezingRate is the rate at which the product should be frozen
	FreezingRate float64 `json:"freezing_rate" validate:"required,gte=0"`
	// RecomFreezTemp is the recommended freezing temperature for the product
	RecomFreezTemp float64 `json:"recommended_freezing_temperature" validate:"required,gte=0"`
	// ProductTypeID is the unique identifier of the product type
	ProductTypeID int `json:"product_type_id" validate:"required"`
	// SellerID is the unique identifier of the seller
	SellerID int `json:"seller_id"`
}
