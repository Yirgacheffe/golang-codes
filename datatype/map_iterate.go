package main

import (
	"fmt"
	"sort"
)

func simple() {
	fMap := make(map[int]string)
	fMap[0] = "test"
	fMap[1] = "sample"
	fMap[2] = "Golang Fun~~"

	for k, v := range fMap {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}
}

func withSlice() {
	sMap := make(map[int]string)
	sMap[0] = "test"
	sMap[1] = "sample"
	sMap[2] = "Golang Fun Again!!"

	var keys []int
	for k, _ := range sMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("Key: %d, Value: %s\n", k, sMap[k])
	}
}

func main() {
	simple()
	withSlice()
}
