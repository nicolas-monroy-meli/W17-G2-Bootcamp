package common

import "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"

func PatchProduct(product *models.Product, newProduct models.ProductPatch) {
	if newProduct.ID != nil {
		product.ID = *newProduct.ID
	}
	if newProduct.ProductCode != nil {
		product.ProductCode = *newProduct.ProductCode
	}
	if newProduct.Description != nil {
		product.Description = *newProduct.Description
	}
	if newProduct.Height != nil {
		product.Height = *newProduct.Height
	}
	if newProduct.Length != nil {
		product.Length = *newProduct.Length
	}
	if newProduct.Width != nil {
		product.Width = *newProduct.Width
	}
	if newProduct.Weight != nil {
		product.Weight = *newProduct.Weight
	}
	if newProduct.ExpirationRate != nil {
		product.ExpirationRate = *newProduct.ExpirationRate
	}
	if newProduct.FreezingRate != nil {
		product.FreezingRate = *newProduct.FreezingRate
	}
	if newProduct.RecomFreezTemp != nil {
		product.RecomFreezTemp = *newProduct.RecomFreezTemp
	}
	if newProduct.ProductTypeID != nil {
		product.ProductTypeID = *newProduct.ProductTypeID
	}
	if newProduct.SellerID != nil {
		product.SellerID = *newProduct.SellerID
	}
}
