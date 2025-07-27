package data

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	TestDb    *sql.DB
	MockDb    sqlmock.Sqlmock
	TestTable *sqlmock.Rows
}

func (suite *TestSuite) SetupTest(table string) {
	var err error
	suite.TestDb, suite.MockDb, err = sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to open mock database: %v", err)
	}

	// Initialize the testTable based on the provided table name
	suite.TestTable = suite.tableType(table)() // Calling the function here to generate the rows
	if suite.TestTable == nil {
		suite.T().Fatalf("No mock data found for table: %s", table)
	}
}

func (suite *TestSuite) tableType(table string) func() *sqlmock.Rows {
	switch table {
	case "sellers":
		return suite.buildSellers
	case "locality":
		return suite.buildLocalities
	case "sel_by_loc":
		return suite.buildSelByLoc
	default:
		return nil
	}
}

func (suite *TestSuite) buildSellers() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"}).
		AddRow(1, 1001, "Alpha Traders Inc.", "123 Alpha St, New York, NY", "+1-212-555-0101", 1).
		AddRow(2, 1008, "Omicron Ventures", "888 Omicron Dr, San Francisco, CA", "+1-415-555-0110", 2).
		AddRow(3, 1002, "Beta Logistics Ltd.", "456 Beta Blvd, Chicago, IL", "+1-312-555-0102", 3)
}

func (suite *TestSuite) buildLocalities() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "locality_name", "province_name", "country_name"}).
		AddRow(1, "Manhattan", "New York", "USA").
		AddRow(2, "Downtown", "California", "USA").
		AddRow(3, "Lakeview", "Illinois", "USA")
}

func (suite *TestSuite) buildSelByLoc() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"locality_id", "locality_name", "sellers_count"}).
		AddRow(1, "Manhattan", 5).
		AddRow(2, "Downtown", 3).
		AddRow(3, "Lakeview", 2)
}
