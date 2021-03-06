package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // return the channel to the caller
}

// simple fanin
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {

	c := boring("boring!")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	results := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-results)
	}

	fmt.Println(2e3)
	fmt.Printf("%q\n", "quit")

	/*
		// this is for all conversation
		var timeout = time.After(5 * time.Second)

		for {
			select {
			case s := <-c:
				fmt.Println(s)
			case <-time.After(1 * time.Second): // only for each message
				return
			case <-timeout:
				return
			}
		}
	*/
}
