package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p
	}
}

func main() {

	ports := make(chan int, 100)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	// 1st time forgot to run in another go routine
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	var openPorts []int

	for i := 1; i <= 1024; i++ {
		p := <-results
		if p != 0 {
			openPorts = append(openPorts, p)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	fmt.Println(openPorts)

}
