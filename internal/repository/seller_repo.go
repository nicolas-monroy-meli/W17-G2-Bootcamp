package repository

import (
	"github.com/smartineztri_meli/W17-G2-Bootcamp/docs"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewSellerRepo creates a new instance of the Seller repository
func NewSellerRepo(sellers map[int]mod.Seller) *SellerDB {
	return &SellerDB{
		db: sellers,
	}
}

// SellerDB is the implementation of the Seller database
type SellerDB struct {
	db map[int]mod.Seller
}

// FindAll returns all sellers from the database
func (r *SellerDB) FindAll() (sellers map[int]mod.Seller, err error) {
	result := r.db
	if len(r.db) == 0 {
		return nil, e.ErrSellerRepositoryNotFound
	}
	return result, nil
}

// FindByID returns a seller from the database by its id
func (r *SellerDB) FindByID(id int) (seller mod.Seller, err error) {
	val, ok := r.db[id]
	if !ok {
		return mod.Seller{}, e.ErrSellerRepositoryNotFound
	}
	return val, nil
}

// Save saves a seller into the database
func (r *SellerDB) Save(seller *mod.Seller) (err error) {
	for _, v := range r.db {
		if v.CID == seller.CID {
			return e.ErrSellerRepositoryDuplicated
		}
	}
	seller.ID = len(r.db) + 1
	r.db[seller.ID] = *seller
	docs.WriterFile("sellers.json", r.db)
	return nil
}

// Update updates a seller in the database
func (r *SellerDB) Update(seller *mod.Seller) (err error) {
	r.db[seller.ID] = *seller
	err = docs.WriterFile("sellers.json", r.db)
	return err
}

// Delete deletes a seller from the database
func (r *SellerDB) Delete(id int) (err error) {
	_, exists := r.db[id]
	if !exists {
		return e.ErrSellerRepositoryNotFound
	}
	delete(r.db, id)
	docs.WriterFile("sellers.json", r.db)
	return
}
