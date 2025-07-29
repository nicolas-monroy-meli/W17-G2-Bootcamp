package repository_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	repo "github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	dt "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/data"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type SellerRepoTestSuite struct {
	dt.TestSuite
	repo *repo.SellerDB
}

func (suite *SellerRepoTestSuite) TestSellers_FindAll() {
	t := suite.T()
	t.Run("#1 - All Success", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		suite.MockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").
			WillReturnRows(suite.TestTable)
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// When
		result, err := suite.repo.FindAll()

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
		suite.SetupTest("sellers")
		suite.MockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").
			WillReturnRows(suite.TestTable.AddRow(1, 1001, "Alpha Traders Inc.", "123 Alpha St, New York, NY", "+1-212-555-0101", nil))
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindAll()

		// then
		expected := e.ErrParseError
		require.ErrorIs(t, err, expected)
	})

	t.Run("#3 - All Query is malformed", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		suite.MockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").
			WillReturnError(e.ErrQueryError)
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindAll()

		// then
		expected := e.ErrQueryError
		require.ErrorIs(t, err, expected)
	})

	t.Run("#4 - All Query is empty", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		suite.MockDb.ExpectQuery("SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers`").
			WillReturnRows(sqlmock.NewRows(suite.TestColumns))
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindAll()

		// then
		expected := e.ErrQueryIsEmpty
		require.ErrorIs(t, err, expected)
	})
}

func (suite *SellerRepoTestSuite) TestSellers_FindById() {
	t := suite.T()
	expectedQuery := "SELECT `id`, `cid`,`company_name`,`address`,`telephone`,`locality_id` FROM `sellers` WHERE `id` = ?"

	t.Run("#1 - ID Success", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		mockRow := sqlmock.NewRows(suite.TestColumns).
			AddRow(1, 1001, "Alpha Traders Inc.", "123 Alpha St, New York, NY", "+1-212-555-0101", 1)

		suite.MockDb.ExpectQuery(expectedQuery).
			WithArgs(1).
			WillReturnRows(mockRow)

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// When
		result, err := suite.repo.FindByID(1)

		// then
		expected := mod.Seller{ID: 1, CID: 1001, CompanyName: "Alpha Traders Inc.", Address: "123 Alpha St, New York, NY", Telephone: "+1-212-555-0101", Locality: 1}
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("#2 - ID Parse Error", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		mockRowWithNil := sqlmock.NewRows(suite.TestColumns).
			AddRow(1, 1001, "Alpha Traders Inc.", "123 Alpha St, New York, NY", "+1-212-555-0101", nil)

		suite.MockDb.ExpectQuery(expectedQuery).
			WithArgs(1).
			WillReturnRows(mockRowWithNil)

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindByID(1)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrParseError)
	})

	t.Run("#3 - ID No Seller found", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectQuery(expectedQuery).
			WithArgs(99).
			WillReturnError(sql.ErrNoRows)

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindByID(99)

		// then
		expected := e.ErrSellerRepositoryNotFound
		require.ErrorIs(t, err, expected)
	})
}

