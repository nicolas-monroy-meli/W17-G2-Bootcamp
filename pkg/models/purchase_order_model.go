package models

import (
	"fmt"
	"strings"
	"time"
)

type PurchaseOrder struct {
	ID              int            `json:"id"`
	OrderNumber     string         `json:"order_number" validate:"required"`
	OrderDate       Date           `json:"order_date" validate:"required"`
	TrackingCode    string         `json:"tracking_code" validate:"required"`
	BuyerId         int            `json:"buyer_id" validate:"required"`
	ProductsDetails []OrderDetails `json:"products_details" validate:"min=1"`
}

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(d).Format("2006-01-02"))
	return []byte(formatted), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}
