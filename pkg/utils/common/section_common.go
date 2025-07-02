package common

import (
	"errors"
	"fmt"
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

func PatchSection(request models.SectionPatch, section models.Section) models.Section {
	fmt.Println(request.WarehouseID)
	switch {
	case request.SectionNumber != nil:
		section.SectionNumber = *request.SectionNumber
		fallthrough
	case request.CurrentTemperature != nil:

		section.CurrentTemperature = *request.CurrentTemperature
		fallthrough
	case request.MinimumTemperature != nil:
		section.MinimumTemperature = *request.MinimumTemperature
		fallthrough
	case request.CurrentCapacity != nil:
		section.CurrentCapacity = *request.CurrentCapacity
		fallthrough
	case request.MinimumCapacity != nil:
		section.MinimumCapacity = *request.MinimumCapacity
		fallthrough
	case request.MaximumCapacity != nil:
		section.MaximumCapacity = *request.MaximumCapacity
		fallthrough
	case request.WarehouseID != nil:
		section.WarehouseID = *request.WarehouseID
		fallthrough
	case request.ProductTypeID != nil:
		section.ProductTypeID = *request.ProductTypeID
	}
	fmt.Println(section)
	return section
}
