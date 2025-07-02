package common

import mod "github.com/smartineztri_meli/W17-G2-Bootcamp/pkg/models"

func PatchEmployees(model mod.Employee, section mod.Employee) mod.Employee {
	patchedEmployee := section
	if model.CardNumberID != "" && model.CardNumberID != patchedEmployee.CardNumberID {
		patchedEmployee.CardNumberID = model.CardNumberID
	}

	if model.FirstName != "" && model.FirstName != patchedEmployee.FirstName {
		patchedEmployee.FirstName = model.FirstName
	}
	if model.LastName != "" && model.LastName != patchedEmployee.LastName {
		patchedEmployee.LastName = model.LastName
	}

	if model.WarehouseID != 0 && model.WarehouseID != patchedEmployee.WarehouseID {
		patchedEmployee.WarehouseID = model.WarehouseID
	}

	return patchedEmployee
}
