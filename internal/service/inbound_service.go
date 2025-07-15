package service

import (
	"errors"
	internal "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/interfaces"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"strings"
)

// NewEmployeeService creates a new instance of the employee service
func NewInboundService(inbound internal.InboundRepository) *InboundService {
	return &InboundService{
		rp: inbound,
	}
}

// EmployeeService is the default implementation of the employee service
type InboundService struct {
	// rp is the repository used by the service
	rp internal.InboundRepository
}

func (s *InboundService) Save(order *mod.InboundOrders) (*mod.InboundOrders, error) {
	if order.OrderNumber == "" || order.EmployeeId == 0 || order.WarehouseId == 0 {
		return nil, e.ErrInboundOrderInvalidData
	}

	// 2. Intentar crear la orden directamente
	createdOrder, err := s.rp.Save(order)
	if err != nil {
		// 3. Interpretar el error de la base de datos
		errStr := err.Error()

		// Error por duplicado de order_number (UNIQUE constraint)
		if strings.Contains(errStr, "Duplicate entry") && strings.Contains(errStr, "order_number") {
			return nil, errors.New("order number already exists")
		}

		// Error de llave foránea (FOREIGN KEY constraint)
		if strings.Contains(errStr, "Cannot add or update a child row") {
			return nil, errors.New("employee not found")
		}

		return nil, err // Retornar el error genérico si no coincide con los esperados
	}

	return createdOrder, nil
}

func (s *InboundService) FindOrdersByEmployee(employeeID int) ([]mod.EmployeeReport, error) {
	return s.rp.FindOrdersByEmployee(employeeID)
}
