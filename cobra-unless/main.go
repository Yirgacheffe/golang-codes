package main

import (
	"fmt"
	"os"
	"unless/cmd"
)

func main() {
	fmt.Println(os.Args[1:])
	cmd.Execute()
}
