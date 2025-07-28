package repository

import (
	m "github.com/smartineztri_meli/W17-G2-Bootcamp/tests/mock"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// Helper
func setupMockProductBatchRepo(t *testing.T) (*ProductBatchDB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	repo := NewProductBatchRepo(db)
	return repo, mock, func() { db.Close() }
}

// -- FIND ALL
func TestProductBatchDB_FindAll(t *testing.T) {
	tests := []struct {
		name        string
		mockQuery   func(mock sqlmock.Sqlmock)
		expected    []mod.ProductBatch
		expectedErr error
	}{
		{
			name: "DB with data",
			mockQuery: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows(m.ProductBatchTableStruct).
					AddRows(m.ProductBatchDataValuesSelect...)
				mock.ExpectQuery(m.ProductBatchSelectExpectedQuery).WillReturnRows(rows)
			},
			expected: []mod.ProductBatch{
				{
					ID:                 1,
					BatchNumber:        1,
					CurrentQuantity:    200,
					InitialQuantity:    200,
					CurrentTemperature: 2,
					MinimumTemperature: -5,
					DueDate:            time.Date(2024, 07, 05, 17, 00, 00, 0, time.UTC),
					ManufacturingDate:  time.Date(2024, 06, 1, 0, 00, 00, 0, time.UTC),
					ManufacturingHour:  "08:00:00",
					ProductId:          1,
					SectionId:          1,
				},
				{
					ID:                 2,
					BatchNumber:        2,
					CurrentQuantity:    310,
					InitialQuantity:    310,
					CurrentTemperature: -2,
					MinimumTemperature: -6,
					DueDate:            time.Date(2024, 8, 01, 12, 00, 00, 0, time.UTC),
					ManufacturingDate:  time.Date(2024, 7, 1, 0, 00, 00, 0, time.UTC),
					ManufacturingHour:  "09:30:00",
					ProductId:          2,
					SectionId:          2,
				},
			},
			expectedErr: nil,
		},
		{
			name: "Err empty DB",
			mockQuery: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{`id`, `batch_number`})
				mock.ExpectQuery(m.ProductBatchSelectExpectedQuery).WillReturnRows(rows)
			},
			expected:    []mod.ProductBatch{},
			expectedErr: e.ErrEmptyDB,
		},
		{
			name: "Query Error",
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(m.ProductBatchSelectExpectedQuery).WillReturnError(e.ErrQueryError)
			},
			expected:    []mod.ProductBatch{},
			expectedErr: e.ErrQueryError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockProductBatchRepo(t)
			defer teardown()
			tc.mockQuery(mock)

			batches, err := repo.FindAll()

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Len(t, batches, len(tc.expected))
				require.Equal(t, tc.expected, batches)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// -- SAVE
func TestProductBatchDB_Save(t *testing.T) {
	baseBatch := func() *mod.ProductBatch {
		return &mod.ProductBatch{
			BatchNumber:        1,
			CurrentQuantity:    2,
			InitialQuantity:    1,
			CurrentTemperature: 3,
			MinimumTemperature: 0,
			DueDate:            time.Date(2024, 8, 01, 12, 00, 00, 0, time.UTC),
			ManufacturingDate:  time.Date(2024, 7, 1, 0, 00, 00, 0, time.UTC),
			ManufacturingHour:  "9:00:00",
			ProductId:          1,
			SectionId:          1,
		}
	}

	type testcase struct {
		name    string
		setup   func(sqlmock.Sqlmock, *mod.ProductBatch)
		wantID  int
		wantErr error
	}

	tests := []testcase{
		{
			name: "HappyPath",
			setup: func(mock sqlmock.Sqlmock, batch *mod.ProductBatch) {
				mock.ExpectExec("INSERT INTO `product_batches`").
					WithArgs(batch.BatchNumber, batch.CurrentQuantity, batch.InitialQuantity,
						batch.CurrentTemperature, batch.MinimumTemperature, batch.DueDate,
						batch.ManufacturingDate, batch.ManufacturingHour, batch.ProductId, batch.SectionId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantID:  1,
			wantErr: nil,
		},
		{
			name: "ErrDuplicated",
			setup: func(mock sqlmock.Sqlmock, batch *mod.ProductBatch) {
				mock.ExpectExec("INSERT INTO `product_batches`").
					WithArgs(batch.BatchNumber, batch.CurrentQuantity, batch.InitialQuantity,
						batch.CurrentTemperature, batch.MinimumTemperature, batch.DueDate,
						batch.ManufacturingDate, batch.ManufacturingHour, batch.ProductId, batch.SectionId).
					WillReturnError(e.DupErr)
			},
			wantID:  0,
			wantErr: e.ErrProductBatchDuplicated,
		},
		{
			name: "ErrForeignKey",
			setup: func(mock sqlmock.Sqlmock, batch *mod.ProductBatch) {
				mock.ExpectExec("INSERT INTO `product_batches`").
					WithArgs(batch.BatchNumber, batch.CurrentQuantity, batch.InitialQuantity,
						batch.CurrentTemperature, batch.MinimumTemperature, batch.DueDate,
						batch.ManufacturingDate, batch.ManufacturingHour, batch.ProductId, batch.SectionId).
					WillReturnError(e.FkErr)
			},
			wantID:  0,
			wantErr: e.ErrForeignKeyError,
		},
		{
			name: "ErrNotFound (LastInsertId error)",
			setup: func(mock sqlmock.Sqlmock, batch *mod.ProductBatch) {
				// FakeResult implements RowsAffected/LastInsertId simulating failure
				mock.ExpectExec("INSERT INTO `product_batches`").
					WithArgs(batch.BatchNumber, batch.CurrentQuantity, batch.InitialQuantity,
						batch.CurrentTemperature, batch.MinimumTemperature, batch.DueDate,
						batch.ManufacturingDate, batch.ManufacturingHour, batch.ProductId, batch.SectionId).
					WillReturnResult(e.FakeResult{})
			},
			wantID:  0,
			wantErr: e.ErrProductBatchNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockProductBatchRepo(t)
			defer teardown()
			batch := baseBatch()
			tc.setup(mock, batch)

			err := repo.Save(batch)
			if tc.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.wantErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantID, batch.ID)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
