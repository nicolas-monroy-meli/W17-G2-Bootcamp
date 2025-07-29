package common

import (
	"errors"
	"testing"

	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestValidateWarehouseUpdate(t *testing.T) {
	// Caso base válido
	baseWarehouse := models.Warehouse{
		Address:            "Calle Falsa 123",
		Telephone:          "123456789",
		WarehouseCode:      "WH001",
		MinimumCapacity:    10,
		MinimumTemperature: 20,
	}

	tests := []struct {
		name        string
		warehouse   models.Warehouse
		expectedErr error
	}{
		// Casos válidos
		{
			name:        "valid_warehouse",
			warehouse:   baseWarehouse,
			expectedErr: nil,
		},
		{
			name: "valid_min_temp_-50",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.MinimumTemperature = -50
				return w
			}(),
			expectedErr: nil,
		},
		{
			name: "valid_min_temp_50",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.MinimumTemperature = 50
				return w
			}(),
			expectedErr: nil,
		},

		// Casos inválidos
		{
			name: "empty_address",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.Address = ""
				return w
			}(),
			expectedErr: errors.New("el campo 'address' es requerido"),
		},
		{
			name: "whitespace_address",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.Address = "   "
				return w
			}(),
			expectedErr: errors.New("el campo 'address' es requerido"),
		},
		{
			name: "empty_telephone",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.Telephone = ""
				return w
			}(),
			expectedErr: errors.New("el campo 'telephone' es requerido"),
		},
		{
			name: "whitespace_telephone",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.Telephone = "   "
				return w
			}(),
			expectedErr: errors.New("el campo 'telephone' es requerido"),
		},
		{
			name: "empty_warehouse_code",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.WarehouseCode = ""
				return w
			}(),
			expectedErr: errors.New("el campo 'warehouse_code' es requerido"),
		},
		{
			name: "whitespace_warehouse_code",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.WarehouseCode = "   "
				return w
			}(),
			expectedErr: errors.New("el campo 'warehouse_code' es requerido"),
		},
		{
			name: "zero_min_capacity",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.MinimumCapacity = 0
				return w
			}(),
			expectedErr: errors.New("el campo 'minimum_capacity' debe ser mayor a 0"),
		},
		{
			name: "negative_min_capacity",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.MinimumCapacity = -5
				return w
			}(),
			expectedErr: errors.New("el campo 'minimum_capacity' debe ser mayor a 0"),
		},
		{
			name: "min_temp_below_-50",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.MinimumTemperature = -51
				return w
			}(),
			expectedErr: errors.New("el campo 'minimum_temperature' debe estar entre -50 y 50"),
		},
		{
			name: "min_temp_above_50",
			warehouse: func() models.Warehouse {
				w := baseWarehouse
				w.MinimumTemperature = 51
				return w
			}(),
			expectedErr: errors.New("el campo 'minimum_temperature' debe estar entre -50 y 50"),
		},
		{
			name: "multiple_errors_returns_first",
			warehouse: models.Warehouse{
				Address:            "",  // Error 1
				Telephone:          "",  // Error 2
				WarehouseCode:      "",  // Error 3
				MinimumCapacity:    0,   // Error 4
				MinimumTemperature: 100, // Error 5
			},
			expectedErr: errors.New("el campo 'address' es requerido"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateWarehouseUpdate(tt.warehouse)
			if tt.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
