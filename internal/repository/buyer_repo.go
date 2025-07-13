package repository

import (
	"database/sql"
	"errors"
	"fmt"
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
		err := rows.Scan(&by.ID, &by.CardNumberID, &by.FirstName, &by.LastName)
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
		fmt.Println("error", err)
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
		fmt.Println(err)
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
	fmt.Println("Entro")
	_, err = r.db.Exec(
		"UPDATE buyers "+
			"SET id_card_number=?, first_name=?, last_name=? WHERE id=?",
		(*buyer).CardNumberID, (*buyer).FirstName, (*buyer).LastName, (*buyer).ID,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
	//id := buyer.ID
	//_, ok := r.db[id]
	//
	//if !ok {
	//	err = e.ErrBuyerRepositoryNotFound
	//	return
	//}
	//
	//for _, b := range r.db {
	//	if b.ID != buyer.ID && b.CardNumberID == buyer.CardNumberID {
	//		err = e.ErrBuyerRepositoryCardDuplicated
	//		return
	//	}
	//}
	//
	//r.db[id] = *buyer
	//err = docs.WriterFile(filePath, r.db)
}

// Delete deletes a buyer from the database by its id
func (r *BuyerDB) Delete(id int) (err error) {
	_, err = r.db.Exec("DELETE FROM buyers WHERE id = ?", id)

	if err != nil {
		return
	}
	return
}

func (r *BuyerDB) GetByCardNumber(cardNumber string) (buyer mod.Buyer, err error) {
	rowBuyer := r.db.QueryRow(
		"SELECT id, id_card_number FROM buyers WHERE id_card_number = ? ", cardNumber,
	)

	err = rowBuyer.Scan(
		&buyer.ID,
		&buyer.CardNumberID,
	)

	if err != nil {
		return
	}

	return buyer, nil
}
