package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	for i := 0; i < 5; i++ {
		v, e := rand.Int(rand.Reader, big.NewInt(1000))
		if e != nil {
			fmt.Print(e)
		}
		fmt.Print(v, " ") // 654 775 288 944 965 - maybe
	}
}
