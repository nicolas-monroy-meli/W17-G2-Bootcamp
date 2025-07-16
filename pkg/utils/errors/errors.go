package errors

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	// Requests
	ErrRequestIdMustBeInt  = errors.New("handler: id must be an integer")
	ErrRequestIdMustBeGte0 = errors.New("handler: id must be greater than 0")
	ErrRequestNoBody       = errors.New("handler: request must have a body")
	ErrRequestWrongBody    = errors.New("handler: body does not meet requirements")
	ErrRequestFailedBody   = errors.New("handler: failed to read body")

	//Query
	ErrQueryError   = errors.New("repository: unable to execute query")
	ErrInsertError  = errors.New("repository: insert is returning an error")
	ErrQueryIsEmpty = errors.New("repository: query returned no info")

	//Insert
	ErrForeignKeyError    = errors.New("repository: unable to execute query due to foreign key error")
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

	//PurchaseOrder
	ErrPORepositoryOrderNumberDuplicated = errors.New("repository: Order number duplicated")

	//Employee
	// ErrEmployeeRepositoryNotFound is returned when the employee is not found
	ErrEmployeeRepositoryNotFound = errors.New("repository: employee not found")
	// ErrEmployeeRepositoryDuplicated is returned when the employee already exists
	// Errores de Employee
	ErrEmployeeRepositoryDuplicated = errors.New("repository: employee already exists")

	// Errores de Product
	ErrProductRepositoryNotFound   = errors.New("repository: product not found")
	ErrProductRepositoryDuplicated = errors.New("repository: product already exists")
	// ErrProductRecordRepositoryNotFound is returned when the product record is not found
	ErrProductRecordRepositoryNotFound = errors.New("repository: product record not found")
	// ErrProductRecordRepositoryDuplicated is returned when the product record already exists
	ErrProductRecordRepositoryDuplicated = errors.New("repository: product record already exists")

	// Errores de Section
	ErrEmptySectionDB              = errors.New("repository: empty DB")
	ErrSectionRepositoryNotFound   = errors.New("repository: section not found")
	ErrSectionRepositoryDuplicated = errors.New("repository: section already exists")

	// Errores de Seller
	ErrSellerRepositoryNotFound   = errors.New("repository: seller not found")
	ErrSellerRepositoryDuplicated = errors.New("repository: seller already exists")

	//Locality
	// ErrLocalityNotFound is returned when the locality is not found
	ErrLocalityRepositoryNotFound = errors.New("repository: locality not found")

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
		case "gtefield":
			customMsg = fmt.Sprintf("%s must be greater than or equal to %s", field, err.Param())
		case "gtfield":
			customMsg = fmt.Sprintf("%s must be greater than %s", field, err.Param())
		case "lte":
			customMsg = fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
		case "lt":
			customMsg = fmt.Sprintf("%s must be less than %s", field, err.Param())
		case "ltefield":
			customMsg = fmt.Sprintf("%s must be less than or equal to %s", field, err.Param())
		case "ltfield":
			customMsg = fmt.Sprintf("%s must be less than %s", field, err.Param())
		default:
			customMsg = fmt.Sprintf("%s failed on %s validation", field, err.Tag())
		}

		errorsList[field] = customMsg
	}
	return errorsList
}
