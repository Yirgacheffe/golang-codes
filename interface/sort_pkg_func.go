package main

import (
	"fmt"
	"sort"
)

type Dog struct {
	Name string
	Size int
}

type Dogs []Dog

func (d Dogs) Len() int {
	return len(d)
}

func (d Dogs) Less(i, j int) bool {
	return d[i].Size < d[j].Size
}

func (d Dogs) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	i := []int{1, 2, 5, 1, 9, 5, 3, 7, 8, 23, 12, 67, 34, 45}
	f := []float64{2.3, 6.7, 3.2, 5.7, 9.2, 2.5, 4.9, 3.1}
	s := []string{"John", "Peter", "Luke", "Watson", "Turing"}

	sort.Ints(i)
	fmt.Println(i)

	sort.Float64s(f)
	fmt.Println(f)

	sort.Strings(s)
	fmt.Println(s)

	ix := []int{1, 2, 3, 4}
	fmt.Println(sort.IntsAreSorted(ix))

	fx := []float64{2.3, 4.5, 1.2}
	fmt.Println(sort.Float64sAreSorted(fx))

	sx := []string{"Anaconda", "Jaguar", "Leopard", "Panda", "Bear"}
	fmt.Println(sort.StringsAreSorted(sx))

	is := []int{1, 3, 2, 3, 4, 7, 4, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(is)))

	fmt.Println(is)

	// data of dogs type
	dogs := []Dog{
		{"Terrier", 8}, {"Pug", 10}, {"Chihuahua", 6},
	}

	sort.Sort(Dogs(dogs))
	fmt.Println(dogs)
}
