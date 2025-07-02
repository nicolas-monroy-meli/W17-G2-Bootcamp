package common

import mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"

func PatchSeller(seller, newSeller mod.Seller) mod.Seller {
	if newSeller.Telephone != "" {
		seller.Telephone = newSeller.Telephone
	}
	if newSeller.Address != "" {
		seller.Address = newSeller.Address
	}
	if newSeller.CompanyName != "" {
		seller.CompanyName = newSeller.CompanyName
	}
	if newSeller.CID != 0 {
		seller.CID = newSeller.CID
	}
	return seller
}