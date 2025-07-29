package repository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	prodData "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/data"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ProductRecordRepoTestSuite struct {
	prodData.TestSuite
	repo *repository.ProductRecordDB
}

func (suite *ProductRecordRepoTestSuite) TestFindAllPR() {
	t := suite.T()
	t.Run("#1 - Success", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).
			AddRow(1, "2025-07-28", 100.0, 150.0, 10).
			AddRow(2, "2025-07-27", 200.0, 250.0, 11)
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records;")).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllPR()
		require.NoError(t, err)
		require.Len(t, result, 2)
	})

	t.Run("#2 - Error en query", func(t *testing.T) {
		suite.SetupTest("product_records")
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records;")).WillReturnError(fmt.Errorf("db error"))
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllPR()
		require.ErrorIs(t, err, e.ErrProductRepositoryNotFound)
		require.Nil(t, result)
	})

	t.Run("#3 - Error en scan", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).AddRow(nil, "2025-07-28", 100.0, 150.0, 10)
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records;")).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllPR()
		require.ErrorIs(t, err, e.ErrProductRecordRepositoryNotFound)
		require.Nil(t, result)
	})

	t.Run("#4 - Error en rows.Err()", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).AddRow(1, "2025-07-28", 100.0, 150.0, 10)
		rows.RowError(0, fmt.Errorf("rows error"))
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records;")).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllPR()
		require.ErrorIs(t, err, e.ErrProductRecordRepositoryNotFound)
		require.Nil(t, result)
	})

	t.Run("#5 - Sin resultados", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"})
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records;")).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllPR()
		require.ErrorIs(t, err, e.ErrProductRecordRepositoryNotFound)
		require.Nil(t, result)
	})
}

func (suite *ProductRecordRepoTestSuite) TestFindAllByProductIDPR() {
	t := suite.T()
	productID := 10
	t.Run("#1 - Success", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).AddRow(1, "2025-07-28", 100.0, 150.0, productID)
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records WHERE product_id = ?;")).WithArgs(productID).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllByProductIDPR(productID)
		require.NoError(t, err)
		require.Len(t, result, 1)
	})

	t.Run("#2 - Error en query", func(t *testing.T) {
		suite.SetupTest("product_records")
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records WHERE product_id = ?;")).WithArgs(productID).WillReturnError(fmt.Errorf("db error"))
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllByProductIDPR(productID)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("#3 - Error en scan", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).AddRow(nil, "2025-07-28", 100.0, 150.0, productID)
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records WHERE product_id = ?;")).WithArgs(productID).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllByProductIDPR(productID)
		require.ErrorIs(t, err, e.ErrProductRecordRepositoryNotFound)
		require.Nil(t, result)
	})

	t.Run("#4 - Error en rows.Err()", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).AddRow(1, "2025-07-28", 100.0, 150.0, productID)
		rows.RowError(0, fmt.Errorf("rows error"))
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records WHERE product_id = ?;")).WithArgs(productID).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllByProductIDPR(productID)
		require.ErrorIs(t, err, e.ErrProductRecordRepositoryNotFound)
		require.Nil(t, result)
	})

	t.Run("#5 - Sin resultados", func(t *testing.T) {
		suite.SetupTest("product_records")
		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"})
		suite.MockDb.ExpectQuery(regexp.QuoteMeta("SELECT `id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id` FROM frescos_db.product_records WHERE product_id = ?;")).WithArgs(productID).WillReturnRows(rows)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		result, err := suite.repo.FindAllByProductIDPR(productID)
		require.ErrorIs(t, err, e.ErrProductRecordRepositoryNotFound)
		require.Nil(t, result)
	})
}

func (suite *ProductRecordRepoTestSuite) TestSavePR() {
	t := suite.T()
	t.Run("#1 - Insert exitoso", func(t *testing.T) {
		suite.SetupTest("product_records")
		pr := &mod.ProductRecord{LastUpdateDate: "2025-07-28", PurchasePrice: 100.0, SalePrice: 150.0, ProductID: 10}
		suite.MockDb.ExpectExec(regexp.QuoteMeta("INSERT INTO frescos_db.product_records (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES(?, ?, ?, ?);")).
			WithArgs(pr.LastUpdateDate, pr.PurchasePrice, pr.SalePrice, pr.ProductID).
			WillReturnResult(sqlmock.NewResult(123, 1))
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		err := suite.repo.SavePR(pr)
		require.NoError(t, err)
		require.Equal(t, 123, pr.ID)
	})

	t.Run("#2 - Error en insert", func(t *testing.T) {
		suite.SetupTest("product_records")
		pr := &mod.ProductRecord{LastUpdateDate: "2025-07-28", PurchasePrice: 100.0, SalePrice: 150.0, ProductID: 10}
		suite.MockDb.ExpectExec(regexp.QuoteMeta("INSERT INTO frescos_db.product_records (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES(?, ?, ?, ?);")).
			WithArgs(pr.LastUpdateDate, pr.PurchasePrice, pr.SalePrice, pr.ProductID).
			WillReturnError(fmt.Errorf("insert error"))
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		err := suite.repo.SavePR(pr)
		require.ErrorContains(t, err, "insert error")
	})

	t.Run("#3 - Error en LastInsertId", func(t *testing.T) {
		suite.SetupTest("product_records")
		pr := &mod.ProductRecord{LastUpdateDate: "2025-07-28", PurchasePrice: 100.0, SalePrice: 150.0, ProductID: 10}
		result := sqlmock.NewErrorResult(fmt.Errorf("lastinsertid error"))
		suite.MockDb.ExpectExec(regexp.QuoteMeta("INSERT INTO frescos_db.product_records (`last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES(?, ?, ?, ?);")).
			WithArgs(pr.LastUpdateDate, pr.PurchasePrice, pr.SalePrice, pr.ProductID).
			WillReturnResult(result)
		suite.repo = repository.NewProductRecordRepo(suite.TestDb)

		err := suite.repo.SavePR(pr)
		require.ErrorContains(t, err, "lastinsertid error")
	})
}

func TestProductRecordRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRecordRepoTestSuite))
}
