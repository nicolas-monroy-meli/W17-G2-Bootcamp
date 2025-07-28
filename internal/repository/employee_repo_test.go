package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// Helper para configurar el mock del repositorio de empleados.
func setupMockEmployeeRepo(t *testing.T) (*EmployeeDB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	repo := NewEmployeeRepo(db)
	return repo, mock, func() { db.Close() }
}

// -- FIND ALL --
func TestEmployeeDB_FindAll(t *testing.T) {
	// Datos de prueba
	employeeCols := []string{"id", "id_card_number", "first_name", "last_name", "wareHouse_id"}
	mockedRows := sqlmock.NewRows(employeeCols).
		AddRow(1, "12345", "John", "Doe", 1).
		AddRow(2, "67890", "Jane", "Smith", 1)

	expectedEmployees := []mod.Employee{
		{ID: 1, CardNumberID: "12345", FirstName: "John", LastName: "Doe", WarehouseID: 1},
		{ID: 2, CardNumberID: "67890", FirstName: "Jane", LastName: "Smith", WarehouseID: 1},
	}

	query := "SELECT id,id_card_number,first_name,last_name, wareHouse_id FROM employees"

	// Casos de prueba
	testCases := []struct {
		name        string
		mockQuery   func(mock sqlmock.Sqlmock)
		expected    []mod.Employee
		expectedErr error
	}{
		{
			name: "HappyPath_FindAll",
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(mockedRows)
			},
			expected:    expectedEmployees,
			expectedErr: nil,
		},
		{
			name: "Err_EmptyDB",
			mockQuery: func(mock sqlmock.Sqlmock) {
				emptyRows := sqlmock.NewRows(employeeCols)
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(emptyRows)
			},
			expected:    nil,
			expectedErr: e.ErrEmployeeRepositoryNotFound,
		},
		{
			name: "Err_QueryError",
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(errors.New("failed to query employees"))
			},
			expected:    nil,
			expectedErr: errors.New("failed to query employees"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockEmployeeRepo(t)
			defer teardown()

			tc.mockQuery(mock)
			employees, err := repo.FindAll()

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, employees)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// -- FIND BY ID --
func TestEmployeeDB_FindByID(t *testing.T) {
	employeeCols := []string{"id", "id_card_number", "first_name", "last_name", "wareHouse_id"}
	query := "SELECT id,id_card_number,first_name,last_name, wareHouse_id FROM employees WHERE id = ?"
	mockedRow := sqlmock.NewRows(employeeCols).AddRow(1, "12345", "John", "Doe", 1)

	expectedEmployee := mod.Employee{ID: 1, CardNumberID: "12345", FirstName: "John", LastName: "Doe", WarehouseID: 1}

	testCases := []struct {
		name        string
		inputID     int
		mockQuery   func(mock sqlmock.Sqlmock)
		expected    mod.Employee
		expectedErr error
	}{
		{
			name:    "HappyPath_FindByID",
			inputID: 1,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(mockedRow)
			},
			expected:    expectedEmployee,
			expectedErr: nil,
		},
		{
			name:    "Err_NotFound",
			inputID: 99,
			mockQuery: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(99).WillReturnError(errors.New("repository: employee not found"))
			},
			expected:    mod.Employee{},
			expectedErr: errors.New("failed to scan employee by ID"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockEmployeeRepo(t)
			defer teardown()

			tc.mockQuery(mock)
			employee, err := repo.FindByID(tc.inputID)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, employee)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// -- SAVE --
func TestEmployeeDB_Save(t *testing.T) {
	query := "INSERT INTO employees (id_card_number,first_name,last_name, wareHouse_id ) VALUES (?, ?,?,?)"
	employeeToSave := &mod.Employee{CardNumberID: "12345", FirstName: "John", LastName: "Doe", WarehouseID: 1}

	testCases := []struct {
		name        string
		mockExec    func(mock sqlmock.Sqlmock)
		expectedID  int
		expectedErr error
	}{
		{
			name: "HappyPath_Save",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(employeeToSave.CardNumberID, employeeToSave.FirstName, employeeToSave.LastName, employeeToSave.WarehouseID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedID:  1,
			expectedErr: nil,
		},
		{
			name: "Err_ExecFailed",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(employeeToSave.CardNumberID, employeeToSave.FirstName, employeeToSave.LastName, employeeToSave.WarehouseID).
					WillReturnError(errors.New("failed to insert employee"))
			},
			expectedID:  0,
			expectedErr: errors.New("failed to insert employee"),
		},
		{
			name: "Err_LastInsertIdFailed",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(employeeToSave.CardNumberID, employeeToSave.FirstName, employeeToSave.LastName, employeeToSave.WarehouseID).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("failed to get last insert ID")))
			},
			expectedID:  0,
			expectedErr: errors.New("failed to get last insert ID"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockEmployeeRepo(t)
			defer teardown()

			tc.mockExec(mock)
			err := repo.Save(employeeToSave)

			if tc.expectedErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedID, employeeToSave.ID)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// -- UPDATE --
func TestEmployeeDB_Update(t *testing.T) {
	query := "UPDATE employees SET id_card_number = ?, first_name = ?, last_name = ?, wareHouse_id = ? WHERE id = ?"
	employeeToUpdate := &mod.Employee{CardNumberID: "54321", FirstName: "John", LastName: "Updated", WarehouseID: 2}
	targetID := 1

	testCases := []struct {
		name        string
		mockExec    func(mock sqlmock.Sqlmock)
		expectedErr error
	}{
		{
			name: "HappyPath_Update",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(employeeToUpdate.CardNumberID, employeeToUpdate.FirstName, employeeToUpdate.LastName, employeeToUpdate.WarehouseID, targetID).
					WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected
			},
			expectedErr: nil,
		},
		{
			name: "Err_NotFound",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(employeeToUpdate.CardNumberID, employeeToUpdate.FirstName, employeeToUpdate.LastName, employeeToUpdate.WarehouseID, targetID).
					WillReturnResult(sqlmock.NewResult(0, 0)) // 0 rows affected
			},
			expectedErr: e.ErrEmployeeRepositoryNotFound,
		},
		{
			name: "Err_ExecFailed",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).
					WithArgs(employeeToUpdate.CardNumberID, employeeToUpdate.FirstName, employeeToUpdate.LastName, employeeToUpdate.WarehouseID, targetID).
					WillReturnError(errors.New("failed to update employee"))
			},
			expectedErr: errors.New("failed to update employee"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockEmployeeRepo(t)
			defer teardown()

			tc.mockExec(mock)
			err := repo.Update(targetID, employeeToUpdate)

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

// -- DELETE --
func TestEmployeeDB_Delete(t *testing.T) {
	query := "DELETE FROM employees WHERE id = ?"
	targetID := 1

	testCases := []struct {
		name        string
		mockExec    func(mock sqlmock.Sqlmock)
		expectedErr error
	}{
		{
			name: "HappyPath_Delete",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(targetID).
					WillReturnResult(sqlmock.NewResult(0, 1)) // 1 row affected
			},
			expectedErr: nil,
		},
		{
			name: "Err_NotFound",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(targetID).
					WillReturnResult(sqlmock.NewResult(0, 0)) // 0 rows affected
			},
			expectedErr: e.ErrEmployeeRepositoryNotFound,
		},
		{
			name: "Err_ExecFailed",
			mockExec: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(targetID).
					WillReturnError(errors.New("failed to delete employee"))
			},
			expectedErr: errors.New("failed to delete employee"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo, mock, teardown := setupMockEmployeeRepo(t)
			defer teardown()

			tc.mockExec(mock)
			err := repo.Delete(targetID)

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
