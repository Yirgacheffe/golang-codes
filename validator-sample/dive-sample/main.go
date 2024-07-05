package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Test struct {
	Array []string          `validate:"required,gt=0,dive,required"`
	Map   map[string]string `validate:"required,gt=0,dive,keys,keymax,endkeys,required,max=100"`
}

var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterAlias("keymax", "max=10")

	var t Test

	doValidate(t)

	t.Array = []string{""}
	t.Map = map[string]string{"test > than 10": ""}

	doValidate(t)

}

func doValidate(t Test) {
	fmt.Println("do ------------ test")
	err := validate.Struct(t)
	fmt.Println(err)
}
