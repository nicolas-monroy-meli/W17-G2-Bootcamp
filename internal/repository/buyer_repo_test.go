package repository

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type TestBuyeRepo struct {
	suite.Suite
	TestDb      *sql.DB
	MockDb      sqlmock.Sqlmock
	TestColumns []string
	TestTable   *sqlmock.Rows
	Repo        *BuyerDB
}

func (s *TestBuyeRepo) SetupTest() {
	s.TestDb, s.MockDb, _ = sqlmock.New()
	s.TestColumns = []string{"id", "id_card_number", "first_name", "last_name"}
	s.TestTable = sqlmock.NewRows(s.TestColumns).
		AddRow(1, "1234567890123456", "Juan", "Pérez").
		AddRow(2, "9876543210987654", "Ana", "González")

	s.Repo = NewBuyerRepo(s.TestDb)

}

func (s *TestBuyeRepo) TestGetAllRepo() {
	t := s.T()
	expectedQuery := "SELECT `id`, `id_card_number`, `first_name`, `last_name` FROM buyers"

	t.Run("Case 1: Success", func(t *testing.T) {
		// given
		s.SetupTest()
		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnRows(s.TestTable)

		// When
		result, err := s.Repo.FindAll()

		// then
		expected := []mod.Buyer{
			{ID: 1, CardNumberID: "1234567890123456", FirstName: "Juan", LastName: "Pérez"},
			{ID: 2, CardNumberID: "9876543210987654", FirstName: "Ana", LastName: "González"},
		}
		require.NoError(t, err)
		require.Len(t, result, 2)
		require.Equal(t, expected, result)
	})

	t.Run("Case 2: Fail - No data", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnRows(sqlmock.NewRows(s.TestColumns)) // sin AddRow

		result, err := s.Repo.FindAll()

		require.NoError(t, err)
		require.Empty(t, result)
	})

	t.Run("Case 3: Error on query", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnError(errors.New("db fail"))

		result, err := s.Repo.FindAll()

		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("Case 4: Scan failure", func(t *testing.T) {
		s.SetupTest()
		s.TestTable = sqlmock.NewRows(s.TestColumns).
			AddRow("STR_NOT_INT", "1234567890123456", "Juan", "Pérez")
		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnRows(s.TestTable)

		result, err := s.Repo.FindAll()

		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("Case 5: Rows.Err() error", func(t *testing.T) {
		s.SetupTest()
		s.TestTable = sqlmock.NewRows(s.TestColumns).
			AddRow(1, "1234567890123456", "Juan", "Pérez")

		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnRows(s.TestTable).RowsWillBeClosed().WillReturnError(nil)
		s.TestTable.RowError(0, errors.New("row error"))

		result, err := s.Repo.FindAll()

		require.Error(t, err)
		require.Nil(t, result)
	})

}

func (s *TestBuyeRepo) TestGetByIdRepo() {
	t := s.T()
	expectedQuery := "SELECT `id`, `id_card_number`, `first_name`, `last_name` FROM buyers WHERE buyers.id = ?"
	t.Run("Case 1: Success", func(t *testing.T) {
		// given
		s.SetupTest()
		s.MockDb.ExpectQuery(expectedQuery).
			WithArgs(1).
			WillReturnRows(s.TestTable)

		// When
		result, err := s.Repo.FindByID(1)

		// then
		expected := mod.Buyer{
			ID: 1, CardNumberID: "1234567890123456", FirstName: "Juan", LastName: "Pérez",
		}
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("Case 2: Not found", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnRows(sqlmock.NewRows(s.TestColumns))

		_, err := s.Repo.FindByID(999)

		require.Error(t, err)
		require.Equal(t, e.ErrBuyerRepositoryNotFound, err)
	})

	t.Run("Case 3: Query error", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnError(errors.New("db error"))

		_, err := s.Repo.FindByID(1)

		require.Error(t, err)
		require.Contains(t, err.Error(), "db error")
	})

	t.Run("Case 4: Scan error", func(t *testing.T) {
		s.SetupTest()
		s.TestTable = sqlmock.NewRows(s.TestColumns).AddRow("NOT_AN_INT", "1234567890123456", "Juan", "Pérez")
		s.MockDb.ExpectQuery(expectedQuery).
			WillReturnRows(s.TestTable)

		_, err := s.Repo.FindByID(1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "converting driver.Value type string")
	})

}

func (s *TestBuyeRepo) TestSaveBuyerRepo() {
	t := s.T()
	newBuyer := mod.Buyer{
		ID:           0,
		FirstName:    "Jorge",
		LastName:     "Casanova",
		CardNumberID: "1032",
	}

	expectedQuery := "INSERT INTO buyers (id_card_number, first_name, last_name) VALUES (?, ?, ?)"

	t.Run("Case 1: Success", func(t *testing.T) {
		// given
		s.SetupTest()

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName).
			WillReturnResult(sqlmock.NewResult(3, 1))

		// When
		err := s.Repo.Save(&newBuyer)

		// then
		require.NoError(t, err)
		require.Equal(t, 3, newBuyer.ID)
	})

	t.Run("Case 2: Exec error", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName).
			WillReturnError(errors.New("db fail"))

		err := s.Repo.Save(&newBuyer)

		require.Error(t, err)
		require.Contains(t, err.Error(), "db fail")
	})

	t.Run("Case 3: Duplicated card number", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName).
			WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry"})

		err := s.Repo.Save(&newBuyer)

		require.ErrorIs(t, err, e.ErrBuyerRepositoryCardDuplicated)
	})

	t.Run("Case 4: Fail on LastInsertId", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("fail id")))

		err := s.Repo.Save(&newBuyer)

		require.Error(t, err)
		require.Contains(t, err.Error(), "fail id")
	})

	t.Run("Case 5: Fail Sql other error", func(t *testing.T) {
		s.SetupTest()

		mysqlErr := &mysql.MySQLError{Number: 2050, Message: "Unknown error"}
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName).
			WillReturnError(mysqlErr)

		err := s.Repo.Save(&newBuyer)

		require.Error(t, err)
		require.Equal(t, mysqlErr, err)
	})
}

