package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	v, e := rand.Prime(rand.Reader, 1024)
	if e != nil {
		fmt.Print(e)
	}
	fmt.Println(v)
}
