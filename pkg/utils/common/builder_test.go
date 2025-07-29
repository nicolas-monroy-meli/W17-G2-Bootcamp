package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuildPatchQuery(t *testing.T) {
	tests := []struct {
		name       string
		fields     map[string]interface{}
		where      string
		fieldOrder []string
		whereArgs  []interface{}
		wantSQL    string
		wantArgs   []interface{}
	}{
		{
			name:       "patch section_number and current_capacity, default order",
			fields:     map[string]interface{}{"section_number": 5, "current_capacity": 44},
			where:      "7",
			fieldOrder: nil, // no specific order
			wantSQL:    "UPDATE sections SET section_number = ?, current_capacity = ? WHERE id=7",
			wantArgs:   []interface{}{5, 44},
		},
		{
			name:       "different field order",
			fields:     map[string]interface{}{"section_number": 5, "current_capacity": 44},
			where:      "8",
			fieldOrder: []string{"current_capacity", "section_number"},
			wantSQL:    "UPDATE sections SET current_capacity = ?, section_number = ? WHERE id=8",
			wantArgs:   []interface{}{44, 5},
		},
		{
			name:       "one field, where args",
			fields:     map[string]interface{}{"section_number": 45},
			where:      "?",
			whereArgs:  []interface{}{9},
			fieldOrder: nil,
			wantSQL:    "UPDATE sections SET section_number = ? WHERE id=?",
			wantArgs:   []interface{}{45, 9},
		},
		{
			name:       "empty fields map",
			fields:     map[string]interface{}{},
			where:      "12",
			fieldOrder: nil,
			wantSQL:    "UPDATE sections SET  WHERE id=12",
			wantArgs:   []interface{}{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sqlStr, args := BuildPatchQuery("sections", tc.fields, tc.where, tc.fieldOrder, tc.whereArgs...)
			require.Equal(t, tc.wantSQL, sqlStr)
			require.Equal(t, tc.wantArgs, args)
		})
	}
}