func (s *TestBuyeRepo) TestUpdateBuyerRepo() {
	t := s.T()
	patchBuyer := mod.Buyer{
		ID:           1,
		FirstName:    "Jorge",
		LastName:     "Casanova",
		CardNumberID: "1032",
	}

	expectedQuery := "UPDATE buyers SET id_card_number=?, first_name=?, last_name=? WHERE id=?"

	t.Run("Case 1: Update Success", func(t *testing.T) {
		// given
		s.SetupTest()

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchBuyer.CardNumberID, patchBuyer.FirstName, patchBuyer.LastName, patchBuyer.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		// when
		err := s.Repo.Update(&patchBuyer)

		// then
		require.NoError(t, err)
	})

	t.Run("#2 - Exec error", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchBuyer.CardNumberID, patchBuyer.FirstName, patchBuyer.LastName, patchBuyer.ID).
			WillReturnError(errors.New("db fail"))

		err := s.Repo.Update(&patchBuyer)

		require.Error(t, err)
		require.Contains(t, err.Error(), "db fail")
	})

	t.Run("#3 - Duplicate key error", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchBuyer.CardNumberID, patchBuyer.FirstName, patchBuyer.LastName, patchBuyer.ID).
			WillReturnError(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry"})

		err := s.Repo.Update(&patchBuyer)

		require.ErrorIs(t, err, e.ErrBuyerRepositoryCardDuplicated)
	})

	t.Run("#4 - Other MySQL error", func(t *testing.T) {
		s.SetupTest()

		mysqlErr := &mysql.MySQLError{Number: 1049, Message: "Unknown database"}
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchBuyer.CardNumberID, patchBuyer.FirstName, patchBuyer.LastName, patchBuyer.ID).
			WillReturnError(mysqlErr)

		err := s.Repo.Update(&patchBuyer)

		require.Error(t, err)
		require.Equal(t, mysqlErr, err)
	})
}

