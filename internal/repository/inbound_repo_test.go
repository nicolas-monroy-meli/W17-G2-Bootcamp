package repository

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

// Helper para configurar el mock del repositorio de inbound orders.
func setupMockInboundRepo(t *testing.T) (*InboundDB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	repo := NewInboundRepo(db)
	return repo, mock, func() { db.Close() }
}

// -- SAVE --
func TestInboundDB_Save(t *testing.T) {
	query := `INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id)
              VALUES (?, ?, ?, ?, ?)`

	baseOrder := func() *mod.InboundOrders {
		return &mod.InboundOrders{
			OrderDate:      "2025-07-28",
			OrderNumber:    "ORD-001",
			EmployeeId:     1,
			ProductBatchId: 1,
			WarehouseId:    1,
		}
	}

	testCases := []struct {
		name        string
		setup       func(mock sqlmock.Sqlmock, order *mod.InboundOrders)
		input       *mod.InboundOrders
		expectedID  int
		expectedErr error
	}{
		{
			name: "HappyPath",
			setup: func(mock sqlmock.Sqlmock, order *mod.InboundOrders) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(order.OrderDate, order.OrderNumber, order.EmployeeId, order.ProductBatchId, order.WarehouseId).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			input:       baseOrder(),
			expectedID:  1,
			expectedErr: nil,
		},
		{
			name: "Err_ExecFailed",
			setup: func(mock sqlmock.Sqlmock, order *mod.InboundOrders) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(order.OrderDate, order.OrderNumber, order.EmployeeId, order.ProductBatchId, order.WarehouseId).
					WillReturnError(sql.ErrConnDone)
			},
			input:       baseOrder(),
			expectedID:  0,
			expectedErr: e.ErrInboundOrderInternal,
		},
		{
			name: "Err_LastInsertIdFailed",
			setup: func(mock sqlmock.Sqlmock, order *mod.InboundOrders) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(order.OrderDate, order.OrderNumber, order.EmployeeId, order.ProductBatchId, order.WarehouseId).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("last id error")))
			},
			input:       baseOrder(),
			expectedID:  0,
			expectedErr: e.ErrInboundOrderInternal,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockInboundRepo(t)
			defer teardown()

			tc.setup(mock, tc.input)
			savedOrder, err := repo.Save(tc.input)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tc.expectedErr))
			} else {
				require.NoError(t, err)
				require.NotNil(t, savedOrder)
				require.Equal(t, tc.expectedID, savedOrder.Id)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// -- FindOrdersByEmployee --
func TestInboundDB_FindOrdersByEmployee(t *testing.T) {
	baseQuery := `SELECT e.id, e.id_card_number, e.first_name, e.last_name, e.wareHouse_id, COUNT(io.id) AS inbound_orders_count
                   FROM employees AS e LEFT JOIN inbound_orders AS io ON e.id = io.employee_id`
	reportCols := []string{"id", "id_card_number", "first_name", "last_name", "wareHouse_id", "inbound_orders_count"}

	testCases := []struct {
		name        string
		employeeID  int
		setup       func(mock sqlmock.Sqlmock)
		expectedLen int
		expectedErr error
	}{
		{
			name:       "HappyPath_SpecificEmployee",
			employeeID: 1,
			setup: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(baseQuery + " WHERE e.id = ? GROUP BY e.id")
				rows := sqlmock.NewRows(reportCols).
					AddRow(1, "12345", "John", "Doe", 1, 5)
				mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			expectedLen: 1,
			expectedErr: nil,
		},
		{
			name:       "HappyPath_AllEmployees",
			employeeID: 0,
			setup: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(baseQuery + " GROUP BY e.id")
				rows := sqlmock.NewRows(reportCols).
					AddRow(1, "12345", "John", "Doe", 1, 5).
					AddRow(2, "67890", "Jane", "Smith", 1, 3)
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			expectedLen: 2,
			expectedErr: nil,
		},
		{
			name:       "Err_EmployeeNotFound",
			employeeID: 99,
			setup: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(baseQuery + " WHERE e.id = ? GROUP BY e.id")
				rows := sqlmock.NewRows(reportCols) // Sin filas
				mock.ExpectQuery(query).WithArgs(99).WillReturnRows(rows)
			},
			expectedLen: 0,
			expectedErr: e.ErrEmployeeNotFound,
		},
		{
			name:       "Err_QueryFailed",
			employeeID: 1,
			setup: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(baseQuery + " WHERE e.id = ? GROUP BY e.id")
				mock.ExpectQuery(query).WithArgs(1).WillReturnError(sql.ErrConnDone)
			},
			expectedLen: 0,
			expectedErr: e.ErrEmployeeInternal,
		},
		{
			name:       "Err_ScanFailed",
			employeeID: 0,
			setup: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(baseQuery + " GROUP BY e.id")
				rows := sqlmock.NewRows(reportCols).
					AddRow(1, "12345", "John", "Doe", 1, 5).
					RowError(0, errors.New("scan error")) // Error en la primera fila
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			expectedLen: 0,
			expectedErr: e.ErrEmployeeInternal,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockInboundRepo(t)
			defer teardown()

			tc.setup(mock)
			reports, err := repo.FindOrdersByEmployee(tc.employeeID)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.True(t, errors.Is(err, tc.expectedErr))
			} else {
				require.NoError(t, err)
				require.Len(t, reports, tc.expectedLen)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
