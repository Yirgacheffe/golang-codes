package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second * 2)
	t1 := time.Now()
	fmt.Println("Time diff:", t1.Sub(t))

	skFormat := t.Format("01 January 2006")
	fmt.Println(skFormat)

	loc, _ := time.LoadLocation("Europe/Paris") // ignore error
	londonTime := t.In(loc)
	fmt.Println("Paris:", londonTime)

	fmt.Println(t.Format("3:04PM"))                        // 8:54PM
	fmt.Println(t.Format("Jan-02-06"))                     // February 27, 2020
	fmt.Println(t.Format("Jan _2 15:04:05.000000"))        // Feb 27 21:07:10.714500
	fmt.Println(t.Format("3:04:05 PM"))                    // 9:07:10 PM
	fmt.Println(t.Format("Mon, 02 Jan 2006 15:04:05 MST")) // Thu, 27 Feb 2020 21:07:10 IST

}
