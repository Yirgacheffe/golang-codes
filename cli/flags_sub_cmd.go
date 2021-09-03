package main

import (
	"flag"
	"fmt"
	"os"
)

/*
 * example: go run filename.go add -a=42 -b=23
 * example: go run filename.go mul -a=25 -b=4
 */
func main() {

	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	a := addcmd.Int("a", 0, "The value of a")
	b := addcmd.Int("b", 0, "The value of b")

	mulcmd := flag.NewFlagSet("mul", flag.ExitOnError)
	c := mulcmd.Int("a", 0, "The value of a")
	d := mulcmd.Int("b", 0, "The value of b")

	switch os.Args[1] {
	case "add":
		addcmd.Parse(os.Args[2:])
		fmt.Println(*a + *b)
	case "mul":
		mulcmd.Parse(os.Args[2:])
		fmt.Println(*(c) * (*d))
	default:
		fmt.Println("expected add or mul command")
		os.Exit(1)
	}

}
