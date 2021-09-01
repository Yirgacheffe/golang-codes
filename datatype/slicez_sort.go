package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height int
	weight int
}

func sortPerson() {

	persons := []Person{
		{"Mihalis", 180, 90},
		{"Bill", 134, 45},
		{"Marietta", 155, 45},
		{"Epifanios", 144, 50},
		{"Athina", 134, 45},
	}

	asc := func(i, j int) bool {
		return persons[i].height < persons[j].height
	}

	dsc := func(i, j int) bool {
		return persons[i].height > persons[j].height
	}

	sort.Slice(persons, asc)
	fmt.Println("<:", persons)

	sort.Slice(persons, dsc)
	fmt.Println(">:", persons)

	/*
		persons := make([]Person, 0)

		persons = append(persons, Person{"Mihalis", 180, 90})
		persons = append(persons, Person{"Bill", 134, 45})
		persons = append(persons, Person{"Marietta", 155, 45})
		persons = append(persons, Person{"Epifanios", 144, 50})
		persons = append(persons, Person{"Athina", 134, 45})
	*/

}

func main() {

	sortPerson()

}
