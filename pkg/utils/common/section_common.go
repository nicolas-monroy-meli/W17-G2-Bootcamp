package common

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"net/http"
	"strconv"
	"strings"
)

// IdRequests encapsulates the process of getting the id parameter and returns an int number and an error if necessary
func IdRequests(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return 0, e.EmptyParams
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
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`
        SELECT 
            s.id, 
            s.section_number, 
            COUNT(p.id) as product_count
        FROM sections s
        LEFT JOIN products p ON p.id = s.product_type_id
        WHERE s.id IN (%s)
        GROUP BY s.id, s.section_number
        ORDER BY s.id
    `, strings.Join(placeholders, ","))

	return query, args
}

func ParseWarehouseIDs(param string) ([]int, error) {
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

	if len(ids) == 0 {
		return nil, fmt.Errorf("no valid warehouse IDs provided")
	}

	return ids, nil
}
