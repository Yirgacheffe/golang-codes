package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculate failed, radius %0.2f is less than 0.", radius)
	}
	return radius * radius * math.Pi, nil
}

func main() {
	radius := -20.1
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle is %0.2f", area)
}
