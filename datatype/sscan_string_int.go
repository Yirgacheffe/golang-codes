package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	s := flag.Arg(0)

	var i int
	_, e := fmt.Sscan(s, &i)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("i is:", i)

	// string parse with format
	k := "user:0456"
	var j int
	if _, err := fmt.Sscanf(k, "user:%4d", &j); err == nil {
		fmt.Println(j)
	}
}
