package main

import (
	"container/list"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var xx = "Hello, world, xx!"

type Person struct {
	Name string
}

type Android struct {
	Person Person
	Model  string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

func f() {
	fmt.Println(xx)
}

func average(xs []float64) float64 {
	// panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f2returnWithName() (r int) {
	r = 1
	return
}

func f2returnMultipleValue() (int, int) {
	return 5, 6
}

func f2VariadicParams(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Factorial
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func sFirstFunc() {
	fmt.Println("1st")
}

func sSecondFunc() {
	fmt.Println("2nd")
}

func zeroIt(x int) {
	x = 0
}

func zeroItWithPointer(xPtr *int) {
	*xPtr = 0
}

func oneItWithPointer(xPtr *int) {
	*xPtr = 1
}

func routineTestF(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// Directional
// func pinger(c chan<- string)
func pinger(c chan string) {
	// for i := 0; ; i++ {
	for i := 0; i < 10; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	// for i := 0; ; i++ {
	for i := 0; i < 10; i++ {
		c <- "pong"
	}
}

// Directional
// func printer(c <-chan string)
func printer(c chan string) {
	for {
		msg := <-c

		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	var xyznn complex128 = complex(1, 5)
	fmt.Printf("%v, %T\n", xyznn, xyznn)

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	fmt.Println(len("Hello, world!"))
	fmt.Println("Hello, world!"[1])
	fmt.Println("Hello, " + "world!")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x = "HeLlO, wOrld!"
	fmt.Println(x)

	var y string
	y = "Hello, wOrld!"
	fmt.Println(y)

	z := "Hello, wORld!"
	fmt.Println(z)

	var a = "hello"
	var b = "world"
	fmt.Println(a == b)

	var c = "hello"
	var d = "hello"
	fmt.Println(c == d)

	f()

	// const testing
	const cy = "Hello, world from const!"
	fmt.Println(cy)

	const (
		name  = "evian"
		stars = "muguruza"
	)

	//
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)

	//
	xxx := 5
	xxx++

	fmt.Println(xxx)

	//
	fmt.Println("--------------------------")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}

		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Unknown number!")
		}

	}

	fmt.Println("--------------------------")
	var xyz [5]int

	xyz[3] = 100
	fmt.Println(xyz)

	var xyy [5]float64
	xyy[0] = 98
	xyy[1] = 93
	xyy[2] = 77
	xyy[3] = 82
	xyy[4] = 83

	var total float64
	for i := 0; i < 5; i++ {
		total += xyy[i]
	}

	fmt.Println(total / 5)
	fmt.Println(total / float64(len(xyy)))

	var totalx float64
	for _, value := range xyy {
		totalx += value
	}

	fmt.Println(totalx / 5)

	xxy := [4]float64{
		98,
		// 93,
		77,
		82,
		83,
	}

	for i, value := range xxy {
		fmt.Printf("%d = %f \n", i, value)
	}

	xxyx := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(xxyx[2:5])

	fmt.Println("--------------------------")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice11 := []int{1, 2, 3}
	slice22 := make([]int, 2)

	copy(slice22, slice11)
	fmt.Println(slice11)
	fmt.Println(slice22)

	fmt.Println("--------------------------")
	var xzmap map[string]int

	xzmap = make(map[string]int)
	xzmap["keya"] = 10

	fmt.Println(xzmap["keya"])

	xymap := make(map[int]int)
	xymap[1] = 12
	fmt.Println(xymap[1])

	delete(xymap, 1)
	fmt.Println(xymap[1])

	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["O"] = "Oxygen"

	fmt.Println(elements["O"])
	fmt.Println(elements["A"])

	namex, ok := elements["A"]
	fmt.Println(namex, ok)

	if namey, ok := elements["A"]; ok {
		fmt.Println(namey, ok)
	} else {
		fmt.Println("Not in Ok")
	}

	elementz := map[string]string{
		"C": "Carbon",
		"F": "Fluorine",
	}

	fmt.Println(elementz["H"])

	elementy := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
	}

	if el, ok := elementy["N"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println("--------------------------")
	xss := []float64{98, 93, 77, 83, 92, 300, 23}
	fmt.Println(average(xss))

	xm, ym := f2returnMultipleValue()
	fmt.Println(xm, ym)

	fmt.Println(f2VariadicParams(1, 2, 3, 4))

	xsx := []int{93, 4, 98, 23, 53, 5}
	fmt.Println(f2VariadicParams(xsx...))

	sAdd := func(x, y int) int {
		return x + y
	}

	fmt.Println(sAdd(5, 6))

	// closure
	xinc := 0
	sIncrement := func() int {
		xinc++
		return xinc
	}

	fmt.Println(sIncrement())
	fmt.Println(sIncrement())

	nextEvent := makeEvenGenerator()
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())

	fmt.Println(factorial(10))

	defer sSecondFunc()
	sFirstFunc()

	// defer
	// f, _ := os.Open(fileName)
	// defer f.Close()
	fmt.Println("--------------------------")
	// defer func() {
	//	 sstr := recover()
	//	 fmt.Println(sstr)
	// }()
	// panic("PANIC here")

	fmt.Println("--------------------------")
	sXx := 5
	fmt.Println(sXx)
	zeroIt(sXx)
	fmt.Println(sXx)

	zeroItWithPointer(&sXx)
	fmt.Println(sXx)

	sXs := new(int)
	fmt.Println(*sXs)

	oneItWithPointer(sXs)
	fmt.Println(*sXs)

	fmt.Println("--------------------------")
	xsAP := new(Android)
	xsAP.Person.Talk()

	fmt.Println("--------------------------")
	errss := errors.New("error message we defined")
	fmt.Println(errss)

	fmt.Println("--------------------------")

	var sxl list.List

	sxl.PushBack(1)
	sxl.PushBack(2)
	sxl.PushBack(3)

	for e := sxl.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	var inputAgain string
	fmt.Scanln(&inputAgain)

	fmt.Println("--------------------------")
	maxps := flag.Int("max", 6, "the max value")
	flag.Parse()

	fmt.Println(rand.Intn(*maxps))

	fmt.Println("--------------------------")
	go routineTestF(0)

	var rInput string
	fmt.Scanln(&rInput)

	for i := 0; i < 10; i++ {
		go routineTestF(i)
	}

	var xInput string
	fmt.Scanln(&xInput)

	fmt.Println("--------------------------")
	var ct = make(chan string)

	go pinger(ct)
	go ponger(ct)
	go printer(ct)

	fmt.Println("--------------------------")
	var pInput string
	fmt.Scanln(&pInput)

	var (
		name111 string
		age111  int
	)

	n, _ := fmt.Sscanf("xyz 8", "%s%d", &name111, age111)
	fmt.Println(n, name111, age111)

	fmt.Println("--------------------------")

	// buffered channel
	// cyx1 := make(chan string, 2)
	// cyx2 := make(chan string, 5)

	cxx1 := make(chan string)
	cxx2 := make(chan string)

	go func() {
		for {
			cxx1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			cxx2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-cxx1:
				fmt.Println("Message 1:", msg1)
			case msg2 := <-cxx2:
				fmt.Println("Message 2:", msg2)
			case <-time.After(time.Second):
				fmt.Println("Timeout")
			default:
				fmt.Println("Nothing ready")
			}
		}
	}()

	var scInput string
	fmt.Scanln(&scInput)

	fmt.Println("--------------------------")

	var x1 uint8 = 1<<1 | 1<<5
	var y1 uint8 = 1<<1 | 1<<2

	fmt.Printf("x1 = %08b\n", x1)
	fmt.Printf("y1 = %08b\n", y1)

	time.Sleep(10 * time.Second)

}
