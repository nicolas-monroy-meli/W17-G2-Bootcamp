package repository

import (
	"database/sql"
	"fmt"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

type carryRepository struct {
	db *sql.DB
}

func NewCarryRepository(db *sql.DB) *carryRepository {
	return &carryRepository{db: db}
}

func (r *carryRepository) GetAll() ([]models.Carry, error) {
	query := `
		SELECT id, cid, locality_id, company_name, address, telephone 
		FROM carries
	`

	rows, err := r.db.Query(query) // Query sin contexto
	if err != nil {
		return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}
	defer rows.Close()

	var carries []models.Carry
	for rows.Next() {
		var c models.Carry
		if err := rows.Scan(
			&c.ID,
			&c.CID,
			&c.LocalityID,
			&c.CompanyName,
			&c.Address,
			&c.Telephone,
		); err != nil {
			return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
		}
		carries = append(carries, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	return carries, nil
}

// GetByID sin contexto
func (r *carryRepository) GetByID(id int) (models.Carry, error) {
	query := `
		SELECT id, cid, locality_id, company_name, address, telephone 
		FROM carries 
		WHERE id = ?
	`

	var c models.Carry
	err := r.db.QueryRow(query, id).Scan(
		&c.ID,
		&c.CID,
		&c.LocalityID,
		&c.CompanyName,
		&c.Address,
		&c.Telephone,
	)

	switch {
	case err == sql.ErrNoRows:
		return models.Carry{}, e.ErrCarryRepositoryNotFound
	case err != nil:
		return models.Carry{}, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	return c, nil
}

// Save sin contexto
func (r *carryRepository) Save(c *models.Carry) error {
	query := `
		INSERT INTO carries 
			(cid, locality_id, company_name, address, telephone) 
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, // Exec sin contexto
		c.CID,
		c.LocalityID,
		c.CompanyName,
		c.Address,
		c.Telephone,
	)
	if err != nil {
		return fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	c.ID = int(id)
	return nil
}

// Update
func (r *carryRepository) Update(c *models.Carry) error {
	existingCarry, err := r.GetByCID(c.CID)
	if err != nil && err != e.ErrCarryRepositoryNotFound {
		return fmt.Errorf("error checking CID: %w", err)
	}

	if existingCarry.ID != 0 && existingCarry.ID != c.ID {
		return e.ErrCarryRepositoryDuplicated
	}

	query := `
			UPDATE carries 
			SET 
				cid = ?, 
				locality_id = ?, 
				company_name = ?, 
				address = ?, 
				telephone = ? 
			WHERE id = ?
		`

	result, err := r.db.Exec(query, // Exec
		c.CID,
		c.LocalityID,
		c.CompanyName,
		c.Address,
		c.Telephone,
		c.ID,
	)
	if err != nil {
		return fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return e.ErrCarryRepositoryNotFound
	}

	return nil
}

func (r *carryRepository) Delete(id int) error {
	query := `DELETE FROM carries WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return e.ErrCarryRepositoryNotFound
	}

	return nil
}

// GetReportByLocality
func (r *carryRepository) GetReportByLocality(localityID int) ([]models.LocalityCarryReport, error) {
	query := `
		SELECT 
			l.id, 
			l.locality_name, 
			COUNT(c.id) AS carries_count
		FROM localities l
		LEFT JOIN carries c ON l.id = c.locality_id
		WHERE l.id = ?
		GROUP BY l.id, l.locality_name;
	`

	rows, err := r.db.Query(query, localityID) // Query
	if err != nil {
		return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}
	defer rows.Close()

	var reports []models.LocalityCarryReport
	for rows.Next() {
		var report models.LocalityCarryReport
		if err := rows.Scan(
			&report.LocalityID,
			&report.LocalityName,
			&report.CarriesCount,
		); err != nil {
			return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	return reports, nil
}

func (r *carryRepository) GetReportByLocalityAll() ([]models.LocalityCarryReport, error) {
	query := `
		SELECT 
			l.id, 
			l.locality_name, 
			COUNT(c.id) AS carries_count
		FROM localities l
		LEFT JOIN carries c ON l.id = c.locality_id
		GROUP BY l.id, l.locality_name;
	`

	rows, err := r.db.Query(query) // Query
	if err != nil {
		return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}
	defer rows.Close()

	var reports []models.LocalityCarryReport
	for rows.Next() {
		var report models.LocalityCarryReport
		if err := rows.Scan(
			&report.LocalityID,
			&report.LocalityName,
			&report.CarriesCount,
		); err != nil {
			return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
		}
		reports = append(reports, report)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	return reports, nil
}

func (r *carryRepository) ExistsLocality(localityID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM localities WHERE id = ?)`
	err := r.db.QueryRow(query, localityID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}
	return exists, nil
}

// ExistsCID
func (r *carryRepository) ExistsCID(cid string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM carries WHERE cid = ?)`
	err := r.db.QueryRow(query, cid).Scan(&exists) // QueryRow
	if err != nil {
		return false, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}
	return exists, nil
}

// GetByCID
func (r *carryRepository) GetByCID(cid string) (models.Carry, error) {
	query := `
		SELECT id, cid, locality_id, company_name, address, telephone 
		FROM carries 
		WHERE cid = ?
	`

	var c models.Carry
	err := r.db.QueryRow(query, cid).Scan(
		&c.ID,
		&c.CID,
		&c.LocalityID,
		&c.CompanyName,
		&c.Address,
		&c.Telephone,
	)

	switch {
	case err == sql.ErrNoRows:
		return models.Carry{}, e.ErrCarryRepositoryNotFound
	case err != nil:
		return models.Carry{}, fmt.Errorf("%w: %v", e.ErrRepositoryDatabase, err)
	}

	return c, nil
}
