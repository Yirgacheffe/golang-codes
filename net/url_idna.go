package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

func main() {

	src := "我礼物"
	ascii, err := idna.TOASCII(src)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s -> %s/n", src, ascii)

}
