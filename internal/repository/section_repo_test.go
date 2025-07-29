package repository

import (
	"database/sql"
	"database/sql/driver"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	m "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/mock"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/require"
)

// Helper
func setupMockSectionRepo(t *testing.T) (*SectionDB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	repo := NewSectionRepo(db)
	return repo, mock, func() { db.Close() }
}

// FindAll
func TestSectionDB_FindAll(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		expected    []mod.Section
		expectedErr error
	}{
		{
			name: "returns data",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(m.SectionTableStruct).
					AddRows(m.SectionDataValuesSelect...)
				mock.ExpectQuery(m.SectionSelectExpectedQuery).WillReturnRows(rows)
			},
			expected: []mod.Section{
				{
					ID:                 1,
					SectionNumber:      1,
					CurrentTemperature: 0,
					MinimumTemperature: -5,
					CurrentCapacity:    50,
					MinimumCapacity:    20,
					MaximumCapacity:    100,
					WarehouseID:        1,
					ProductTypeID:      1,
				},
				{
					ID:                 2,
					SectionNumber:      2,
					CurrentTemperature: -2,
					MinimumTemperature: -6,
					CurrentCapacity:    60,
					MinimumCapacity:    30,
					MaximumCapacity:    110,
					WarehouseID:        2,
					ProductTypeID:      2,
				},
			},
			expectedErr: nil,
		},
		{
			name: "returns empty db error",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "section_number"})
				mock.ExpectQuery(m.SectionSelectExpectedQuery).WillReturnRows(rows)
			},
			expected:    []mod.Section{},
			expectedErr: e.ErrEmptyDB,
		},
		{
			name: "Query Error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(m.SectionSelectExpectedQuery).
					WillReturnError(e.ErrQueryError)
			},
			expected:    []mod.Section{},
			expectedErr: e.ErrQueryError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockSectionRepo(t)
			defer teardown()
			tc.setupMock(mock)
			sections, err := repo.FindAll()
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, sections)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// FindByID
func TestSectionDB_FindByID(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock, id int)
		id          int
		expected    mod.Section
		expectedErr error
	}{
		{
			name: "existing section",
			id:   2,
			setupMock: func(mock sqlmock.Sqlmock, id int) {
				rows := sqlmock.NewRows(m.SectionTableStruct).
					AddRow(m.SectionDataValuesSelectByID...)
				mock.ExpectQuery(m.SectionSelectWhereExpectedQuery).
					WithArgs(id).WillReturnRows(rows)
			},
			expected: mod.Section{
				ID:                 2,
				SectionNumber:      2,
				CurrentTemperature: -2,
				MinimumTemperature: -6,
				CurrentCapacity:    60,
				MinimumCapacity:    30,
				MaximumCapacity:    110,
				WarehouseID:        2,
				ProductTypeID:      2,
			},
			expectedErr: nil,
		},
		{
			name: "not found",
			id:   2,
			setupMock: func(mock sqlmock.Sqlmock, id int) {
				mock.ExpectQuery(m.SectionSelectExpectedQuery).
					WithArgs(id).WillReturnError(sql.ErrNoRows)
			},
			expected:    mod.Section{},
			expectedErr: e.ErrSectionRepositoryNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockSectionRepo(t)
			defer teardown()
			tc.setupMock(mock, tc.id)
			section, err := repo.FindByID(tc.id)
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, section)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// Save
func TestSectionDB_Save(t *testing.T) {
	baseSection := &mod.Section{
		SectionNumber:      1,
		CurrentTemperature: -2,
		MinimumTemperature: -6,
		CurrentCapacity:    60,
		MinimumCapacity:    30,
		MaximumCapacity:    110,
		WarehouseID:        2,
		ProductTypeID:      2,
	}
	tests := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock, s *mod.Section)
		expectedID  int
		expectedErr error
	}{
		{
			name: "happy path",
			setupMock: func(mock sqlmock.Sqlmock, s *mod.Section) {
				mock.ExpectExec("INSERT INTO `sections`").
					WithArgs(s.SectionNumber, s.CurrentTemperature, s.MinimumTemperature, s.CurrentCapacity, s.MinimumCapacity, s.MaximumCapacity, s.WarehouseID, s.ProductTypeID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedID:  1,
			expectedErr: nil,
		},
		{
			name: "duplicated",
			setupMock: func(mock sqlmock.Sqlmock, s *mod.Section) {
				mock.ExpectExec("INSERT INTO `sections`").
					WithArgs(s.SectionNumber, s.CurrentTemperature, s.MinimumTemperature, s.CurrentCapacity, s.MinimumCapacity, s.MaximumCapacity, s.WarehouseID, s.ProductTypeID).
					WillReturnError(e.DupErr)
			},
			expectedID:  0,
			expectedErr: e.ErrSectionRepositoryDuplicated,
		},
		{
			name: "foreign key error",
			setupMock: func(mock sqlmock.Sqlmock, s *mod.Section) {
				mock.ExpectExec("INSERT INTO `sections`").
					WithArgs(s.SectionNumber, s.CurrentTemperature, s.MinimumTemperature, s.CurrentCapacity, s.MinimumCapacity, s.MaximumCapacity, s.WarehouseID, s.ProductTypeID).
					WillReturnError(e.FkErr)
			},
			expectedID:  0,
			expectedErr: e.ErrForeignKeyError,
		},
		{
			name: "last insert id error",
			setupMock: func(mock sqlmock.Sqlmock, s *mod.Section) {
				mock.ExpectExec("INSERT INTO `sections`").
					WithArgs(s.SectionNumber, s.CurrentTemperature, s.MinimumTemperature, s.CurrentCapacity, s.MinimumCapacity, s.MaximumCapacity, s.WarehouseID, s.ProductTypeID).
					WillReturnResult(e.FakeResult{})
			},
			expectedID:  0,
			expectedErr: e.ErrSectionRepositoryNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockSectionRepo(t)
			defer teardown()
			tc.setupMock(mock, baseSection)
			err := repo.Save(baseSection)
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedID, baseSection.ID)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// Delete
func TestSectionDB_Delete(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock, id int)
		id          int
		expectedErr error
	}{
		{
			name: "happy path",
			id:   2,
			setupMock: func(mock sqlmock.Sqlmock, id int) {
				mock.ExpectExec("DELETE FROM `sections`").
					WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: nil,
		},
		{
			name: "not found",
			id:   2,
			setupMock: func(mock sqlmock.Sqlmock, id int) {
				mock.ExpectExec("DELETE FROM `sections`").
					WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: e.ErrSectionRepositoryNotFound,
		},
		{
			name: "query error",
			id:   2,
			setupMock: func(mock sqlmock.Sqlmock, id int) {
				mock.ExpectExec("DELETE FROM `sections`").
					WithArgs(id).
					WillReturnError(e.ErrQueryError)
			},
			expectedErr: e.ErrQueryError,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockSectionRepo(t)
			defer teardown()
			tc.setupMock(mock, tc.id)
			err := repo.Delete(tc.id)
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestSectionDB_Update(t *testing.T) {
	toDriverValueSlice := func(args []interface{}) []driver.Value {
		result := make([]driver.Value, len(args))
		for i, v := range args {
			result[i] = v
		}
		return result
	}
	tests := []struct {
		name        string
		id          int
		fields      map[string]interface{}
		setupMock   func(mock sqlmock.Sqlmock, id int, fields map[string]interface{}, mockSec mod.Section)
		expected    *mod.Section
		expectedErr error
	}{
		{
			//this was failing because, go doesn't sort maps
			name:   "successful update",
			id:     1,
			fields: map[string]interface{}{"current_capacity": 77, "minimum_capacity": 35},
			setupMock: func(mock sqlmock.Sqlmock, id int, fields map[string]interface{}, mockSec mod.Section) {
				_, args := common.BuildPatchQuery("sections", fields, strconv.Itoa(id), nil)
				mock.ExpectExec("UPDATE sections SET").
					WithArgs(toDriverValueSlice(args)...).
					WillReturnResult(sqlmock.NewResult(0, 1))

				findRows := sqlmock.NewRows([]string{
					"id", "section_number", "current_temperature", "minimum_temperature",
					"current_capacity", "minimum_capacity", "maximum_capacity",
					"warehouse_id", "product_type_id"}).
					AddRow(&mockSec.ID, &mockSec.SectionNumber, &mockSec.CurrentTemperature,
						&mockSec.MinimumTemperature, &mockSec.CurrentCapacity,
						&mockSec.MinimumCapacity, &mockSec.MaximumCapacity,
						&mockSec.WarehouseID, &mockSec.ProductTypeID)

				mock.ExpectQuery(m.SectionSelectWhereExpectedQuery).
					WithArgs(id).
					WillReturnRows(findRows)
			},
			expected: &mod.Section{
				ID:                 1,
				SectionNumber:      1,
				CurrentTemperature: -2,
				MinimumTemperature: -6,
				CurrentCapacity:    77,
				MinimumCapacity:    35,
				MaximumCapacity:    110,
				WarehouseID:        2,
				ProductTypeID:      2,
			},
			expectedErr: nil,
		},
		{
			name:   "foreign key error",
			id:     1,
			fields: map[string]interface{}{"minimum_temperature": -12},
			setupMock: func(mock sqlmock.Sqlmock, id int, fields map[string]interface{}, mockSec mod.Section) {
				_, args := common.BuildPatchQuery("sections", fields, strconv.Itoa(id), nil)
				mock.ExpectExec("UPDATE sections SET").
					WithArgs(toDriverValueSlice(args)...).
					WillReturnError(e.FkErr)
			},
			expected:    &mod.Section{},
			expectedErr: e.ErrForeignKeyError,
		},
		{
			name:   "duplicated key error",
			id:     1,
			fields: map[string]interface{}{"minimum_temperature": -12},
			setupMock: func(mock sqlmock.Sqlmock, id int, fields map[string]interface{}, mockSec mod.Section) {
				_, args := common.BuildPatchQuery("sections", fields, strconv.Itoa(id), nil)
				mock.ExpectExec("UPDATE sections SET").
					WithArgs(toDriverValueSlice(args)...).
					WillReturnError(e.DupErr)
			},
			expected:    &mod.Section{},
			expectedErr: e.ErrSectionRepositoryDuplicated,
		},
		{
			name:   "no rows affected error",
			id:     1,
			fields: map[string]interface{}{"minimum_temperature": -12},
			setupMock: func(mock sqlmock.Sqlmock, id int, fields map[string]interface{}, mockSec mod.Section) {
				_, args := common.BuildPatchQuery("sections", fields, strconv.Itoa(id), nil)
				mock.ExpectExec("UPDATE sections SET").
					WithArgs(toDriverValueSlice(args)...).
					WillReturnResult(sqlmock.NewResult(0, 0))
				findRows := sqlmock.NewRows([]string{
					"id", "section_number", "current_temperature", "minimum_temperature",
					"current_capacity", "minimum_capacity", "maximum_capacity",
					"warehouse_id", "product_type_id"}).
					AddRow(&mockSec.ID, &mockSec.SectionNumber, &mockSec.CurrentTemperature,
						&mockSec.MinimumTemperature, &mockSec.CurrentCapacity,
						&mockSec.MinimumCapacity, &mockSec.MaximumCapacity,
						&mockSec.WarehouseID, &mockSec.ProductTypeID)

				mock.ExpectQuery(m.SectionSelectWhereExpectedQuery).
					WithArgs(id).
					WillReturnRows(findRows)
			},
			expected:    &mod.Section{},
			expectedErr: e.ErrNoRowsAffected,
		}, {
			name:   "not found after update error",
			id:     1,
			fields: map[string]interface{}{"minimum_temperature": -12},
			setupMock: func(mock sqlmock.Sqlmock, id int, fields map[string]interface{}, mockSec mod.Section) {
				_, args := common.BuildPatchQuery("sections", fields, strconv.Itoa(id), nil)
				mock.ExpectExec("UPDATE sections SET").
					WithArgs(toDriverValueSlice(args)...).
					WillReturnResult(sqlmock.NewResult(0, 0))

				mock.ExpectQuery(m.SectionSelectWhereExpectedQuery).
					WithArgs(id).
					WillReturnError(sql.ErrNoRows)
			},
			expected:    &mod.Section{},
			expectedErr: e.ErrSectionRepositoryNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockSectionRepo(t)
			defer teardown()
			tc.setupMock(mock, tc.id, tc.fields, *tc.expected)
			result, err := repo.Update(tc.id, tc.fields)
			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, tc.expectedErr, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, tc.expected)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestSectionDB_ReportProducts(t *testing.T) {
	toDriverValueSlice := func(args []int) []driver.Value {
		result := make([]driver.Value, len(args))
		for i, v := range args {
			result[i] = v
		}
		return result
	}

	type testCase struct {
		name       string
		setupMock  func(sqlmock.Sqlmock, []int)
		ids        []int
		wantResult []mod.ReportProductsResponse
		wantErr    error
	}

	tests := []testCase{
		{
			name: "happy path - all IDs found",
			ids:  []int{1, 2},
			setupMock: func(mock sqlmock.Sqlmock, ids []int) {
				// Row with products for id=1 and id=2
				rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
					AddRow(1, 101, 7).
					AddRow(2, 202, 3)
				mock.ExpectQuery("SELECT (.+) FROM (.+)").
					WithArgs(toDriverValueSlice(ids)...).
					WillReturnRows(rows)
			},
			wantResult: []mod.ReportProductsResponse{
				{SectionId: 1, SectionNumber: 101, ProductsCount: 7},
				{SectionId: 2, SectionNumber: 202, ProductsCount: 3},
			},
			wantErr: nil,
		},
		{
			name: "error from db.Query",
			ids:  []int{1, 2},
			setupMock: func(mock sqlmock.Sqlmock, ids []int) {
				mock.ExpectQuery("SELECT (.+) FROM (.+)").
					WithArgs(toDriverValueSlice(ids)...).
					WillReturnError(e.ErrQueryError)
			},
			wantResult: nil,
			wantErr:    e.ErrQueryError,
		},
		{
			name: "some IDs not found",
			ids:  []int{1, 2, 3},
			setupMock: func(mock sqlmock.Sqlmock, ids []int) {
				// Only id 1 and 2 found in query result, id 3 is missing!
				rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
					AddRow(1, 101, 7).
					AddRow(2, 202, 3)
				mock.ExpectQuery("SELECT (.+) FROM (.+)").
					WithArgs(toDriverValueSlice(ids)...).
					WillReturnRows(rows)
			},
			wantResult: nil,
			wantErr:    e.ErrSectionRepositoryNotFound,
		},
		{
			name: "no ids filter, returns all",
			ids:  nil,
			setupMock: func(mock sqlmock.Sqlmock, ids []int) {
				rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
					AddRow(2, 202, 2).
					AddRow(4, 404, 0)
				mock.ExpectQuery("SELECT (.+) FROM (.+)").
					WithoutArgs().
					WillReturnRows(rows)
			},
			wantResult: []mod.ReportProductsResponse{
				{SectionId: 2, SectionNumber: 202, ProductsCount: 2},
				{SectionId: 4, SectionNumber: 404, ProductsCount: 0},
			},
			wantErr: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockSectionRepo(t)
			defer teardown()
			tc.setupMock(mock, tc.ids)
			got, err := repo.ReportProducts(tc.ids)
			if tc.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.wantErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantResult, got)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
