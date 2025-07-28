package errors

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	ErrNoRowsAffected = errors.New("repository: no rows were affected")
	// DataRetrievedSuccess string that tells the data was retrieved
	// Requests
	ErrRequestIdMustBeInt  = errors.New("handler: id must be an integer")
	ErrRequestIdMustBeGte0 = errors.New("handler: id must be greater than 0")
	ErrRequestNoBody       = errors.New("handler: request must have a body")
	ErrRequestWrongBody    = errors.New("handler: body does not meet requirements")
	ErrRequestFailedBody   = errors.New("handler: failed to read body")
	ErrNothingToUpdate     = errors.New("handler: nothing to update")
	//Query
	ErrQueryError   = errors.New("repository: unable to execute query")
	ErrParseError   = errors.New("repository: unable to parse row")
	ErrInsertError  = errors.New("repository: insert is returning an error")
	ErrQueryIsEmpty = errors.New("repository: query returned no info")

	//Insert
	ErrForeignKeyError    = errors.New("repository: unable to execute query due to foreign key error")
	ErrRepositoryDatabase = errors.New("repository: database operation failed")

	// Mensajes exitosos
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

	//Inbound
	ErrInboundOrderNotFound      = errors.New("inbound order not found")
	ErrInboundOrderAlreadyExists = errors.New("inbound order with this ID already exists") // Si el ID fuera generado por la app y no AUTO_INCREMENT
	ErrEmployeeInternal          = errors.New("internal server error for employee")
	ErrInboundOrderInternal      = errors.New("internal server error for inbound order")
	ErrEmployeeNotFound          = errors.New("employee not found")
	ErrInboundOrderInvalidData   = errors.New("invalid inbound order data")

	//Product
	// ErrProductRepositoryNotFound is returned when the product is not found
	ErrProductRepositoryNotFound = errors.New("repository: product not found")
	// ErrProductRepositoryDuplicated is returned when the product already exists
	ErrProductRepositoryDuplicated = errors.New("repository: product already exists")
	// ErrProductRecordRepositoryNotFound is returned when the product record is not found
	ErrProductRecordRepositoryNotFound = errors.New("repository: product record not found")
	// ErrProductRecordRepositoryDuplicated is returned when the product record already exists
	ErrProductRecordRepositoryDuplicated = errors.New("repository: product record already exists")

	// Errores de Section
	ErrEmptyDB                     = errors.New("repository: empty DB")
	ErrSectionRepositoryNotFound   = errors.New("repository: section not found")
	ErrSectionRepositoryDuplicated = errors.New("repository: section already exists")

	ErrProductBatchNotFound   = errors.New("repository: Product Batch not found")
	ErrProductBatchDuplicated = errors.New("repository: Product Batch already exists")

	//Seller
	// ErrSellerRepositoryNotFound is returned when the seller is not found
	ErrSellerRepositoryNotFound = errors.New("repository: seller not found")
	// ErrSellerRepositoryDuplicated is returned when the seller already exists

	ErrSellerRepositoryDuplicated = errors.New("repository: seller already exists")

	//Locality
	// ErrLocalityNotFound is returned when the locality is not found
	ErrLocalityRepositoryNotFound   = errors.New("repository: locality not found")
	ErrLocalityRepositoryDuplicated = errors.New("repository: locality already exists")

	// Errores de Warehouse
	ErrWarehouseRepositoryNotFound   = errors.New("repository: warehouse not found")
	ErrWarehouseRepositoryDuplicated = errors.New("repository: warehouse already exists")

	// Errores de Carry (Nuevos)
	ErrCarryRepositoryNotFound         = errors.New("repository: carry not found")
	ErrCarryRepositoryDuplicated       = errors.New("repository: carry already exists")
	ErrCarryRepositoryLocalityNotFound = errors.New("repository: locality not found for carry")
)

func validTime(fl validator.FieldLevel) bool {
	_, err := time.Parse("15:04:05", fl.Field().String())
	return err == nil
}

var DupErr = &mysql.MySQLError{
	Number:  1062,
	Message: "Duplicate Entry",
}

var FkErr = &mysql.MySQLError{
	Number:  1452,
	Message: "Foreign Key Constrain Error",
}

// FakeResult created in order to raise Fail on lastInsertId
type FakeResult struct{}

func (r FakeResult) LastInsertId() (int64, error) { return 0, errors.New("fail on last insert id") }
func (r FakeResult) RowsAffected() (int64, error) { return 1, nil }

// ValidateStruct returns a string map of formatted errors
func ValidateStruct(s interface{}) map[string]string {
	v := validator.New()
	v.RegisterValidation("hhmmss", validTime)
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
		case "hhmmss":
			customMsg = fmt.Sprintf("%s must follow HH:MM:SS format", field)
		default:
			customMsg = fmt.Sprintf("%s failed on %s validation", field, err.Tag())
		}

		errorsList[field] = customMsg
	}
	return errorsList
}
