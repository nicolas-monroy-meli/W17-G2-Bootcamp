package repository

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	//e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewBuyerRepo creates a new instance of the Buyer repository
func NewBuyerRepo(db *sql.DB) *BuyerDB {
	return &BuyerDB{
		db: db,
	}
}

// BuyerDB is the implementation of the Buyer database
type BuyerDB struct {
	db *sql.DB
}

// FindAll returns all buyers from the database
func (r *BuyerDB) FindAll() (buyers []mod.Buyer, err error) {
	rows, err := r.db.Query("SELECT `id`, `id_card_number`, `first_name`, `last_name` FROM buyers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var by mod.Buyer
		// scan the row into the customer
		err = rows.Scan(&by.ID, &by.CardNumberID, &by.FirstName, &by.LastName)
		if err != nil {
			return nil, err
		}
		buyers = append(buyers, by)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// FindByID returns a buyer from the database by its id
func (r *BuyerDB) FindByID(id int) (buyer mod.Buyer, err error) {
	row := r.db.QueryRow(""+
		"SELECT "+
		"`id`, `id_card_number`, `first_name`, `last_name` "+
		"FROM buyers "+
		"WHERE buyers.id = ?", id)

	if err = row.Err(); err != nil {
		return
	}

	err = row.Scan(
		&buyer.ID,
		&buyer.CardNumberID,
		&buyer.FirstName,
		&buyer.LastName,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = e.ErrBuyerRepositoryNotFound
			return
		}

		return buyer, err
	}

	return buyer, nil
}

// Save saves the given buyer in the database
func (r *BuyerDB) Save(buyer *mod.Buyer) (err error) {
	result, err := r.db.Exec(
		"INSERT INTO buyers (id_card_number, first_name, last_name) "+
			"VALUES (?, ?, ?)",
		(*buyer).CardNumberID, (*buyer).FirstName, (*buyer).LastName,
	)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				return e.ErrBuyerRepositoryCardDuplicated
			default:
				return err
			}
		}

		return err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return err
	}
	(*buyer).ID = int(lastInsertId)
	return
}

// Update updates the given buyer in the database
func (r *BuyerDB) Update(buyer *mod.Buyer) (err error) {
	_, err = r.db.Exec(
		"UPDATE buyers "+
			"SET id_card_number=?, first_name=?, last_name=? WHERE id=?",
		(*buyer).CardNumberID, (*buyer).FirstName, (*buyer).LastName, (*buyer).ID,
	)

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				return e.ErrBuyerRepositoryCardDuplicated
			default:
				return err
			}
		}

		return err
	}

	return nil
}

// Delete deletes a buyer from the database by its id
func (r *BuyerDB) Delete(id int) (err error) {
	rows, err := r.db.Exec("DELETE FROM buyers WHERE id = ?", id)
	if err != nil {
		return err
	}

	result, _ := rows.RowsAffected()
	if result == 0 {
		return e.ErrBuyerRepositoryNotFound
	}

	return
}

func (r *BuyerDB) GetPurchaseOrderReport(id *int) (reports []mod.BuyerReportPO, err error) {
	var rows *sql.Rows
	query := "" +
		"SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(p.buyer_id) as purchase_orders_count " +
		"FROM buyers b " +
		"INNER JOIN purchase_orders p " +
		"ON p.buyer_id = b.id "

	if id != nil {
		query += "WHERE b.id = ? GROUP BY b.id"
		rows, err = r.db.Query(query, *id)
	} else {
		query += "GROUP BY b.id"
		rows, err = r.db.Query(query)
	}

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	found := false
	for rows.Next() {
		found = true
		var report mod.BuyerReportPO
		err = rows.Scan(
			&report.ID,
			&report.CardNumberID,
			&report.FirstName,
			&report.LastName,
			&report.PurchaseOrderCount,
		)

		if err != nil {
			return
		}

		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return
	}

	if !found && id != nil {
		return nil, e.ErrBuyerRepositoryNotFound
	}

	return reports, nil
}
