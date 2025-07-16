package models

type OrderDetails struct {
	ID                int     `json:"id"`
	CleanLinessStatus string  `json:"clean_liness_status" validate:"min=1"`
	Quantity          int     `json:"quantity" validate:"min=1"`
	Temperature       float64 `json:"temperature" validate:"min=1"`
	ProductRecordId   int     `json:"product_record_id"`
	PurchaseOrderId   int     `json:"purchase_order_id"`
}
