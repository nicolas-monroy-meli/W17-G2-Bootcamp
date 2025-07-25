package data

import "github.com/DATA-DOG/go-sqlmock"

var (
	SellerTestDb, SellerMockDb,  _= sqlmock.New()
	SellerColumns                 = []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
	SellerTestTable               = RebuildSellers()
)

func RebuildSellers() *sqlmock.Rows {
	return sqlmock.NewRows(SellerColumns).
		AddRow(1, 1001, "Alpha Traders Inc.", "123 Alpha St, New York, NY", "+1-212-555-0101", 1).
		AddRow(2, 1008, "Omicron Ventures", "888 Omicron Dr, San Francisco, CA", "+1-415-555-0110", 2).
		AddRow(3, 1002, "Beta Logistics Ltd.", "456 Beta Blvd, Chicago, IL", "+1-312-555-0102", 3)
}