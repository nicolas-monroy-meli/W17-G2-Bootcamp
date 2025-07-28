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
	tableFunc := suite.tableType(table)
	if tableFunc == nil {
		suite.T().Fatalf("No mock data found for table: %s", table)
	}
	suite.TestColumns, suite.TestTable = tableFunc()
}

func (suite *TestSuite) tableType(table string) func() ([]string, *sqlmock.Rows) {
	switch table {
	case "products":
		return suite.buildProducts
	case "product_records":
		return suite.buildProductsRecords
	case "sellers":
		return suite.buildSellers
	case "localities":
		return suite.buildLocalities
	case "sel_by_loc":
		return suite.buildSelByLoc
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

func (suite *TestSuite) buildSellers() ([]string, *sqlmock.Rows) {
	column := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
	return column, sqlmock.NewRows(column).
		AddRow(1, 1001, "Alpha Traders Inc.", "123 Alpha St, New York, NY", "+1-212-555-0101", 1).
		AddRow(2, 1008, "Omicron Ventures", "888 Omicron Dr, San Francisco, CA", "+1-415-555-0110", 2).
		AddRow(3, 1002, "Beta Logistics Ltd.", "456 Beta Blvd, Chicago, IL", "+1-312-555-0102", 3)
}

func (suite *TestSuite) buildLocalities() ([]string, *sqlmock.Rows) {
	column := []string{"id", "locality_name", "province_name", "country_name"}
	return column, sqlmock.NewRows(column).
		AddRow(1, "Manhattan", "New York", "USA").
		AddRow(2, "Downtown", "California", "USA").
		AddRow(3, "Lakeview", "Illinois", "USA")
}

func (suite *TestSuite) buildSelByLoc() ([]string, *sqlmock.Rows) {
	column := []string{"locality_id", "locality_name", "sellers_count"}
	return column, sqlmock.NewRows(column).
		AddRow(1, "Manhattan", 5).
		AddRow(2, "Downtown", 3).
		AddRow(3, "Lakeview", 2)
}
