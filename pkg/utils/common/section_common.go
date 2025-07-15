package common

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"
	e "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/utils/errors"
	"net/http"
	"strconv"
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

func PatchSection(request models.SectionPatch, section models.Section) models.Section {
	fmt.Println(request.CurrentCapacity)
	if request.SectionNumber != nil {
		section.SectionNumber = *request.SectionNumber
	}
	if request.CurrentTemperature != nil {
		section.CurrentTemperature = *request.CurrentTemperature
	}
	if request.MinimumTemperature != nil {
		section.MinimumTemperature = *request.MinimumTemperature
	}
	if request.CurrentCapacity != nil {
		section.CurrentCapacity = *request.CurrentCapacity
	}
	if request.MinimumCapacity != nil {
		section.MinimumCapacity = *request.MinimumCapacity
	}
	if request.MaximumCapacity != nil {
		section.MaximumCapacity = *request.MaximumCapacity
	}
	if request.WarehouseID != nil {
		section.WarehouseID = *request.WarehouseID
	}
	if request.ProductTypeID != nil {
		section.ProductTypeID = *request.ProductTypeID
	}
	return section
}
