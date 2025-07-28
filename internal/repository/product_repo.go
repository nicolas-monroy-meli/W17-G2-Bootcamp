package repository

import (
	"database/sql"
	"strings"

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

// FindAll returns all products from the database - TESTED
func (r *ProductDB) FindAll() (products []mod.Product, err error) {
	rows, err := r.db.Query("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM products;")
	if err != nil {
		return nil, e.ErrProductRepositoryNotFound
	}
	defer rows.Close()
	for rows.Next() {
		var product mod.Product
		if err := rows.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Height, &product.Length, &product.Width, &product.Weight, &product.ExpirationRate, &product.FreezingRate, &product.RecomFreezTemp, &product.ProductTypeID, &product.SellerID); err != nil {
			return nil, e.ErrProductRepositoryNotFound
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, e.ErrProductRepositoryNotFound
	}
	if len(products) == 0 {
		return nil, e.ErrProductRepositoryNotFound
	}
	return products, nil
}

// FindByID returns a product from the database by its id - TESTED
func (r *ProductDB) FindByID(id int) (product mod.Product, err error) {
	row := r.db.QueryRow("SELECT `id`, `product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id` FROM frescos_db.products WHERE id = ?;", id)
	if err := row.Scan(&product.ID, &product.ProductCode, &product.Description, &product.Height, &product.Length, &product.Width, &product.Weight, &product.ExpirationRate, &product.FreezingRate, &product.RecomFreezTemp, &product.ProductTypeID, &product.SellerID); err != nil {
		return mod.Product{}, e.ErrProductRepositoryNotFound
	}
	if product.ID == 0 {
		return mod.Product{}, e.ErrProductRepositoryNotFound
	}
	return product, nil
}

// Save saves a product into the database - TESTED
func (r *ProductDB) Save(product *mod.Product) (err error) {
	if _, exists := r.FindByID(product.ID); exists == nil {
		err = e.ErrProductRepositoryDuplicated
		return
	}
	result, err := r.db.Exec("INSERT INTO frescos_db.products (`product_code`, `description`, `height`, `length`, `width`, `net_weight`, `expiration_rate`, `freezing_rate`, `recommended_freezing_temperature`, `product_type_id`, `seller_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
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
		if strings.Contains(err.Error(), "foreign key constraint fails") && strings.Contains(err.Error(), "products_ibfk_1") {
			return e.ErrSellerRepositoryNotFound
		}
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
	_, err = r.db.Exec("UPDATE frescos_db.products SET `product_code` = ?, `description` = ?, `height` = ?, `length` = ?, `width` = ?, `net_weight` = ?, `expiration_rate` = ?, `freezing_rate` = ?, `recommended_freezing_temperature` = ?, `product_type_id` = ?, `seller_id` = ? WHERE id = ?;",
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
		if strings.Contains(err.Error(), "foreign key constraint fails") && strings.Contains(err.Error(), "products_ibfk_1") {
			return e.ErrSellerRepositoryNotFound
		}
		return
	}
	return
}

// Delete deletes a product from the database by its id
func (r *ProductDB) Delete(id int) (err error) {
	if _, exists := r.FindByID(id); exists != nil {
		return e.ErrProductRepositoryNotFound
	}
	_, err = r.db.Exec("DELETE FROM frescos_db.products WHERE id = ?;", id)
	if err != nil {
		return
	}
	return nil
}
