package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Data struct {
	Name    string
	Email   string
	Details *Details
}

type Details struct {
	FamilyMembers *FamilyMembers
	Salary        string
}

type FamilyMembers struct {
	NameOfMum string
	NameOfDad string
}

type Data2 struct {
	Name string
	Age  uint8
}

var validate *validator.Validate

func main() {
	validate = validator.New()

	validateStruct()
	validateStructNested()
}

func validateStruct() {

	data := Data2{
		Name: "leo",
		Age:  130,
	}

	rules := map[string]string{
		"Name": "min=4,max=6",
		"Age":  "min=4,max=6",
	}

	validate.RegisterStructValidationMapRules(rules, Data2{})

	err := validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			fmt.Println(e)
		}
	}

}

func validateStructNested() {

	data := Data{
		Name:  "112fdfhj343",
		Email: "xyz3884@mail.com",
		Details: &Details{
			Salary: "1000",
		},
	}

	rules1 := map[string]string{
		"Name":    "min=4,max=6",
		"Email":   "required,email",
		"Details": "required",
	}

	rules2 := map[string]string{
		"Salary":        "number",
		"FamilyMembers": "required",
	}

	rules3 := map[string]string{
		"NameOfMum": "required,min=4,max=32",
		"NameOfDad": "required,min=4,max=32",
	}

	validate.RegisterStructValidationMapRules(rules1, Data{})
	validate.RegisterStructValidationMapRules(rules2, Details{})
	validate.RegisterStructValidationMapRules(rules3, FamilyMembers{})

	err := validate.Struct(data)
	if err != nil {
		fmt.Println(err)
	}

}
