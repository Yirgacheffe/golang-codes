package main

import (
	"fmt"

	"hello-tls/pb/person"
	"hello-tls/pb/team"
)

func main() {

	p1 := person.Person{}
	fmt.Println(p1)

	p2 := person.Person{}
	fmt.Println(p2)

	ps := []*person.Person{&p1, &p2}

	t := team.Team{Person: ps}
	fmt.Println(t)
}
