package main

import "fmt"

type Painting struct {
	string
	Artist
}

type Artist struct {
	string
}

func main() {
	var p Painting
	p.string = "Starry Night"
	p.Artist.string = "Van Goph"

	fmt.Println(p)
}
