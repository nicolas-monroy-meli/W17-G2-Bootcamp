package service

import (
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
)

func NewPurchaseOrderService(
	purchaseRp internal.PurchaseOrderRepository,
) *PurchaseOrderService {
	return &PurchaseOrderService{
		rp: purchaseRp,
	}
}

// BuyerService is the default implementation of the buyer service
type PurchaseOrderService struct {
	// rp is the repository used by the service
	rp internal.PurchaseOrderRepository
}

func (s *PurchaseOrderService) Save(purchaseOrder *mod.PurchaseOrder) error {
	return s.rp.Save(purchaseOrder)
}
