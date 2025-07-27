package repository

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewSellerRepo creates a new instance of the Seller repository
func NewSellerRepo(sellers *sql.DB) *SellerDB {
	return &SellerDB{
		db: sellers,
	}
}

// SellerDB is the implementation of the Seller database
type SellerDB struct {
	db *sql.DB
}

// FindAll returns all sellers from the database -TESTED
func (r *SellerDB) FindAll() (sellers []mod.Seller, err error) {
	rows, err := r.db.Query("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`")
	if err != nil {
		return nil, e.ErrQueryError
	}
	defer rows.Close()
	for rows.Next() {
		var seller mod.Seller
		err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.Locality)
		if err != nil {
			return nil, e.ErrParseError
		}
		sellers = append(sellers, seller)
	}
	if len(sellers) == 0 {
		return nil, e.ErrQueryIsEmpty
	}
	return
}

// FindByID returns a seller from the database by its id -TESTED
func (r *SellerDB) FindByID(id int) (seller mod.Seller, err error) {
	row := r.db.QueryRow("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers` WHERE `id` = ?", id)
	err = row.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone, &seller.Locality)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mod.Seller{}, e.ErrSellerRepositoryNotFound
		}
		return mod.Seller{}, errors.Join(e.ErrParseError, err)
	}
	return seller, nil
}

// Save saves a seller into the database -TESTED
func (r *SellerDB) Save(seller *mod.Seller) (id int, err error) {
	result, err := r.db.Exec("INSERT INTO `sellers`(`cid`,`company_name`,`address`,`telephone`,`locality_id`) VALUES(?,?,?,?,?)", seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality)
	if err != nil {
		var mySQLErr *mysql.MySQLError
		if errors.As(err, &mySQLErr) {
			if mySQLErr.Number == 1062 {
				return 0, e.ErrSellerRepositoryDuplicated
			}
			if mySQLErr.Number == 1452 {
				return 0, e.ErrForeignKeyError
			}
		}
		return 0, errors.Join(e.ErrInsertError, err)
	}
	id64, _ := result.LastInsertId()
	id = int(id64)
	return id, nil
}

// Update updates a seller in the database -TESTED
func (r *SellerDB) Update(seller *mod.Seller) (err error) {
	_, err = r.db.Exec("UPDATE `sellers` SET `cid`=?,`company_name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id`= ?", seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality, seller.ID)
	if err != nil {
		var mySQLErr *mysql.MySQLError
		if errors.As(err, &mySQLErr) {
			if mySQLErr.Number == 1452 {
				return e.ErrForeignKeyError
			}
			if mySQLErr.Number == 1062 {
				return e.ErrSellerRepositoryDuplicated
			}
		}
		return errors.Join(e.ErrRepositoryDatabase, err)
	}
	return nil
}

// Delete deletes a seller from the database -TESTED
func (r *SellerDB) Delete(id int) (err error) {
	rows, err := r.db.Exec("DELETE FROM `sellers` WHERE `id`=?", id)
	if err != nil {
		return errors.Join(e.ErrRepositoryDatabase, err)
	}
	result, _ := rows.RowsAffected()
	if result == 0 {
		return e.ErrSellerRepositoryNotFound
	}
	return nil
}
