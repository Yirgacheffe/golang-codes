package main

import (
	"fmt"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	// en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type User struct {
	FirstName string     `validate:"required"`
	LastName  string     `validate:"required"`
	Age       uint8      `validate:"gte=0,lte=150"`
	Email     string     `validate:"required,email"`
	Favourite string     `validate:"iscolor"`
	Addresses []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func main() {
	en := en.New()
	zh := zh.New()
	uni = ut.New(en, zh)

	trans, _ := uni.GetTranslator("zh")

	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)

	translateAll(trans)
	translateIndividual(trans)
	translateOverride(trans)
}

func translateAll(trans ut.Translator) {

	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10"`
		Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "kk's blog",
		Tagline:  "the 1st line of tag is a long line",
		Tagline2: "1",
	}

	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs.Translate(trans))
	}

}

func translateIndividual(trans ut.Translator) {

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}

}

func translateOverride(trans ut.Translator) {

	validate.RegisterTranslation(
		"required",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value!", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fmt.Println(e.Translate(trans))
		}
	}

}
