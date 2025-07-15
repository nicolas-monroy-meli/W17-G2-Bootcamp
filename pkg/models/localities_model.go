package models

// Locality is a struct that contains the locality's information
type Locality struct {
	// ID is the unique identifier of the seller
	ID int `json:"id"`
	// Name is the locality's name
	Name string `json:"locality_name" validate:"required"`
	// Province is the province's name
	Province string `json:"province_name" validate:"required"`
	// Country is the country's name
	Country string `json:"country_name" validate:"required"`
}

type SelByLoc struct {
	ID    int    `json:"locality_id"`
	Name  string `json:"locality_name"`
	Count int    `json:"sellers_count"`
}
