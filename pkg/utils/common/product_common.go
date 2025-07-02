package common

import "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"

func PatchProduct(product, newProduct models.Product) models.Product {
	if newProduct.ID != 0 {
		product.ID = newProduct.ID
	}
	if newProduct.ProductCode != "" {
		product.ProductCode = newProduct.ProductCode
	}
	if newProduct.Description != "" {
		product.Description = newProduct.Description
	}
	if newProduct.Height != 0 {
		product.Height = newProduct.Height
	}
	if newProduct.Length != 0 {
		product.Length = newProduct.Length
	}
	if newProduct.Width != 0 {
		product.Width = newProduct.Width
	}
	if newProduct.Weight != 0 {
		product.Weight = newProduct.Weight
	}
	if newProduct.ExpirationRate != 0 {
		product.ExpirationRate = newProduct.ExpirationRate
	}
	if newProduct.FreezingRate != 0 {
		product.FreezingRate = newProduct.FreezingRate
	}
	if newProduct.RecomFreezTemp != 0 {
		product.RecomFreezTemp = newProduct.RecomFreezTemp
	}
	if newProduct.ProductTypeID != 0 {
		product.ProductTypeID = newProduct.ProductTypeID
	}
	if newProduct.SellerID != 0 {
		product.SellerID = newProduct.SellerID
	}
	return product
}
