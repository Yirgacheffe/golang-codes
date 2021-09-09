package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

type Point struct {
	x, y float64
}

type Circle struct {
	origin Point
	r      float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleAreaOverload(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func totalArea(shape ...Shape) float64 {
	var area float64
	for _, v := range shape {
		area += v.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	fmt.Println(math.Abs(-12.3)) // 12.3
	fmt.Println(math.Sin(43))    // -0.8317747426285983
	fmt.Println(math.Cos(57))    // 0.8998668269691938
	fmt.Println(math.Log(89))    // 4.48863636973214
	fmt.Println(math.Log10(90))  // 1.9542425094393248

	fmt.Println(math.Mod(14, 5))  // 4
	fmt.Println(math.Ceil(12.5))  // 13
	fmt.Println(math.Floor(23.7)) // 23
	fmt.Println(math.Min(23, 45)) // 23
	fmt.Println(math.Max(23, 45)) // 45
	fmt.Println(math.Pow(2, 16))  // 65536

	fmt.Println(math.IsNaN(33)) // false

	fmt.Println("--------------------------")
	fmt.Println(distance(0, 0, 3, 4))
	fmt.Println(circleArea(0, 0, 3))

	fmt.Println("--------------------------")
	var xsP Point
	var xsC Circle
	var xsI float32

	fmt.Println(xsP)
	fmt.Println(xsC)
	fmt.Println(xsI)

	xxP := new(Point)
	xxC := new(Circle)

	fmt.Println(xxP) // Pointer
	fmt.Println(xxC) // Pointer

	xyP := Point{x: 0, y: 0}
	xyC := Circle{origin: xyP, r: 5}

	fmt.Println(xyC)
	fmt.Println(&xyC) // Pointer

	xyP.x = 5
	xyP.y = 5
	fmt.Println(xyP)

	xyC.r = 5
	fmt.Println("The origin is:", xyC.origin)
	fmt.Println(circleAreaOverload(xyC))

	r := Rectangle{0, 0, 10, 20}
	fmt.Println(r.area())

	xifRect1 := Rectangle{0, 0, 4, 3}
	xifRect2 := Rectangle{1, 1, 8, 9}

	xifRects := [2]Rectangle{xifRect1, xifRect2}
	fmt.Println(totalArea(&xifRects[0], &xifRects[1]))

	multiShapeXs := MultiShape{
		shapes: []Shape{
			&Rectangle{0, 0, 4, 3},
			&Rectangle{1, 1, 8, 9},
		},
	}

	fmt.Println(multiShapeXs.area())

}
