package common_test

import (
	"testing"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	"github.com/stretchr/testify/require"
)

func TestPatchSeller(t *testing.T) {
	cid := 99
	companyName := "Acme Inc."
	address := "123 Market Street"
	telephone := "555-1234"
	locality := 10

	baseSeller := mod.Seller{
		ID:          1,
		CID:         1,
		CompanyName: "Original Co",
		Address:     "Old Address",
		Telephone:   "000-0000",
		Locality:    1,
	}

	tests := []struct {
		name       string
		patch      mod.SellerPatch
		wantSeller mod.Seller
	}{
		{
			name:       "#1 Nothing to Patch - empty patch",
			patch:      mod.SellerPatch{},
			wantSeller: baseSeller, // unchanged
		},
		{
			name:  "#2 One Patch - company name",
			patch: mod.SellerPatch{CompanyName: &companyName},
			wantSeller: mod.Seller{
				ID:          1,
				CID:         1,
				CompanyName: "Acme Inc.",
				Address:     "Old Address",
				Telephone:   "000-0000",
				Locality:    1,
			},
		},
		{
			name: "#3 Multiple Patches - all fields",
			patch: mod.SellerPatch{
				CID:         &cid,
				CompanyName: &companyName,
				Address:     &address,
				Telephone:   &telephone,
				Locality:    &locality,
			},
			wantSeller: mod.Seller{
				ID:          1,
				CID:         99,
				CompanyName: "Acme Inc.",
				Address:     "123 Market Street",
				Telephone:   "555-1234",
				Locality:    10,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := common.PatchSeller(baseSeller, tc.patch)
			require.Equal(t, &tc.wantSeller, result)
		})
	}
}
