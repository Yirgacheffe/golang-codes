package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mux  = sync.Mutex{}
	cond = sync.NewCond((&mux))
)

var queue []int

func producer() {
	i := 0

	for {
		mux.Lock()
		queue = append(queue, i)
		i++
		mux.Unlock()

		cond.Signal()
		time.Sleep(1 * time.Second)
	}
}

func consumer(name string) {
	for {
		mux.Lock()
		for len(queue) == 0 {
			cond.Wait()
		}

		fmt.Println(name, queue[0])
		queue = queue[1:]
		mux.Unlock()
	}
}

func main() {
	go producer()

	go consumer("consumer-1")
	go consumer("consumer-2")

	for {
		time.Sleep(1 * time.Minute)
	}
}
