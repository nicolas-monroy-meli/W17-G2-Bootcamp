package common

import (
	"fmt"
	"strings"
)

var SectionPatchFieldOrder = []string{
	"section_number",
	"current_temperature",
	"minimum_temperature",
	"current_capacity",
	"minimum_capacity",
	"maximum_capacity",
	"warehouse_id",
	"product_type_id",
}

func BuildPatchQuery(table string, fields map[string]interface{}, where string, fieldOrder []string, whereArgs ...interface{}) (string, []interface{}) {
	setClauses := []string{}
	args := []interface{}{}
	if fieldOrder == nil {
		fieldOrder = SectionPatchFieldOrder
	}
	for _, k := range fieldOrder {
		if v, ok := fields[k]; ok {
			setClauses = append(setClauses, fmt.Sprintf("%s = ?", k))
			args = append(args, v)
		}
	}
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%s",
		table, strings.Join(setClauses, ", "), where)
	args = append(args, whereArgs...)
	return query, args
}
