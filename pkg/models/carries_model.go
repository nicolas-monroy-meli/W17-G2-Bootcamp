package models

type Carry struct {
	ID          int    `json:"id" validate:"-"`
	CID         string `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required,numeric"`
	LocalityID  int    `json:"locality_id" validate:"required,min=1"`
}

type LocalityCarryReport struct {
	LocalityID   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	CarriesCount int    `json:"carries_count"`
}
