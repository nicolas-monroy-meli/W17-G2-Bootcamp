package common

import (
	"context"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"github.com/stretchr/testify/require"
)

func TestIdRequests(t *testing.T) {
	tests := []struct {
		name       string
		idParam    string
		wantID     int
		wantErrStr string
	}{
		{"valid", "42", 42, ""},
		{"empty", "", 0, e.ErrQueryIsEmpty.Error()},
		{"not a number", "bob", 0, "unexpected:"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			// set up chi url param
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.idParam)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			got, err := IdRequests(req)
			if tc.wantErrStr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.wantErrStr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantID, got)
			}
		})
	}
}

func TestPatchSection(t *testing.T) {
	sectionNumber := 1
	temp := 10.3
	MinimumTemperature := 10.2
	CurrentCapacity := 1
	MinimumCapacity := 1
	MaximumCapacity := 3
	WarehouseID := 2
	ProductTypeID := 2
	patchTest := mod.SectionPatch{
		SectionNumber:      &sectionNumber,
		CurrentTemperature: &temp,
		MinimumTemperature: &MinimumTemperature,
		CurrentCapacity:    &CurrentCapacity,
		MinimumCapacity:    &MinimumCapacity,
		MaximumCapacity:    &MaximumCapacity,
		WarehouseID:        &WarehouseID,
		ProductTypeID:      &ProductTypeID,
	}
	tests := []struct {
		name    string
		patch   mod.SectionPatch
		wantMap map[string]interface{}
	}{
		{
			"all nil", mod.SectionPatch{}, map[string]interface{}{},
		},
		{
			"one field", mod.SectionPatch{SectionNumber: &sectionNumber}, map[string]interface{}{"section_number": 1},
		},
		{
			"multiple fields", patchTest,
			map[string]interface{}{"section_number": 1, "current_temperature": 10.3, "minimum_temperature": 10.2, "current_capacity": 1, "minimum_capacity": 1, "maximum_capacity": 3, "warehouse_id": 2, "product_type_id": 2},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PatchSection(tc.patch)
			require.Equal(t, tc.wantMap, got)
		})
	}
}

func TestGetQueryReport(t *testing.T) {
	tests := []struct {
		name      string
		ids       []int
		wantWhere string // part of SQL to match for IDs filter
		wantArgs  []interface{}
	}{
		{"empty (all)", nil, "WHERE", nil},
		{"one id", []int{3}, "WHERE s.id IN (?)", []interface{}{3}},
		{"many ids", []int{2, 4, 8}, "WHERE s.id IN (?,?,?)\n", []interface{}{2, 4, 8}},
		{"no filter", []int{}, "WHERE", nil}, // this maybe generates no WHERE clause!
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sql, args := GetQueryReport(tc.ids)
			if len(tc.ids) == 0 {
				require.NotContains(t, sql, "WHERE s.id IN")
				require.Equal(t, 0, len(args))
			} else {
				require.Contains(t, sql, tc.wantWhere)
				require.Equal(t, tc.wantArgs, args)
			}
			require.Contains(t, sql, "LEFT JOIN products p ON")
		})
	}
}

func TestParseIDs(t *testing.T) {
	tests := []struct {
		name       string
		param      string
		want       []int
		wantErrStr string
	}{
		{"empty", "", []int{}, ""},
		{"empty-with coma", " , ", []int{}, ""},
		{"single id", "5", []int{5}, ""},
		{"trimmed csv", "3, 9,22", []int{3, 9, 22}, ""},
		{"invalid int", "8,bob,3", nil, "must be an integer"},
		{"zero", "0,2", nil, "must be greater than 0"},
		{"all whitespace", "  ", []int{}, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseIDs(tc.param)
			if tc.wantErrStr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.wantErrStr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}
