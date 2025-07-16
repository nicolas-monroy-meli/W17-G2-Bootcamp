package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"time"
)

func NewPurchaseOrderRepo(db *sql.DB) *PurchaseOrderDB {
	return &PurchaseOrderDB{
		db: db,
	}
}

// BuyerDB is the implementation of the Buyer database
type PurchaseOrderDB struct {
	db *sql.DB
}

func (r *PurchaseOrderDB) Save(purchaseOrder *mod.PurchaseOrder) (err error) {

	result, err := r.db.Exec(
		"INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id) "+
			"VALUES (?, ?, ?, ?)",
		(*purchaseOrder).OrderNumber, time.Time((*purchaseOrder).OrderDate),
		(*purchaseOrder).TrackingCode, (*purchaseOrder).BuyerId,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				return e.ErrPORepositoryOrderNumberDuplicated
			case 1452:
				return fmt.Errorf("Foreign key error: %w", err)
			default:
				return err
			}
		}
	}

	lastInsertId, err := result.LastInsertId()

	for idx, od := range purchaseOrder.ProductsDetails {
		od.PurchaseOrderId = int(lastInsertId)
		err = r.insertOrderDetail(&od)
		if err != nil {
			break
		}
		purchaseOrder.ProductsDetails[idx] = od
	}

	if err != nil {
		return err
	}

	(*purchaseOrder).ID = int(lastInsertId)

	return err
}

func (r *PurchaseOrderDB) insertOrderDetail(orderDetails *mod.OrderDetails) (err error) {
	result, err := r.db.Exec(
		"INSERT INTO order_details (clean_liness_status, quantity, temperature, product_record_id, purchase_order_id) "+
			"VALUES (?, ?, ?, ?, ?)",
		(*orderDetails).CleanLinessStatus, (*orderDetails).Quantity, (*orderDetails).Temperature,
		(*orderDetails).ProductRecordId, (*orderDetails).PurchaseOrderId,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1452:
				return fmt.Errorf("Foreign key error: %w", err)
			default:
				return err
			}
		}
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	(*orderDetails).ID = int(lastInsertId)

	return nil
}

func (r *PurchaseOrderDB) GetByOrderNumber(orderNumber string) (purchaseOrder mod.PurchaseOrder, err error) {
	rowBuyer := r.db.QueryRow(
		"SELECT id, order_number FROM purchase_orders WHERE order_number = ? ", orderNumber,
	)

	err = rowBuyer.Scan(
		&purchaseOrder.ID,
		&purchaseOrder.OrderNumber,
	)

	if err != nil {
		return
	}

	return purchaseOrder, nil
}
