package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	Gender         string     `validate:"oneof=male female prefer_not_to"`
	FavouriteColor string     `validate:"iscolor"`
	Addresses      []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate.New(validator.WithRequiredStructEnabled())

	validateStruct()
	validateVariable()
}

func validateStruct() {
	address := &Address{
		Street: "Wujiayao",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "George",
		LastName:       "Smith",
		Age:            "30",
		Gender:         "male",
		Email:          "george.smith@volvo.com",
		FavouriteColor: "#000-",
		Address:        []*Address{address},
	}

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return
	}
	// save to database if no error happened
}

func validateVariable() {
	myEmail := "jobyevlogs.outllook.com"

	errs := validate.Var(myEmail, "required,email")
	if errs != nil {
		fmt.Println(errs)
		return
	}

	// email ok, keep going
}
