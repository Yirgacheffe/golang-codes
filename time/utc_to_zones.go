package main

import (
	"fmt"
	"time"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	loc, e := time.LoadLocation("America/Los_Angeles")
	CheckError(e)

	fmt.Println(time.Now().UTC().In(loc))
}
