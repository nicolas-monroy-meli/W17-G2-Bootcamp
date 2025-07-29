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
	"time"
)

type TestPurchaseOrderRepo struct {
	suite.Suite
	TestDb      *sql.DB
	MockDb      sqlmock.Sqlmock
	TestColumns []string
	Repo        *PurchaseOrderDB
}

func (s *TestPurchaseOrderRepo) SetupTest() {
	s.TestDb, s.MockDb, _ = sqlmock.New()
	s.TestColumns = []string{"id", "order_number", "order_date", "tracking_code", "buyer_id"}

	s.Repo = NewPurchaseOrderRepo(s.TestDb)

}

func (s *TestPurchaseOrderRepo) TestSaveBuyerRepo() {
	t := s.T()

	newPurchaseOrder := mod.PurchaseOrder{
		//ID:           1,
		OrderNumber:  "PO-2024-001",
		OrderDate:    mod.Date(time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)),
		TrackingCode: "TRACK-123",
		BuyerId:      42,
		ProductsDetails: []mod.OrderDetails{
			{
				//ID:                100,
				CleanLinessStatus: "Clean",
				Quantity:          10,
				Temperature:       6.5,
				ProductRecordId:   3001,
				PurchaseOrderId:   1,
			},
			{
				//ID:                101,
				CleanLinessStatus: "Clean",
				Quantity:          5,
				Temperature:       4.2,
				ProductRecordId:   3002,
				PurchaseOrderId:   1,
			},
		},
	}

	expectedQueryPurchaseOrder := `
        INSERT INTO purchase_orders (order_number, order_date, tracking_code, buyer_id) 
        VALUES (?, ?, ?, ?)`

	expectedQueryOrderDetail := `
        INSERT INTO order_details (clean_liness_status, quantity, temperature, product_record_id, purchase_order_id) 
        VALUES (?, ?, ?, ?, ?)`

	t.Run("Case 1: Success", func(t *testing.T) {
		// given
		s.SetupTest()

		s.MockDb.ExpectBegin()

		//Query purchase order
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnResult(sqlmock.NewResult(21, 1))

		//query order detail
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryOrderDetail)).
			WithArgs(newPurchaseOrder.ProductsDetails[0].CleanLinessStatus, newPurchaseOrder.ProductsDetails[0].Quantity, newPurchaseOrder.ProductsDetails[0].Temperature, newPurchaseOrder.ProductsDetails[0].ProductRecordId, 21).
			WillReturnResult(sqlmock.NewResult(31, 1))

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryOrderDetail)).
			WithArgs(newPurchaseOrder.ProductsDetails[1].CleanLinessStatus, newPurchaseOrder.ProductsDetails[1].Quantity, newPurchaseOrder.ProductsDetails[1].Temperature, newPurchaseOrder.ProductsDetails[1].ProductRecordId, 21).
			WillReturnResult(sqlmock.NewResult(32, 1))

		s.MockDb.ExpectCommit()

		// When
		err := s.Repo.Save(&newPurchaseOrder)

		// then
		require.NoError(s.T(), err)
		require.Equal(s.T(), 21, newPurchaseOrder.ID)
		require.Equal(s.T(), 31, newPurchaseOrder.ProductsDetails[0].ID)
		require.Equal(s.T(), 32, newPurchaseOrder.ProductsDetails[1].ID)

		// 8. Verifica todas las expectativas
		require.NoError(s.T(), s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 2: Fail - General error", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnError(errors.New("db failed"))
		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)
		require.Error(t, err)
		require.Contains(t, err.Error(), "db failed")
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 3: Fail - Order Number duplicated", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		mysqlErr := &mysql.MySQLError{Number: 1062, Message: "Duplicate"}
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnError(mysqlErr)
		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)
		require.ErrorIs(t, err, e.ErrPORepositoryOrderNumberDuplicated)
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 4: Fail - Purchase Order Foreign Key Error", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		mysqlErr := &mysql.MySQLError{Number: 1452, Message: "Foreign key"}
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnError(mysqlErr)
		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)
		require.ErrorIs(t, err, e.ErrForeignKeyError)
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 5: Order detail fails - No sql error", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnResult(sqlmock.NewResult(21, 1))

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryOrderDetail)).
			WillReturnError(errors.New("detail fail"))
		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)
		require.Error(t, err)
		require.Contains(t, err.Error(), "detail fail")
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 6: Fail - Order details foregin Key", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		mysqlErr := &mysql.MySQLError{Number: 1452, Message: "Foreign key"}
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnResult(sqlmock.NewResult(21, 1))

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryOrderDetail)).
			WithArgs(newPurchaseOrder.ProductsDetails[0].CleanLinessStatus, newPurchaseOrder.ProductsDetails[0].Quantity, newPurchaseOrder.ProductsDetails[0].Temperature, newPurchaseOrder.ProductsDetails[0].ProductRecordId, 21).
			WillReturnError(mysqlErr)
		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)
		require.ErrorIs(t, err, e.ErrForeignKeyError)
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 7: Fail - Order details other mysql error", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		mysqlErr := &mysql.MySQLError{Number: 1049, Message: "Unknown database"}
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnResult(sqlmock.NewResult(21, 1))

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryOrderDetail)).
			WithArgs(newPurchaseOrder.ProductsDetails[0].CleanLinessStatus, newPurchaseOrder.ProductsDetails[0].Quantity, newPurchaseOrder.ProductsDetails[0].Temperature, newPurchaseOrder.ProductsDetails[0].ProductRecordId, 21).
			WillReturnError(mysqlErr)
		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)

		require.ErrorIs(t, err, mysqlErr)
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 8: Fail - Order details failed in lastInsertId", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnResult(sqlmock.NewResult(21, 1))

		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryOrderDetail)).
			WithArgs(newPurchaseOrder.ProductsDetails[0].CleanLinessStatus, newPurchaseOrder.ProductsDetails[0].Quantity, newPurchaseOrder.ProductsDetails[0].Temperature, newPurchaseOrder.ProductsDetails[0].ProductRecordId, 21).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("fail id")))

		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)
		require.Error(t, err)
		require.Contains(t, err.Error(), "fail id")
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})

	t.Run("Case 9: Fail - Purchase Order Other mysql error", func(t *testing.T) {
		s.SetupTest()
		s.MockDb.ExpectBegin()
		mysqlErr := &mysql.MySQLError{Number: 1049, Message: "Unknown database"}
		s.MockDb.ExpectExec(regexp.QuoteMeta(expectedQueryPurchaseOrder)).
			WithArgs(newPurchaseOrder.OrderNumber, time.Time(newPurchaseOrder.OrderDate), newPurchaseOrder.TrackingCode, newPurchaseOrder.BuyerId).
			WillReturnError(mysqlErr)

		s.MockDb.ExpectRollback()

		err := s.Repo.Save(&newPurchaseOrder)

		require.ErrorIs(t, err, mysqlErr)
		require.NoError(t, s.MockDb.ExpectationsWereMet())
	})
}

func TestRepoPOSuite(t *testing.T) {
	suite.Run(t, new(TestPurchaseOrderRepo))
}
