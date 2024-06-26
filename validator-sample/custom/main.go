package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type DBBackedUser struct {
	Name sql.NullString `validate:"required"`
	Age  sql.NullInt16  `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New()
	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt16{}, sql.NullBool{}, sql.NullFloat64{})

	x := DBBackedUser{Name: sql.NullString{String: "", Valid: false}, Age: sql.NullInt16{Int16: 30, Valid: true}}

	err := validate.Struct(x)
	if err != nil {
		fmt.Printf("Errs(s): \n%+v\n", err)
	}
}

// ValidateValuer convert custom value into primitive value
func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}

		// this line was useless
		return errors.Wrap(err, "if see this")
	}
	return nil
}
