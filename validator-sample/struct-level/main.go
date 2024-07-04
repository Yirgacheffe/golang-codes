package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Gender uint

const (
	Male Gender = iota + 1
	Female
	Intersex
)

func (gender Gender) String() string {
	terms := []string{"", "Male", "Female", "Intersex"}
	if gender < Male || gender > Intersex {
		return "unknown"
	}
	return terms[gender]
}

type User struct {
	FirstName      string     `json:"fname"`
	LastName       string     `json:"lname"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `json:"e-mail" validate:"required,email"`
	FavouriteColor string     `validate:"hexcolor|rgb|rgba"`
	Addresses      []*Address `validate:"required,dive,required"`
	Gender         Gender     `json:"gender" validate:"required,gender_custom_validation"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

type validationError struct {
	Namespace       string `json:"namespace"`
	Field           string `json:"fidle"`
	StructNamespace string `json:"struct_namespace"`
	StructField     string `json:"struct_field"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actual_tag"`
	Kind            string `json:"kind"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	Param           string `json:"param"`
	Message         string `json:"message"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	err := validate.RegisterValidation("gender_custom_validation", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(Gender)
		return value.String() != "unknown"
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	address := &Address{
		Street: "Wu Jia Yao Street",
		Planet: "Persphone",
		Phone:  "02284985844",
		City:   "Unknown",
	}

	user := &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "xyz.dica@gmail",
		FavouriteColor: "#000",
		Addresses:      []*Address{address},
	}

	err = validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			e := validationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
			}

			indent, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			fmt.Println(string(indent))
		}
	}

	fmt.Println(Male.String())

}

func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(
			user.FirstName,
			"fname",
			"FirstName",
			"fnameorlname",
			"",
		)

		sl.ReportError(
			user.LastName,
			"lname",
			"LastName",
			"fnameorlname",
			"",
		)
	}

	// check other tags here ...
}