func (suite *SellerRepoTestSuite) TestSellers_Save() {
	t := suite.T()

	newSeller := &mod.Seller{
		CID:         1009,
		CompanyName: "Zeta Innovations",
		Address:     "900 Zeta Ln, Austin, TX",
		Telephone:   "+1-512-555-0120",
		Locality:    4,
	}
	expectedQuery := "INSERT INTO `sellers`(`cid`,`company_name`,`address`,`telephone`,`locality_id`) VALUES(?,?,?,?,?)"

	t.Run("#1 - Save Success", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newSeller.CID, newSeller.CompanyName, newSeller.Address, newSeller.Telephone, newSeller.Locality).
			WillReturnResult(sqlmock.NewResult(4, 1))

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		insertedID, err := suite.repo.Save(newSeller)

		// then
		t.Log(insertedID, err)
		require.NoError(t, err)
		require.Equal(t, 4, insertedID)
	})

	t.Run("#2 - Save Duplicated Entry", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newSeller.CID, newSeller.CompanyName, newSeller.Address, newSeller.Telephone, newSeller.Locality).
			WillReturnError(&mysql.MySQLError{Number: 1062})

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		_, err := suite.repo.Save(newSeller)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrSellerRepositoryDuplicated)
	})

	t.Run("#3 - Save Foreign Key not found", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newSeller.CID, newSeller.CompanyName, newSeller.Address, newSeller.Telephone, newSeller.Locality).
			WillReturnError(&mysql.MySQLError{Number: 1452})

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		_, err := suite.repo.Save(newSeller)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrForeignKeyError)
	})

	t.Run("#4 - Save Unknown error", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newSeller.CID, newSeller.CompanyName, newSeller.Address, newSeller.Telephone, newSeller.Locality).
			WillReturnError(e.ErrRepositoryDatabase)

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		_, err := suite.repo.Save(newSeller)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrRepositoryDatabase)
	})
}

func (suite *SellerRepoTestSuite) TestSellers_Update() {
	t := suite.T()

	patchedSeller := &mod.Seller{
		ID:          1,
		CID:         1000,
		CompanyName: "Omega Traders Inc.",
		Address:     "321 Omega St, New York, NY",
		Telephone:   "+1-212-555-0101",
		Locality:    1,
	}
	expectedQuery := "UPDATE `sellers` SET `cid`=?,`company_name`=?,`address`=?,`telephone`=?,`locality_id`=? WHERE `id`= ?"

	t.Run("#1 - Update Success", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchedSeller.CID, patchedSeller.CompanyName, patchedSeller.Address, patchedSeller.Telephone, patchedSeller.Locality, patchedSeller.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		err := suite.repo.Update(patchedSeller)

		// then
		require.NoError(t, err)
	})

	t.Run("#2 - Update Duplicated Entry", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchedSeller.CID, patchedSeller.CompanyName, patchedSeller.Address, patchedSeller.Telephone, patchedSeller.Locality, patchedSeller.ID).
			WillReturnError(&mysql.MySQLError{Number: 1062})

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		err := suite.repo.Update(patchedSeller)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrSellerRepositoryDuplicated)
	})

	t.Run("#3 - Update Foreign Key not found", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchedSeller.CID, patchedSeller.CompanyName, patchedSeller.Address, patchedSeller.Telephone, patchedSeller.Locality, patchedSeller.ID).
			WillReturnError(&mysql.MySQLError{Number: 1452})

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		err := suite.repo.Update(patchedSeller)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrForeignKeyError)
	})

	t.Run("#4 - Update Unknown DB Error", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(patchedSeller.CID, patchedSeller.CompanyName, patchedSeller.Address, patchedSeller.Telephone, patchedSeller.Locality, patchedSeller.ID).
			WillReturnError(errors.New("unexpected db error"))

		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		err := suite.repo.Update(patchedSeller)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrRepositoryDatabase)
	})
}

func (suite *SellerRepoTestSuite) TestSellers_Delete() {
	t := suite.T()

	expectedQuery := "DELETE FROM `sellers` WHERE `id`=?"

	t.Run("#1 - Delete Success", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		err := suite.repo.Delete(1)

		// then
		require.NoError(t, err)
	})

	t.Run("#2 - Delete seller not found", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 0))
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		err := suite.repo.Delete(1)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrSellerRepositoryNotFound)
	})

	t.Run("#3 - Update Unknown DB Error", func(t *testing.T) {
		// given
		suite.SetupTest("sellers")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(1).
			WillReturnError(errors.New("unexpected db error"))
		suite.repo = repo.NewSellerRepo(suite.TestDb)

		// when
		err := suite.repo.Delete(1)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrRepositoryDatabase)
	})
}

func TestSellerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(SellerRepoTestSuite))
}
