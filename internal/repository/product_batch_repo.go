package repository

import (
	"database/sql"
	"fmt"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
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
	}

	if len(batches) == 0 {
		return nil, errors.ErrEmptySectionDB
	}
	return batches, nil
}

func (r *ProductBatchDB) BatchExists(id int, batchNumber *int) (res bool, err error) {
	var exists bool
	if batchNumber != nil {
		query := `SELECT EXISTS (SELECT 1 FROM product_batches WHERE id = ? or batch_number=?)`
		err = r.db.QueryRow(query, id, *batchNumber).Scan(&exists)

	} else {
		query := `SELECT EXISTS (SELECT 1 FROM product_batches WHERE id = ?)`
		err = r.db.QueryRow(query, id).Scan(&exists)
	}
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *ProductBatchDB) Save(batch *mod.ProductBatch) (err error) {
	if exists, _ := r.BatchExists(batch.ID, &batch.BatchNumber); !exists {
		return errors.ErrProductBatchDuplicated
	}
	result, err := r.db.Exec("INSERT INTO `product_batches` (`batch_number`,`current_quantity`,`initial_quantity`,`current_temperature`, `minimum_temperature`, `due_date`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id`) VALUES(?,?,?,?,?,?,?,?,?,?)",
		(*batch).BatchNumber, (*batch).CurrentQuantity, (*batch).InitialQuantity, (*batch).CurrentTemperature, (*batch).MinimumTemperature, (*batch).DueDate, (*batch).ManufacturingDate, (*batch).ManufacturingHour, (*batch).ProductId, (*batch).SectionId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	(*batch).ID = int(id)
	return
}
