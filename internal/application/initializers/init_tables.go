package application

import (
	mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"gorm.io/gorm"
)

func CreateTables(db *gorm.DB) {
	err := db.AutoMigrate(&mod.Section{}, &mod.ProductBatch{}, &mod.Product{})
	if err != nil {
		panic(err)
		return
	}
}
