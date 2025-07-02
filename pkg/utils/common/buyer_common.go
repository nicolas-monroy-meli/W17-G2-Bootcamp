package common

import mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"

func ValidatePatchRequest(buyer *mod.Buyer, body mod.BuyerPatch) {

	if body.CardNumberID != nil {
		(*buyer).CardNumberID = *body.CardNumberID
	}
	if body.FirstName != nil {
		(*buyer).FirstName = *body.FirstName
	}
	if body.LastName != nil {
		(*buyer).LastName = *body.LastName
	}
}
