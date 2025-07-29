package common_test

import (
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPatchEmployees(t *testing.T) {
	// Define la tabla de pruebas
	tests := []struct {
		name     string
		existing models.Employee
		updates  models.Employee
		expected models.Employee
	}{
		{
			name: "should update all provided fields",
			existing: models.Employee{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
			updates: models.Employee{
				CardNumberID: "456",
				FirstName:    "Jane",
				LastName:     "Smith",
				WarehouseID:  20,
			},
			expected: models.Employee{
				ID:           1,
				CardNumberID: "456",
				FirstName:    "Jane",
				LastName:     "Smith",
				WarehouseID:  20,
			},
		},
		{
			name: "should update only CardNumberID",
			existing: models.Employee{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
			updates: models.Employee{
				CardNumberID: "456",
			},
			expected: models.Employee{
				ID:           1,
				CardNumberID: "456",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
		},
		{
			name: "should return original employee if no updates are provided",
			existing: models.Employee{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
			updates: models.Employee{},
			expected: models.Employee{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
		},
		{
			name: "should not change fields if update value is same as existing",
			existing: models.Employee{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
			updates: models.Employee{
				CardNumberID: "123",  // Mismo que existente
				FirstName:    "John", // Mismo que existente
				WarehouseID:  10,     // Mismo que existente
			},
			expected: models.Employee{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
		},
		{
			name: "should handle partial updates correctly",
			existing: models.Employee{
				ID:           1,
				CardNumberID: "ABC",
				FirstName:    "Mike",
				LastName:     "Tyson",
				WarehouseID:  50,
			},
			updates: models.Employee{
				FirstName:   "Michael",
				WarehouseID: 60,
			},
			expected: models.Employee{
				ID:           1,
				CardNumberID: "ABC", // Sin cambios
				FirstName:    "Michael",
				LastName:     "Tyson", // Sin cambios
				WarehouseID:  60,
			},
		},
		{
			name: "should not change ID field", // Asegura que el ID no se modifica
			existing: models.Employee{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "John",
				LastName:     "Doe",
				WarehouseID:  10,
			},
			updates: models.Employee{
				ID:        99, // Intentando cambiar el ID, que no está en la lógica de parcheo
				FirstName: "Jane",
			},
			expected: models.Employee{
				ID:           1, // El ID debe permanecer el original
				CardNumberID: "123",
				FirstName:    "Jane",
				LastName:     "Doe",
				WarehouseID:  10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			patched := common.PatchEmployees(tt.updates, tt.existing)
			assert.Equal(t, tt.expected, patched, "The patched employee did not match the expected output")
		})
	}
}
