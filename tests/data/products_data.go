package data

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	TestDb      *sql.DB
	MockDb      sqlmock.Sqlmock
	TestColumns []string
	TestTable   *sqlmock.Rows
}

func (suite *TestSuite) SetupTest(table string) {
	suite.TestDb, suite.MockDb, _ = sqlmock.New()
	suite.TestColumns, suite.TestTable = suite.tableType(table)()
	if suite.TestTable == nil {
		suite.T().Fatalf("No mock data found for table: %s", table)
	}
}

func (suite *TestSuite) tableType(table string) func() ([]string, *sqlmock.Rows) {
	switch table {
	case "products":
		return suite.buildProducts
	case "product_records":
		return suite.buildProductsRecords
	default:
		return nil
	}
}

func (suite *TestSuite) buildProducts() ([]string, *sqlmock.Rows) {
	column := []string{
		"id",
		"product_code",
		"description",
		"height",
		"length",
		"width",
		"net_weight",
		"expiration_rate",
		"freezing_rate",
		"recommended_freezing_temperature",
		"product_type_id",
		"seller_id",
	}
	return column, sqlmock.NewRows(column).
		AddRow(1, "P001", "Product 1", 10.0, 20.0, 5.0, 2.0, 0.1, 0.05, -18.0, 1, 101).
		AddRow(2, "P002", "Product 2", 15.0, 25.0, 7.0, 3.0, 0.2, 0.06, -20.0, 2, 102).
		AddRow(3, "P003", "Product 3", 12.0, 22.0, 6.0, 2.5, 0.15, 0.07, -19.0, 1, 103)
}

func (suite *TestSuite) buildProductsRecords() ([]string, *sqlmock.Rows) {
	column := []string{
		"id",
		"last_update_date",
		"purchase_price",
		"sale_price",
		"product_id",
	}
	return column, sqlmock.NewRows(column).
		AddRow(1, "2023-10-01", 100.0, 150.0, 1).
		AddRow(2, "2023-10-02", 200.0, 250.0, 2).
		AddRow(3, "2023-10-03", 300.0, 350.0, 3)
}
