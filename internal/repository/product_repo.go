package repository

import (
	"database/sql"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewProductRepo creates a new instance of the Product repository
func NewProductRepo(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

// ProductDB is the implementation of the Product database
type ProductDB struct {
	db *sql.DB
}

// FindAll returns all products from the database
func (r *ProductDB) FindAll() (products map[int]mod.Product, err error) {
	rows, err := r.db.Query("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM fresco_db.products;")
	if err != nil {
		return nil, e.ErrProductRepositoryNotFound
	}
	defer rows.Close()
	for rows.Next() {
		var product mod.Product
		if err := rows.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Height, &product.Length, &product.Width, &product.Weight, &product.ExpirationRate, &product.FreezingRate, &product.RecomFreezTemp, &product.ProductTypeID, &product.SellerID); err != nil {
			return nil, e.ErrProductRepositoryNotFound
		}
		if products == nil {
			products = make(map[int]mod.Product)
		}
		products[product.ID] = product
	}
	if err := rows.Err(); err != nil {
		return nil, e.ErrProductRepositoryNotFound
	}
	if len(products) == 0 {
		return nil, e.ErrProductRepositoryNotFound
	}
	return products, nil
}

// FindByID returns a product from the database by its id
func (r *ProductDB) FindByID(id int) (product mod.Product, err error) {
	row := r.db.QueryRow("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM fresco_db.products WHERE id = ?;", id)
	if err != nil {
		return mod.Product{}, e.ErrProductRepositoryNotFound
	}
	if err := row.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Height, &product.Length, &product.Width, &product.Weight, &product.ExpirationRate, &product.FreezingRate, &product.RecomFreezTemp, &product.ProductTypeID, &product.SellerID); err != nil {
		return mod.Product{}, e.ErrProductRepositoryNotFound
	}
	if product.ID == 0 {
		return mod.Product{}, e.ErrProductRepositoryNotFound
	}
	return product, nil
}

// Save saves a product into the database
func (r *ProductDB) Save(product *mod.Product) (err error) {
	if _, exists := r.FindByID(product.ID); exists == nil {
		err = e.ErrProductRepositoryDuplicated
		return
	}
	result, err := r.db.Exec("INSERT INTO fresco_db.products (`product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
		(*product).ProductCode,
		(*product).Description,
		(*product).Height,
		(*product).Length,
		(*product).Width,
		(*product).Weight,
		(*product).ExpirationRate,
		(*product).FreezingRate,
		(*product).RecomFreezTemp,
		(*product).ProductTypeID,
		(*product).SellerID,
	)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	(*product).ID = int(id)
	return
}

// Update updates a product in the database
func (r *ProductDB) Update(product *mod.Product) (err error) {
	/* if _, exists := r.FindByID(product.ID); exists != nil {
		return e.ErrProductRepositoryNotFound
	} */
	_, err = r.db.Exec("UPDATE fresco_db.products SET `product_code` = ?, `description` = ?, `height` = ?, `length` = ?, `width` = ?, `net_weight` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `recommended_freezing_temperature` = ?, `product_type_id` = ?, `seller_id` = ? WHERE id = ?;",
		(*product).ProductCode,
		(*product).Description,
		(*product).Height,
		(*product).Length,
		(*product).Width,
		(*product).Weight,
		(*product).ExpirationRate,
		(*product).FreezingRate,
		(*product).RecomFreezTemp,
		(*product).ProductTypeID,
		(*product).SellerID,
		(*product).ID,
	)
	if err != nil {
		return
	}
	return
}

// Delete deletes a product from the database by its id
func (r *ProductDB) Delete(id int) (err error) {
	if _, exists := r.FindByID(id); exists != nil {
		return e.ErrProductRepositoryNotFound
	}
	_, err = r.db.Exec("DELETE FROM fresco_db.products WHERE id = ?;", id)
	if err != nil {
		return
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////////////////////

// FindAllPR returns all product records from the database
func (r *ProductDB) FindAllPR() (productRecords map[int]mod.ProductRecord, err error) {
	rows, err := r.db.Query("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM fresco_db.product_records;")
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
func (r *ProductDB) FindAllByProductIDPR(productID int) (productRecords map[int]mod.ProductRecord, err error) {
	rows, err := r.db.Query("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM fresco_db.product_records WHERE product_id = ?;", productID)
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
func (r *ProductDB) SavePR(productRecord *mod.ProductRecord) (err error) {
	// If product does not exist, return error
	_, err = r.FindByID(productRecord.ProductID)
	if err != nil {
		err = e.ErrProductRepositoryNotFound
		return
	}
	result, err := r.db.Exec("INSERT INTO fresco_db.product_records (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES(?, ?, ?, ?);",
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
