package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func work() {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}

	wg.Done()
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go work()
	}

	wg.Wait() // GODEBUG=schedtrace=1000 go run deepdive_queue.go
}
