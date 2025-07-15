package repository

import (
	"database/sql"
	"fmt"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
)

type warehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) *warehouseRepository {
	return &warehouseRepository{db: db}
}

// GetAll devuelve un slice de warehouses
func (r *warehouseRepository) GetAll() ([]models.Warehouse, error) {
	query := `
		SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
		FROM warehouses
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var wh models.Warehouse
		if err := rows.Scan(
			&wh.ID,
			&wh.WarehouseCode,
			&wh.Address,
			&wh.Telephone,
			&wh.MinimumCapacity,
			&wh.MinimumTemperature,
		); err != nil {
			return nil, fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
		}
		warehouses = append(warehouses, wh)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}

	return warehouses, nil
}

// GetByID sin contexto
func (r *warehouseRepository) GetByID(id int) (models.Warehouse, error) {
	query := `
		SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
		FROM warehouses 
		WHERE id = ?
	`

	var wh models.Warehouse
	err := r.db.QueryRow(query, id).Scan(
		&wh.ID,
		&wh.WarehouseCode,
		&wh.Address,
		&wh.Telephone,
		&wh.MinimumCapacity,
		&wh.MinimumTemperature,
	)

	switch {
	case err == sql.ErrNoRows:
		return models.Warehouse{}, utils.ErrWarehouseRepositoryNotFound
	case err != nil:
		return models.Warehouse{}, fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}

	return wh, nil
}

// Save sin validaciones (hechas en el servicio)
func (r *warehouseRepository) Save(wh *models.Warehouse) error {
	exists, err := r.ExistsWarehouseCode(wh.WarehouseCode)
	if err != nil {
		return err
	}
	if exists {
		return utils.ErrWarehouseRepositoryDuplicated
	}

	query := `
		INSERT INTO warehouses 
			(warehouse_code, address, telephone, minimum_capacity, minimum_temperature) 
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(query,
		wh.WarehouseCode,
		wh.Address,
		wh.Telephone,
		wh.MinimumCapacity,
		wh.MinimumTemperature,
	)
	if err != nil {
		return fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}

	wh.ID = int(id)
	return nil
}

// Update sin contexto
func (r *warehouseRepository) Update(wh *models.Warehouse) error {
	query := `
		UPDATE warehouses 
		SET 
			warehouse_code = ?, 
			address = ?, 
			telephone = ?, 
			minimum_capacity = ?,
			minimum_temperature = ?
		WHERE id = ?
	`

	result, err := r.db.Exec(query,
		wh.WarehouseCode,
		wh.Address,
		wh.Telephone,
		wh.MinimumCapacity,
		wh.MinimumTemperature,
		wh.ID,
	)
	if err != nil {
		return fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.ErrWarehouseRepositoryNotFound
	}

	return nil
}

// Delete sin contexto
func (r *warehouseRepository) Delete(id int) error {
	query := `DELETE FROM warehouses WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return utils.ErrWarehouseRepositoryNotFound
	}

	return nil
}

// ExistsWarehouseCode verifica si el código ya existe
func (r *warehouseRepository) ExistsWarehouseCode(code string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`
	err := r.db.QueryRow(query, code).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}
	return exists, nil
}
func (r *warehouseRepository) GetByWarehouseCode(code string) (models.Warehouse, error) {
	query := `
		SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
		FROM warehouses 
		WHERE warehouse_code = ?
	`

	var wh models.Warehouse
	err := r.db.QueryRow(query, code).Scan(
		&wh.ID,
		&wh.WarehouseCode,
		&wh.Address,
		&wh.Telephone,
		&wh.MinimumCapacity,
		&wh.MinimumTemperature,
	)

	switch {
	case err == sql.ErrNoRows:
		return models.Warehouse{}, utils.ErrWarehouseRepositoryNotFound
	case err != nil:
		return models.Warehouse{}, fmt.Errorf("%w: %v", utils.ErrRepositoryDatabase, err)
	}

	return wh, nil
}
