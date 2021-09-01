package main

import "fmt"

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	var fileSize float64 = 4 * 1024 * 1024 * 1024
	fmt.Printf("%.2fGB\n", fileSize/GB)
}
