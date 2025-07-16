package internal

import mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"

type PurchaseOrderRepository interface {
	// Save saves the given purchase order
	Save(purhcaseOrder *mod.PurchaseOrder) error

	//Get By Order Number
	GetByOrderNumber(orderNumber string) (purchaseOrder mod.PurchaseOrder, err error)
}

type PurchaseOrderService interface {
	// Save saves the given purchase order
	Save(purhcaseOrder *mod.PurchaseOrder) error
}
