package common_test

import (
	"testing"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	"github.com/stretchr/testify/assert"
)

func TestPatchProduct_AllFieldsPatched(t *testing.T) {
	base := models.Product{
		ID:             1,
		ProductCode:    "A",
		Description:    "desc",
		Height:         1.0,
		Length:         2.0,
		Width:          3.0,
		Weight:         4.0,
		ExpirationRate: 5.0,
		FreezingRate:   6.0,
		RecomFreezTemp: 7.0,
		ProductTypeID:  8,
		SellerID:       9,
	}
	patch := models.ProductPatch{
		ID:             intPtr(10),
		ProductCode:    strPtr("B"),
		Description:    strPtr("newdesc"),
		Height:         floatPtr(11.0),
		Length:         floatPtr(12.0),
		Width:          floatPtr(13.0),
		Weight:         floatPtr(14.0),
		ExpirationRate: floatPtr(15.0),
		FreezingRate:   floatPtr(16.0),
		RecomFreezTemp: floatPtr(17.0),
		ProductTypeID:  intPtr(18),
		SellerID:       intPtr(19),
	}
	common.PatchProduct(&base, patch)
	assert.Equal(t, 10, base.ID)
	assert.Equal(t, "B", base.ProductCode)
	assert.Equal(t, "newdesc", base.Description)
	assert.Equal(t, 11.0, base.Height)
	assert.Equal(t, 12.0, base.Length)
	assert.Equal(t, 13.0, base.Width)
	assert.Equal(t, 14.0, base.Weight)
	assert.Equal(t, 15.0, base.ExpirationRate)
	assert.Equal(t, 16.0, base.FreezingRate)
	assert.Equal(t, 17.0, base.RecomFreezTemp)
	assert.Equal(t, 18, base.ProductTypeID)
	assert.Equal(t, 19, base.SellerID)
}

func TestPatchProduct_NoFieldsPatched(t *testing.T) {
	base := models.Product{
		ID:             1,
		ProductCode:    "A",
		Description:    "desc",
		Height:         1.0,
		Length:         2.0,
		Width:          3.0,
		Weight:         4.0,
		ExpirationRate: 5.0,
		FreezingRate:   6.0,
		RecomFreezTemp: 7.0,
		ProductTypeID:  8,
		SellerID:       9,
	}
	patch := models.ProductPatch{} // todos nil
	common.PatchProduct(&base, patch)
	assert.Equal(t, 1, base.ID)
	assert.Equal(t, "A", base.ProductCode)
	assert.Equal(t, "desc", base.Description)
	assert.Equal(t, 1.0, base.Height)
	assert.Equal(t, 2.0, base.Length)
	assert.Equal(t, 3.0, base.Width)
	assert.Equal(t, 4.0, base.Weight)
	assert.Equal(t, 5.0, base.ExpirationRate)
	assert.Equal(t, 6.0, base.FreezingRate)
	assert.Equal(t, 7.0, base.RecomFreezTemp)
	assert.Equal(t, 8, base.ProductTypeID)
	assert.Equal(t, 9, base.SellerID)
}

func TestPatchProduct_SomeFieldsPatched(t *testing.T) {
	base := models.Product{
		ID:             1,
		ProductCode:    "A",
		Description:    "desc",
		Height:         1.0,
		Length:         2.0,
		Width:          3.0,
		Weight:         4.0,
		ExpirationRate: 5.0,
		FreezingRate:   6.0,
		RecomFreezTemp: 7.0,
		ProductTypeID:  8,
		SellerID:       9,
	}
	patch := models.ProductPatch{
		Description: strPtr("patched"),
		Weight:      floatPtr(99.0),
	}
	common.PatchProduct(&base, patch)
	assert.Equal(t, 1, base.ID)
	assert.Equal(t, "A", base.ProductCode)
	assert.Equal(t, "patched", base.Description)
	assert.Equal(t, 1.0, base.Height)
	assert.Equal(t, 2.0, base.Length)
	assert.Equal(t, 3.0, base.Width)
	assert.Equal(t, 99.0, base.Weight)
	assert.Equal(t, 5.0, base.ExpirationRate)
	assert.Equal(t, 6.0, base.FreezingRate)
	assert.Equal(t, 7.0, base.RecomFreezTemp)
	assert.Equal(t, 8, base.ProductTypeID)
	assert.Equal(t, 9, base.SellerID)
}

// helpers
func intPtr(i int) *int           { return &i }
func strPtr(s string) *string     { return &s }
func floatPtr(f float64) *float64 { return &f }
