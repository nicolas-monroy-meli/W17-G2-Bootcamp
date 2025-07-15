package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// Errores generales
	ErrRequestIdMustBeInt = errors.New("handler: id must be an integer")
	ErrRequestNoBody      = errors.New("handler: request must have a body")
	ErrRequestWrongBody   = errors.New("handler: body does not meet requirements")
	ErrRequestFailedBody  = errors.New("handler: failed to read body")
	ErrRepositoryDatabase = errors.New("repository: database operation failed")

	// Mensajes exitosos
	EmptyParams          = "handler: empty parameters"
	DataRetrievedSuccess = "handler: data retrieved successfully"
	SectionDeleted       = "handler: section deleted successfully"
	SectionCreated       = "handler: section successfully created"
	SectionUpdated       = "handler: section successfully updated"

	// Errores de Buyer
	ErrBuyerRepositoryNotFound       = errors.New("repository: buyer not found")
	ErrBuyerRepositoryDuplicated     = errors.New("repository: buyer already exists")
	ErrBuyerRepositoryCardDuplicated = errors.New("repository: Card id duplicated")

	// Errores de Employee
	ErrEmployeeRepositoryNotFound   = errors.New("repository: employee not found")
	ErrEmployeeRepositoryDuplicated = errors.New("repository: employee already exists")

	// Errores de Product
	ErrProductRepositoryNotFound   = errors.New("repository: product not found")
	ErrProductRepositoryDuplicated = errors.New("repository: product already exists")

	// Errores de Section
	ErrEmptySectionDB              = errors.New("repository: empty DB")
	ErrSectionRepositoryNotFound   = errors.New("repository: section not found")
	ErrSectionRepositoryDuplicated = errors.New("repository: section already exists")

	// Errores de Seller
	ErrSellerRepositoryNotFound   = errors.New("repository: seller not found")
	ErrSellerRepositoryDuplicated = errors.New("repository: seller already exists")

	// Errores de Warehouse
	ErrWarehouseRepositoryNotFound   = errors.New("repository: warehouse not found")
	ErrWarehouseRepositoryDuplicated = errors.New("repository: warehouse already exists")

	// Errores de Carry (Nuevos)
	ErrCarryRepositoryNotFound         = errors.New("repository: carry not found")
	ErrCarryRepositoryDuplicated       = errors.New("repository: carry already exists")
	ErrCarryRepositoryLocalityNotFound = errors.New("repository: locality not found for carry")
)

// ValidateStruct retorna un mapa de errores de validación
func ValidateStruct(s interface{}) map[string]string {
	v := validator.New()
	errorsList := make(map[string]string)

	err := v.Struct(s)
	if err == nil {
		return nil
	}

	for _, err := range err.(validator.ValidationErrors) {
		customMsg := "unexpected error"
		field := err.Field()

		switch err.Tag() {
		case "required":
			customMsg = fmt.Sprintf("%s is required", field)
		case "gte":
			customMsg = fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
		case "gt":
			customMsg = fmt.Sprintf("%s must be greater than %s", field, err.Param())
		case "lte":
			customMsg = fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
		case "lt":
			customMsg = fmt.Sprintf("%s must be less than %s", field, err.Param())
		default:
			customMsg = fmt.Sprintf("%s failed on %s validation", field, err.Tag())
		}

		errorsList[field] = customMsg
	}
	return errorsList
}