func (s *TestBuyeRepo) TestBuyerRepo_Delete() {
	t := s.T()

	expectedQuery := "DELETE FROM buyers WHERE id = ?"

	t.Run("#1 - Delete Success", func(t *testing.T) {
		// given
		s.SetupTest()

		s.MockDb.ExpectExec(expectedQuery).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		// when
		err := s.Repo.Delete(1)

		// then
		require.NoError(t, err)
	})

	t.Run("#2 - Exec Error", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectExec(expectedQuery).
			WithArgs(1).
			WillReturnError(errors.New("db fail"))

		err := s.Repo.Delete(1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "db fail")
	})

	t.Run("#3 - Not Found", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectExec(expectedQuery).
			WithArgs(99).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := s.Repo.Delete(99)
		require.ErrorIs(t, err, e.ErrBuyerRepositoryNotFound)
	})

}

func (s *TestBuyeRepo) TestBuyerRepo_GetReport() {
	t := s.T()
	expectedQuery := "SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(p.buyer_id) as purchase_orders_count " +
		"FROM buyers b " +
		"INNER JOIN purchase_orders p " +
		"ON p.buyer_id = b.id GROUP BY b.id"

	expectedQueryId := "SELECT b.id, b.id_card_number, b.first_name, b.last_name, COUNT(p.buyer_id) as purchase_orders_count " +
		"FROM buyers b " +
		"INNER JOIN purchase_orders p " +
		"ON p.buyer_id = b.id WHERE b.id = ? GROUP BY b.id"

	t.Run("Case 1: Success - All report", func(t *testing.T) {
		s.SetupTest()

		rows := sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "123", "Juan", "Pérez", 4).
			AddRow(2, "999", "Ana", "Gómez", 2)

		s.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
			WillReturnRows(rows)

		result, err := s.Repo.GetPurchaseOrderReport(nil)
		require.NoError(t, err)
		require.Len(t, result, 2)
		require.Equal(t, 4, result[0].PurchaseOrderCount)
		require.Equal(t, "Ana", result[1].FirstName)
	})

	t.Run("Case 2: Success - Single report by id", func(t *testing.T) {
		s.SetupTest()
		id := 1

		rows := sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "123", "Juan", "Pérez", 6)

		s.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQueryId)).
			WithArgs(id).
			WillReturnRows(rows)

		result, err := s.Repo.GetPurchaseOrderReport(&id)
		require.NoError(t, err)
		require.Len(t, result, 1)
		require.Equal(t, 6, result[0].PurchaseOrderCount)
		require.Equal(t, "Pérez", result[0].LastName)
	})

	t.Run("Case 3: Query error", func(t *testing.T) {
		s.SetupTest()

		s.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
			WillReturnError(errors.New("db fail"))

		result, err := s.Repo.GetPurchaseOrderReport(nil)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("Case 4: Scan error", func(t *testing.T) {
		s.SetupTest()

		rows := sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "purchase_orders_count"}).
			AddRow("NOINT", "123", "Juan", "Pérez", 4)

		s.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
			WillReturnRows(rows)

		result, err := s.Repo.GetPurchaseOrderReport(nil)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("Case 5: rows.Err after Next", func(t *testing.T) {
		s.SetupTest()

		rows := sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "purchase_orders_count"}).
			AddRow(1, "123", "Juan", "Pérez", 4).
			RowError(0, errors.New("row error"))

		s.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
			WillReturnRows(rows)

		result, err := s.Repo.GetPurchaseOrderReport(nil)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("Case 6: Not found by id", func(t *testing.T) {
		s.SetupTest()
		id := 300

		s.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQueryId)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "id_card_number", "first_name", "last_name", "purchase_orders_count"}))

		result, err := s.Repo.GetPurchaseOrderReport(&id)
		require.ErrorIs(t, err, e.ErrBuyerRepositoryNotFound)
		require.Nil(t, result)
	})
}

func TestRepoBuyerSuite(t *testing.T) {
	suite.Run(t, new(TestBuyeRepo))
}
