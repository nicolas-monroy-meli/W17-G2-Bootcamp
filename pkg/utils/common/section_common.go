package common

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils"
	"net/http"
	"strconv"
)

// IdRequests encapsulates the process of getting the id parameter and returns an int number and an error if necessary
func IdRequests(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		return 0, errors.New(utils.EmptyParams)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return id, nil
}

func PatchSection(model models.Section, section models.Section) models.Section {
	switch {
	case model.SectionNumber != 0 && model.SectionNumber != section.SectionNumber:
		section.SectionNumber = model.SectionNumber
	case model.CurrentTemperature != 0 && model.CurrentTemperature != section.CurrentTemperature:
		section.CurrentTemperature = model.CurrentTemperature
	case model.MinimumTemperature != 0 && model.MinimumTemperature != section.MinimumTemperature:
		section.MinimumTemperature = model.MinimumTemperature
	case model.CurrentCapacity != 0 && model.CurrentCapacity != section.CurrentCapacity:
		section.CurrentCapacity = model.CurrentCapacity
	case model.MinimumCapacity != 0 && model.MinimumCapacity != section.MinimumCapacity:
		section.MinimumCapacity = model.MinimumCapacity
	case model.MaximumCapacity != 0 && model.MaximumCapacity != section.MaximumCapacity:
		section.MaximumCapacity = model.MaximumCapacity
	case model.WarehouseID != 0 && model.WarehouseID != section.WarehouseID:
		section.WarehouseID = model.WarehouseID
	case model.ProductTypeID != 0 && model.ProductTypeID != section.ProductTypeID:
		section.ProductTypeID = model.ProductTypeID
	}
	return section
}
