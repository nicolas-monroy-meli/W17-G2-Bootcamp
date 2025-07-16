package repository

import (
	"database/sql"

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
func (r *LocalityDB) FindAllLocalities() (result []models.Locality, err error) {
	rows, err := r.db.Query("SELECT l.id, l.locality_name, l.province_name, l.country_name FROM localities AS l")
	if err != nil {
		return nil, e.ErrQueryError
	}
	defer rows.Close()

	for rows.Next() {
		var locality models.Locality
		err = rows.Scan(&locality.ID, &locality.Name, &locality.Province, &locality.Country)
		if err != nil {
			return nil, e.ErrQueryError
		}
		result = append(result, locality)
	}
	return result, nil
}

func (r *LocalityDB) FindSellersByLocID(id int) (result []models.SelByLoc, err error) {
	var rows *sql.Rows
	if id == -1 {
		rows, err = r.db.Query("SELECT l.id, l.locality_name, count(*) FROM localities AS l INNER JOIN sellers as s ON l.id=s.locality_id GROUP BY l.id")
		if err != nil {
			return nil, e.ErrQueryError
		}
		defer rows.Close()
	} else {
		rows, err = r.db.Query("SELECT l.id, l.locality_name, count(*) FROM localities AS l INNER JOIN sellers as s ON l.id=s.locality_id GROUP BY l.id HAVING l.id= ?", id)
		if err != nil {
			return nil, e.ErrQueryError
		}
		defer rows.Close()
	}
	for rows.Next() {
		var locality models.SelByLoc
		err = rows.Scan(&locality.ID, &locality.Name, &locality.Count)
		if err != nil {
			return nil, e.ErrQueryError
		}
		result = append(result, locality)
	}
	if len(result) == 0 {
		return nil, e.ErrLocalityRepositoryNotFound
	}
	return result, nil
}

// Save saves a locality into the database
func (r *LocalityDB) Save(locality *models.Locality) (id int, err error) {
	result, err := r.db.Exec("INSERT INTO `localities`(`locality_name`,`province_name`,`country_name`) VALUES(?,?,?)", locality.Name, locality.Province, locality.Country)
	if err != nil {
		return 0, e.ErrInsertError
	}
	id64, _ := result.LastInsertId()
	id = int(id64)
	return id, nil
}
