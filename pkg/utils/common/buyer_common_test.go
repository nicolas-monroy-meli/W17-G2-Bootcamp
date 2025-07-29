package common

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPatchBuyer(t *testing.T) {

	baseBuyer := mod.Buyer{
		CardNumberID: "1000",
		FirstName:    "Juan",
		LastName:     "Peralta",
	}

	cardNumberId := "2000"
	firstName := "Luis"
	lastName := "Vargas"

	tests := []struct {
		name       string
		patch      mod.BuyerPatch
		buyerToMap mod.Buyer
		expected   mod.Buyer
	}{
		{
			name:       "Case 1: Empty patch",
			patch:      mod.BuyerPatch{},
			buyerToMap: baseBuyer,
			expected:   baseBuyer,
		},
		{
			name: "Case 2: Success patch - Card Number id",
			patch: mod.BuyerPatch{
				CardNumberID: &cardNumberId,
			},
			buyerToMap: baseBuyer,
			expected: mod.Buyer{
				CardNumberID: cardNumberId,
				FirstName:    baseBuyer.FirstName,
				LastName:     baseBuyer.LastName,
			},
		},
		{
			name: "Case 3: Multiple patch",
			patch: mod.BuyerPatch{
				CardNumberID: &cardNumberId,
				FirstName:    &firstName,
				LastName:     &lastName,
			},
			buyerToMap: baseBuyer,
			expected: mod.Buyer{
				CardNumberID: cardNumberId,
				FirstName:    firstName,
				LastName:     lastName,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ValidatePatchRequest(test.buyerToMap, test.patch)
			require.Equal(t, test.expected, result)
		})
	}
}
