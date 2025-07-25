package repository_tests

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	dt "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/data"
	"github.com/stretchr/testify/require"
)

func TestSellers_GetAll(t *testing.T) {
	t.Run("#1 - Success", func(t *testing.T) {
		// given
		repo := repository.NewSellerRepo(dt.SellerTestDb)
		dt.SellerMockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").
			WillReturnRows(dt.RebuildSellers())

		// When
		result, err := repo.FindAll()

		// then
		expected := []mod.Seller{
			{ID: 1, CID: 1001, CompanyName: "Alpha Traders Inc.", Address: "123 Alpha St, New York, NY", Telephone: "+1-212-555-0101", Locality: 1},
			{ID: 2, CID: 1008, CompanyName: "Omicron Ventures", Address: "888 Omicron Dr, San Francisco, CA", Telephone: "+1-415-555-0110", Locality: 2},
			{ID: 3, CID: 1002, CompanyName: "Beta Logistics Ltd.", Address: "456 Beta Blvd, Chicago, IL", Telephone: "+1-312-555-0102", Locality: 3},
		}
		require.NoError(t, err)
		require.Len(t, result, 3)
		require.Equal(t, expected, result)
	})
	t.Run("#2 - Unable to parse DB info", func(t *testing.T) {
		// given
		repo := repository.NewSellerRepo(dt.SellerTestDb)
		dt.SellerMockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").
			WillReturnRows(dt.RebuildSellers().AddRow(1, 1001, "Alpha Traders Inc.", "123 Alpha St, New York, NY", "+1-212-555-0101", nil))

		// When
		_, err := repo.FindAll()

		// then
		expected := e.ErrParseError
		require.ErrorIs(t, err, expected)
	})
	t.Run("#3 - Query is malformed", func(t *testing.T) {
		// given
		dt.SellerTestTable = dt.RebuildSellers()
		repo := repository.NewSellerRepo(dt.SellerTestDb)
		dt.SellerMockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").WillReturnError(e.ErrQueryError)

		// When
		_, err := repo.FindAll()

		// then
		expected := e.ErrQueryError
		require.ErrorIs(t, err, expected)
	})
	t.Run("#4 - Query is empty", func(t *testing.T) {
		// given
		dt.SellerTestTable = dt.RebuildSellers()
		repo := repository.NewSellerRepo(dt.SellerTestDb)
		dt.SellerMockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").
			WillReturnRows(sqlmock.NewRows(dt.SellerColumns))

		// When
		_, err := repo.FindAll()

		// then
		expected := e.ErrQueryIsEmpty
		require.ErrorIs(t, err, expected)
	})
}

func TestSellerCreate_HappyPath(t *testing.T) {
}
