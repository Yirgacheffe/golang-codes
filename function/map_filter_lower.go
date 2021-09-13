package main

import (
	"fmt"
	"strings"
)

// Filter - Map --> with lambda function -----------------
type Work struct {
	Data    string
	Version int
}

func Filter(ws []Work, f func(w Work) bool) []Work {
	result := make([]Work, 0)
	for _, v := range ws {
		if f(v) {
			result = append(result, v)
		}
	}
	return result // re, return new filtered list
}

func Map(ws []Work, f func(w Work) Work) []Work {
	result := make([]Work, len(ws))
	for i, v := range ws {
		result[i] = f(v)
	}
	return result // re, return new mapped work list
}

func LowerCaseData(w Work) Work {
	w.Data = strings.ToLower(w.Data)
	return w
}

func IncrementVersion(w Work) Work {
	w.Version++
	return w
}

func OlderThen(v int) func(w Work) bool {
	return func(w Work) bool { return w.Version >= v }
}

func main() {

	ws := []Work{
		{
			Data: "Example 1", Version: 1,
		},
		{
			Data: "Example 2", Version: 2,
		},
	}

	fmt.Printf("Initial--: %#v\n", ws)
	fmt.Println("--")

	ws = Map(ws, LowerCaseData)
	fmt.Printf("LowerCase: %#v\n", ws)
	fmt.Println("--")

	ws = Map(ws, IncrementVersion)
	fmt.Printf("Version++: %#v\n", ws)
	fmt.Println("--")

	ws = Filter(ws, OlderThen(3))
	fmt.Printf("RemoveVer: %#v\n", ws)

}
