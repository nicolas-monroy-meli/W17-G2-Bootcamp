package repository

import (
	"github.com/smartineztri_meli/W17-G2-Bootcamp/docs"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
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

var filePath = "/buyers.json"

// FindAll returns all buyers from the database
func (r *BuyerDB) FindAll() (buyers map[int]mod.Buyer, err error) {
	buyers = r.db
	if len(buyers) == 0 {

	}

	return
}

// FindByID returns a buyer from the database by its id
func (r *BuyerDB) FindByID(id int) (buyer mod.Buyer, err error) {
	_, ok := r.db[id]

	if !ok {
		err = utils.ErrBuyerRepositoryNotFound
		return
	}

	return r.db[id], nil
}

// Save saves the given buyer in the database
func (r *BuyerDB) Save(buyer *mod.Buyer) (err error) {
	_, ok := r.db[buyer.ID]
	id := 1

	if ok {
		err = utils.ErrBuyerRepositoryDuplicated
		return
	}

	for _, b := range r.db {
		if b.CardNumberID == buyer.CardNumberID {
			err = utils.ErrBuyerRepositoryCardDuplicated
			return
		}
		if b.ID > id {
			id = b.ID
		}
	}

	(*buyer).ID = id + 1
	r.db[buyer.ID] = *buyer
	err = docs.WriterFile(filePath, r.db)
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
