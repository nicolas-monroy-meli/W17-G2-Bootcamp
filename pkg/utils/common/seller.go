package common

import (
	"fmt"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

func PatchSeller(seller mod.Seller, newSeller mod.SellerPatch) (*mod.Seller, error) {
	c := 0
	if newSeller.Telephone != nil {
		seller.Telephone = *newSeller.Telephone
		c++
	}
	if newSeller.Address != nil {
		seller.Address = *newSeller.Address
		c++
	}
	if newSeller.CompanyName != nil {
		seller.CompanyName = *newSeller.CompanyName
		c++
	}
	if newSeller.CID != nil {
		seller.CID = *newSeller.CID
		c++
	}
	if c == 0 {
		return nil, fmt.Errorf("error")
	}
	return &seller, nil
}
