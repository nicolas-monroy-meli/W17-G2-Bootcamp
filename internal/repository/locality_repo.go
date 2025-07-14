package repository

import (
	"database/sql"
	"errors"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// NewLocalityRepo creates a new instance of the Locality repository
func NewLocalityRepo(localities *sql.DB) *LocalityDB {
	return &LocalityDB{
		db: localities,
	}
}

// LocalityDB is the implementation of the Locality database
type LocalityDB struct {
	db *sql.DB
}

// FindByID returns a seller from the database by its id
func (r *LocalityDB) FindSellersByLocality() (result []models.SelByLoc, err error) {
	rows, err := r.db.Query("SELECT l.id, l.locality_name, count(*) FROM localities AS l INNER JOIN sellers as s ON l.id=s.locality_id GROUP BY l.id")
	if err != nil {
		return nil, e.ErrQueryError
	}
	defer rows.Close()

	for rows.Next() {
		var locality models.SelByLoc
		err = rows.Scan(&locality.ID, &locality.Name, &locality.Count)
		if err != nil {
			return nil, e.ErrQueryError
		}
		result = append(result, locality)
	}
	return result, nil
}

func (r *LocalityDB) FindSellersByLocID(id int) (result models.SelByLoc, err error) {
	row := r.db.QueryRow("SELECT l.id, l.locality_name, count(*) FROM localities AS l INNER JOIN sellers as s ON l.id=s.locality_id GROUP BY l.id HAVING l.id= ?", id)
	if row.Err() != nil {
		return models.SelByLoc{}, e.ErrQueryError
	}
	err = row.Scan(&result.ID, &result.Name, &result.Count)
	if err != nil {
		if result.ID == 0 {
			return models.SelByLoc{}, e.ErrForeignKeyError
		}
		if errors.Is(err, sql.ErrNoRows) {
			return result, nil
		}
	}

	return result, nil
}

/*
// Save saves a seller into the database
func (r *LocalityDB) Save(seller *mod.Locality) (id int, err error) {
	err = r.findByCID(seller.CID)
	if err == nil {
		return 0, e.ErrLocalityRepositoryDuplicated
	}

	result, err := r.db.Exec("INSERT INTO `sellers`(`cid`,`company_name`,`address`,`telephone`,`locality_id`) VALUES(?,?,?,?,?)", seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality)
	if err != nil {
		return 0, e.ErrForeignKeyError
	}
	id64, _ := result.LastInsertId()
	id = int(id64)
	return id, nil
}
*/
