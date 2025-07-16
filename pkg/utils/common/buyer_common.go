package common

import mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"

func ValidatePatchRequest(buyer mod.Buyer, body mod.BuyerPatch) (buyerMapped mod.Buyer) {

	buyerMapped = buyer

	if body.CardNumberID != nil {
		buyerMapped.CardNumberID = *body.CardNumberID
	}
	if body.FirstName != nil {
		buyerMapped.FirstName = *body.FirstName
	}
	if body.LastName != nil {
		buyerMapped.LastName = *body.LastName
	}

	return buyerMapped
}
