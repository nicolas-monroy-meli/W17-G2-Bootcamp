package repository

import (
	"fmt"
	"testing"

	"database/sql"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/require"
)

func setupWareHouseRepo(t *testing.T) (*warehouseRepository, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("SqlMock error")

	}
	repo := NewWarehouseRepository(db)
	close := func() {
		db.Close()
	}

	return repo, mock, close
}

func TestGetAll(t *testing.T) {
	repo, mock, close := setupWareHouseRepo(t)
	defer close()

	t.Run("success", func(t *testing.T) {
		expect := []models.Warehouse{
			{
				ID:                 1,
				WarehouseCode:      "WH001",
				Address:            "Calle Falsa 123",
				Telephone:          "123456789",
				MinimumCapacity:    100,
				MinimumTemperature: 25,
			},
			{
				ID:                 2,
				WarehouseCode:      "WH002",
				Address:            "Avenida Siempreviva 456",
				Telephone:          "987654321",
				MinimumCapacity:    200,
				MinimumTemperature: 30,
			},
		}

		// Configura el mock para retornar 2 filas
		rows := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature",
		}).
			AddRow(1, "WH001", "Calle Falsa 123", "123456789", 100, 25).
			AddRow(2, "WH002", "Avenida Siempreviva 456", "987654321", 200, 30)

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses
        `)).
			WillReturnRows(rows)

		result, err := repo.GetAll()
		require.NoError(t, err)
		require.Equal(t, expect, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("empty_result", func(t *testing.T) {
		// Configura el mock para retornar 0 filas
		rows := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature",
		})

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses
        `)).
			WillReturnRows(rows)

		result, err := repo.GetAll()
		require.NoError(t, err)
		require.Empty(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database_error", func(t *testing.T) {
		// Simula un error en la consulta
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses
        `)).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetAll()
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetByID(t *testing.T) {
	repo, mock, close := setupWareHouseRepo(t)
	defer close()

	t.Run("get_by_id_ok", func(t *testing.T) {
		expected := models.Warehouse{
			ID:                 1,
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		// Configura el mock para retornar una fila
		rows := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature",
		}).AddRow(
			expected.ID,
			expected.WarehouseCode,
			expected.Address,
			expected.Telephone,
			expected.MinimumCapacity,
			expected.MinimumTemperature,
		)

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses 
            WHERE id = ?
        `)).
			WithArgs(1).
			WillReturnRows(rows)

		warehouse, err := repo.GetByID(1)
		require.NoError(t, err)
		require.Equal(t, expected, warehouse)
		require.NoError(t, mock.ExpectationsWereMet()) // Verifica que todas las expectativas se cumplieron
	})

	t.Run("get_by_id_not_found", func(t *testing.T) {
		// Configura el mock para retornar "no rows"
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses 
            WHERE id = ?
        `)).
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		result, err := repo.GetByID(999)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrWarehouseRepositoryNotFound)
		require.Equal(t, models.Warehouse{}, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("get_by_id_database_error", func(t *testing.T) {
		// Simula un error genérico de la base de datos
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses 
            WHERE id = ?
        `)).
			WithArgs(1).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetByID(1)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Equal(t, models.Warehouse{}, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestSave(t *testing.T) {
	repo, mock, close := setupWareHouseRepo(t)
	defer close()

	t.Run("create_ok", func(t *testing.T) {
		res := models.Warehouse{
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		// Mock de SELECT EXISTS (no existe)
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs("WH001").
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		// Mock de INSERT exitoso
		mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO warehouses 
                (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) 
            VALUES (?, ?, ?, ?, ?)
        `)).
			WithArgs(
				res.WarehouseCode,
				res.Address,
				res.Telephone,
				res.MinimumCapacity,
				res.MinimumTemperature,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Save(&res)
		require.NoError(t, err)
		require.Equal(t, 1, res.ID)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create_conflict", func(t *testing.T) {
		req := models.Warehouse{
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		// Mock de SELECT EXISTS (ya existe)
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs("WH001").
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

		err := repo.Save(&req)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrWarehouseRepositoryDuplicated)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create_database_error_on_exists_check", func(t *testing.T) {
		req := models.Warehouse{
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		// Mock de SELECT EXISTS con error
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs("WH001").
			WillReturnError(fmt.Errorf("database error"))

		err := repo.Save(&req)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create_database_error_on_insert", func(t *testing.T) {
		req := models.Warehouse{
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		// Mock de SELECT EXISTS (no existe)
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs("WH001").
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		// Mock de INSERT con error
		mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO warehouses 
                (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) 
            VALUES (?, ?, ?, ?, ?)
        `)).
			WithArgs(
				req.WarehouseCode,
				req.Address,
				req.Telephone,
				req.MinimumCapacity,
				req.MinimumTemperature,
			).
			WillReturnError(fmt.Errorf("database error"))

		err := repo.Save(&req)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create_error_on_last_insert_id", func(t *testing.T) {
		req := models.Warehouse{
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		// Mock de SELECT EXISTS (no existe)
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs("WH001").
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		// Mock de INSERT exitoso, pero error al obtener el ID
		mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO warehouses 
                (warehouse_code, address, telephone, minimum_capacity, minimum_temperature) 
            VALUES (?, ?, ?, ?, ?)
        `)).
			WithArgs(
				req.WarehouseCode,
				req.Address,
				req.Telephone,
				req.MinimumCapacity,
				req.MinimumTemperature,
			).
			WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error getting last insert id")))

		err := repo.Save(&req)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestUpdate(t *testing.T) {
	repo, mock, close := setupWareHouseRepo(t)
	defer close()

	t.Run("update_ok", func(t *testing.T) {
		res := models.Warehouse{
			ID:                 1,
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		mock.ExpectExec(regexp.QuoteMeta(`
            UPDATE warehouses 
            SET 
                warehouse_code = ?, 
                address = ?, 
                telephone = ?, 
                minimum_capacity = ?,
                minimum_temperature = ?
            WHERE id = ?
        `)).
			WithArgs(
				res.WarehouseCode,
				res.Address,
				res.Telephone,
				res.MinimumCapacity,
				res.MinimumTemperature,
				res.ID,
			).
			WillReturnResult(sqlmock.NewResult(1, 1)) // 1 fila afectada

		err := repo.Update(&res)
		require.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("update_not_found", func(t *testing.T) {
		req := models.Warehouse{
			ID:                 999, // ID inexistente
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		mock.ExpectExec(regexp.QuoteMeta(`
            UPDATE warehouses 
            SET 
                warehouse_code = ?, 
                address = ?, 
                telephone = ?, 
                minimum_capacity = ?,
                minimum_temperature = ?
            WHERE id = ?
        `)).
			WithArgs(
				req.WarehouseCode,
				req.Address,
				req.Telephone,
				req.MinimumCapacity,
				req.MinimumTemperature,
				req.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 0)) // 0 filas afectadas

		err := repo.Update(&req)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrWarehouseRepositoryNotFound)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("update_foreignKeyViolation", func(t *testing.T) {
		req := models.Warehouse{
			ID:                 1,
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		mock.ExpectExec(regexp.QuoteMeta(`
            UPDATE warehouses 
            SET 
                warehouse_code = ?, 
                address = ?, 
                telephone = ?, 
                minimum_capacity = ?,
                minimum_temperature = ?
            WHERE id = ?
        `)).
			WithArgs(
				req.WarehouseCode,
				req.Address,
				req.Telephone,
				req.MinimumCapacity,
				req.MinimumTemperature,
				req.ID,
			).
			WillReturnError(&mysql.MySQLError{
				Number:  1452,
				Message: "FOREIGN KEY constraint fails",
			})

		err := repo.Update(&req)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("update_database_error", func(t *testing.T) {
		req := models.Warehouse{
			ID:                 1,
			WarehouseCode:      "WH001",
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		mock.ExpectExec(regexp.QuoteMeta(`
            UPDATE warehouses 
            SET 
                warehouse_code = ?, 
                address = ?, 
                telephone = ?, 
                minimum_capacity = ?,
                minimum_temperature = ?
            WHERE id = ?
        `)).
			WithArgs(
				req.WarehouseCode,
				req.Address,
				req.Telephone,
				req.MinimumCapacity,
				req.MinimumTemperature,
				req.ID,
			).
			WillReturnError(fmt.Errorf("database error"))

		err := repo.Update(&req)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	rp, mock, close := setupWareHouseRepo(t)
	defer close()

	t.Run("delete_ok", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM warehouses WHERE id = ?`)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := rp.Delete(1)
		require.NoError(t, err)
	})

	t.Run("delete_not found", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM warehouses WHERE id = ?`)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 0)).
			WillReturnError(sql.ErrNoRows)

		err := rp.Delete(1)

		require.Error(t, err)
	})

	t.Run("delete_no rows affected", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM warehouses WHERE id = ?`)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := rp.Delete(1)

		require.Error(t, err)
	})
}

