package main

import "fmt"

func main() {

	in1 := make(chan struct{})
	in2 := make(chan struct{})

	go func(ch1, ch2 chan struct{}) {

		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
				} else {
					fmt.Println(v)
				}
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
				} else {
					fmt.Println(v)
				}
			}

			// multiple channel closed, return from goroutine
			if ch1 == nil && ch2 == nil {
				return
			}
		}

	}(in1, in2)

	close(in1)
	close(in2)

}
