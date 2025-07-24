package repository_tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/internal/repository"
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestProductBatch_FindAll_DB_with_Data(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	rows := sqlmock.NewRows([]string{"id", "batch_number", "current_quantity", "initial_quantity", "current_temperature", "minimum_temperature", "due_date", "manufacturing_date", "manufacturing_hour", "product_id", "section_id"}).
		AddRow(1, 1, 200, 200, 2, -5, time.Date(2024, 07, 05, 17, 00, 00, 0, time.UTC), time.Date(2024, 06, 1, 0, 00, 00, 0, time.UTC), "08:00:00", 1, 1).
		AddRow(2, 2, 310, 310, -2, -6, time.Date(2024, 8, 01, 12, 00, 00, 0, time.UTC), time.Date(2024, 7, 1, 0, 00, 00, 0, time.UTC), "09:30:00", 2, 2)
	mock.ExpectQuery("SELECT `id`,`batch_number`, `current_quantity`, `initial_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` ").
		WillReturnRows(rows)
	repo := repository.NewProductBatchRepo(db)
	batches, err := repo.FindAll()
	expectedResult := []mod.ProductBatch{
		{
			ID:                 1,
			BatchNumber:        1,
			CurrentQuantity:    200,
			InitialQuantity:    200,
			CurrentTemperature: 2,
			MinimumTemperature: -5,
			DueDate:            time.Date(2024, 07, 05, 17, 00, 00, 0, time.UTC),
			ManufacturingDate:  time.Date(2024, 06, 1, 0, 00, 00, 0, time.UTC),
			ManufacturingHour:  "08:00:00",
			ProductId:          1,
			SectionId:          1,
		},
		{
			ID:                 2,
			BatchNumber:        2,
			CurrentQuantity:    310,
			InitialQuantity:    310,
			CurrentTemperature: -2,
			MinimumTemperature: -6,
			DueDate:            time.Date(2024, 8, 01, 12, 00, 00, 0, time.UTC),
			ManufacturingDate:  time.Date(2024, 7, 1, 0, 00, 00, 0, time.UTC),
			ManufacturingHour:  "09:30:00",
			ProductId:          2,
			SectionId:          2,
		},
	}
	require.NoError(t, err)
	require.Len(t, batches, 2)
	require.Equal(t, expectedResult, batches)
}

func TestProductBatch_Create_HappyPath(t *testing.T) {

}
