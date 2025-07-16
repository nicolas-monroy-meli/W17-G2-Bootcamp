package repository

import (
	"database/sql"
	"fmt"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

type InboundDB struct {
	db *sql.DB
}

func NewInboundRepo(database *sql.DB) *InboundDB {
	return &InboundDB{
		db: database,
	}
}

func (r *InboundDB) Save(order *mod.InboundOrders) (*mod.InboundOrders, error) {
	query := `INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id)
	          VALUES (?, ?, ?, ?, ?)`

	res, err := r.db.Exec(query, order.OrderDate, order.OrderNumber, order.EmployeeId, order.ProductBatchId, order.WarehouseId)
	if err != nil {
		// La base de datos es la que nos dirá el tipo de error
		// Aquí deberías inspeccionar el error para saber qué ha fallado
		return nil, fmt.Errorf("%w: %v", e.ErrInboundOrderInternal, err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%w: failed to get last insert ID: %v", e.ErrInboundOrderInternal, err)
	}
	order.Id = int(lastID)

	return order, nil
}

func (r *InboundDB) FindOrdersByEmployee(employeeID int) ([]mod.EmployeeReport, error) {
	query := `
        SELECT
            e.id,
            e.id_card_number,
            e.first_name,
            e.last_name,
            e.wareHouse_id,
            COUNT(io.id) AS inbound_orders_count
        FROM employees AS e
        LEFT JOIN inbound_orders AS io ON e.id = io.employee_id`

	args := []interface{}{}
	if employeeID > 0 {
		query += " WHERE e.id = ?"
		args = append(args, employeeID)
	}

	query += " GROUP BY e.id"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: failed to query report: %v", e.ErrEmployeeInternal, err)
	}
	defer rows.Close()

	var reports []mod.EmployeeReport
	for rows.Next() {
		var report mod.EmployeeReport
		if err := rows.Scan(
			&report.ID,
			&report.CardNumberID,
			&report.FirstName,
			&report.LastName,
			&report.WarehouseID,
			&report.InboundOrdersCount,
		); err != nil {
			return nil, fmt.Errorf("%w: failed to scan report row: %v", e.ErrEmployeeInternal, err)
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: rows iteration error: %v", e.ErrEmployeeInternal, err)
	}

	if len(reports) == 0 && employeeID > 0 {
		return nil, e.ErrEmployeeNotFound
	}

	return reports, nil
}
