package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

type Person struct {
	Age int
}

func not_right() {
	persons := make([]Person, 10)

	for _, p := range persons {
		mutex.Lock()
		defer mutex.Unlock() // bad
		p.Age = 13
		mutex.Unlock() // right
	}
}

func maybe_right() {
	persons := make([]Person, 10)

	for _, p := range persons {
		func() {
			mutex.Lock()
			defer mutex.Unlock()
			p.Age = 13
		}()
	}
}

func main() {
	fmt.Println("--------")
}
