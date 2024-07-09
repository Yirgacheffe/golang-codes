package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	validate = validator.New()

	validateMap()
	validateNestedMap()
}

func validateMap() {
	user := map[string]interface{}{"name": "aya kIAni", "email": "333923xyz@gmail.com"}

	rules := map[string]interface{}{"name": "required,min=8,max=32", "email": "omitempty,required,email"}
	errs := validate.ValidateMap(user, rules)
	if len(errs) > 0 {
		fmt.Println(errs)
	}
}

func validateNestedMap() {
	data := map[string]interface{}{
		"name":  "atenate",
		"email": "kanjilu@gmail.com",
		"details": map[string]interface{}{
			"family_members": map[string]interface{}{"father_name": "nate sanro", "mother_name": "nate hibiki"},
			"salary":         "3000",
			"phones": []map[string]interface{}{
				{
					"number": "86-3234-3232", "remark": "home",
				},
				{
					"number": "86-1234-4321", "remark": "work",
				},
			},
		},
	}

	rules := map[string]interface{}{
		"name":  "min=4,max=32",
		"email": "required,email",
		"details": map[string]interface{}{
			"family_members": map[string]interface{}{
				"father_name": "required,min=4,max=32",
				"mother_name": "required,min=4,max=32",
			},
			"salary": "number",
			"phones": map[string]interface{}{
				"number": "required,min=4,max=32",
				"remark": "required,min=1,max=32",
			},
		},
	}

	errs := validate.ValidateMap(data, rules)

	if len(errs) != 0 {
		fmt.Println(errs)
	}

}
