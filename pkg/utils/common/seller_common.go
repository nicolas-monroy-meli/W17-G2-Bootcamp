package common

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

func PatchSeller(seller mod.Seller, newSeller mod.SellerPatch) *mod.Seller {
	if newSeller.Telephone != nil {
		seller.Telephone = *newSeller.Telephone
	}
	if newSeller.Address != nil {
		seller.Address = *newSeller.Address
	}
	if newSeller.CompanyName != nil {
		seller.CompanyName = *newSeller.CompanyName
	}
	if newSeller.CID != nil {
		seller.CID = *newSeller.CID
	}
	if newSeller.Locality != nil {
		seller.Locality = *newSeller.Locality
	}
	return &seller
}
