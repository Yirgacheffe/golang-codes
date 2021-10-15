package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	v := rand.Perm(8)

	fmt.Println(v) // [rand 0 ~ 7] an array
}
