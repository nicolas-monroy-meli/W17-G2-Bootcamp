package models

type OrderDetails struct {
	ID                int     `json:"id"`
	CleanLinessStatus string  `json:"clean_liness_status" validate:"required"`
	Quantity          int     `json:"quantity" validate:"required"`
	Temperature       float64 `json:"temperature" validate:"required"`
	ProductRecordId   int     `json:"product_record_id" validate:"required"`
	PurchaseOrderId   int     `json:"purchase_order_id"`
}
