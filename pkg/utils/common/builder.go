package common

import (
	"fmt"
	"strings"
)

func BuildPatchQuery(table string, fields map[string]interface{}, where string, whereArgs ...interface{}) (string, []interface{}) {
	setClauses := []string{}
	args := []interface{}{}

	for col, val := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
		args = append(args, val)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		table,
		strings.Join(setClauses, ", "),
		where,
	)
	args = append(args, whereArgs...)

	return query, args
}
