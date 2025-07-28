package repository_tests

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	dt "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/data"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type LocalityRepoTestSuite struct {
	dt.TestSuite
	repo *repository.LocalityDB
}

func (suite *LocalityRepoTestSuite) TestLocalities_FindAll() {
	t := suite.T()
	t.Run("#1 - All Success", func(t *testing.T) {
		// given
		suite.SetupTest("localities")
		suite.MockDb.ExpectQuery("SELECT l.id, l.locality_name, l.province_name, l.country_name FROM localities AS l").
			WillReturnRows(suite.TestTable)
		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		result, err := suite.repo.FindAllLocalities()

		// then
		expected := []mod.Locality{
			{ID: 1, Name: "Manhattan", Province: "New York", Country: "USA"},
			{ID: 2, Name: "Downtown", Province: "California", Country: "USA"},
			{ID: 3, Name: "Lakeview", Province: "Illinois", Country: "USA"},
		}
		require.NoError(t, err)
		require.Len(t, result, 3)
		require.Equal(t, expected, result)
	})

	t.Run("#2 - All Unable to parse DB info", func(t *testing.T) {
		// given
		suite.SetupTest("localities")
		suite.MockDb.ExpectQuery("SELECT l.id, l.locality_name, l.province_name, l.country_name FROM localities AS l").
			WillReturnRows(suite.TestTable.AddRow(4, "Medellin", "Antioquia", nil))
		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindAllLocalities()

		// then
		expected := e.ErrParseError
		require.ErrorIs(t, err, expected)
	})

	t.Run("#3 - All Query is malformed", func(t *testing.T) {
		// given
		suite.SetupTest("localities")
		suite.MockDb.ExpectQuery("SELECT l.id, l.locality_name, l.province_name, l.country_name FROM localities AS l").
			WillReturnError(e.ErrQueryError)
		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindAllLocalities()

		// then
		expected := e.ErrQueryError
		require.ErrorIs(t, err, expected)
	})

	t.Run("#4 - All Query is empty", func(t *testing.T) {
		// given
		suite.SetupTest("localities")
		suite.MockDb.ExpectQuery("SELECT l.id, l.locality_name, l.province_name, l.country_name FROM localities AS l").
			WillReturnRows(sqlmock.NewRows(suite.TestColumns))
		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindAllLocalities()

		// then
		expected := e.ErrQueryIsEmpty
		require.ErrorIs(t, err, expected)
	})
}

func (suite *LocalityRepoTestSuite) TestLocalities_Save() {
	t := suite.T()
	newLocality := &mod.Locality{
		Name:     "Medellin",
		Province: "Antioquia",
		Country:  "Colombia",
	}
	expectedQuery := "INSERT INTO `localities`(`locality_name`,`province_name`,`country_name`) VALUES(?,?,?)"

	t.Run("#1 - Save Success", func(t *testing.T) {
		// given
		suite.SetupTest("localities")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newLocality.Name, newLocality.Province, newLocality.Country).
			WillReturnResult(sqlmock.NewResult(1, 1))

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// when
		insertedID, err := suite.repo.Save(newLocality)

		// then
		t.Log(insertedID, err)
		require.NoError(t, err)
		require.Equal(t, 1, insertedID)
	})

	t.Run("#2 - Save Duplicated Entry", func(t *testing.T) {
		// given
		suite.SetupTest("localities")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newLocality.Name, newLocality.Province, newLocality.Country).
			WillReturnError(&mysql.MySQLError{Number: 1062})

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// when
		_, err := suite.repo.Save(newLocality)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrLocalityRepositoryDuplicated)
	})

	t.Run("#3 - Save Unknown error", func(t *testing.T) {
		// given
		suite.SetupTest("localities")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectExec(regexp.QuoteMeta(expectedQuery)).
			WithArgs(newLocality.Name, newLocality.Province, newLocality.Country).
			WillReturnError(errors.New("unexpected db error"))

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// when
		_, err := suite.repo.Save(newLocality)

		// then
		require.Error(t, err)
		t.Log(err)
		require.ErrorIs(t, err, e.ErrInsertError)
	})
}

func (suite *LocalityRepoTestSuite) TestLocalities_FindSellerByLocalityID() {
	t := suite.T()
	expectedQuery := "SELECT l.id, l.locality_name, count(s.id) FROM localities AS `l` LEFT JOIN `sellers` as `s` ON l.id=s.locality_id GROUP BY l.id"

	t.Run("#1 - ID All Success", func(t *testing.T) {
		// given
		suite.SetupTest("sel_by_loc")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
			WillReturnRows(suite.TestTable)

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		result, err := suite.repo.FindSellersByLocID(-1)

		// then
		expected := []mod.SelByLoc{
			{ID: 1, Name: "Manhattan", Count: 5},
			{ID: 2, Name: "Downtown", Count: 3},
			{ID: 3, Name: "Lakeview", Count: 2},
		}

		require.NoError(t, err)
		require.Equal(t, expected, result)
	})
	t.Run("#2 - ID One Success", func(t *testing.T) {
		// given
		suite.SetupTest("sel_by_loc")
		defer suite.TestDb.Close()

		mockRow := sqlmock.NewRows(suite.TestColumns).
			AddRow(1, "Manhattan", 5)

		suite.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery + " HAVING l.id= ?")).
			WithArgs(1).
			WillReturnRows(mockRow)

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		result, err := suite.repo.FindSellersByLocID(1)

		// then
		expected := []mod.SelByLoc{{ID: 1, Name: "Manhattan", Count: 5}}
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("#3 - ID Parse Error", func(t *testing.T) {
		// given
		suite.SetupTest("sel_by_loc")
		defer suite.TestDb.Close()

		mockRowWithNil := sqlmock.NewRows(suite.TestColumns).
			AddRow(nil, "Manhattan", 5)

		suite.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery + " HAVING l.id= ?")).
			WithArgs(1).
			WillReturnRows(mockRowWithNil)

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindSellersByLocID(1)

		// then
		require.Error(t, err)
		require.ErrorIs(t, err, e.ErrParseError)
	})

	t.Run("#4 - ID Query is empty", func(t *testing.T) {
		// given
		suite.SetupTest("sel_by_loc")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
			WithArgs(1).WillReturnRows(sqlmock.NewRows(suite.TestColumns))

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindSellersByLocID(1)

		// then
		expected := e.ErrLocalityRepositoryNotFound
		require.ErrorIs(t, err, expected)
	})

	t.Run("#5 - ID Unknown error", func(t *testing.T) {
		// given
		suite.SetupTest("sel_by_loc")
		defer suite.TestDb.Close()

		suite.MockDb.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
			WithArgs(1).
			WillReturnError(errors.New("unexpected db error"))

		suite.repo = repository.NewLocalityRepo(suite.TestDb)

		// When
		_, err := suite.repo.FindSellersByLocID(1)

		// then
		expected := e.ErrQueryError
		require.ErrorIs(t, err, expected)
	})
}

func TestLocalityRepoTestSuite(t *testing.T) {
	suite.Run(t, new(LocalityRepoTestSuite))
}
