package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/require"
)

func setupCarriesRepo(t *testing.T) (*carryRepository, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("SqlMock error")

	}
	repo := NewCarryRepository(db)
	close := func() {
		db.Close()
	}

	return repo, mock, close
}
func TestCarryRepository_GetAll(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("success", func(t *testing.T) {
		expect := []models.Carry{
			{
				ID:          1,
				CID:         "CID#100",
				CompanyName: "Company A",
				Address:     "Address 123",
				Telephone:   "123456789",
				LocalityID:  6700,
			},
			{
				ID:          2,
				CID:         "CID#200",
				CompanyName: "Company B",
				Address:     "Address 456",
				Telephone:   "987654321",
				LocalityID:  6701,
			},
		}

		// Configura el mock para retornar 2 filas
		rows := sqlmock.NewRows([]string{
			"id", "cid", "locality_id", "company_name", "address", "telephone",
		}).
			AddRow(1, "CID#100", 6700, "Company A", "Address 123", "123456789").
			AddRow(2, "CID#200", 6701, "Company B", "Address 456", "987654321")

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries
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
			"id", "cid", "locality_id", "company_name", "address", "telephone",
		})

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries
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
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries
        `)).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetAll()
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestCarryRepository_GetByID(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("success", func(t *testing.T) {
		expected := models.Carry{
			ID:          1,
			CID:         "CID#100",
			LocalityID:  6700,
			CompanyName: "Fast Logistics",
			Address:     "Calle Falsa 123",
			Telephone:   "123456789",
		}

		// Configura el mock para retornar 1 fila
		rows := sqlmock.NewRows([]string{
			"id", "cid", "locality_id", "company_name", "address", "telephone",
		}).AddRow(
			expected.ID,
			expected.CID,
			expected.LocalityID,
			expected.CompanyName,
			expected.Address,
			expected.Telephone,
		)

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE id = ?
        `)).
			WithArgs(1).
			WillReturnRows(rows)

		result, err := repo.GetByID(1)
		require.NoError(t, err)
		require.Equal(t, expected, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("not_found", func(t *testing.T) {
		// Simula que no se encuentra el registro
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE id = ?
        `)).
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		result, err := repo.GetByID(999)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrCarryRepositoryNotFound)
		require.Equal(t, models.Carry{}, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("database_error", func(t *testing.T) {
		// Simula un error genérico de la base de datos
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE id = ?
        `)).
			WithArgs(1).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetByID(1)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Equal(t, models.Carry{}, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCarryRepository_Save(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("save_success", func(t *testing.T) {
		carry := &models.Carry{
			CID:         "CID#100",
			LocalityID:  6700,
			CompanyName: "Fast Logistics",
			Address:     "Calle Falsa 123",
			Telephone:   "123456789",
		}

		// Configura el mock para una inserción exitosa
		mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO carries 
                (cid, locality_id, company_name, address, telephone) 
            VALUES (?, ?, ?, ?, ?)
        `)).
			WithArgs(
				carry.CID,
				carry.LocalityID,
				carry.CompanyName,
				carry.Address,
				carry.Telephone,
			).
			WillReturnResult(sqlmock.NewResult(1, 1)) // ID generado = 1

		err := repo.Save(carry)
		require.NoError(t, err)
		require.Equal(t, 1, carry.ID) // Verifica que el ID se actualizó
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("save_database_error", func(t *testing.T) {
		carry := &models.Carry{
			CID:         "CID#200",
			LocalityID:  6701,
			CompanyName: "Quick Delivery",
			Address:     "Avenida Siempreviva 456",
			Telephone:   "987654321",
		}

		// Simula un error en la inserción
		mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO carries 
                (cid, locality_id, company_name, address, telephone) 
            VALUES (?, ?, ?, ?, ?)
        `)).
			WithArgs(
				carry.CID,
				carry.LocalityID,
				carry.CompanyName,
				carry.Address,
				carry.Telephone,
			).
			WillReturnError(fmt.Errorf("database error"))

		err := repo.Save(carry)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("save_last_insert_id_error", func(t *testing.T) {
		carry := &models.Carry{
			CID:         "CID#300",
			LocalityID:  6702,
			CompanyName: "Speedy Transports",
			Address:     "Boulevard Los Olivos 789",
			Telephone:   "555555555",
		}

		// Simula un error al obtener el LastInsertId
		mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO carries 
                (cid, locality_id, company_name, address, telephone) 
            VALUES (?, ?, ?, ?, ?)
        `)).
			WithArgs(
				carry.CID,
				carry.LocalityID,
				carry.CompanyName,
				carry.Address,
				carry.Telephone,
			).
			WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error getting last insert id")))

		err := repo.Save(carry)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestCarryRepository_Update(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("update_success_new_cid", func(t *testing.T) {
		// Configurar el mock para GetByCID (no existe otro CID)
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE cid = ?
        `)).
			WithArgs("CID#100").
			WillReturnError(sql.ErrNoRows)

		// Mock para la actualización exitosa
		mock.ExpectExec(regexp.QuoteMeta(`
            UPDATE carries 
            SET 
                cid = ?, 
                locality_id = ?, 
                company_name = ?, 
                address = ?, 
                telephone = ? 
            WHERE id = ?
        `)).
			WithArgs("CID#100", 6700, "Fast Logistics", "Calle Falsa 123", "123456789", 1).
			WillReturnResult(sqlmock.NewResult(0, 1)) // 1 fila afectada

		carry := &models.Carry{
			ID:          1,
			CID:         "CID#100",
			LocalityID:  6700,
			CompanyName: "Fast Logistics",
			Address:     "Calle Falsa 123",
			Telephone:   "123456789",
		}

		err := repo.Update(carry)
		require.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("update_cid_conflict", func(t *testing.T) {
		// Mock para GetByCID: Retorna un carry con ID diferente
		rows := sqlmock.NewRows([]string{"id", "cid", "locality_id", "company_name", "address", "telephone"}).
			AddRow(2, "CID#100", 6700, "Conflict Logistics", "Address 456", "987654321")

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE cid = ?
        `)).
			WithArgs("CID#100").
			WillReturnRows(rows)

		carry := &models.Carry{
			ID:          1, // ID diferente al encontrado (2)
			CID:         "CID#100",
			LocalityID:  6700,
			CompanyName: "Fast Logistics",
			Address:     "Calle Falsa 123",
			Telephone:   "123456789",
		}

		err := repo.Update(carry)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrCarryRepositoryDuplicated)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("update_not_found", func(t *testing.T) {
		// Mock para GetByCID: No hay conflicto (mismo ID)
		rows := sqlmock.NewRows([]string{"id", "cid", "locality_id", "company_name", "address", "telephone"}).
			AddRow(1, "CID#100", 6700, "Fast Logistics", "Calle Falsa 123", "123456789")

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE cid = ?
        `)).
			WithArgs("CID#100").
			WillReturnRows(rows)

		// Mock para la actualización que no afecta filas
		mock.ExpectExec(regexp.QuoteMeta(`
            UPDATE carries 
            SET 
                cid = ?, 
                locality_id = ?, 
                company_name = ?, 
                address = ?, 
                telephone = ? 
            WHERE id = ?
        `)).
			WithArgs("CID#100", 6700, "Fast Logistics", "Calle Falsa 123", "123456789", 1).
			WillReturnResult(sqlmock.NewResult(0, 0)) // 0 filas afectadas

		carry := &models.Carry{
			ID:          1,
			CID:         "CID#100",
			LocalityID:  6700,
			CompanyName: "Fast Logistics",
			Address:     "Calle Falsa 123",
			Telephone:   "123456789",
		}

		err := repo.Update(carry)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrCarryRepositoryNotFound)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("update_database_error", func(t *testing.T) {
		// Mock para GetByCID: No hay conflicto
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE cid = ?
        `)).
			WithArgs("CID#100").
			WillReturnError(sql.ErrNoRows)

		// Mock para error en la actualización
		mock.ExpectExec(regexp.QuoteMeta(`
            UPDATE carries 
            SET 
                cid = ?, 
                locality_id = ?, 
                company_name = ?, 
                address = ?, 
                telephone = ? 
            WHERE id = ?
        `)).
			WithArgs("CID#100", 6700, "Fast Logistics", "Calle Falsa 123", "123456789", 1).
			WillReturnError(fmt.Errorf("database error"))

		carry := &models.Carry{
			ID:          1,
			CID:         "CID#100",
			LocalityID:  6700,
			CompanyName: "Fast Logistics",
			Address:     "Calle Falsa 123",
			Telephone:   "123456789",
		}

		err := repo.Update(carry)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCarryRepository_Delete(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("delete_success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`
            DELETE FROM carries 
            WHERE id = ?
        `)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1)) // 1 fila afectada

		err := repo.Delete(1)
		require.NoError(t, err)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("delete_not_found", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`
            DELETE FROM carries 
            WHERE id = ?
        `)).
			WithArgs(999).
			WillReturnResult(sqlmock.NewResult(0, 0)) // 0 filas afectadas

		err := repo.Delete(999)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrCarryRepositoryNotFound)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("delete_database_error", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(`
            DELETE FROM carries 
            WHERE id = ?
        `)).
			WithArgs(1).
			WillReturnError(fmt.Errorf("database error"))

		err := repo.Delete(1)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestCarryRepository_GetReportByLocality(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("report_by_locality_success", func(t *testing.T) {
		localityID := 6700
		expected := []models.LocalityCarryReport{
			{
				LocalityID:   localityID,
				LocalityName: "CABA",
				CarriesCount: 5,
			},
		}

		// Mock de la consulta con JOIN
		rows := sqlmock.NewRows([]string{"id", "locality_name", "carries_count"}).
			AddRow(localityID, "CABA", 5)

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT 
                l.id, 
                l.locality_name, 
                COUNT(c.id) AS carries_count
            FROM localities l
            LEFT JOIN carries c ON l.id = c.locality_id
            WHERE l.id = ?
            GROUP BY l.id, l.locality_name;
        `)).
			WithArgs(localityID).
			WillReturnRows(rows)

		result, err := repo.GetReportByLocality(localityID)
		require.NoError(t, err)
		require.Equal(t, expected, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("report_by_locality_not_found", func(t *testing.T) {
		localityID := 999
		rows := sqlmock.NewRows([]string{"id", "locality_name", "carries_count"})

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT 
                l.id, 
                l.locality_name, 
                COUNT(c.id) AS carries_count
            FROM localities l
            LEFT JOIN carries c ON l.id = c.locality_id
            WHERE l.id = ?
            GROUP BY l.id, l.locality_name;
        `)).
			WithArgs(localityID).
			WillReturnRows(rows)

		result, err := repo.GetReportByLocality(localityID)
		require.NoError(t, err)
		require.Empty(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("report_by_locality_database_error", func(t *testing.T) {
		localityID := 6700
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT 
                l.id, 
                l.locality_name, 
                COUNT(c.id) AS carries_count
            FROM localities l
            LEFT JOIN carries c ON l.id = c.locality_id
            WHERE l.id = ?
            GROUP BY l.id, l.locality_name;
        `)).
			WithArgs(localityID).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetReportByLocality(localityID)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCarryRepository_GetReportByLocalityAll(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("report_all_success", func(t *testing.T) {
		expected := []models.LocalityCarryReport{
			{
				LocalityID:   6700,
				LocalityName: "CABA",
				CarriesCount: 10,
			},
			{
				LocalityID:   6701,
				LocalityName: "Rosario",
				CarriesCount: 5,
			},
		}

		rows := sqlmock.NewRows([]string{"id", "locality_name", "carries_count"}).
			AddRow(6700, "CABA", 10).
			AddRow(6701, "Rosario", 5)

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT 
                l.id, 
                l.locality_name, 
                COUNT(c.id) AS carries_count
            FROM localities l
            LEFT JOIN carries c ON l.id = c.locality_id
            GROUP BY l.id, l.locality_name;
        `)).
			WillReturnRows(rows)

		result, err := repo.GetReportByLocalityAll()
		require.NoError(t, err)
		require.Equal(t, expected, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("report_all_empty", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "locality_name", "carries_count"})

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT 
                l.id, 
                l.locality_name, 
                COUNT(c.id) AS carries_count
            FROM localities l
            LEFT JOIN carries c ON l.id = c.locality_id
            GROUP BY l.id, l.locality_name;
        `)).
			WillReturnRows(rows)

		result, err := repo.GetReportByLocalityAll()
		require.NoError(t, err)
		require.Empty(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("report_all_database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT 
                l.id, 
                l.locality_name, 
                COUNT(c.id) AS carries_count
            FROM localities l
            LEFT JOIN carries c ON l.id = c.locality_id
            GROUP BY l.id, l.locality_name;
        `)).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetReportByLocalityAll()
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Nil(t, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestCarryRepository_ExistsCID(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("exists_cid_true", func(t *testing.T) {
		cid := "CID#100"
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT EXISTS(SELECT 1 FROM carries WHERE cid = ?)
        `)).
			WithArgs(cid).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

		exists, err := repo.ExistsCID(cid)
		require.NoError(t, err)
		require.True(t, exists)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("exists_cid_false", func(t *testing.T) {
		cid := "CID#999"
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT EXISTS(SELECT 1 FROM carries WHERE cid = ?)
        `)).
			WithArgs(cid).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		exists, err := repo.ExistsCID(cid)
		require.NoError(t, err)
		require.False(t, exists)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("exists_cid_database_error", func(t *testing.T) {
		cid := "CID#100"
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT EXISTS(SELECT 1 FROM carries WHERE cid = ?)
        `)).
			WithArgs(cid).
			WillReturnError(fmt.Errorf("database error"))

		exists, err := repo.ExistsCID(cid)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.False(t, exists)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCarryRepository_GetByCID(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("get_by_cid_success", func(t *testing.T) {
		cid := "CID#100"
		expected := models.Carry{
			ID:          1,
			CID:         cid,
			LocalityID:  6700,
			CompanyName: "Fast Logistics",
			Address:     "Calle Falsa 123",
			Telephone:   "123456789",
		}

		rows := sqlmock.NewRows([]string{"id", "cid", "locality_id", "company_name", "address", "telephone"}).
			AddRow(expected.ID, expected.CID, expected.LocalityID, expected.CompanyName, expected.Address, expected.Telephone)

		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE cid = ?
        `)).
			WithArgs(cid).
			WillReturnRows(rows)

		result, err := repo.GetByCID(cid)
		require.NoError(t, err)
		require.Equal(t, expected, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("get_by_cid_not_found", func(t *testing.T) {
		cid := "CID#999"
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE cid = ?
        `)).
			WithArgs(cid).
			WillReturnError(sql.ErrNoRows)

		result, err := repo.GetByCID(cid)
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrCarryRepositoryNotFound)
		require.Equal(t, models.Carry{}, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("get_by_cid_database_error", func(t *testing.T) {
		cid := "CID#100"
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT id, cid, locality_id, company_name, address, telephone 
            FROM carries 
            WHERE cid = ?
        `)).
			WithArgs(cid).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetByCID(cid)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.Equal(t, models.Carry{}, result)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestCarryRepository_ExistsLocality(t *testing.T) {
	repo, mock, close := setupCarriesRepo(t)
	defer close()

	t.Run("exists_locality_true", func(t *testing.T) {
		localityID := 6700

		// Mock: Retorna true (la localidad existe)
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT EXISTS(SELECT 1 FROM localities WHERE id = ?)
        `)).
			WithArgs(localityID).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

		exists, err := repo.ExistsLocality(localityID)
		require.NoError(t, err)
		require.True(t, exists)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("exists_locality_false", func(t *testing.T) {
		localityID := 999

		// Mock: Retorna false (la localidad no existe)
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT EXISTS(SELECT 1 FROM localities WHERE id = ?)
        `)).
			WithArgs(localityID).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		exists, err := repo.ExistsLocality(localityID)
		require.NoError(t, err)
		require.False(t, exists)
		require.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("exists_locality_database_error", func(t *testing.T) {
		localityID := 6700

		// Mock: Retorna error de base de datos
		mock.ExpectQuery(regexp.QuoteMeta(`
            SELECT EXISTS(SELECT 1 FROM localities WHERE id = ?)
        `)).
			WithArgs(localityID).
			WillReturnError(fmt.Errorf("database connection failed"))

		exists, err := repo.ExistsLocality(localityID)
		require.Error(t, err)
		require.Contains(t, err.Error(), e.ErrRepositoryDatabase.Error())
		require.False(t, exists)
		require.NoError(t, mock.ExpectationsWereMet())
	})
}
