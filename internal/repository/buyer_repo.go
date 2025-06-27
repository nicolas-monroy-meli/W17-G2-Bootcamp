package repository

import (
	mod "github.com/nicolas-monroy-meli/W17-G2-Bootcamp/pkg/models"
)

// NewBuyerRepo creates a new instance of the Buyer repository
func NewBuyerRepo(buyers map[int]mod.Buyer) *BuyerDB {
	return &BuyerDB{
		db: buyers,
	}
}

// BuyerDB is the implementation of the Buyer database
type BuyerDB struct {
	db map[int]mod.Buyer
}

// FindAll returns all buyers from the database
func (r *BuyerDB) FindAll() (buyers map[int]mod.Buyer, err error) {

	return
}

// FindByID returns a buyer from the database by its id
func (r *BuyerDB) FindByID(id int) (buyer mod.Buyer, err error) {

	return
}

// Save saves the given buyer in the database
func (r *BuyerDB) Save(buyer *mod.Buyer) (err error) {

	return
}

// Update updates the given buyer in the database
func (r *BuyerDB) Update(buyer *mod.Buyer) (err error) {

	return
}

// Delete deletes a buyer from the database by its id
func (r *BuyerDB) Delete(id int) (err error) {

	return
}
