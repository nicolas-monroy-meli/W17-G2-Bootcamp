package repository

import (
	"database/sql"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewProductRecordRepo creates a new instance of the Product Record repository
func NewProductRecordRepo(db *sql.DB) *ProductRecordDB {
	return &ProductRecordDB{
		db: db,
	}
}

// ProductRecordDB is the implementation of the Product Record database
type ProductRecordDB struct {
	db *sql.DB
}

// FindAllPR returns all product records from the database
func (r *ProductRecordDB) FindAllPR() (productRecords map[int]mod.ProductRecord, err error) {
	rows, err := r.db.Query("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records;")
	if err != nil {
		return nil, e.ErrProductRepositoryNotFound
	}
	defer rows.Close()
	for rows.Next() {
		var productRecord mod.ProductRecord
		if err := rows.Scan(&productRecord.ID, &productRecord.LastUpdateDate, &productRecord.PurchasePrice, &productRecord.SalePrice, &productRecord.ProductID); err != nil {
			return nil, e.ErrProductRecordRepositoryNotFound
		}
		if productRecords == nil {
			productRecords = make(map[int]mod.ProductRecord)
		}
		productRecords[productRecord.ID] = productRecord
	}
	if err := rows.Err(); err != nil {
		return nil, e.ErrProductRecordRepositoryNotFound
	}
	if len(productRecords) == 0 {
		return nil, e.ErrProductRecordRepositoryNotFound
	}
	return
}

// FindAllByProductIDPR returns all product records from the database by product id
func (r *ProductRecordDB) FindAllByProductIDPR(productID int) (productRecords map[int]mod.ProductRecord, err error) {
	rows, err := r.db.Query("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records WHERE product_id = ?;", productID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var productRecord mod.ProductRecord
		if err := rows.Scan(&productRecord.ID, &productRecord.LastUpdateDate, &productRecord.PurchasePrice, &productRecord.SalePrice, &productRecord.ProductID); err != nil {
			return nil, e.ErrProductRecordRepositoryNotFound
		}
		if productRecords == nil {
			productRecords = make(map[int]mod.ProductRecord)
		}
		productRecords[productRecord.ID] = productRecord
	}
	if err := rows.Err(); err != nil {
		return nil, e.ErrProductRecordRepositoryNotFound
	}
	if len(productRecords) == 0 {
		return nil, e.ErrProductRecordRepositoryNotFound
	}
	return
}

// SavePR saves a product record into the database
func (r *ProductRecordDB) SavePR(productRecord *mod.ProductRecord) (err error) {
	result, err := r.db.Exec("INSERT INTO frescos_db.product_records (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES(?, ?, ?, ?);",
		(*productRecord).LastUpdateDate,
		(*productRecord).PurchasePrice,
		(*productRecord).SalePrice,
		(*productRecord).ProductID,
	)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	(*productRecord).ID = int(id)
	return
}
