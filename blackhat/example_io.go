package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (r *FooReader) Read(b []byte) (int, error) {
	fmt.Println("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (w *FooWriter) Write(b []byte) (int, error) {
	fmt.Println("out > ")
	return os.Stdout.Write(b)
}

func main() {

	var (
		r FooReader
		w FooWriter
	)

	var buf = make([]byte, 1024)

	s, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Unable to read data")
	}

	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = w.Write(buf)
	if err != nil {
		log.Fatalln("Unable to write data")
	}

	fmt.Printf("write %d bytes to stdout\n", s)

	// High level API to copy data
	/*
		if _, err := io.Copy(&w, &r); err != nil {
			log.Fatalln("Unable to read/write data, copy failed!")
		}
	*/
}
