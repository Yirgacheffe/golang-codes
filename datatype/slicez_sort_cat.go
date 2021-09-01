package main

import (
	"fmt"
	"sort"
)

type Cat struct {
	Name string
	Age  int
}

type ByName []Cat

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

type ByAge []Cat

func (ts ByAge) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (ts ByAge) Len() int {
	return len(ts)
}

func (ts ByAge) Less(i, j int) bool {
	return ts[i].Age < ts[j].Age
}

func main() {

	// Sort
	kids := []Cat{
		{"MaDudu", 6}, {"HuangQiangQiang", 4}, {"MaDaGui", 9},
	}

	sort.Sort(ByAge(kids))
	fmt.Println(kids)

	sort.Sort(ByName(kids))
	fmt.Println(kids)

}
