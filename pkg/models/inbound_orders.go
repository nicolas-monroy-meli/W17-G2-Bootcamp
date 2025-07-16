package models

type InboundOrders struct {
	Id             int    `json:"id" validate:"numeric,min=1"`
	OrderDate      string `json:"order_date"`
	OrderNumber    string `json:"order_number"`
	EmployeeId     int    `json:"employee_id" validate:"numeric"`
	ProductBatchId int    `json:"product_batch_id" validate:"numeric"`
	WarehouseId    int    `json:"warehouse_id" validate:"numeric"`
}
