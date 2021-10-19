package main

import "fmt"

type Person struct {
	name string
}

type StoreKeeper struct {
	Person
}

func main() {
	var p = StoreKeeper{}
	p.Person = Person{"Jane Doe"}

	fmt.Println(p.name)
}
