package common

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
)

// IdRequests encapsulates the process of getting the id parameter and returns an int number and an error if necessary
func IdRequests(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return 0, e.ErrQueryIsEmpty
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("unexpected: %s", err.Error()))
	}
	return id, nil
}

func PatchSection(request models.SectionPatch) map[string]interface{} {
	fields := map[string]interface{}{}
	if request.SectionNumber != nil {
		fields["section_number"] = *request.SectionNumber
	}
	if request.CurrentTemperature != nil {
		fields["current_temperature"] = *request.CurrentTemperature
	}
	if request.MinimumTemperature != nil {
		fields["minimum_temperature"] = *request.MinimumTemperature
	}
	if request.CurrentCapacity != nil {
		fields["current_capacity"] = *request.CurrentCapacity
	}
	if request.MinimumCapacity != nil {
		fields["minimum_capacity"] = *request.MinimumCapacity
	}
	if request.MaximumCapacity != nil {
		fields["maximum_capacity"] = *request.MaximumCapacity
	}
	if request.WarehouseID != nil {
		fields["warehouse_id"] = *request.WarehouseID
	}
	if request.ProductTypeID != nil {
		fields["product_type_id"] = *request.ProductTypeID
	}
	return fields
}

func GetQueryReport(ids []int) (string, []interface{}) {
	// Common SQL
	base := `
        SELECT 
            s.id, 
            s.section_number, 
            COUNT(p.id) as product_count
        FROM sections s
        LEFT JOIN products p ON p.id = s.product_type_id
    `
	whereClause := ""
	args := []interface{}{}

	// If ID list is provided, filter, else show all
	if len(ids) > 0 {
		placeholders := make([]string, len(ids))
		for i := range ids {
			placeholders[i] = "?"
		}
		whereClause = fmt.Sprintf("WHERE s.id IN (%s)", strings.Join(placeholders, ","))
		for _, id := range ids {
			args = append(args, id)
		}
	}

	sql := fmt.Sprintf(`
        %s
        %s
        GROUP BY s.id, s.section_number
        ORDER BY s.id
    `, base, whereClause)
	return sql, args
}

func ParseIDs(param string) ([]int, error) {
	param = strings.TrimSpace(param)
	if param == "" {
		// No ids provided, means "select all"
		return []int{}, nil
	}
	idStrings := strings.Split(param, ",")
	ids := make([]int, 0, len(idStrings))

	for _, idStr := range idStrings {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, fmt.Errorf("'%s' is not a valid integer", idStr)
		}
		if id <= 0 {
			return nil, fmt.Errorf("section ID must be positive, got %d", id)
		}
		ids = append(ids, id)
	}
	// Do not error if ids are empty: caller interprets [] as "fetch all"
	return ids, nil
}
