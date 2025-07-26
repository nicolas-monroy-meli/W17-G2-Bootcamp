package tests

import (
	"database/sql/driver"
	"time"
)

var ProductBatchTableStruct = []string{"id", "batch_number", "current_quantity", "initial_quantity", "current_temperature", "minimum_temperature", "due_date", "manufacturing_date", "manufacturing_hour", "product_id", "section_id"}

var ProductBatchDataValuesSelect = [][]driver.Value{
	{
		1, 1, 200, 200, 2, -5, time.Date(2024, 07, 05, 17, 00, 00, 0, time.UTC), time.Date(2024, 06, 1, 0, 00, 00, 0, time.UTC), "08:00:00", 1, 1,
	},
	{
		2, 2, 310, 310, -2, -6, time.Date(2024, 8, 01, 12, 00, 00, 0, time.UTC), time.Date(2024, 7, 1, 0, 00, 00, 0, time.UTC), "09:30:00", 2, 2,
	},
}

var ProductBatchSelectExpectedQuery = "SELECT `id`,`batch_number`, `current_quantity`, `initial_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` "
