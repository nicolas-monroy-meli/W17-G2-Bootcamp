package internal

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"net/http"
)

type InboundRepository interface {
	Save(inbound *mod.InboundOrders) (*mod.InboundOrders, error)
	FindOrdersByEmployee(id int) ([]mod.EmployeeReport, error)
}

type InboundService interface {
	Save(inbound *mod.InboundOrders) (*mod.InboundOrders, error)
	FindOrdersByEmployee(id int) ([]mod.EmployeeReport, error)
}

type InboundHandler interface {
	Create() http.HandlerFunc
	GetOrdersByEmployee() http.HandlerFunc
}
