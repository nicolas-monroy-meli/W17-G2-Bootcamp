package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

type ProductBatchDB struct {
	db *sql.DB
}

func NewProductBatchRepo(db *sql.DB) *ProductBatchDB {
	return &ProductBatchDB{
		db: db,
	}
}

func (r *ProductBatchDB) FindAll() (batches []mod.ProductBatch, err error) {
	rows, err := r.db.Query("SELECT `id`,`batch_number`, `current_quantity`, `initial_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` ")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var batch mod.ProductBatch
		err = rows.Scan(&batch.ID, &batch.BatchNumber, &batch.CurrentQuantity, &batch.InitialQuantity, &batch.CurrentTemperature, &batch.MinimumTemperature, &batch.DueDate, &batch.ManufacturingDate, &batch.ManufacturingHour, &batch.ProductId, &batch.SectionId)
		if err != nil {
			return nil, err
		}
		batches = append(batches, batch)
	}

	if len(batches) == 0 {
		return nil, e.ErrEmptySectionDB
	}
	return batches, nil
}

func (r *ProductBatchDB) Save(batch *mod.ProductBatch) (err error) {
	result, err := r.db.Exec("INSERT INTO `product_batches` (`batch_number`,`current_quantity`,`initial_quantity`,`current_temperature`, `minimum_temperature`, `due_date`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id`) VALUES(?,?,?,?,?,?,?,?,?,?)",
		(*batch).BatchNumber, (*batch).CurrentQuantity, (*batch).InitialQuantity, (*batch).CurrentTemperature, (*batch).MinimumTemperature, (*batch).DueDate, (*batch).ManufacturingDate, (*batch).ManufacturingHour, (*batch).ProductId, (*batch).SectionId)
	if err != nil {
		var mySQLErr *mysql.MySQLError
		if errors.As(err, &mySQLErr) {
			if mySQLErr.Number == 1452 {
				return e.ErrForeignKeyError
			}
			if mySQLErr.Number == 1062 {
				return e.ErrProductBatchDuplicated
			}
		}
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	(*batch).ID = int(id)
	return
}
