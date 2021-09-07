package main

import (
	"crypto/rand"
	"fmt"
)

func main() {

	b := make([]byte, 10)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

}
