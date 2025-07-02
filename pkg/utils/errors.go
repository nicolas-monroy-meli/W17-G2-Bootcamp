package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ErrRequestIdMustBeInt = errors.New("handler: id must be an integer")
	ErrRequestNoBody      = errors.New("handler: request must have a body")
	ErrRequestWrongBody   = errors.New("handler: body does not meet requirements")
	ErrRequestFailedBody  = errors.New("handler: failed to read body")

	// EmptyParams string telling the parameters are empty
	EmptyParams = "handler: empty parameters"
	// DataRetrievedSuccess string that tells the data was retrieved
	DataRetrievedSuccess = "handler: data retrieved successfully"
	// SectionDeleted string that tells the section was deleted successfully
	SectionDeleted = "handler: section deleted successfully"
	// SectionCreated string to show a successful creation
	SectionCreated = "handler: section successfully created"
	// SectionUpdated string to show a successful update
	SectionUpdated = "handler: section successfully updated"

	//Buyer
	// ErrBuyerRepositoryNotFound is returned when the buyer is not found
	ErrBuyerRepositoryNotFound = errors.New("repository: buyer not found")
	// ErrBuyerRepositoryDuplicated is returned when the buyer already exists
	ErrBuyerRepositoryDuplicated = errors.New("repository: buyer already exists")

	//Employee
	// ErrEmployeeRepositoryNotFound is returned when the employee is not found
	ErrEmployeeRepositoryNotFound = errors.New("repository: employee not found")
	// ErrEmployeeRepositoryDuplicated is returned when the employee already exists
	ErrEmployeeRepositoryDuplicated = errors.New("repository: employee already exists")

	//Product
	// ErrProductRepositoryNotFound is returned when the product is not found
	ErrProductRepositoryNotFound = errors.New("repository: product not found")
	// ErrProductRepositoryDuplicated is returned when the product already exists
	ErrProductRepositoryDuplicated = errors.New("repository: product already exists")

	//Section

	//ErrEmptySectionDB returned when there aren't any sections to show due db emptiness
	ErrEmptySectionDB = errors.New("repository: empty DB")
	// ErrSectionRepositoryNotFound is returned when the section is not found
	ErrSectionRepositoryNotFound = errors.New("repository: section not found")
	// ErrSectionRepositoryDuplicated is returned when the section already exists
	ErrSectionRepositoryDuplicated = errors.New("repository: section already exists")

	//Seller
	// ErrSellerRepositoryNotFound is returned when the seller is not found
	ErrSellerRepositoryNotFound = errors.New("repository: seller not found")
	// ErrSellerRepositoryDuplicated is returned when the seller already exists
	ErrSellerRepositoryDuplicated = errors.New("repository: seller already exists")

	//Warehouse
	// ErrWarehouseRepositoryNotFound is returned when the warehouse is not found
	ErrWarehouseRepositoryNotFound = errors.New("repository: warehouse not found")
	// ErrWarehouseRepositoryDuplicated is returned when the warehouse already exists
	ErrWarehouseRepositoryDuplicated = errors.New("repository: warehouse already exists")
)

// ValidateStruct returns a string map of formatted errors
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