func TestExistsWarehouseCode(t *testing.T) {
	repo, mock, close := setupWareHouseRepo(t)
	defer close()

	t.Run("exists_true", func(t *testing.T) {
		code := "WH001"

		// Mock: Retorna true
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs(code).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

		exists, err := repo.ExistsWarehouseCode(code)
		require.NoError(t, err)
		require.True(t, exists)
	})

	t.Run("exists_false", func(t *testing.T) {
		code := "WH002"

		// Mock: Retorna false
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs(code).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		exists, err := repo.ExistsWarehouseCode(code)
		require.NoError(t, err)
		require.False(t, exists)
	})

	t.Run("database_error", func(t *testing.T) {
		code := "WH003"

		// Mock: Retorna error
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`,
		)).
			WithArgs(code).
			WillReturnError(fmt.Errorf("database error"))

		exists, err := repo.ExistsWarehouseCode(code)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.False(t, exists)
	})
}

func TestGetByWarehouseCode(t *testing.T) {
	repo, mock, close := setupWareHouseRepo(t)
	defer close()

	t.Run("get_by_code_ok", func(t *testing.T) {
		code := "WH001"
		expected := models.Warehouse{
			ID:                 1,
			WarehouseCode:      code,
			Address:            "Calle Falsa 123",
			Telephone:          "123456789",
			MinimumCapacity:    100,
			MinimumTemperature: 25,
		}

		// Mock: Retorna una fila
		rows := sqlmock.NewRows([]string{
			"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature",
		}).
			AddRow(
				expected.ID,
				expected.WarehouseCode,
				expected.Address,
				expected.Telephone,
				expected.MinimumCapacity,
				expected.MinimumTemperature,
			)

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses 
            WHERE warehouse_code = ?
        `)).
			WithArgs(code).
			WillReturnRows(rows)

		result, err := repo.GetByWarehouseCode(code)
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("get_by_code_not_found", func(t *testing.T) {
		code := "WH999"

		// Mock: Retorna error "no rows"
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses 
            WHERE warehouse_code = ?
        `)).
			WithArgs(code).
			WillReturnError(sql.ErrNoRows)

		result, err := repo.GetByWarehouseCode(code)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrWarehouseRepositoryNotFound)
		require.Equal(t, models.Warehouse{}, result)
	})

	t.Run("get_by_code_database_error", func(t *testing.T) {
		code := "WH001"

		// Mock: Retorna error genérico
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature 
            FROM warehouses 
            WHERE warehouse_code = ?
        `)).
			WithArgs(code).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetByWarehouseCode(code)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Equal(t, models.Warehouse{}, result)
	})
}
