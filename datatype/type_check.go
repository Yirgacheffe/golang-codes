package main

import "fmt"

type ID string

type Person struct {
	name string
	age  int
}

func CheckType(s interface{}) {
	switch s.(type) {
	case bool:
		fmt.Println("boolean value: ", s.(bool))
	case string:
		fmt.Println("It's string.")
	case int:
		fmt.Println("It's int.")
	case float64:
		fmt.Println(s.(float64))
	case Person:
		fmt.Println(s.(Person))
	case chan int:
		fmt.Println(s.(chan int))
	default:
		fmt.Println("Not sure what it is...")
	}
}

func PrintlnType(x interface{}) {
	fmt.Printf("type is '%T', value: %v\n", x, x)
}

func PrintlnNot(x interface{}) {
	if v, ok := x.(ID); ok {
		fmt.Printf("x type ID, value is: %v\n", v)
	} else {
		fmt.Printf("'%T' is not type I want\n", x)
	}
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, Value = %v\n", i, i)
}

func main() {
	strt := struct{ name string }{name: "Kurosaki"}
	describe(strt)

	CheckType("test")
	CheckType(64)
	CheckType(false)

	var ix interface{}
	ix = "test again"

	if s, ok := ix.(string); ok {
		fmt.Println("value is:", s)
	}

	if _, ok := ix.(int); !ok {
		fmt.Println("I am happy to handle this.")
	}
}
