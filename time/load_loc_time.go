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
	loc, e := time.LoadLocation("EST")
	CheckError(e)

	fmt.Println(time.Now().In(loc))
}
