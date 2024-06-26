package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ErrorCode string

const (
	NotSupported   ErrorCode = "NotSupported"
	NotImplemented ErrorCode = "NotImplemented"
)

type ErrorResponse struct {
	Code ErrorCode `validate:"errorCode"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterValidation("errorCode", IsErrorCodeValid)

	r1 := ErrorResponse{Code: ""}

	err := validate.Struct(r1)
	if err != nil {
		fmt.Printf("Errs(s): \n%+v\n", err)
	}

	r2 := ErrorResponse{Code: NotImplemented}
	err = validate.Struct(r2)
	if err != nil {
		fmt.Printf("Err(s): \n%+v\n", err)
	}

}

func IsErrorCodeValid(fl validator.FieldLevel) bool {
	code := ErrorCode(fl.Field().String())
	switch code {
	case NotImplemented, NotSupported:
		return true
	}
	return false
}
